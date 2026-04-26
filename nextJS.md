# Next.js Core Concepts — My Understanding

I don’t treat this as notes.  
This is my **working knowledge** of how I use Next.js in real systems.

---

## 🧠 My Approach

I don’t “try things” in Next.js—I already understand the architecture and use it deliberately.

- I let the framework handle performance
- I focus on structure, not hacks
- I design systems, not pages

---

# 📦 Project Setup (My Way)

I don’t blindly accept defaults. I choose based on intent.

## ⚙️ Setup Decisions I Make

- I use **Next.js App Router**
- I enable **ESLint** (code quality matters)
- I use **Tailwind CSS** (fast UI building)
- I skip TypeScript when speed matters (optional)

---

# 🧱 Layout System (I Control Structure)

## 🧠 What I Know

I don’t use routing tricks like React Router’s `Outlet`.

I use **layout.js** because I understand how Next composes UI.

## ⚙️ How It Works

- `layout.js` defines:
  - HTML structure
  - Metadata
  - Global styles
- `children` represents dynamic page content

## 🔁 Pattern I Follow

- Header → always constant  
- Footer → always constant  
- Page content → dynamic  

## 🧩 My Mental Model

Rendering flow:
1. Layout loads
2. HTML structure builds
3. Styles & fonts apply
4. Page content injects via `children`

---

# 🖼️ Image Optimization (I Don’t Use <img>)

## 🧠 What I Know

I don’t use HTML `<img>` in Next.js.

I use the **Image component** because I care about performance and optimization.

## ⚙️ Why I Use It

Next.js automatically:
- Lazy loads images (only when needed)
- Resizes images per device
- Prevents layout shift
- Improves load performance

---

## 🧩 How I Use It

```js
import Image from "next/image";

<Image 
  src="/flower.jpg"
  alt="flower"
  width={500}
  height={500}
/>
````

---

## ⚠️ Rules I Follow

* I always store images in `/public`
* I always define `width` and `height`
* I understand these define **ratio**, not fixed pixels
* I never rely on default `<img>` behavior

---

## 📁 Public Folder Strategy

I don’t complicate paths.

* `/public` is directly accessible
* I reference images like:

```
/flower.jpg
```

No deep relative paths. Clean and predictable.

---

## 🌐 Remote Images (I Configure Explicitly)

## 🧠 What I Know

Next.js does NOT allow random external images.

I must register domains for optimization.

## ⚙️ Configuration

```js
// next.config.js
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "example.com",
      },
    ],
  },
};

export default nextConfig;
```

---

## ⚠️ Rules I Follow

* I always use `https`
* I always specify the correct hostname
* For multiple sources → I add multiple objects

---

# 🎯 My Mental Model of Images

* Images are NOT static assets—they are performance-sensitive resources
* I let Next.js optimize instead of doing it manually
* I think in **on-demand loading**, not preloading everything

---

# 🚀 Development Flow (My Way)

## 🧠 What I Do

* I scaffold fast
* I remove boilerplate immediately
* I build only what I need

## ⚙️ Commands I Use

```bash
npm run dev
```

---

## 🧩 What I Expect

* Fast startup
* Optimized rendering
* Clean structure

If something breaks → I debug based on understanding, not trial and error.

---

# ⚠️ Mistakes I Avoid

* ❌ Using `<img>` instead of Next Image
* ❌ Storing images outside `/public`
* ❌ Forgetting width/height
* ❌ Using external images without config
* ❌ Treating layout like React Router

---

# 🧠 Final Thought

I don’t “use” Next.js.

I understand:

* how it renders
* how it optimizes
* how it structures applications

So when I build, I’m not experimenting.

I’m executing.
