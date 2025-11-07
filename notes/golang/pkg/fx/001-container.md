# Container

Container is the abstraction responsible for holding all constructors and values. It’s the primary means by which an application interacts with Fx. You teach the container about the needs of your application, how to perform certain operations, and then you let it handle actually running your application.

Fx does not provide direct access to the container. Instead, you specify operations to perform on the container by providing fx.Options to the fx.New constructor.

```go
package fx

type App
  func New(opts ...Option) *App
  func (app *App) Run()

type Option
  func Provide(constructors ...interface{}) Option
  func Invoke(funcs ...interface{}) Option
```

## Providing values

You must provide values to the container before you can use them. Fx provides two ways to provide values to the container:

### provide

```go
fx.Provide(
  func(cfg *Config) *Logger { /* ... */ },
)
```

### supply

```go
fx.Provide(
  fx.Supply(&Config{
    Name: "my-app",
  }),
)

// is the same as

fx.Provide(func() *Config { return &Config{Name: "my-app"} })

```

*However, even then, fx.Supply comes with a caveat: it can only be used for non-interface values.*

### transient

#### todo

### invoke

`fx.Invoke` is typically used ***for root-level invocations***, like starting a server or running a main loop. It's also useful for invoking ***functions that have side effects***.

```go
fx.New(
  fx.Provide(newHTTPServer),
  fx.Invoke(startHTTPServer),
).Run()
```
