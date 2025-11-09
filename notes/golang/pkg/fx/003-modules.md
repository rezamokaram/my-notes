# Modules

An Fx module is a shareable Go library or package that provides self-contained functionality to an Fx application.

## Writing modules

### example

```go
package main

import (
    "context"
    "fmt"
    "go.uber.org/fx"
)

// --- Service definitions ---

type Config struct {
    AppName string
}

func NewConfig() *Config {
    return &Config{AppName: "MyFXApp"}
}

type Database struct {
    DSN string
}

func NewDatabase(cfg *Config) *Database {
    return &Database{DSN: "postgres://localhost:5432/" + cfg.AppName}
}

func NewHandler(db *Database) *Handler {
    return &Handler{db: db}
}

type Handler struct {
    db *Database
}

func (h *Handler) Serve() {
    fmt.Println("Serving with DB:", h.db.DSN)
}

// --- Define a module ---

var AppModule = fx.Module("app",
    fx.Provide(
        NewConfig,
        NewDatabase,
        NewHandler,
    ),
    fx.Invoke(func(h *Handler) {
        h.Serve()
    }),
)

// --- Main app using the module ---

func main() {
    app := fx.New(
        AppModule, // Include the module
        fx.Invoke(func(lc fx.Lifecycle) {
            lc.Append(fx.Hook{
                OnStart: func(ctx context.Context) error {
                    fmt.Println("Application starting...")
                    return nil
                },
                OnStop: func(ctx context.Context) error {
                    fmt.Println("Application stopping...")
                    return nil
                },
            })
        }),
    )

    app.Run()
}
```

#### Key Takeaways

- fx.Module("name", ...) groups multiple fx.Provide, fx.Invoke, and other options.

- You can import modules from other packages and compose them.

- This keeps your app modular and testable — e.g., one module for “database,” one for “http,” one for “auth,” etc.
