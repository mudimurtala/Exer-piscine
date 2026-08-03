# Day 2 Notes: JWT Authentication for the Inventory API

## What we built today

Starting point: every endpoint in the API was open to anyone. By the end of today, creating, updating, and deleting products and categories requires a valid login, and deleting a product specifically requires the admin role. This directly covers the "secure backend systems" and "support database design" parts of the Grazac job description, since access control is a core piece of backend security.

Concretely, we added:

1. A `users` table storing hashed passwords
2. A registration endpoint (`POST /api/auth/register`)
3. A login endpoint (`POST /api/auth/login`) that returns a JWT
4. A `/api/auth/me` endpoint that returns the logged in user's details
5. Middleware that checks for a valid token on protected routes
6. Middleware that checks for a specific role on the delete product route

## Core concepts explained

### Password hashing with bcrypt

You never store a user's real password in the database. If the database were ever leaked or stolen, plain text passwords would immediately compromise every user's account, likely on other sites too, since people reuse passwords.

Instead, `bcrypt.hash(password, 10)` runs the password through a one way scrambling algorithm. One way means there is no function to reverse it and get the original password back. The output looks like `$2b$10$ddhOrl9FlJ7Ph3Wbl1TuBO8GAcI8kHNgz/VqJCmrE7pHMWR4RYDWe`.

Breaking that string down:
`2b` is the bcrypt algorithm version. `10` is the cost factor, meaning how many rounds of internal hashing happen, higher costs are slower to compute but harder to brute force. The rest is a random salt followed by the actual hash.

The salt matters because it means two users with the identical password `password123` will end up with completely different hashes in the database. Without a salt, an attacker could precompute hashes for common passwords once and match them against every row in every leaked database ever. The salt makes that precomputation useless.

To check a login attempt, you never unhash anything. You take the password the user just typed, hash it the same way, and use `bcrypt.compare(plainPassword, storedHash)`, which handles pulling the salt back out of the stored hash and redoing the same computation, then checks if the results match.

### JSON Web Tokens (JWT)

A JWT is a compact, self contained way for your server to say "this request came from a logged in user" without having to look anything up in a database on every single request.

A JWT is one string, but really three parts joined by dots, each base64 encoded:

`header.payload.signature`

**Header**: a small object saying which algorithm was used to sign the token, for example HMAC SHA256.

**Payload**: the actual data you chose to embed. In our case, that is `userId`, `username`, `role`, `iat` (issued at time), and `exp` (expiry time). Anyone can decode this part and read it. Base64 is an encoding, not encryption, so never put secrets like a raw password inside a JWT payload.

**Signature**: this is the important part. It is created by taking the header and payload, and running them through a signing algorithm together with a secret key that only your server knows, `process.env.JWT_SECRET` in our case. Anyone can read the payload, but nobody can produce a valid signature without knowing the secret. If someone tampers with the payload, even changing one character, the signature will no longer match when the server checks it, and `jwt.verify` will reject it.

This is why the JWT_SECRET must never be committed to GitHub or shared publicly. We generated ours with:

```
node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
```

and stored it in `.env`, which is already listed in `.gitignore`.

### Signing a token

```javascript
const token = jwt.sign(
  { userId: user.id, username: user.username, role: user.role },
  process.env.JWT_SECRET,
  { expiresIn: '1h' }
);
```

This creates the token after a successful login. `expiresIn: '1h'` means the token stops being valid one hour after issue, regardless of whether the user is still active. Short expiry times limit the damage if a token is ever stolen, though real applications often pair this with a longer lived refresh token to avoid making users log in every hour. We saw this expiry happen for real today when an old token came back as expired after enough time had passed between test calls.

### Verifying a token

```javascript
const decoded = jwt.verify(token, process.env.JWT_SECRET);
```

This checks two things in one call: that the signature is valid, meaning the token really was issued by your server and has not been altered, and that the token has not expired. If either check fails, it throws an error, which we catch and turn into a 401 response.

### Middleware in Express

Middleware is a function that runs in between a request arriving and the final route handler responding to it. It receives `req`, `res`, and a third argument, `next`, which is a function you call to pass control forward to whatever comes next in the chain.

```javascript
const authenticate = (req, res, next) => {
  const authHeader = req.headers.authorization;
  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return res.status(401).json({ error: 'No token provided' });
  }
  const token = authHeader.split(' ')[1];
  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    req.user = decoded;
    next();
  } catch (err) {
    return res.status(401).json({ error: 'Invalid or expired token' });
  }
};
```

If the check fails, we send a response directly and never call `next()`, so the request stops there and the real route handler never runs. If the check passes, we attach the decoded payload to `req.user`, a spot on the request object, so every function downstream can see who made the request, then call `next()` to let the request continue.

The header convention `Authorization: Bearer <token>` is a widely used standard. "Bearer" means whoever holds, or bears, this token is treated as authenticated, no other proof needed. That is also why tokens must be protected in transit, typically by always using HTTPS in production.

### Applying middleware selectively

```javascript
router.get('/', getProducts);
router.post('/', authenticate, createProduct);
router.delete('/:id', authenticate, authorizeRoles('admin'), deleteProduct);
```

Express lets you chain as many middleware functions as you want before the final handler, and they run left to right. GET requests here have no middleware, so anyone can view products. POST requires `authenticate` to pass first. DELETE requires both `authenticate` and `authorizeRoles('admin')` to pass, in that order, before `deleteProduct` ever runs.

### Authentication vs authorization

These sound similar but answer different questions, and mixing them up is a common junior developer mistake to avoid saying out loud in an interview.

Authentication asks: who are you? This is what `authenticate` does, verifying the token proves the request comes from a real, logged in user. Failure here returns 401 Unauthorized, meaning "we don't know who you are."

Authorization asks: are you allowed to do this? This is what `authorizeRoles` does, checking that the known, authenticated user actually has permission for this specific action. Failure here returns 403 Forbidden, meaning "we know exactly who you are, but you are not allowed to do this."

We proved both today: a request with no token at all got 401, and a request with a perfectly valid token belonging to a staff account trying to delete a product got 403, since only admin is in the allowed roles list for that route.

### The middleware factory pattern

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

`authorizeRoles` is not itself middleware. It is a function that takes role names and returns a middleware function customized to check for those roles. This is why it is called with parentheses and arguments directly in the route definition, `authorizeRoles('admin')`, rather than being passed by name like `authenticate` is. This pattern lets the same piece of logic be reused for different rules elsewhere, for example `authorizeRoles('admin', 'staff')` on a route that should allow either role.

### Why we re-query the database in /me instead of trusting the token directly

```javascript
const getMe = async (req, res) => {
  const result = await db.query(
    'SELECT id, username, role, created_at FROM users WHERE id = $1',
    [req.user.userId]
  );
  ...
};
```

The token's payload already contains `userId`, `username`, and `role`, decoded and available on `req.user`. We could have just returned that directly with no database call at all. We chose to query fresh instead, because a token stays valid for up to an hour, and if an admin's role changed to something else during that window, the token would still claim the old role until it expired. Querying the database on `/me` shows the true, current state rather than a possibly stale snapshot from login time.

## Terminal workflow notes

Running three terminals side by side for this kind of work is a genuinely good habit:

- Terminal 1: running `npm run dev` and editing files
- Terminal 2: a standing `psql -U ibnmuhyideen -d inventory_api` session, left open, for checking table contents any time with `SELECT * FROM users;` or `\dt`
- Terminal 3: curl commands for testing endpoints

One practical curl lesson from today: pasting a long JWT directly into a curl command risks the terminal wrapping or truncating the string, which produces a confusing "Invalid or expired token" error that looks like a code bug but is actually just an incomplete token. The fix is capturing the token into a shell variable straight from the login response instead of copy pasting it:

```bash
TOKEN=$(curl -s -X POST http://localhost:5000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "teststaff", "password": "staffpass123"}' \
  | node -e "let d='';process.stdin.on('data',c=>d+=c);process.stdin.on('end',()=>console.log(JSON.parse(d).token))")
```

Then every following curl call just references `$TOKEN`, which is shorter, safer, and always accurate.

## Files created or changed today

- `src/db/schema.sql` — added the `users` table
- `src/controllers/authController.js` — new, `register`, `login`, `getMe`
- `src/routes/authRoutes.js` — new, mounts the three auth endpoints
- `src/middleware/authMiddleware.js` — new, `authenticate` and `authorizeRoles`
- `src/app.js` — mounted `authRoutes` at `/api/auth`
- `src/routes/productRoutes.js` — added `authenticate` to POST/PUT, added `authenticate` plus `authorizeRoles('admin')` to DELETE
- `src/routes/categoryRoutes.js` — added `authenticate` to POST/PUT/DELETE
- `.env` — added `JWT_SECRET`

## Sample interview questions and answers

**Q: Why do you hash passwords instead of encrypting them?**
A: Encryption is reversible by design, if you can encrypt something you can also decrypt it back to the original, which means anyone holding the key, including an attacker who steals it, could recover every real password. Hashing with bcrypt is one way, there is no operation that turns the hash back into the original password. To check a login, you hash the attempt and compare hashes, you never need to recover the original.

**Q: What's inside a JWT, and is it safe to put a password inside one?**
A: A JWT has three parts, a header, a payload, and a signature, joined by dots. The header and payload are base64 encoded, which is easily reversible by anyone, so it is not encryption. The signature is what protects the token from being tampered with, produced using a secret key only the server knows. You should never put sensitive data like a password inside a JWT payload, since anyone with the token can decode and read it. What you can safely put there is non sensitive identifying information the server needs quickly, like a user id and role.

**Q: What happens if a JWT is stolen?**
A: Whoever holds it can act as that user until it expires, since JWTs use a "bearer" model, meaning possession alone is proof enough. This is why we set a fairly short expiry, one hour here, and why in production this should always run over HTTPS so tokens cannot be intercepted in transit. Larger systems often pair a short lived access token with a longer lived refresh token stored more securely, so a stolen access token has a small time window of usefulness.

**Q: What is the difference between authentication and authorization?**
A: Authentication confirms who is making a request, that is what verifying the JWT signature does. Authorization confirms whether that specific, known user is allowed to perform this specific action, that is what checking their role does. A request can pass authentication and still fail authorization, for example a logged in staff account trying to delete a product when only admins are allowed to.

**Q: Why is your authorization middleware written as a function that returns a function?**
A: It's a pattern for making middleware configurable. `authorizeRoles('admin')` needs to know which role or roles to check for, but Express expects to be handed a middleware function with the standard `(req, res, next)` signature, not a function call with extra arguments. So `authorizeRoles` takes the allowed roles as input and returns a new function matching that exact signature, closing over those roles so the returned function remembers them when it eventually runs.

**Q: Why does GET stay open while POST, PUT, and DELETE require a token?**
A: It reflects the actual risk of each action. Viewing data is generally safe to expose, especially for something like a product catalog that might reasonably be public facing. Creating, changing, or deleting data can cause real damage or be abused, so those actions need proof of who is making the request. Applying the same blanket rule to every route regardless of risk would either lock out legitimate anonymous browsing or leave dangerous actions unprotected.

## What's next: Day 3

Third party API integration for low stock email alerts, automated tests with Jest and Supertest, and technical documentation.
