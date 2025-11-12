# Uber Fx: Export Boundaries, Private, and Related Concepts

## 🧩 1. The Big Picture: Fx Modules and Boundaries

When an app grows, you don’t want *everything* in one big `fx.New(...)` block.  
You want **modules** — independent pieces (like “database”, “API”, “metrics”, etc.) that can be reused or replaced.

Fx supports this idea through **export boundaries** — a way to define what a module *exports* and what stays *private* inside it.

---

## ⚙️ 2. What’s an Export Boundary?

An **export boundary** defines what a module *exposes* (exports) to the outside world, and what stays internal.

In simpler words:
> It’s a contract between a module and the rest of the app.

Modules can:

- **Provide** internal dependencies only used inside the module.
- **Export** selected types for use by the rest of the app.

This helps you control *dependency leakage* — preventing unrelated packages from depending on your internals.

---

## 📦 3. How Fx Defines These Boundaries

Uber Fx uses two special annotations / helpers to define boundaries:

### 🔸 `fx.Private`

Marks constructors or invocations as *internal* to the module.

```go
fx.Private(
    fx.Provide(newDBConnection),
    fx.Provide(newCache),
)
```

These dependencies are **not visible** outside the module.

---

### 🔹 `fx.Export`

Marks specific constructors or types as *exported* (public from the module).

```go
fx.Export(
    fx.Provide(NewUserService),
)
```

Now, other modules can depend on `UserService`.

---

## 🧱 4. Example: A Module with a Clear Boundary

Let’s say you have a **User Module** that handles user logic.

```go
// user/module.go
package user

import "go.uber.org/fx"

type Repo struct{}
type Service struct {
    repo *Repo
}

func NewRepo() *Repo {
    return &Repo{}
}

func NewService(repo *Repo) *Service {
    return &Service{repo: repo}
}

var Module = fx.Module("user",
    fx.Private( // only visible inside "user"
        fx.Provide(NewRepo),
    ),
    fx.Export( // visible outside "user"
        fx.Provide(NewService),
    ),
)
```

And in your main app:

```go
package main

import (
    "fmt"
    "go.uber.org/fx"
    "yourapp/user"
)

func main() {
    fx.New(
        user.Module,
        fx.Invoke(func(s *user.Service) {
            fmt.Println("User service loaded:", s)
        }),
    ).Run()
}
```

Here:

- The `Repo` is **private** (cannot be injected outside `user.Module`).
- The `Service` is **exported** (visible to the main app).

---

## 🧠 5. Why Export Boundaries Matter

Without boundaries, Fx modules can get tangled — any component can depend on any other.

Boundaries enforce:

- **Encapsulation:** hide implementation details.
- **Clear contracts:** define what a module offers.
- **Replaceability:** you can swap modules without breaking others.

---

## ⚙️ 6. Relation Between Private, Export, and Module

| Concept | Description | Visibility |
|----------|--------------|-------------|
| `fx.Private` | Internal-only components | Inside module |
| `fx.Export` | Public API of the module | Outside module |
| `fx.Module` | Groups related Provides/Invokes under a name | Depends on private/export |

---

## 🧩 7. Export Boundary Functions (Advanced)

Internally, Fx calls your module's **boundary functions** when it wires dependencies.  
An **export boundary function** describes how to “bridge” from the module to the outside world — i.e., which types are exported.

You don’t usually call these directly; instead, Fx’s `fx.Export` and `fx.Private` APIs *define* these boundaries automatically.

Conceptually, Fx’s wiring graph looks like this:

```tree
App
 ├── user.Module
 │     ├── Private: Repo
 │     └── Export:  Service
 └── api.Module
       └── depends on user.Service
```

The *export boundary function* is the interface Fx uses internally to say:
> “The only thing you can see from user.Module is Service.”

---

## ✅ Summary

| Term | Purpose |
|------|----------|
| **Export boundary** | Defines what a module exposes to the rest of the app |
| **`fx.Private`** | Keeps constructors/invokes internal to the module |
| **`fx.Export`** | Makes constructors/invokes visible outside the module |
| **`fx.Module`** | Groups related provides/invokes under a named module |
| **Boundary function** | Fx’s internal representation of a module’s public contract |

---

**In short:**  
Export boundaries in Uber Fx help you design modular, maintainable systems by defining clear, explicit interfaces between parts of your application.
