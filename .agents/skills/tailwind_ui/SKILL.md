---
name: Tailwind UI
description: A collection of modern, beautiful, and dynamic Tailwind CSS UI components and patterns for ShakhaNow.
---

# 🚀 Tailwind UI Skill: Premium Aesthetics for ShakhaNow

This skill provides a standardized approach to building a "WOW" UI for ShakhaNow using Tailwind CSS. It focuses on premium design principles, glassmorphism, and smooth animations.

## 🎨 Design Principles

1.  **Dark Mode First**: The theme is designed as a deep dark mode (#0f172a / #020617) with vibrant accent colors.
2.  **Glassmorphism**: Use `backdrop-blur-xl`, `bg-slate-900/60`, and subtle borders (`border-slate-700/50`) to create a layered, modern look.
3.  **Vibrant Gradients**: Use Indigo, Violet, and Rose for primary actions and accents.
4.  **Micro-animations**: Use subtle transitions (`transition-all duration-300`) and the custom `.animate-float` utility for a living interface.
5.  **Modern Typography**: Use `Outfit` for headings and `Inter` for body text.

## 🧱 Core Components

### 🪟 Glass Card
Use this for containers, panels, and sections.
```html
<div class="glass-card p-6 rounded-2xl">
  <!-- Content -->
</div>
```

### 🔘 Premium Button
A high-impact button with a gradient background and hover scale effect.
```html
<button class="primary-btn">
  Get Started
</button>
```

### 🏷️ Status Badges (Role Colors)
Use these for different user roles in ShakhaNow.
-   **Mukhya Shishak**: Blue (`bg-blue-500/10 text-blue-400 border-blue-500/20`)
-   **Karyavah**: Amber/Orange (`bg-amber-500/10 text-amber-400 border-amber-500/20`)
-   **Gatnayak**: Emerald (`bg-emerald-500/10 text-emerald-400 border-emerald-500/20`)
-   **Swayamsevak**: Indigo (`bg-indigo-500/10 text-indigo-400 border-indigo-500/20`)

Example:
```html
<span class="px-2.5 py-0.5 rounded-full text-xs font-medium border bg-blue-500/10 text-blue-400 border-blue-500/20">
  Mukhya Shishak
</span>
```

## 📐 Layout Patterns

-   **Dashboard Grid**: Use `grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6`.
-   **Header**: Sticky header with `backdrop-blur-md` and `bg-slate-950/80`.

## ✨ Animations

-   **Floating Action**: Add `.animate-float` for subtle, organic movement.
-   **Hover Lift**: Use `hover:-translate-y-1 hover:shadow-2xl transition-all duration-300`.

## 🛠️ Usage

When tasked with building or updating a page, ALWAYS ensure:
1.  All elements follow the glassmorphism style where appropriate.
2.  Typography follows the `font-heading` and `font-sans` classes.
3.  Inputs use the `.glass-input` component from `input.css`.
