# Inventory API — Day 1 Learning Notes and Reference

A complete beginner friendly breakdown of everything we touched today, why we touched it, and how to talk about it in an interview. Read this slowly, more than once. You do not need to memorize it, you need to understand the shape of it.

---

## 1. The Big Picture First

Before any terminology, hold this one sentence in your head, everything else hangs off it:

**A backend receives a request, talks to a database, and sends back a response.**

That is the entire job. Express is the tool that receives the request. PostgreSQL is where the data lives. Everything we installed and every command we ran exists only to make that one loop work smoothly and safely.

---

## 2. Core Terminology, Explained Simply

**Node.js**
Not a framework, an engine. Normally JavaScript only runs inside a browser. Node.js lets that same JavaScript run directly on a computer or server, outside the browser. This is what makes a JavaScript backend possible at all.

**Express**
A library built on top of Node.js that makes handling web requests easy. Without Express you would have to manually parse every incoming request byte by byte. Express gives you simple tools like `app.get()` and `app.post()` instead.

**PostgreSQL (often shortened to Postgres)**
A relational database. "Relational" means data is stored in tables that can reference each other, like your `products` table referencing your `categories` table. This is different from something like Firebase Firestore, which stores loose documents with no enforced structure between them.

**npm**
Node Package Manager. It is how you download and manage external code libraries (called packages) that other people wrote, so you do not have to write everything from scratch.

**package.json**
A file that lists your project's name, version, dependencies, and the scripts you can run. Think of it as your project's ID card and instruction sheet combined.

**node_modules**
The actual folder where all your installed packages physically live. It gets huge and is never uploaded to GitHub, which is exactly why we put it in `.gitignore`.

**Middleware**
A function that runs in between a request arriving and your route logic handling it. `express.json()` is middleware, it reads incoming JSON data and makes it usable before your code even sees the request. Think of middleware as a checkpoint every request passes through first.

**API (Application Programming Interface)**
A defined set of URLs your server responds to, so other programs (a frontend, a mobile app, another service) can talk to your backend in a predictable way.

**REST API**
An API that organizes its URLs around resources (like `/products`) and uses standard HTTP verbs to act on them: GET to read, POST to create, PUT to update, DELETE to remove. This is the pattern we used all day.

**Environment variables**
Settings and secrets (like a database password) that live outside your code, in a `.env` file, and get loaded in at runtime. This means your code never has secrets typed directly into it.

**dotenv**
The npm package that reads your `.env` file and loads those values into `process.env` so your code can use them.

**nodemon**
A development tool that watches your files and automatically restarts your server every time you save a change. Without it you would have to manually stop and restart the server after every single edit.

**pg**
The official Node.js driver for PostgreSQL. It is the translator that lets your JavaScript code send SQL commands to Postgres and get results back.

**Connection pool**
Instead of opening a brand new connection to the database for every single request (slow, wasteful), a pool keeps a small set of connections open and ready to reuse. This is why we created `pool` in `src/db/index.js`.

---

## 3. What Actually Happened During Setup, Command by Command

**`sudo apt install postgresql postgresql-contrib -y`**
This installed the actual Postgres database engine onto your WSL Ubuntu machine. `postgresql-contrib` adds some extra official utilities. The `-y` flag just means "yes, don't ask me to confirm, proceed automatically."

**`sudo service postgresql start`**
This started the Postgres process running in the background so it can accept connections. On WSL specifically, this does not happen automatically when you open a new terminal, so remember to run it each fresh session.

**Where does the database actually live?**
On disk, at `/var/lib/postgresql/16/main`, inside your WSL Linux filesystem. You never touch that folder directly, you always go through Postgres itself using commands like `psql`.

**`sudo -u postgres createuser --superuser ibnmuhyideen`**
Postgres has its own separate list of users, completely different from your Linux login. This command told Postgres "create a user (called a role) named ibnmuhyideen, and give it full admin rights." Without this, Postgres would not know who you are.

**`sudo -u postgres psql -c "ALTER USER ibnmuhyideen PASSWORD 'devpassword123';"`**
This set a password on that new Postgres role. We need this because our Node.js code will need to log in with a username and password, it cannot use the same shortcut your terminal uses.

**Why did `psql -d postgres` work with no password prompt?**
Because Postgres on Ubuntu defaults to something called peer authentication for local connections, if your Linux username matches your Postgres username exactly, it trusts you automatically when connecting from the same machine. Your Node.js app later will not get this shortcut, it will always need the actual password from your `.env` file.

**`createdb inventory_api`**
Created an empty, brand new database inside Postgres, ready to hold our tables.

**`npm init -y`**
Generated a fresh `package.json` file. The `-y` flag skips the usual round of questions (project name, version, license) and just accepts all the defaults instantly.

**`npm install express pg dotenv`**
Downloaded three packages and added them into `node_modules`, and recorded them in `package.json` under `"dependencies"`, meaning your project needs these to actually run.

**`npm install --save-dev nodemon`**
Same idea, but `--save-dev` records it under `"devDependencies"` instead, meaning it is a tool you need while building, but it is not required for the finished app to run in production.

---

## 4. Databases and Tables, Explained

**Table**
A structured grid of data, made of rows and columns, similar in concept to a spreadsheet tab. We created three: `categories`, `products`, `stock_movements`.

**Schema**
The blueprint that defines what each table looks like, its columns, their types, and the rules attached to them. `src/db/schema.sql` is our blueprint file.

**Primary key**
A column that uniquely identifies each row, no two rows can ever share one. Our `id SERIAL PRIMARY KEY` does this, `SERIAL` means Postgres auto generates a new incrementing number every time a row is added, so you never have to supply it yourself.

**Foreign key**
A column in one table that points to the primary key of another table, this is what creates a relationship. `category_id INTEGER REFERENCES categories(id)` in our `products` table means "this number must match an actual id that exists in categories."

**Constraints**
Rules the database enforces automatically, no matter what code tries to insert data. Examples from our schema:
- `NOT NULL` — this value can never be empty
- `UNIQUE` — no two rows can share this value (used on `sku` and category `name`)
- `CHECK (price >= 0)` — rejects any attempt to insert a negative price
- `DEFAULT NOW()` — automatically fills in the current timestamp if none is given

**`ON DELETE SET NULL` vs `ON DELETE CASCADE`**
These define what happens to related rows when something they depend on is deleted. `SET NULL` (used on category deletion) leaves the product alive but clears its category. `CASCADE` (used on product deletion) deletes the product's stock history along with it, since history for a product that no longer exists is meaningless.

**Confirming your tables exist**
`psql -d inventory_api -c "\dt"` connects to your database and runs the special psql command `\dt`, which lists every table. This is your quick sanity check any time you want to confirm your schema actually applied.

**Loading a schema file**
`psql -d inventory_api -f src/db/schema.sql` tells psql "connect to inventory_api and run every SQL command inside this file." The `-f` flag means "from file."

---

## 5. Writing SQL Safely From Node.js

**Parameterized queries**
Instead of building a query by gluing text together, we write placeholders like `$1`, `$2` and pass the actual values in a separate array:

```javascript
pool.query('SELECT * FROM categories WHERE id = $1', [id]);
```

This matters enormously. If we instead wrote `` `SELECT * FROM categories WHERE id = ${id}` `` and someone passed in a malicious string as `id`, they could potentially rewrite our entire query and access or destroy data that was never meant to be exposed. This attack is called SQL injection, and it is one of the most well known security vulnerabilities in all of web development. Parameterized queries are Postgres's own built in defense against it.

**Transactions**
A transaction groups several database operations so they succeed or fail together, as a single unit. Our stock movement logic uses:
- `BEGIN` — start the group
- `COMMIT` — everything succeeded, make it permanent
- `ROLLBACK` — something went wrong, undo everything in this group as if none of it happened

Without this, if your server crashed halfway through updating a product's quantity and recording its history, you could end up with a quantity number that no longer matches its own history, permanently broken data.

**`FOR UPDATE`**
Added to a SELECT inside a transaction, it locks that specific row until the transaction finishes. This solves what is called a race condition: imagine two sales happening on the same product within the same second. Without the lock, both could read the same starting quantity, both calculate independently, and the second write would silently erase the first. `FOR UPDATE` forces the second request to politely wait its turn.

---

## 6. Understanding curl

`curl` is a command line tool for sending HTTP requests, exactly the kind of request a browser or a frontend app would normally send, except you are typing it directly. It is how we tested our API without needing a frontend at all.

Breaking down a command we used:

```
curl -X POST http://localhost:5000/api/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Wireless Mouse", "sku": "WM-001"}'
```

- `-X POST` — tells curl which HTTP method to use, in this case, POST (create something new). If you omit `-X`, curl defaults to GET.
- `-H "Content-Type: application/json"` — adds a header, this specific header tells the server "the data I'm sending you is in JSON format, please parse it that way."
- `-d '{...}'` — the actual data being sent in the request body.

A plain `curl http://localhost:5000/api/categories` with no flags automatically sends a GET request, which is why we could just read data with the bare command.

---

## 7. Project Structure, and Why We Split Files This Way

```
src/
  db/            → talks to the database
  routes/        → defines which URL goes to which function
  controllers/   → holds the actual logic for each route
  app.js         → builds the Express app and connects everything
  server.js      → actually starts the server listening
```

This separation is not decoration, it is a real, common professional pattern. If a senior developer asks you to change how a specific database query behaves, you know to look in `controllers`. If they want a new URL added, you look in `routes`. Keeping concerns separated like this is exactly what "maintainable code" means in a job description.

`app.js` versus `server.js` specifically: `app.js` builds the Express application itself, `server.js` is the file that actually starts it listening on a port. Separating these two lets you later write automated tests that use the app without needing a live running server, which becomes relevant once we add testing.

---

## 8. What You Actually Built Today, In Plain English

You built a working backend with a real database behind it. Specifically:

- A PostgreSQL database with three related tables, enforcing real business rules (unique SKUs, non negative prices)
- Full create, read, update, delete for categories and products, including a proper join so products show their category name
- A stock movement system that safely updates quantity, keeps a permanent history, and protects against two sales happening at the same time corrupting your data
- Meaningful error responses (404 for not found, 409 for conflicts, 400 for bad input) instead of the server just crashing

---

## 9. Interview Preparation, Beginner Level

Below are likely questions, explained with the specific terminology broken down first, then a sample answer in plain language you could actually say out loud.

**Q: Walk me through what happens when a request hits your API.**
*Terminology check: "request" just means the incoming ask, like "give me all products." "Response" is what you send back.*
Sample answer: "The request first passes through Express, which checks the route, so for example GET slash products. That matches to a specific controller function. Inside that function, I use the pg package to run a SQL query against Postgres, wait for the result, and then send it back to the client as JSON."

**Q: How do you prevent SQL injection?**
*Terminology check: SQL injection is when someone sneaks malicious code into a value they submit, trying to trick your database into running commands you never intended.*
Sample answer: "I use parameterized queries, so instead of building the SQL string by directly inserting user input, I use placeholders like dollar one, dollar two, and pass the actual values separately. Postgres handles escaping them safely, so user input can never be interpreted as part of the SQL command itself."

**Q: What is a database transaction, and when would you use one?**
Sample answer: "A transaction groups multiple database operations so they either all succeed or none of them apply. I used one when recording a stock movement, because I'm updating the product's quantity and inserting a history record at the same time. If one of those failed partway through, I roll back the whole thing so the data never ends up in an inconsistent state."

**Q: How would you handle two people trying to buy the last item in stock at the same time?**
*Terminology check: this is called a race condition, when timing between two operations causes an unexpected or incorrect result.*
Sample answer: "I lock the specific row using FOR UPDATE inside the transaction. That way, if two requests come in nearly simultaneously, the second one has to wait until the first transaction finishes, so they can't both read the same stale quantity and both proceed as if the stock still existed."

**Q: What's the difference between SQL and NoSQL databases?**
Sample answer: "SQL databases like Postgres store data in structured tables with defined relationships between them, enforced by the database itself. NoSQL databases like Firestore, which I've used before on the frontend, store more flexible documents without enforced structure between them. SQL tends to fit data with clear relationships, like an inventory system, better."

**Q: Why did you separate your code into routes, controllers, and a database layer?**
Sample answer: "It keeps each part of the code responsible for one thing. Routes just decide which URL maps to which function. Controllers hold the actual business logic. The database layer only knows how to talk to Postgres. It makes the codebase easier to navigate and change without breaking unrelated parts."

**Q: What is middleware?**
Sample answer: "It's a function that runs in between a request coming in and it actually reaching my route logic. For example, express dot json is middleware that reads incoming JSON data before my controller even sees the request, so I can use it directly as req dot body."

---

## 10. What's Coming Next

Day two will cover JWT authentication (so not just anyone can modify your data), and the third party email alert integration for low stock, which directly matches a specific line in Grazac's job description. Day three covers automated testing and documentation.

You'll get a notes file like this one at the end of each day. Keep them together, they'll build into a full personal reference by Tuesday.
