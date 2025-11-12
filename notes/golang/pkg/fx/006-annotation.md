# 🧩 Uber Fx Annotations — Complete Guide

This document explains both **struct-based annotations** (`fx.In` / `fx.Out`) and **functional annotations** (`fx.Annotate`) in **Uber Fx**.

---

## 🔹 1. What Are Annotations?

Annotations in Uber Fx let you differentiate between multiple dependencies of the same type, or group them together.
They make dependency injection more flexible and explicit.

| Annotation     | Purpose                                                  |
| -------------- | -------------------------------------------------------- |
| **`name`**     | Distinguish between different instances of the same type |
| **`group`**    | Collect multiple dependencies into a slice               |
| **`optional`** | Allow missing dependencies without errors                |

---

## ⚙️ 2. Struct-Based Annotations (`fx.In` / `fx.Out`)

Annotations can be declared via struct tags using Fx's `fx.In` and `fx.Out` types.

### 🧱 2-Example — Named Dependencies

```go
type DB struct {
    Role string
}

// Provide two DBs with names
type DBOutputs struct {
    fx.Out
    ReadDB  *DB `name:"read"`
    WriteDB *DB `name:"write"`
}

func NewDatabases() DBOutputs {
    return DBOutputs{
        ReadDB:  &DB{Role: "read"},
        WriteDB: &DB{Role: "write"},
    }
}

// Consume named dependency
type ServiceParams struct {
    fx.In
    ReadDB *DB `name:"read"`
}

func NewService(p ServiceParams) {
    fmt.Println("Using DB:", p.ReadDB.Role)
}
```

---

### 🧱 2-Example — Grouped Dependencies

```go
type Middleware struct {
    Name string
}

type MiddlewareResult struct {
    fx.Out
    Middleware *Middleware `group:"middlewares"`
}

func NewAuth() MiddlewareResult {
    return MiddlewareResult{Middleware: &Middleware{Name: "auth"}}
}

func NewLogger() MiddlewareResult {
    return MiddlewareResult{Middleware: &Middleware{Name: "logger"}}
}

type ServerParams struct {
    fx.In
    Middlewares []*Middleware `group:"middlewares"`
}

func NewServer(p ServerParams) {
    for _, m := range p.Middlewares {
        fmt.Println("Middleware:", m.Name)
    }
}
```

---

### 🧩 Optional Dependencies

```go
type Params struct {
    fx.In
    Logger *Logger `optional:"true"`
}
```

If `Logger` isn’t provided, Fx won’t fail — it just sets `Logger = nil`.

---

### ✅ Summary — Struct-Based Annotations

| Tag                   | Meaning                      | Pattern |
| --------------------- | ---------------------------- | ------- |
| `name:"read"`         | Provide a named value        | 1:1     |
| `group:"middlewares"` | Contribute to a shared slice | N:1     |
| `optional:"true"`     | Allow missing dependency     | —       |

---

## ⚙️ 3. Functional Annotations (`fx.Annotate`)

Instead of using structs, you can annotate a function directly with `fx.Annotate`.

It’s perfect for small constructors where you don’t want to define `fx.In` / `fx.Out` structs.

---

### 🧱 3-Example — Named Dependencies

```go
type DB struct {
    Role string
}

func NewReadDB() *DB { return &DB{Role: "read"} }
func NewWriteDB() *DB { return &DB{Role: "write"} }

type Params struct {
    fx.In
    ReadDB  *DB `name:"read"`
    WriteDB *DB `name:"write"`
}

func NewService(p Params) {
    fmt.Println("Using:", p.ReadDB.Role, p.WriteDB.Role)
}

func main() {
    app := fx.New(
        fx.Provide(
            fx.Annotate(NewReadDB, fx.ResultTags(`name:"read"`)),
            fx.Annotate(NewWriteDB, fx.ResultTags(`name:"write"`)),
        ),
        fx.Invoke(NewService),
    )
    app.Run()
}
```

---

### 🧱 3-Example — Grouped Dependencies

```go
type Middleware struct {
    Name string
}

func Auth() *Middleware    { return &Middleware{Name: "auth"} }
func Logging() *Middleware { return &Middleware{Name: "logging"} }

type ServerParams struct {
    fx.In
    Middlewares []*Middleware `group:"middlewares"`
}

func NewServer(p ServerParams) {
    for _, m := range p.Middlewares {
        fmt.Println(m.Name)
    }
}

func main() {
    app := fx.New(
        fx.Provide(
            fx.Annotate(Auth, fx.ResultTags(`group:"middlewares"`)),
            fx.Annotate(Logging, fx.ResultTags(`group:"middlewares"`)),
        ),
        fx.Invoke(NewServer),
    )
    app.Run()
}
```

---

### 🧩 Annotating Parameters (Inputs)

```go
func NewService(readDB, writeDB *DB) *Service {
    return &Service{Read: readDB, Write: writeDB}
}

fx.Provide(
    fx.Annotate(
        NewService,
        fx.ParamTags(`name:"read"`, `name:"write"`),
    ),
)
```

---

## ✅ Summary — Functional Annotations

| Function          | Purpose                        | Example                                            |
| ----------------- | ------------------------------ | -------------------------------------------------- |
| `fx.ParamTags()`  | Annotate inputs                | `fx.ParamTags('name:"read"')`                      |
| `fx.ResultTags()` | Annotate outputs               | `fx.ResultTags('group:"middlewares"')`             |
| `fx.Annotate()`   | Attach metadata to constructor | `fx.Annotate(NewDB, fx.ResultTags('name:"read"'))` |

---

## ⚖️ 4. Comparison — `fx.In` / `fx.Out` vs `fx.Annotate`

| Feature                          | `fx.In` / `fx.Out` | `fx.Annotate` |
| -------------------------------- | ------------------ | ------------- |
| **Best for complex types**       | ✅                  | ❌             |
| **Best for simple providers**    | ❌                  | ✅             |
| **Supports name/group/optional** | ✅                  | ✅             |
| **Explicit field names**         | ✅                  | ❌             |
| **Less boilerplate**             | ❌                  | ✅             |

---

## 🧩 Example — Full Mix

```go
fx.Provide(
    fx.Annotate(
        NewHandler,
        fx.ParamTags(`group:"middlewares"`, `name:"logger"`),
        fx.ResultTags(`group:"routes"`),
    ),
)
```

Meaning:

* Takes a group of middlewares and a named logger.
* Provides a route in the `"routes"` group.

---

## ✅ TL;DR Summary

| Concept                | Description                         |
| ---------------------- | ----------------------------------- |
| **`fx.In` / `fx.Out`** | Struct-based dependency metadata    |
| **`fx.Annotate`**      | Functional, lightweight alternative |
| **`fx.ParamTags()`**   | Annotate inputs                     |
| **`fx.ResultTags()`**  | Annotate outputs                    |
| **`name`**             | Label a dependency uniquely         |
| **`group`**            | Combine multiple dependencies       |
| **`optional`**         | Make dependency optional            |

---

**Use `fx.In` / `fx.Out`** for clarity in large components.
**Use `fx.Annotate`** for small, simple constructors.
Both can coexist seamlessly in the same Fx application.
