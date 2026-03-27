# 🚀 Frontend Interview Preparation Guide (JS + React + DSA)

This document summarizes a **complete frontend interview round** including JavaScript fundamentals, coding questions, React concepts, and Redux architecture — along with key learnings and preparation strategy.

---

# 📌 Interview Structure

The interview typically follows this flow:

1. Introduction
2. JavaScript Fundamentals
3. Output-Based Questions
4. Coding Round (DSA)
5. React Concepts
6. Redux Basics

---

# ⚙️ JavaScript Theory

### 🔹 null vs undefined

* `undefined`: Variable declared but not assigned
* `null`: Intentionally empty value
* Note: `typeof null === "object"` (known JS quirk)

---

### 🔹 Shallow Copy vs Deep Copy

* **Shallow Copy**: Copies reference for nested objects
* **Deep Copy**: Fully independent clone

```js
// Shallow copy
const obj2 = { ...obj }

// Deep copy
const obj2 = JSON.parse(JSON.stringify(obj))
```

---

### 🔹 Currying

A function that takes arguments one at a time.

```js
function add(a) {
  return function(b) {
    return a + b
  }
}
```

---

### 🔹 Prototypal Inheritance

Objects inherit properties from other objects via prototype chain.

---

### 🔹 Higher-Order Functions

Functions that:

* Accept another function OR
* Return a function

Examples: `map`, `filter`, `reduce`

---

### 🔹 Memoization

Caching function results to optimize performance.

---

# 🧪 Output-Based Questions

| Expression         | Output        |
| ------------------ | ------------- |
| `typeof undefined` | `"undefined"` |
| `true + false`     | `1`           |
| `"5" + 1`          | `"51"`        |
| `"1" + 2 + 3`      | `"123"`       |
| `0 == false`       | `true`        |
| `0 === false`      | `false`       |
| `Math.max()`       | `-Infinity`   |
| `!!null`           | `false`       |

---

### 🔹 Temporal Dead Zone (TDZ)

```js
console.log(a) // ❌ Error
let a = 5
```

---

### 🔹 Object Reference Behavior

```js
const user = { name: "A" }
const user2 = user
user2.name = "B"

console.log(user.name) // "B"
```

---

# 💻 Coding Questions

---

### 🔹 1. Second Largest Element (O(n))

```js
function secondLargest(arr) {
  let max = -Infinity, second = -Infinity

  for (let num of arr) {
    if (num > max) {
      second = max
      max = num
    } else if (num > second && num !== max) {
      second = num
    }
  }
  return second
}
```

---

### 🔹 2. Count Character Frequency

```js
function countChars(str) {
  const map = {}
  for (let ch of str) {
    map[ch] = (map[ch] || 0) + 1
  }
  return map
}
```

---

### 🔹 3. Remove Duplicates (Sorted Array - Two Pointer)

```js
function removeDuplicates(arr) {
  let j = 0
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] !== arr[j]) {
      j++
      arr[j] = arr[i]
    }
  }
  return arr.slice(0, j + 1)
}
```

---

### 🔹 4. Rotate Array by K Steps

```js
function rotate(arr, k) {
  k = k % arr.length

  reverse(arr, 0, arr.length - 1)
  reverse(arr, 0, k - 1)
  reverse(arr, k, arr.length - 1)

  return arr
}

function reverse(arr, l, r) {
  while (l < r) {
    [arr[l], arr[r]] = [arr[r], arr[l]]
    l++
    r--
  }
}
```

---

### 🔹 5. Check if Array is Sorted

```js
function isSorted(arr) {
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] < arr[i - 1]) return false
  }
  return true
}
```

---

### 🔹 6. Merge and Sort Arrays

```js
function mergeAndSort(a, b) {
  return [...a, ...b].sort((x, y) => x - y)
}
```

---

# ⚛️ React Concepts

---

### 🔹 React.memo

* Prevents unnecessary re-renders
* Re-renders only when props change

---

### 🔹 useMemo vs useCallback

| Hook          | Purpose                  |
| ------------- | ------------------------ |
| `useMemo`     | Memoizes computed values |
| `useCallback` | Memoizes functions       |

---

### 🔹 Server Components

* Introduced in React 18
* Run on server → improved performance

---

### 🔹 Middleware in React

* Not native
* Used via Redux middleware

---

### 🔹 Code Splitting

```js
const Component = React.lazy(() => import('./Component'))
```

---

### 🔹 State Management

* Context API → small applications
* Redux → large-scale applications

---

# 🔄 Redux Architecture

---

### 🔹 Core Components

* **Store** → Holds global state
* **Reducer** → Updates state
* **Dispatch** → Sends actions

---

### 🔹 API Calls in Redux

Handled using middleware:

* `redux-thunk`
* `redux-saga`

---

# 📊 Interview Feedback

### ✅ Strengths

* Strong JavaScript fundamentals
* Good conceptual clarity

### ❌ Weaknesses

* Mistakes during coding
* Weak DSA problem-solving
* React concepts not deep enough

---

# 🧠 Key Takeaways

This is a **typical fresher frontend interview pattern**:

* JS fundamentals are heavily tested
* Output-based questions check deep understanding
* Coding focuses on arrays, strings, hashing
* React basics are expected

---

# 🎯 Preparation Strategy

---

## 1. JavaScript (Core)

* Types & coercion
* Closures
* Prototypes
* Functions

---

## 2. Data Structures

Practice daily:

* Arrays
* Strings
* HashMaps

---

## 3. React (Must Know)

* Hooks (`useMemo`, `useCallback`)
* Performance optimization
* State management

---

# 🚀 Final Advice

> Knowing theory is not enough — execution matters.

Focus on:

* Writing clean code under time pressure
* Explaining your approach clearly
* Handling edge cases

---

# ⭐ Bonus Tip

Practice in this order:

1. Concepts
2. Output questions
3. Coding problems
4. Mock interviews

---

**All the best for your interviews 💪**
