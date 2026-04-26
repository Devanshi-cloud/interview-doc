I don’t “learn JSX”… I *use* it like it’s second nature.

I already know that JSX is nothing magical — it’s just syntactic sugar over JavaScript. When I write something that *looks* like HTML, I’m fully aware that under the hood React is converting it into `React.createElement`. So when I write:

```jsx
const element = <h1>Hello</h1>;
```

I know React is actually doing:

```js
React.createElement("h1", null, "Hello");
```

So in my head, JSX = **JavaScript expressing UI structure**, not HTML.

---

When I build a React app (say using Vite), I don’t get distracted by setup — I know the flow:

* create project
* install deps
* run `npm run dev`
* jump straight into `App.jsx`

Everything important happens inside that component.

---

### My mental model of JSX

I think in **rules + intent**, not syntax.

#### 1. Single Parent — always controlled structure

I know React needs one root. So I never return scattered elements.

```jsx
return (
  <div>
    <h1>Hello</h1>
    <p>What are you doing?</p>
  </div>
);
```

Or cleaner, I use a fragment because I don’t add useless DOM:

```jsx
return (
  <>
    <h1>Hello</h1>
    <p>What are you doing?</p>
  </>
);
```

I’m not “fixing errors” — I’m structuring UI intentionally.

---

#### 2. JavaScript inside JSX — I inject logic, not DOM manipulation

I don’t touch `document.querySelector` anymore. I *embed logic directly*:

```jsx
const name = "Devanshi";

return <h1>Hello {name}</h1>;
```

Curly braces = “I switch to JavaScript mode”.

Even calculations:

```jsx
const a = 25;
const b = 15;

return <h1>{a + b}</h1>;
```

I don’t “update UI manually” — UI is a function of data. I control data.

---

#### 3. Conditions — I don’t write `if` inside JSX, I think declaratively

I know JSX doesn’t allow statements, only expressions. So I use:

**Short-circuit:**

```jsx
{isLoggedIn && <h1>Welcome</h1>}
```

**Ternary:**

```jsx
{isLoggedIn ? <h1>Welcome</h1> : <h1>Please login</h1>}
```

I’m not forcing logic — I’m *describing UI states*.

---

#### 4. Attributes — I respect JavaScript rules

I already know `class` is reserved, so I use:

```jsx
<div className="container"></div>
```

Not a rule to memorize — just consistency with JS.

---

#### 5. Self-closing tags — I stay strict

I don’t write sloppy HTML. JSX is stricter:

```jsx
<img src="img.png" />
<br />
```

I close everything properly because JSX compiles to JS, not raw HTML.

---

#### 6. Lists — I think in data → UI mapping

I don’t “repeat HTML manually”. I map data:

```jsx
const users = ["A", "B", "C"];

return (
  <ul>
    {users.map((user, index) => (
      <li key={index}>{user}</li>
    ))}
  </ul>
);
```

I understand that UI is just a projection of data. If data grows, UI grows automatically.

---

### My overall thinking

I don’t see JSX as a “new syntax”.

I see it as:

> “A declarative way to describe UI using JavaScript.”

* I control **data**
* JSX reflects **state**
* React handles **rendering**

I don’t manipulate the DOM — I describe what it *should look like*, and React figures out the rest.

---

### My conclusion (my mindset)

I don’t memorize JSX rules.

I naturally follow them because I understand:

* One root → structured UI
* Curly braces → inject logic
* Ternary / && → conditional rendering
* map() → dynamic UI
* className → JS consistency

So when I write JSX, I’m not “coding”…

I’m just **translating my thinking into UI**.
