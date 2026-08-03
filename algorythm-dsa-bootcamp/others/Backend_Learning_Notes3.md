# Day 3 Notes: Third Party Email Integration, Automated Testing, and Documentation

## What we built today

Today covered three of the Grazac job responsibilities directly: "support integration of third party APIs and external services," "write unit, integration, and automated tests," and "write and maintain technical documentation."

Concretely:

1. Integrated Resend, a third party transactional email API, to send real low stock alert emails
2. Fixed a WSL DNS resolution issue that was blocking the integration
3. Added a "crossed threshold" guard so alerts only fire once per low stock event, not on every subsequent movement
4. Set up Jest and Supertest with an isolated test database, and wrote 16 passing automated tests covering authentication and product endpoints
5. Wrote a complete README documenting setup, environment variables, every API endpoint, and how to run tests

## Core concepts explained

### Third party API integration, the general shape

Almost every third party integration follows the same basic shape, worth recognizing as a pattern rather than something specific to email:

1. Sign up with the provider, get an API key
2. Store that key as an environment variable, never hardcoded in source, never committed to git
3. Install their SDK, or in some cases just call their REST API directly with `fetch`
4. Wrap the actual call in your own function, isolated from the rest of your business logic
5. Decide how failures in that integration should, or should not, affect the rest of your application

We did all five steps today with Resend.

### Why the email logic lives in its own service file

```javascript
// src/services/emailService.js
const { Resend } = require('resend');
const resend = new Resend(process.env.RESEND_API_KEY);

const sendLowStockAlert = async (product) => {
  try {
    await resend.emails.send({ ... });
    console.log(`Low stock alert email sent for ${product.name}`);
  } catch (err) {
    console.error('Failed to send low stock alert email:', err.message);
  }
};

module.exports = { sendLowStockAlert };
```

`stockController.js` never talks to Resend directly, it just calls `sendLowStockAlert(product)`. This separation, sometimes called a service layer, means the "how do we send an email" logic is isolated from the "what counts as low stock" logic. If Resend were ever swapped for a different provider, only this one file would need to change, nothing in the controller would need to know or care.

### Why the email call doesn't get awaited, and why its errors are swallowed

```javascript
if (justCrossedThreshold) {
  sendLowStockAlert({ ...product, quantity: newQuantity });
}
res.status(201).json({ ... });
```

Two deliberate choices here.

No `await` in front of `sendLowStockAlert(...)`, sometimes called "fire and forget." The response to the client goes out immediately after the database transaction commits, without waiting for Resend's network call to finish. The stock update is already safely saved by this point, that's the important, must succeed part of the request. The email is a secondary notification layered on top, so there's no reason to make the caller wait an extra second or two for it.

Inside `sendLowStockAlert`, the `try/catch` means that even if Resend is down, or the network fails, the error is caught and logged quietly, never thrown back up to crash the request. This is a genuinely important pattern: a non critical side effect should never be able to take down a core operation that already succeeded. We proved this today for real, an actual Resend API failure occurred mid session due to the WSL DNS issue, and the stock movement still completed successfully and returned a normal response, exactly as designed.

### The WSL DNS problem, what actually happened

The symptom was `Could not resolve host: api.resend.com`, while `ping 8.8.8.8` (a raw IP address, no name lookup needed) worked fine. That gap is the exact signature of a DNS problem specifically, not a general internet problem: the network connection itself was fine, but turning a domain name into an IP address was failing.

The root cause: WSL auto generates `/etc/resolv.conf` on every startup, and it was pointing at a DNS server, `10.255.255.254`, that wasn't resolving names correctly. The fix had two parts:

```bash
# stop WSL from auto regenerating resolv.conf on every boot
sudo tee /etc/wsl.conf > /dev/null << 'EOF'
[network]
generateResolvConf = false
EOF

# replace it with Google's public DNS
sudo tee /etc/resolv.conf > /dev/null << 'EOF'
nameserver 8.8.8.8
nameserver 8.8.4.4
EOF
```

Then, critically, a full `wsl --shutdown` from Windows PowerShell, not just closing the terminal tab. Closing a terminal window only ends that one shell session, the WSL virtual machine itself keeps running in the background with the old broken network config until it's explicitly shut down. This is a fair thing to know as a general WSL fact, not just for this one bug: any change to WSL level system config (not just app config) usually needs `wsl --shutdown` to actually take effect.

### The "crossed threshold" guard

```javascript
const isLowStock = newQuantity <= product.low_stock_threshold;
const wasAboveThreshold = product.quantity > product.low_stock_threshold;
const justCrossedThreshold = isLowStock && wasAboveThreshold;

if (justCrossedThreshold) {
  sendLowStockAlert({ ...product, quantity: newQuantity });
}
```

`product.quantity` here is still the quantity from before this particular movement, since that variable was never reassigned after being fetched at the top of the function. `newQuantity` is the value after this movement is applied. So the guard asks a precise question: was this product fine a moment ago, and is it a problem right now? Only a genuine new crossing into low stock territory sends an email. A product that's already low and gets sold further down won't trigger a repeat alert, but if it's restocked back above threshold and later drops low again, that later drop correctly triggers a fresh alert, since `product.quantity` at that point will again be above threshold. This is a good pattern to describe as "notify on state change, not on every check."

### Integration testing with Jest and Supertest

Jest is the test runner and assertion library, it finds test files, runs them, and reports pass or fail. Supertest lets you fire real HTTP requests at your Express app directly in memory, without actually starting a server on a real port.

```javascript
const request = require('supertest');
const app = require('../app');

const res = await request(app)
  .post('/api/auth/login')
  .send({ username: 'jane', password: 'secret' });

expect(res.statusCode).toBe(200);
```

This is called an integration test rather than a unit test, because it exercises the whole request path, routing, middleware, controller, database query, all together, closer to how a real client actually experiences the API, rather than testing one isolated function with mocked inputs.

### Test structure: describe, it, beforeEach, afterAll

`describe` groups related tests under a readable label. `it` (or `test`) is one individual case with one specific expectation. `beforeEach` runs before every single test inside its scope, commonly used to reset data so tests never leak state into each other, for example our `DELETE FROM users` before every auth test. `afterAll` runs once, after every test in the file finishes, we used it to call `pool.end()` and properly close the database connection, otherwise Jest hangs after finishing because Node sees a connection still technically open.

### Why tests run against a separate database

Tests create, update, and delete real rows. Running them against the same database used for actual development work would mean every test run wipes or pollutes real data. The fix was a second database, `inventory_api_test`, with the identical schema, and a `.env.test` file with `DB_NAME=inventory_api_test` instead of the real name.

A small setup file tells Jest to load that file before any test runs:

```javascript
// jest.setup.js
require('dotenv').config({ path: '.env.test' });
```

wired in via `package.json`:

```json
"jest": {
  "setupFiles": ["<rootDir>/jest.setup.js"]
}
```

This works cleanly with the existing `db/index.js` because dotenv, by default, does not override environment variables that are already set. Since `jest.setup.js` runs first and sets `DB_NAME` from `.env.test`, by the time `db/index.js` calls its own `dotenv.config()` (which normally loads `.env`), the test values are already in place and dotenv leaves them alone, no code changes needed in `db/index.js` itself.

### Why `--runInBand`

```json
"test": "jest --runInBand"
```

By default Jest runs test files in parallel across multiple processes for speed. Our tests all share one real Postgres database, running in parallel risks race conditions, two test files both trying to insert the same username at the same moment, for example. `--runInBand` forces tests to run one after another instead, trading a bit of speed for correctness and predictability, worth being able to explain if asked why parallelism was deliberately disabled.

### What the test suite actually proves

16 tests passing across two files:

`auth.test.js`: registration succeeds and excludes the password hash from the response, duplicate usernames are rejected, missing passwords are rejected, login returns a token for correct credentials, login rejects wrong passwords, login rejects unknown usernames.

`products.test.js`: products list correctly, viewing products requires no authentication, creating a product requires a token, both staff and admin can create products, duplicate SKUs are rejected, missing required fields are rejected, staff accounts are correctly blocked from deleting products with a 403, admin accounts can delete successfully, deleting a nonexistent product returns 404.

That last pair, the staff 403 versus admin success on delete, is the most interview relevant pair in the whole suite, it's automated proof of the authentication versus authorization distinction actually working, not just a one off curl call from Day 2.

### Documentation: what makes a README actually useful

A README is often the first thing anyone, including an interviewer, looks at before reading any code. The structure we used follows a common, recognizable pattern: overview and features first, so a reader knows in ten seconds what the project does, then prerequisites and setup, so someone can actually get it running, then a reference section for environment variables and every endpoint, then testing instructions, then project structure.

One deliberate detail worth remembering: the setup instructions use generic placeholders like `your_postgres_username` and `your_email@example.com`, not real personal values. A README is a public facing template for anyone setting the project up fresh, not a record of one specific machine's configuration. Documentation that lists real personal secrets is treated as a genuine security problem, especially in a public GitHub repository.

## Terminal and workflow notes

The three terminal setup continued to prove useful today: one for running the dev server and editing files, one parked in psql for checking table state, one for firing curl requests and now `npm test`.

One real thing that happened worth remembering for next time: nodemon crashed with a `SyntaxError` mid session, caused by the file being saved in a half written state, likely an editor autosave triggering a restart before a multi line edit was fully typed out. Nodemon recovered automatically on the next successful save. This is a normal, harmless quirk of live reload tooling, not a bug in the actual logic, worth recognizing quickly rather than panicking over.

## Files created or changed today

- `src/services/emailService.js` — new, wraps the Resend SDK, sends low stock alert emails, catches its own errors
- `src/controllers/stockController.js` — added the low stock alert trigger with the "just crossed threshold" guard
- `.env` — added `RESEND_API_KEY`, `ALERT_EMAIL_TO`
- `/etc/wsl.conf`, `/etc/resolv.conf` — WSL level DNS fix, outside the project folder
- `.env.test` — new, test environment configuration pointing at `inventory_api_test`
- `jest.setup.js` — new, loads `.env.test` before tests run
- `package.json` — added `test` script and Jest config block
- `src/__tests__/auth.test.js` — new, 6 tests
- `src/__tests__/products.test.js` — new, 10 tests
- `README.md` — new, full project documentation
- `inventory_api_test` — new PostgreSQL database, schema loaded from `src/db/schema.sql`

## Sample interview questions and answers

**Q: Walk me through what happens when a third party service you depend on, like your email provider, goes down.**
A: The core operation it's attached to still succeeds. In my case, a stock update is wrapped in a database transaction and commits fully before the email is even attempted. The email call itself is wrapped in its own try/catch, so if it fails, the error is caught and logged, it never bubbles up and never causes the actual request to fail. The client gets a normal successful response either way. I don't even await the email call, since there's no reason to make the caller wait on a non critical side effect.

**Q: Why did you choose to only send an alert when stock first crosses the threshold, rather than every time a low stock movement happens?**
A: Repeated identical alerts for a product that's already known to be low add noise without adding new information, and in a real system that leads to alert fatigue, where people start ignoring notifications altogether. I compare the quantity before and after the movement, and only fire the alert when it moves from above threshold to at or below it, a genuine state change. If it gets restocked and later drops low again, that's treated as a new event and does trigger a fresh alert.

**Q: What's the difference between a unit test and an integration test, and which did you write?**
A: A unit test isolates one small piece of logic, often mocking out its dependencies like the database, and checks that piece alone. An integration test exercises multiple layers working together, in my case using Supertest to fire real HTTP requests at my actual Express app, going through routing, middleware, the controller, and a real test database, closer to how an actual client experiences the API. I wrote integration tests, since for an API, proving the whole request path behaves correctly end to end felt like the more valuable signal for the time I had.

**Q: How do you make sure your tests don't interfere with your real data?**
A: I set up a completely separate PostgreSQL database with the identical schema, and a separate `.env.test` file pointing at it. A Jest setup file loads that test environment before any test file runs, so the app connects to the test database instead of the real one automatically, with no risk of a test accidentally deleting real development data.

**Q: Why do your tests run with `--runInBand` instead of Jest's default parallel mode?**
A: All my tests share one real database connection pool and one real test database. Running them in parallel risks race conditions, for example two test files both trying to register the same username at the same moment. Running them sequentially trades some speed for predictable, conflict free results, which mattered more than speed for a suite this size.

**Q: What do you consider when writing a README for a project?**
A: I think about someone who has never seen the project before and needs to get it running with as little friction as possible. That means a clear description up front, exact setup steps in order, a reference for every environment variable and what it does, and documentation for every API endpoint including what it expects and what it returns. I also use generic placeholders instead of my own real credentials in any example, since a README is a public template, not a personal configuration record.

## What's next: Day 4

Mock interview, explaining the code and every responsibility mapped to it out loud.
