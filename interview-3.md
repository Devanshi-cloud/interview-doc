# 🚀 Backend & JavaScript Interview Preparation

This document covers important concepts asked in backend interviews, including **Go (Golang), APIs, concurrency, and JavaScript fundamentals**.

---

## 🌐 What is an API?

An **API (Application Programming Interface)** allows different software systems to communicate with each other.

💡 Example:
A frontend application fetches data from a backend server using an API.

---

## 🔗 Types of APIs

### 1. REST API

* Uses HTTP protocol
* Stateless
* Most widely used

### 2. GraphQL

* Client requests only required data
* Prevents over-fetching

### 3. SOAP

* XML-based protocol
* Strict and secure
* Used in enterprise systems

### 4. gRPC

* High performance
* Uses Protocol Buffers
* Common in microservices

---

## 🔁 HTTP Methods (CRUD)

| Method | Description     |
| ------ | --------------- |
| GET    | Retrieve data   |
| POST   | Create new data |
| PUT    | Update data     |
| DELETE | Remove data     |

---

## ⚙️ Concurrency in Go

Go supports concurrency using **goroutines** and **channels**.

### 🔹 Goroutines

Lightweight threads managed by Go runtime.

```go
go func() {
    fmt.Println("Running concurrently")
}()
```

---

## 🔐 Mutex (Mutual Exclusion)

Used to prevent **race conditions** when multiple goroutines access shared data.

```go
var mu sync.Mutex

mu.Lock()
// critical section
mu.Unlock()
```

---

## ⚠️ Deadlock

### What is Deadlock?

A situation where all goroutines are waiting indefinitely and no progress is made.

### Example:

```go
ch := make(chan int)
ch <- 1 // deadlock (no receiver)
```

### Error:

```
fatal error: all goroutines are asleep - deadlock!
```

### Causes:

* Sending without receiver
* Waiting indefinitely
* Improper locking

---

## 🧵 Goroutines vs Threads

| Feature    | Goroutines  | Threads |
| ---------- | ----------- | ------- |
| Weight     | Lightweight | Heavy   |
| Managed by | Go runtime  | OS      |
| Cost       | Low         | High    |

---

## ⚡ JavaScript Interview Questions

### 1. Closure Issue

```javascript
for (var i = 0; i < 3; i++) {
  setTimeout(() => console.log(i), 1000);
}
```

**Output:**

```
3 3 3
```

**Fix:**

```javascript
for (let i = 0; i < 3; i++) {
  setTimeout(() => console.log(i), 1000);
}
```

---

### 2. Equality Check

```javascript
console.log(0 == false);  // true
console.log(0 === false); // false
```

---

### 3. Arrow Function Return

```javascript
const add = (a, b) => {
  a + b;
};
```

**Output:** `undefined`

**Fix:**

```javascript
const add = (a, b) => a + b;
```

---

### 4. `this` Keyword Issue

```javascript
const obj = {
  name: "Devanshi",
  greet: () => console.log(this.name),
};
```

**Output:** `undefined`

**Fix:**

```javascript
greet() {
  console.log(this.name);
}
```

---

### 5. Async/Await Mistake

```javascript
async function getData() {
  const data = fetch("url");
  console.log(data);
}
```

**Output:** `Promise {<pending>}`

**Fix:**

```javascript
const data = await fetch("url");
```

---

## 🎯 Key Takeaways

* APIs enable communication between systems
* REST APIs are most commonly used
* Goroutines enable concurrency in Go
* Mutex prevents race conditions
* Deadlocks occur when processes wait indefinitely
* JavaScript has quirks in closures, `this`, and async behavior

---

## 📌 Interview Tips

* Always explain **why** something is wrong
* Discuss **edge cases**
* Mention **performance & concurrency**
* Use **real-world examples**

---

## 🚀 Next Steps

* Build a REST API using Go
* Practice concurrency problems
* Learn JavaScript async deeply
* Explore system design basics

---

⭐ If you found this helpful, consider starring the repository!
