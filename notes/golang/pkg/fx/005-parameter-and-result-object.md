# Parameter and Result Objects

This document summarizes three progressively advanced examples demonstrating how to use **parameter objects (`fx.In`)** and **result objects (`fx.Out`)** in Uber Fx for clean, declarative dependency injection.

---

## 🧩 Example 1: Provide DB and Logger, Inject Both (Basic Example)

$Code$

```go
package main

import (
    "fmt"
    "go.uber.org/fx"
)

type DB struct{}
type Logger struct{}

// Provide DB
func NewDB() *DB {
    fmt.Println("DB created")
    return &DB{}
}

// Provide Logger
func NewLogger() *Logger {
    fmt.Println("Logger created")
    return &Logger{}
}

// Consume both using fx.In
type ServiceParams struct {
    fx.In
    DB     *DB
    Logger *Logger
}

func NewService(p ServiceParams) {
    fmt.Println("Service started with DB:", p.DB, "and Logger:", p.Logger)
}

func main() {
    app := fx.New(
        fx.Provide(NewDB, NewLogger),
        fx.Invoke(NewService),
    )
    app.Run()
}
```

$Output$

```output
DB created
Logger created
Service started with DB: 0xc000010200 and Logger: 0xc000010208
```

---

## ⚙️ Example 2: Using `fx.Out` for DB and Logger (Single Function Provides Both)

$Code$

```go
package main

import (
    "fmt"
    "go.uber.org/fx"
)

type DB struct{}
type Logger struct{}

// Provide both DB and Logger using fx.Out
type AppResources struct {
    fx.Out
    DB     *DB
    Logger *Logger
}

func NewAppResources() AppResources {
    fmt.Println("Creating DB and Logger...")
    return AppResources{
        DB:     &DB{},
        Logger: &Logger{},
    }
}

// Consume both using fx.In
type ServiceParams struct {
    fx.In
    DB     *DB
    Logger *Logger
}

func NewService(p ServiceParams) {
    fmt.Println("Service started with:")
    fmt.Println(" → DB:", p.DB)
    fmt.Println(" → Logger:", p.Logger)
}

func main() {
    app := fx.New(
        fx.Provide(NewAppResources),
        fx.Invoke(NewService),
    )
    app.Run()
}
```

$Output$

```output
Creating DB and Logger...
Service started with:
 → DB: 0xc000010200
 → Logger: 0xc000010208
```

### Why Use `fx.Out`?

- One function can provide multiple outputs
- Explicit declaration of what is being provided
- Works perfectly with named/grouped dependencies
- Clean and flexible for scaling modules

---

## 🧱 Example 3: Separate DB and Logger Providers with Individual Result Objects

$Code$

```go
package main

import (
    "fmt"
    "go.uber.org/fx"
)

type DB struct{}
type Logger struct{}

// Provide DB using fx.Out
type DBResult struct {
    fx.Out
    DB *DB
}

func NewDB() DBResult {
    fmt.Println("Creating DB...")
    return DBResult{DB: &DB{}}
}

// Provide Logger using fx.Out
type LoggerResult struct {
    fx.Out
    Logger *Logger
}

func NewLogger() LoggerResult {
    fmt.Println("Creating Logger...")
    return LoggerResult{Logger: &Logger{}}
}

// Consume both with fx.In
type ServiceParams struct {
    fx.In
    DB     *DB
    Logger *Logger
}

func NewService(p ServiceParams) {
    fmt.Println("Service started with:")
    fmt.Println(" → DB:", p.DB)
    fmt.Println(" → Logger:", p.Logger)
}

func main() {
    app := fx.New(
        fx.Provide(NewDB, NewLogger),
        fx.Invoke(NewService),
    )
    app.Run()
}
```

$Output$

```output
Creating DB...
Creating Logger...
Service started with:
 → DB: 0xc000010200
 → Logger: 0xc000010208
```

---

## ✅ Summary Table

| Example | Technique | Functions | Purpose |
|----------|------------|------------|----------|
| **1** | Basic `fx.In` | `NewDB`, `NewLogger`, `NewService` | Inject multiple dependencies using parameter struct |
| **2** | Combined `fx.Out` | `NewAppResources`, `NewService` | Provide multiple outputs from one function |
| **3** | Separate `fx.Out` per resource | `NewDB`, `NewLogger`, `NewService` | Split providers for clarity and modularity |

---
---
---

### 🧩 Key Concepts Recap

| Concept | Description |
|----------|-------------|
| **`fx.In`** | Used in struct parameters to declare dependencies to be injected |
| **`fx.Out`** | Used in struct return types to declare what the function provides |
| **Named Values (`name:"..."`)** | Used to differentiate multiple instances of the same type |
| **Grouped Values (`group:"..."`)** | Used to collect multiple instances into a slice |
| **Optional Fields (`optional:"true"`)** | Allows dependency to be absent without causing startup failure |

---

**In short:**  

- `fx.In` makes *consuming dependencies* explicit and safe.  
- `fx.Out` makes *providing dependencies* explicit and modular.  
- Both make your Uber Fx applications easier to understand, refactor, and extend.
