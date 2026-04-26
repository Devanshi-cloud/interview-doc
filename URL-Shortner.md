I already know what I’m doing here. I’m not just building a random project — I’m **intentionally designing a scalable system**.

So what did I do?

I started with a simple **URL Shortener**, but in my mind, it was never “simple.” I treated it like a **real production system** from day one.

---

### 🔹 Foundation (I already nailed this)

I chose my stack consciously:

* **Backend:** Node.js + Express
* **Database:** MongoDB
* **Frontend:** React
* **ID generation:** nanoid

I didn’t just code blindly — I understood *why* each piece exists.

---

### 🔹 Then I moved into SYSTEM DESIGN thinking

I didn’t stop at “it works.”

I asked myself:

> “How would this behave with real users?”

That’s where I started upgrading it step by step.

---

### ✅ Step 1: Duplicate URL Handling

I already knew that storing the same URL multiple times is wasteful.

So I enforced:

* If URL already exists → return existing short URL
* Not create duplicates

That’s **efficiency + consistency**.

---

### ✅ Step 2: URL Validation

I don’t trust user input blindly.

So I validated:

* Proper format (http/https)
* No invalid or broken URLs

That’s **input sanitization** — basic but critical.

---

### ✅ Step 3: Click Counter

I wanted analytics, not just functionality.

So I added:

* `clicks` field in DB
* Increment on every redirect

Now I can answer:

> “How many times was this link used?”

That’s **data-driven thinking**.

---

### ✅ Step 4: Expiry System

Now I thought like a product builder.

I implemented:

* Expiry date (e.g., 1 day, 7 days)
* Auto-invalid links after expiry

This enables:

* Temporary links
* Better control over storage

That’s **lifecycle management**.

---

### ✅ Step 5: Custom Alias

Now I moved into **user experience**.

Instead of random IDs, I allowed:

* Custom names like `/rahul`

But I also handled:

* Collision (if alias already taken → reject)

That’s:

> usability + constraint handling

---

### 🔥 Step 6: Redis Caching (This is where I leveled up)

Now I stopped thinking like a coder.

I started thinking like a **system designer**.

I knew:

> “Database calls are expensive.”

So I implemented **Cache-Aside Pattern**.

---

### 💡 My flow (I know this clearly):

1. User hits short URL
2. I check Redis first

   * If found → return instantly (**cache hit**)
   * If not → fetch from MongoDB (**cache miss**)
3. Save result in Redis
4. Return response

---

This is exactly the pattern:

> Cache → DB → Cache → Response

---

Now I understand performance:

* Redis = in-memory → super fast
* MongoDB = persistent → slower

So I reduced DB load.

---

### 🔥 What I proved to myself

I didn’t just “learn Redis”

I proved:

* I can integrate it
* I understand cache hits vs misses
* I can optimize real systems

---

### 🔹 What I’m doing next (I already see the roadmap)

I’m not stopping here.

I’m going deeper:

#### ➤ Rate Limiting

I will control:

* How many requests per user

To prevent:

* abuse
* spam

---

#### ➤ Load Balancer

I’ll simulate:

* Multiple servers
* Traffic distribution

So the system doesn’t crash under load.

---

#### ➤ Enterprise-Level Thinking

I’ll design:

* scalable architecture
* bottleneck handling
* distributed systems

---

### 🧠 My mindset (this is the real difference)

I’m not coding for syntax.

I’m thinking:

> “If this had 1 million users, would it survive?”

That’s why:

* I added validation
* I prevented duplication
* I added expiry
* I optimized using Redis

---

### 🔥 Final truth (I know this)

Anyone can build a URL shortener.

But I didn’t build “a project.”

I built:

> a **mini scalable system with real-world design decisions**
