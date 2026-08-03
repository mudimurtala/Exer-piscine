# Day 4: Grazac Backend Interview, Complete Mock Interview Guide

## How to use this document

This is organized as a real interviewer would move through a conversation, starting from foundational understanding and building up to the specific decisions you made in the Inventory API project. Every answer follows the same shape: what the concept actually is, why it exists, then how it shows up in your project, with a real file or line you can point to. Read a section, cover the answer with your hand, try answering out loud in your own words, then check yourself. That active recall is far more useful than passive reading.

Sections are grouped by theme, not by day, since a real interview jumps around. At the end there is a rapid fire terminology section for last minute review the morning of.

---

## Section A: Node.js and the JavaScript Runtime

### Q: What is Node.js, and why use it for a backend?

Start with the foundation: JavaScript was originally built to run inside a browser only, reacting to clicks, updating pages, nothing more. Node.js is a runtime that takes the same JavaScript language and lets it run outside the browser, directly on a computer or server, so it can do things a browser never could, like reading files, opening database connections, or listening for network requests.

The specific thing that makes Node well suited to backend work is that it is non blocking and event driven. When Node asks the database for data, it does not freeze and wait, it registers "let me know when this is done" and moves on to handle other incoming requests in the meantime. When the database responds, Node comes back and finishes that piece of work. This means one Node process can juggle many simultaneous requests efficiently, without needing a separate thread for every single one.

In the project: every route handler in your controllers is `async`, and every database call uses `await`. That combination is exactly this non blocking model in practice, your server can be handling ten different people's requests to `/api/products` at once, each one picking up again the moment its own database query resolves, without one slow request freezing everyone else out.

### Q: What is the difference between synchronous and asynchronous code, and why does it matter here?

Synchronous code runs one line at a time, each line waiting for the previous one to fully finish before starting. Asynchronous code allows a slow operation, like a network call or database query, to happen in the background while the rest of the program keeps going, then comes back to that result later.

A useful way to hold this in your head: imagine a POS agent serving one customer's transfer, then standing there doing absolutely nothing until the bank confirms it, before serving the next customer in line. That is synchronous. Now imagine the same agent starting several transfers for different customers, and just handling each one the moment its confirmation comes back, out of order, whichever finishes first. That is asynchronous, and it is why one Node server can serve many people without needing many separate agents.

In the project: your `createStockMovement` function does several sequential awaited database calls inside one transaction, on purpose, they must happen in order since each depends on the last. But across different requests hitting your server at the same time, Node handles them concurrently, not sequentially.

### Q: What is npm, and what is the difference between dependencies and devDependencies?

npm, node package manager, is both a tool and a public registry of reusable code packages. Instead of writing your own password hashing algorithm or your own HTTP framework from scratch, you install a package someone already built, tested, and maintains.

`package.json` tracks exactly which packages your project needs, split into two lists. `dependencies` are packages your application needs to actually run, in production, right now, things like `express` and `pg`. `devDependencies` are tools only needed while you are developing, never on the live server, things like `jest`, `supertest`, and `nodemon`.

In the project: `npm install bcrypt jsonwebtoken` went into `dependencies`, since your running app genuinely needs them to hash passwords and issue tokens. `npm install --save-dev jest supertest` went into `devDependencies`, since a production server never needs to run your test suite, only your own machine or your CI pipeline does.

### Q: What does nodemon do, and why not just use `node` directly while developing?

Without nodemon, every time you change a line of code, you would have to manually stop the server (Ctrl+C) and restart it (`node src/server.js`) to see that change take effect. Nodemon watches your files, and the moment it detects a save, it automatically restarts the server for you.

In the project: `"dev": "nodemon src/server.js"` versus `"start": "node src/server.js"`. You use `dev` while building, `start` is what a production hosting platform would actually run, since production never needs the auto restart on save behavior, and you generally want production restarts to be deliberate, not automatic on every file change.

---

## Section B: Express.js and REST API Fundamentals

### Q: What is Express, and what problem does it solve?

Node on its own can technically handle web requests, but the raw tools for that are low level and repetitive, you would be manually parsing URLs, handling different HTTP methods, and writing response headers by hand for every single route. Express is a framework built on top of Node that gives you a clean, structured way to define routes, handle requests, and send responses, without reinventing that plumbing every time.

### Q: What is REST, and what makes an API RESTful?

REST, representational state transfer, is a set of conventions for designing APIs around resources, nouns like "products" or "categories", accessed through standard HTTP methods that describe the action.

`GET` retrieves data without changing anything. `POST` creates something new. `PUT` updates an existing thing, typically replacing it. `DELETE` removes something. The URL identifies which resource, the method describes what to do with it.

In the project: `GET /api/products` reads the list, `POST /api/products` creates one, `PUT /api/products/:id` updates one specific product, `DELETE /api/products/:id` removes one specific product. Same base URL, different meaning depending on the method, that consistency is the whole point of REST, anyone who understands the convention can guess how your API behaves without reading documentation first.

### Q: What is middleware in Express, and how does the request/response cycle actually work?

A request coming into an Express app does not go straight to its final handler. It passes through a chain of functions first, each one called middleware, each with the power to inspect the request, modify it, reject it outright, or pass it along to the next function in line.

Every middleware function receives three things: `req` (the incoming request), `res` (the tool for sending a response), and `next` (a function you call to hand control to whatever comes after you in the chain). If you never call `next()` and never send a response, the request simply hangs forever, which is a real, common bug to watch for.

In the project: `app.use(express.json())` is middleware that runs on every request, parsing incoming JSON request bodies into a usable JavaScript object before any route handler sees it. Your `authenticate` function is middleware that checks for a valid token before letting a request continue to a controller. Middleware can be global (applied to every request via `app.use`) or scoped to specific routes (like `router.post('/', authenticate, createProduct)`, where it only runs for that one route).

### Q: Explain the request/response cycle end to end, using an actual endpoint from your project.

Take `POST /api/products` with a valid token. First, the request arrives at Express. `express.json()` middleware parses the request body. The router matches the URL and method to `router.post('/', authenticate, createProduct)`. `authenticate` runs first, checking the `Authorization` header, verifying the JWT, and either rejecting the request with a 401 or attaching the decoded user to `req.user` and calling `next()`. Control passes to `createProduct`, which pulls fields out of `req.body`, runs a validation check, then queries PostgreSQL through the connection pool to insert a new row. Once the database confirms the insert, the controller builds a JSON response and sends it back with `res.status(201).json(...)`. The client receives that response, and the cycle for this request is complete.

### Q: What is the difference between a path parameter and a query parameter, and where does your project use each?

A path parameter is part of the URL structure itself, identifying a specific resource, like the `:id` in `/api/products/:id`. A query parameter comes after a `?` in the URL, typically used for filtering, sorting, or optional options, like `?sort=price`.

In the project, you use path parameters throughout, `/api/products/:id`, `/api/stock-movements/:product_id`, since every one of those routes is about one specific, identified resource, not a filtered search across many.

### Q: What HTTP status codes did you use in this project, and what does each actually communicate?

`200 OK`, a request succeeded, typically for GET, PUT, or DELETE. `201 Created`, something new was successfully created, used specifically after POST requests that add a new row. `400 Bad Request`, the client sent something invalid, like a missing required field, or a stock movement that would take quantity below zero. `401 Unauthorized`, the server does not know who is making this request, no token, or an invalid or expired one. `403 Forbidden`, the server knows exactly who you are, but you are not allowed to perform this specific action, your staff versus admin delete restriction. `404 Not Found`, the requested resource does not exist, a product id that was never created. `409 Conflict`, the request conflicts with existing data, a duplicate username or duplicate SKU. `500 Internal Server Error`, something broke on the server's side that was not the client's fault, an unexpected database error.

Knowing the difference between 401 and 403 specifically is a very common interview probe, since a lot of junior developers use them interchangeably. 401 means "prove who you are." 403 means "I know who you are, the answer is still no."

---

## Section C: PostgreSQL and Relational Database Fundamentals

### Q: What is a relational database, and why choose PostgreSQL over something like MongoDB for this project?

A relational database organizes data into tables with a fixed structure, columns with defined types, and enforces relationships between tables through keys. This is a strong fit when your data naturally has clear relationships and rules that must always hold true, a stock movement always belongs to exactly one product, a product's price must never be negative, and so on.

MongoDB, by contrast, is a document database, more flexible, less rigid structure, often a better fit when your data shape varies a lot between records or you are optimizing for very high write throughput over strict consistency.

For an inventory system specifically, the relationships and rules matter enormously, you genuinely never want a stock movement pointing at a product that does not exist, or a negative price slipping through. PostgreSQL enforces those rules for you at the database level, not just in your application code, which is a meaningfully stronger guarantee.

### Q: What is a primary key and a foreign key?

A primary key uniquely identifies one row in a table, no two rows can share one, and it can never be empty. In your schema, `id SERIAL PRIMARY KEY` on every table, Postgres auto increments this for you.

A foreign key is a column in one table that references the primary key of another table, and that is exactly how relationships between tables are enforced. `products.category_id` references `categories.id`. `stock_movements.product_id` references `products.id`. The database itself will refuse to let you insert a stock movement pointing at a product id that does not exist, that is a rule enforced at the data layer, not something your application code has to remember to check.

### Q: Explain `ON DELETE SET NULL` versus `ON DELETE CASCADE`, and why you chose differently for each relationship.

Both describe what should happen to rows in a child table when the row they reference in the parent table gets deleted.

`ON DELETE SET NULL`: the child row survives, but the reference gets cleared to null. You used this for `products.category_id`. If a category is deleted, the products that belonged to it are not deleted, they still exist, they simply become uncategorized. That reflects reality, deleting the "Electronics" category should not delete every electronic product you own.

`ON DELETE CASCADE`: the child row is deleted along with the parent. You used this for `stock_movements.product_id`. If a product is deleted entirely, its historical stock movement records no longer mean anything on their own, they exist only in relation to that product, so it makes sense for them to go too.

This distinction, choosing the right deletion behavior per relationship rather than applying one rule everywhere, is a genuinely good thing to be able to explain, it shows you thought about what each relationship actually means, not just how to make foreign keys work syntactically.

### Q: What are constraints, and what constraints did you actually use?

Constraints are rules the database enforces automatically, so invalid data is rejected at the point of insertion, rather than relying on every piece of application code remembering to check.

`NOT NULL`: a column can never be left empty, used on things like `name`, `sku`, `password_hash`, values that genuinely must always exist. `UNIQUE`: no two rows can share the same value in that column, used on `username`, `sku`, `categories.name`. `CHECK`: a custom rule, used on `price NUMERIC(10,2) NOT NULL CHECK (price >= 0)`, meaning the database itself will refuse to store a negative price, even if a bug in the application code somehow tried to insert one.

### Q: What is a JOIN, and what is the difference between an INNER JOIN and a LEFT JOIN?

A JOIN combines rows from two related tables into one result, based on a matching condition, typically a foreign key relationship.

`INNER JOIN` only returns rows where a match exists in both tables. `LEFT JOIN` returns every row from the left, "main," table regardless of whether a match exists on the right, filling in null for any columns that have no match.

In the project, `getProducts` uses `LEFT JOIN categories ON products.category_id = categories.id`. This is deliberate, since `category_id` is nullable, a product with no category assigned should still show up in the product list, just with `category_name` coming back as null, rather than silently disappearing from the results. An INNER JOIN here would have been a subtle bug, quietly hiding uncategorized products from every product listing.

### Q: What is parameterized querying, and why does it matter?

A parameterized query separates the SQL structure from the actual data values, using placeholders like `$1`, `$2` that get safely substituted in by the database driver, rather than directly concatenating user input into a raw SQL string.

```javascript
pool.query('SELECT * FROM users WHERE username = $1', [username]);
```

This is the standard defense against SQL injection, an attack where a malicious user crafts input designed to be interpreted as SQL commands rather than plain data, for example typing something like `' OR '1'='1` into a login field, which could trick a naively built query into returning every row in the table. Because the value is passed separately from the query structure, the database always treats it strictly as data, never as executable SQL, no matter what it contains.

---

## Section D: Transactions, Concurrency, and Race Conditions

### Q: What is a database transaction, and why did you wrap stock movements in one?

A transaction groups multiple database operations into a single all or nothing unit. Either every operation inside it succeeds and gets permanently saved with `COMMIT`, or if anything fails partway through, everything is undone with `ROLLBACK`, as if none of it had ever happened.

Recording a stock movement genuinely involves two separate writes, updating the product's quantity, and inserting a new row into `stock_movements`. Without a transaction, it is possible for the first write to succeed and the second to fail, for example if the server crashed at exactly the wrong moment, leaving the product's quantity changed with no record of why. The transaction guarantees that either both writes happen together, or neither does, the data can never be left in a half updated, inconsistent state.

### Q: What is a race condition, and how did you actually prevent one?

A race condition happens when two operations run at nearly the same time and interfere with each other in a way that produces an incorrect result, because each one read the data before the other had finished writing its update.

Imagine a product has 5 units left, and two customers both buy the last item at almost the same instant. Both requests read "quantity is 5" before either write happens. Both calculate 5 minus 1 equals 4. Both write "4" back. The real answer should have been 3, since two units actually sold, but the second write just overwrote the first, and one sale silently vanished from the truth.

You prevented this using `SELECT * FROM products WHERE id = $1 FOR UPDATE` inside the transaction. `FOR UPDATE` places a row level lock on that specific product the moment it is read. Any other transaction trying to read that same row with `FOR UPDATE` has to wait until the first transaction fully commits or rolls back. This forces the two competing requests to happen one at a time for that specific row, not actually in parallel, so the second request's calculation is always based on the true, already updated quantity, never a stale value.

### Q: Is `FOR UPDATE` going to slow down your whole application under load?

This is a smart thing to preempt in an interview even if not asked directly. The lock is scoped to one specific row, the one product being updated, not the whole table. Two different customers buying two different products at the same time are not blocked by each other at all, only two requests trying to touch the exact same product row at the exact same moment briefly queue behind each other, which is precisely the scenario where correctness genuinely matters more than raw speed.

---

## Section E: Password Security and JWT Authentication

### Q: Why hash passwords instead of encrypting them?

Encryption is reversible by design, if you can encrypt data you can also decrypt it back to the original given the right key. That means anyone who ever gets hold of that key, including an attacker who steals it, could recover every real password. Hashing, specifically with bcrypt, is one way, there is no operation that turns a hash back into the original password. To check a login attempt, you never decrypt anything, you hash the new attempt the same way and compare the two hashes.

### Q: Walk through exactly what `bcrypt.hash(password, 10)` produces and why the number 10 matters.

The output looks like `$2b$10$ddhOrl9FlJ7Ph3Wbl1TuBO8GAcI8kHNgz/VqJCmrE7pHMWR4RYDWe`. `2b` identifies the bcrypt algorithm version. `10` is the cost factor, the number of internal hashing rounds performed, and it is deliberately configurable, higher means slower to compute but harder to brute force. The remainder is a randomly generated salt followed by the actual hash output.

The salt is the part most people forget to mention, and it matters a lot: it means two different users who happen to choose the identical password will end up with completely different stored hashes. Without a salt, an attacker could precompute hashes for common passwords once, then instantly match them against every row in any leaked database ever, a technique called a rainbow table attack. The random salt makes that precomputation worthless.

### Q: What is a JWT, structurally, and what goes in each part?

A JSON Web Token is one string made of three parts joined by dots, each base64 encoded: `header.payload.signature`.

The header states which signing algorithm was used. The payload holds the actual data you chose to embed, in your case `userId`, `username`, `role`, `iat` (issued at time), `exp` (expiry time). The signature is produced by combining the header and payload with a secret key only the server knows, and it is the part that actually protects the token, since anyone can decode and read the payload, base64 is an encoding, not encryption, but nobody can produce a valid matching signature without knowing the secret.

A very common trap question: "is it safe to put a password inside a JWT payload?" No, never, since the payload is trivially readable by anyone who has the token.

### Q: What actually happens inside `jwt.verify`, and what does it protect against?

`jwt.verify` performs two checks at once. First, it recomputes the signature using the server's secret and confirms it matches the signature on the token, proving the token was genuinely issued by this server and has not been altered in any way since. Second, it checks the `exp` field against the current time and rejects the token if it has expired. If either check fails, it throws an error, which your `authenticate` middleware catches and turns into a 401 response.

This protects against forgery, someone crafting a fake token claiming to be an admin, since they would need your secret to produce a matching signature, and against stale access, an old, possibly compromised token being usable forever.

### Q: Why did you set the token expiry to one hour, and what's the tradeoff?

A shorter expiry limits the window of damage if a token is ever stolen, an attacker with a stolen token can only impersonate that user for up to an hour before it stops working. The tradeoff is user convenience, a very short expiry means people have to log in more often. Real production systems often solve this with a two token system, a short lived access token like this one for actual API requests, paired with a longer lived refresh token, stored more carefully, used only to silently obtain a new access token without asking the user to log in again.

---

## Section F: Middleware Patterns and Authorization

### Q: What is the difference between authentication and authorization, precisely?

Authentication answers "who are you," confirming the identity behind a request is genuine. Authorization answers "are you allowed to do this," a separate question asked only after identity is already established.

In the project, `authenticate` handles the first question, verifying the JWT signature and expiry. `authorizeRoles('admin')` handles the second, checking whether the now known, verified user's role permits this specific action. A request can pass authentication completely and still fail authorization, exactly what happens when a valid, logged in staff account attempts to delete a product, a route reserved for admin.

The status codes reflect this distinction precisely: 401 for a failed authentication ("we don't know who you are"), 403 for a failed authorization ("we know exactly who you are, the answer is still no").

### Q: Explain the middleware factory pattern you used for `authorizeRoles`, and why it's written that way instead of a plain function.

```javascript
const authorizeRoles = (...allowedRoles) => {
  return (req, res, next) => {
    if (!allowedRoles.includes(req.user.role)) {
      return res.status(403).json({ error: 'Insufficient permissions for this action' });
    }
    next();
  };
};
```

Express expects every middleware function to have the exact signature `(req, res, next)`. But `authorizeRoles` needs to know something extra, which specific role or roles are allowed on this particular route, and that information is different for every route it gets used on. So `authorizeRoles` itself is not middleware, it is a function that accepts the allowed roles and returns a brand new middleware function, one that "remembers" those roles through closure, and matches Express's expected signature exactly.

This is what makes `authorizeRoles('admin')` and, hypothetically, `authorizeRoles('admin', 'manager')` both work cleanly on different routes using the exact same underlying logic, just configured differently at the point each route is defined.

### Q: Where does `req.user` come from, and why do you trust it?

It is set inside `authenticate`, `req.user = decoded`, where `decoded` is the payload extracted from a JWT that has just been cryptographically verified. Nothing about `req.user` comes from anything the client directly typed into the request body, it comes only from data that was signed by the server itself at login time and just proven authentic by `jwt.verify`. This is why `getMe` and `authorizeRoles` can both safely trust `req.user.userId` and `req.user.role`, without needing to re verify anything themselves, the verification already happened earlier in the middleware chain.

---

## Section G: Third Party API Integration

### Q: Walk through the general shape of integrating any third party service, using your Resend integration as the example.

Sign up with the provider and obtain an API key. Store that key as an environment variable, `RESEND_API_KEY` in `.env`, never hardcoded directly in source code and never committed to version control. Install the provider's SDK, `npm install resend`. Wrap the actual call to their service inside your own dedicated function, `sendLowStockAlert` in `src/services/emailService.js`, isolated from your core business logic. Finally, deliberately decide how a failure in that integration should, or should not, affect the rest of your application.

### Q: Why does the email sending logic live in its own file, separate from the stock controller?

This is a service layer, a deliberate separation of concerns. `stockController.js` is responsible for the business rule of what counts as a low stock event. `emailService.js` is responsible only for the mechanics of how an email actually gets sent through Resend. The controller calls `sendLowStockAlert(product)` without knowing or caring how that function accomplishes its job. If Resend were ever swapped for a different provider, or a second notification channel like Slack were added, only `emailService.js` would need to change, the controller's logic stays completely untouched.

### Q: Why don't you `await` the call to `sendLowStockAlert`, and why does it swallow its own errors instead of throwing?

The stock update itself has already fully committed to the database by the time this call happens, that is the critical, must succeed part of the request. The email is a secondary, non critical notification layered on top of an already successful operation. Not awaiting it, sometimes called fire and forget, means the client gets their response the moment the important work is done, without waiting an extra second or two for an outbound network call to a third party service to finish.

The internal `try/catch` inside `sendLowStockAlert` means that if Resend fails for any reason, a network issue, an outage, an invalid key, the error is caught and logged quietly rather than thrown back up the call stack. This matters a lot: a failure in a non critical side effect should never be able to break or roll back a core operation that already genuinely succeeded. This is not a hypothetical either, you saw it happen for real during a WSL DNS outage, the stock update still completed and returned normally while the email silently failed in the background and logged the error.

### Q: What is idempotency, and does your low stock alert logic relate to it?

Idempotency describes an operation that produces the same end result no matter how many times it is repeated. Your threshold crossing guard is closely related in spirit, though not a textbook idempotency example, it is really about avoiding redundant side effects for a state that hasn't meaningfully changed. By comparing the quantity before and after a movement, you only trigger a new alert on a genuine transition from above threshold to at or below it, not on every single low stock movement afterward. This is the same instinct behind idempotency, avoiding unnecessary repeated effects for what is functionally the same underlying event.

---

## Section H: Testing

### Q: What is the difference between a unit test and an integration test?

A unit test isolates one small, specific piece of logic, typically a single function, and tests it alone, often mocking out anything it depends on, like a database or an external API, so the test is fast and only ever fails because of that one function's own logic.

An integration test exercises multiple layers working together, closer to how a real user or client actually experiences the system. Your tests are integration tests, using Supertest to send real HTTP requests through your actual Express app, exercising routing, middleware, controllers, and a real test database together, in one pass.

### Q: Why did you choose integration tests over unit tests for this project?

For an API specifically, the most valuable thing to prove is that the whole request path behaves correctly end to end, that a request with no token really does get rejected by the actual middleware chain, that a duplicate SKU really does get caught by the actual database constraint, not just that one isolated function returns the right value in theory. Given the time available, integration tests gave the strongest, most realistic confidence per test written.

### Q: Explain `describe`, `it`, `beforeEach`, and `afterAll` and what role each plays.

`describe` groups related tests under one readable label, so output reads like "POST /api/auth/login should reject an incorrect password" rather than an unlabeled flat list. `it` (or `test`) is one individual test case, checking one specific expectation. `beforeEach` runs before every single test within its scope, most often used to reset data so no test can leak state into another, your `DELETE FROM users` before every auth test is a direct example. `afterAll` runs exactly once, after every test in the file has finished, used for final cleanup, your `pool.end()` call, which properly closes the database connection so Jest can exit cleanly instead of hanging.

### Q: Why does your test suite use a completely separate database instead of your development database?

Tests genuinely create, modify, and delete real rows as part of proving behavior. Running that against the same database holding your actual development data would mean every single test run risks wiping or corrupting real data you're relying on elsewhere. A second, identically structured database, `inventory_api_test`, loaded from the same `schema.sql`, gives tests a completely safe space to freely create and destroy data without any risk to real information.

### Q: How does Jest actually know to use the test database instead of the real one?

A `jest.setup.js` file, wired in through the `jest` config block in `package.json`, runs before any test file executes and explicitly loads `.env.test` instead of the default `.env`. This works cleanly alongside the existing `db/index.js`, which still calls a plain `dotenv.config()` with no path, because dotenv, by default, never overrides an environment variable that has already been set. Since the test values are already loaded first, `db/index.js`'s own later call to load `.env` simply has nothing left to override.

### Q: Why does your `test` script include `--runInBand`?

Jest normally runs test files in parallel, across multiple processes, for speed. All of your tests share one single real Postgres test database, and running them in parallel risks race conditions between test files, for example two different test files both trying to insert a user with the same username at nearly the same moment. `--runInBand` runs test files strictly one after another instead, deliberately trading some raw speed for predictable, conflict free results, which matters more here than execution time.

### Q: Give a concrete example of a test that proves your authorization logic actually works, not just your authentication.

The DELETE product tests are the clearest example. One test logs in as a staff account, attempts to delete a product, and asserts the response is `403`. A second test logs in as an admin account, performs the identical delete, and asserts it succeeds with `200`. Both requests carry a completely valid, correctly authenticated token, the only difference is the role attached to that token, which is exactly what proves the authorization layer, not just the authentication layer, is functioning correctly and independently.

---

## Section I: Documentation

### Q: What makes documentation, like your README, actually good rather than just present?

Good documentation respects the reader's time and assumes nothing about what they already know regarding this specific project. It should let someone go from "I've never seen this repository before" to "I have it running locally and understand what it does" using nothing but the document itself, no separate conversation required. Structurally, that means leading with a clear description of what the project actually does, followed by exact, in order setup steps, then a reference section for anything someone might need to look up later, environment variables, every endpoint's expected input and output, rather than making them dig through source code to find those answers.

### Q: Why does your README use generic placeholders instead of your own real values in setup instructions?

A README is a template meant for anyone setting the project up fresh on their own machine, not a personal record of one specific configuration. Using real values, and especially real secrets like an actual password or API key, in a public facing document is a genuine security exposure, particularly once the repository lives on GitHub where anyone can view it. Placeholders like `your_postgres_username` communicate the required shape of the value without leaking anything real.

---

## Section J: Environment and Tooling (WSL, npm, git)

### Q: What is a `.env` file, and why is it excluded from git?

A `.env` file holds configuration values and secrets specific to one environment or machine, database credentials, API keys, signing secrets, kept separate from the actual source code. It is excluded from version control via `.gitignore` because committing it would mean pushing real secrets into a shared, often public, repository, permanently visible in that repository's history even if later deleted. Anyone cloning the project is expected to create their own `.env` locally, following the README, with their own values.

### Q: Describe a real debugging problem you hit in this project and how you diagnosed it, not just fixed it.

The Resend integration initially failed with "Unable to fetch data, the request could not be resolved." Rather than guessing, the diagnosis was methodical: first confirmed raw internet connectivity worked at all with `ping 8.8.8.8`, a raw IP address requiring no name lookup, and it succeeded. Then tested the actual failing case, `curl -v https://api.resend.com`, a domain name requiring DNS resolution, and that failed with "Could not resolve host." That gap, raw IP connectivity working while domain name resolution failed, isolated the problem precisely to DNS, not a general network outage.

The root cause turned out to be WSL's auto generated `/etc/resolv.conf` pointing at a DNS server that was not resolving names correctly. The fix involved creating `/etc/wsl.conf` to stop WSL from regenerating that file automatically, replacing it manually with Google's public DNS servers, then performing a full `wsl --shutdown` from outside WSL entirely, since this was a system level networking change, not an application level one, and simply closing a terminal window leaves the underlying WSL virtual machine still running with the old configuration in the background.

This is a strong story to tell in an interview specifically because it demonstrates a real debugging methodology, isolate what works from what doesn't, form a specific hypothesis from that gap, and verify the fix, rather than just describing a fix that happened to work.

---

## Section K: Project Walkthrough Questions

These are the "tell me about a project you built" style questions, where the goal is a clear, confident narrative, not a feature list.

### Q: Tell me about a backend project you've built recently.

A strong shape for this answer: what it does, why you built it that way, one specific technical decision you're proud of, and one real problem you hit and solved.

"I built a small inventory and stock management API using Node, Express, and PostgreSQL, covering categories, products, and stock movements. The part I'm most proud of is the stock movement logic, it runs inside a database transaction with row level locking, specifically to prevent two simultaneous sales of the same low stock item from corrupting the quantity, which is a real concurrency problem, not just a CRUD exercise. I also added JWT authentication with role based authorization, so only admins can delete products, and integrated a real third party email API that alerts automatically when stock crosses below a threshold, with a guard so it only fires once per event rather than spamming repeated alerts. I backed all of it with an automated Jest and Supertest suite running against an isolated test database, and wrote a full README so someone else could set it up from scratch. Along the way I hit a real WSL DNS issue that blocked outbound API calls entirely, which I diagnosed methodically rather than guessing, and fixed at the system configuration level."

### Q: What was the hardest part of building this?

Being honest and specific here lands far better than a generic answer. A genuinely strong, true answer available to you: the race condition protection in the stock movement logic, since it required actually understanding what could go wrong with two simultaneous requests, not just writing code that looked correct in isolation, and the WSL DNS debugging, since the error message alone didn't point directly at the actual cause, it took a structured process of elimination to find it.

### Q: If you had another week, what would you add or improve?

Good answers to have ready: a refresh token system alongside the current short lived access token, for a smoother login experience without shortening security. Rate limiting on the login endpoint, to slow down brute force password guessing attempts. Input validation using a library like Zod or Joi, to catch malformed requests earlier and more thoroughly than manual `if` checks. Pagination on the products list endpoint, since a real catalog could eventually hold thousands of rows. Structured logging instead of plain `console.log`, for easier debugging in a real production environment.

---

## Section L: Rapid Fire Terminology Glossary

Use this the morning of the interview for a fast final pass. Cover the right column and see if the left term alone brings back a full, confident sentence.

**API**: a defined way for two pieces of software to communicate, in this project specifically an HTTP based REST API.

**Endpoint**: one specific URL and method combination that performs one specific action, `POST /api/auth/login`.

**Middleware**: a function that runs between an incoming request and its final handler, able to inspect, modify, reject, or pass along the request.

**JWT**: a signed token proving a user's identity, without the server needing to look anything up on every single request.

**bcrypt**: a one way password hashing algorithm, includes a random salt, never reversible.

**Salt**: random data mixed into a password before hashing, so identical passwords never produce identical hashes.

**Transaction**: a group of database operations that either all succeed together or all get undone together.

**Race condition**: an error caused by two operations happening at nearly the same time, each unaware of the other's in progress change.

**Row level lock**: a lock placed on one specific database row, preventing other transactions from reading or writing it until the first is done.

**Foreign key**: a column referencing another table's primary key, enforcing that relationships in the data are always valid.

**CASCADE vs SET NULL**: what happens to related child rows when a parent row is deleted, delete them too, or just clear the reference.

**Parameterized query**: separating SQL structure from user supplied data, the core defense against SQL injection.

**Authentication**: confirming who is making a request.

**Authorization**: confirming whether that known, verified requester is allowed to do this specific thing.

**Environment variable**: a configuration value kept outside source code, different per machine or environment, never committed to git.

**Unit test**: tests one isolated piece of logic alone.

**Integration test**: tests multiple layers working together, closer to real usage.

**Fire and forget**: starting an asynchronous operation without waiting for it to finish, used when its result isn't critical to the current response.

**Idempotency**: an operation that produces the same end result no matter how many times it's repeated.

**Service layer**: a dedicated file or module isolating how a specific integration or capability works, kept separate from core business logic.

---

## A note on how to actually use this before Tuesday

Read through it once fully today, out loud where you can, since that engages recall differently than silent reading. Tomorrow, go through it again, but cover the answers and try to speak each one from memory first, only checking after you've made a genuine attempt. The goal isn't to memorize these exact sentences, it's to own the underlying understanding well enough that you could explain any of this to someone completely new to the concept, in your own words, with the calm, composed pacing you've been building through Toastmasters. That's what will actually come through in the interview.
