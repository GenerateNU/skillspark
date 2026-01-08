# SkillSpark Frontend Development Guide

This document outlines the standards, conventions, and workflows for frontend development for SkillSpark. _Please read this guide in its entirety before contributing to the Frontend_.

As always if you have any questions please reach out to your TLs <3

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [GettingLogin Started](#getting-started)
3. [Project Structure](#project-structure)
4. [Adding a New Page](#adding-a-new-page)
5. [Component Development](#component-development)
6. [Hooks and Business Logic](#hooks-and-business-logic)
7. [Styling Guidelines](#styling-guidelines)
8. [Best Practices](#best-practices)

---

## Prerequisites

### Install Bun

This project uses [Bun](https://bun.sh/) as the JavaScript runtime and package manager. Install Bun before proceeding:

```bash
curl -fsSL https://bun.sh/install | bash
```

---

## Getting Started

### Installing Dependencies

From the root of the project run

```bash
cd frontend/web/
bun install
```

### Running the Development Server

To start the Vite development server with hot reload:

```bash
bun run dev
```

The application will be available at `http://localhost:5173`. Changes to your code will automatically refresh in the browser.

---

## Project Structure

The frontend source code resides in `frontend/web/src/`. The following general structure should be maintained:

```
frontend/web/src/
├── main.tsx                    # Application entry point and routing
├── assets/                     # Stores the images/static files for the website
├── index.css                   # Global styles
├── App.tsx                     # Root component
├── components/                 # Shared component library
│   ├── button/
│   │   ├── Button.tsx
│   │   └── useButton.ts
│   ├── input/
│   │   ├── Input.tsx
│   │   └── useInput.ts
│   └── ...
├── login/                      # Login page
│   ├── login.tsx               # Page entry point
│   ├── components/             # Page-specific components
│   │   └── LoginForm.tsx
│   └── hooks/                  # Page-specific hooks
│       └── useLoginForm.ts
├── dashboard/                  # Example page
│   ├── dashboard.tsx           # Page entry point
│   ├── components/
│   └── hooks/
└── ...
```

### Key Principles

- Each page has a dedicated folder containing a single entry point file (e.g., `login.tsx`).
- Page-specific components reside in a `components/` subfolder.
- Page-specific hooks reside in a `hooks/` subfolder.
- Shared components belong in the top level `components/` directory of `web/` and should be reusable across the application.

---

## Adding a New Page

Follow these steps to add a new page to the application.

### Step 1: Create the Page Folder Structure

Create a new folder for your page within `frontend/web/src/`:

```
frontend/web/src/
└── my-new-page/
    ├── my-new-page.tsx         # Page entry point
    ├── components/             # Page-specific components
    └── hooks/                  # Page-specific hooks (if applicable)
```

### Step 2: Create the Page Entry Point

Create the main page file (e.g., `my-new-page.tsx`):

```tsx
import { useMyNewPage } from "./hooks/useMyNewPage";

function MyNewPage() {
  const { data, isLoading } = useMyNewPage();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold">My New Page</h1>
      {/* Page content */}
    </div>
  );
}

export default MyNewPage;
```

### Step 3: Register the Route in main.tsx

Add the new route to `frontend/web/src/main.tsx`:

```tsx
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./login/login.tsx";
import MyNewPage from "./my-new-page/my-new-page.tsx"; // Import the new page

createRoot(document.getElementById("root")!).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/login" element={<Login />} />
      <Route path="/my-new-page" element={<MyNewPage />} />
      {/* Add the route */}
    </Routes>
  </BrowserRouter>
);
```

### Step 4: Add Page-Specific Components and Hooks

Create components and hooks as needed within the respective subfolders:

```tsx
import { useState, useEffect } from "react";

export function useMyNewPage() {
  const [data, setData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Fetch data or perform initialization logic
    setIsLoading(false);
  }, []);

  return { data, isLoading };
}
```

---

## Component Development

### Shared Components

Components that are used across multiple pages must be placed in the `components/` directory at the root of `src/`. Each shared component should have its own folder containing both the component and its associated hook.

```
frontend/web/src/components/
└── button/
    ├── Button.tsx              # Component implementation
    └── useButton.ts            # Component hook (if applicable)
```

Example shared component:

```tsx
import { useButton } from "./useButton";

interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  variant?: "primary" | "secondary";
  disabled?: boolean;
}

export function Button({
  children,
  onClick,
  variant = "primary",
  disabled = false,
}: ButtonProps) {
  const { handleClick, isProcessing } = useButton({ onClick });

  const baseStyles = "px-4 py-2 rounded font-medium transition-colors";
  const variantStyles = {
    primary: "bg-blue-600 text-white hover:bg-blue-700",
    secondary: "bg-gray-200 text-gray-800 hover:bg-gray-300",
  };

  return (
    <button
      onClick={handleClick}
      disabled={disabled || isProcessing}
      className={`${baseStyles} ${variantStyles[variant]}`}
    >
      {children}
    </button>
  );
}
```

```ts
import { useState, useCallback } from "react";

interface UseButtonProps {
  onClick?: () => void;
}

export function useButton({ onClick }: UseButtonProps) {
  const [isProcessing, setIsProcessing] = useState(false);

  const handleClick = useCallback(() => {
    if (onClick) {
      setIsProcessing(true);
      onClick();
      setIsProcessing(false);
    }
  }, [onClick]);

  return { handleClick, isProcessing };
}
```

### Page-Specific Components

Components that are only used within a single page should reside in that page's `components/` subfolder. If a page-specific component later becomes needed elsewhere, refactor it into the shared `components/` directory.

### When to Create a New Component

Create a new component when:

- A piece of UI is used in multiple places.
- A section of a page has distinct, self-contained logic.
- The component improves code readability and maintainability.

---

## Hooks and Business Logic

All business logic should be extracted into custom hooks. This separation ensures that components remain focused on presentation while hooks manage state, side effects, and data transformations.

### Guidelines

- Hooks should be placed in a `hooks/` folder at the same level as the components they serve.
- Name hooks with the `use` prefix (e.g., `useLoginForm`, `useButton`).
- Keep hooks focused on a single responsibility.
- Hooks should return only the values and functions that the component requires.

### Example

```ts
import { useState, FormEvent } from "react";

export function useLoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    setError(null);

    try {
      // Authentication logic
    } catch (err) {
      setError("Invalid credentials");
    } finally {
      setIsSubmitting(false);
    }
  };

  return {
    email,
    setEmail,
    password,
    setPassword,
    error,
    isSubmitting,
    handleSubmit,
  };
}
```

---

## Styling Guidelines

This project uses [Tailwind CSS](https://tailwindcss.com/) for styling. Please skim the docs.

### Raw CSS as a Last Resort

Raw CSS should only be used when Tailwind cannot achieve the desired result. If raw CSS is necessary, document the reason in a comment:

```css
.custom-animation {
  animation: customKeyframe 2s ease-in-out infinite;
}
```

---

## Best Practices

Before contributing to the codebase, please read or skim through the [Bulletproof React](https://github.com/alan2207/bulletproof-react/tree/master) repository. It provides an excellent foundation for scalable React architecture.

### Summary of Key Principles

**Component Design**

- Keep components small and focused on a single responsibility.
- Prefer composition over inheritance.
- Use TypeScript interfaces to define prop types explicitly.

**Code Organization**

- Colocate related code (components, hooks, and utilities that work together should be near each other).
- Avoid deeply nested folder structures; prefer flat hierarchies where practical.

**State Management**

- Lift state only as high as necessary.
- Keep state as close to where it is used as possible.

**Performance**

- Use `useMemo` and `useCallback` judiciously to prevent unnecessary re-renders.
- Avoid premature optimization; measure before optimizing.

**TypeScript**

- Define explicit types for props, state, and function return values.
- Avoid using `any`; prefer `unknown` when the type is truly unknown.

**Code Quality**

- Write self-documenting code with clear variable and function names.
- Add comments only when the "why" is not obvious from the code.
- Keep functions and components small enough to understand at a glance.

---

## Docker Caching

If you encounter issues when with the fronend when you try to run the frontend through the docker container that do not exist when running the web server in dev mode, then please try to invalidate any previous cache and rebuild the container by running the following commands from the root of the project.

```bash
docker compose down
docker system prune -f
docker compose build --no-cache
docker compose up
```

## Quick Reference

| Task                 | Command         |
| -------------------- | --------------- |
| Install dependencies | `bun install`   |
| Start dev server     | `bun run dev`   |
| Build for production | `bun run build` |
