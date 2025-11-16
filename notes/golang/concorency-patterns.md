# Golang Concurrency Patterns

Golang provides built-in support for concurrency through goroutines and channels, enabling a variety of concurrency patterns. Below is a detailed list of commonly used concurrency patterns in Golang.

---

## 1. Worker Pool Pattern

### Description

A set of worker goroutines processes tasks from a shared channel.

### Use Case

Efficiently handling a large number of tasks with limited resources (e.g., processing HTTP requests).

### Example

```go
func worker(id int, tasks <-chan int, results chan<- int) {
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task)
        results <- task * 2
    }
}

func main() {
    tasks := make(chan int, 100)
    results := make(chan int, 100)

    for i := 1; i <= 4; i++ {
        go worker(i, tasks, results)
    }

    for j := 1; j <= 10; j++ {
        tasks <- j
    }
    close(tasks)

    for k := 1; k <= 10; k++ {
        fmt.Println(<-results)
    }
}
```

---

## 2. Fan-Out, Fan-In

### Fan-Out

Spawns multiple goroutines to process tasks concurrently.

### Fan-In

Combines results from multiple goroutines into a single channel.

### Fan-Out, Fan-In Use Case

Distributing workload and aggregating results.

### Fan-Out, Fan-In Example

```go
func generator(numbers ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range numbers {
            out <- n
        }
        close(out)
    }()
    return out
}

func squareWorker(in <-chan int, out chan<- int) {
    for n := range in {
        out <- n * n
    }
}

func main() {
    in := generator(1, 2, 3, 4, 5)

    out := make(chan int)
    for i := 0; i < 3; i++ {
        go squareWorker(in, out)
    }

    for i := 0; i < 5; i++ {
        fmt.Println(<-out)
    }
}
```

---

## 3. Pipeline Pattern

### Pipeline Pattern Description

Passes data through a series of stages, each performing a specific operation.

### Pipeline Pattern Use Case

Sequential processing where each stage operates independently.

### Pipeline Pattern Example

```go
func gen(numbers ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range numbers {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    numbers := gen(2, 3, 4)
    squares := square(numbers)

    for n := range squares {
        fmt.Println(n)
    }
}
```

---

## 4. Publish-Subscribe (Pub-Sub)

### Pub-Sub Description

One goroutine (publisher) sends messages to multiple subscribers via channels.

### Pub-Sub Use Case

Event-driven systems.

### Pub-Sub Example

```go
func publish(ch chan string, messages []string) {
    for _, msg := range messages {
        ch <- msg
    }
    close(ch)
}

func subscribe(ch chan string, name string) {
    for msg := range ch {
        fmt.Printf("%s received: %s\n", name, msg)
    }
}

func main() {
    ch := make(chan string)
    messages := []string{"hello", "world", "golang"}

    go publish(ch, messages)

    go subscribe(ch, "Subscriber1")
    go subscribe(ch, "Subscriber2")

    time.Sleep(1 * time.Second)
}
```

---

## 5. Mutex for Shared Resources

### Shared Resources Description

Protects shared data with a `sync.Mutex` to prevent race conditions.

### Shared Resources Use Case

Safe access to shared variables.

### Shared Resources Example

```go
import "sync"

var count int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    mu.Lock()
    count++
    mu.Unlock()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println("Final Count:", count)
}
```

---

## 6. Rate Limiting

### Rate Limiting Description

Controls the rate of processing tasks using a `time.Ticker`.

### Rate Limiting Use Case

Limiting API calls or task processing.

### Rate Limiting Example

```go
func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.NewTicker(200 * time.Millisecond)
    defer limiter.Stop()

    for req := range requests {
        <-limiter.C
        fmt.Println("Processing request", req)
    }
}
```

---

## 7. Select Statement for Multiplexing

### Multiplexing Description

Waits on multiple channels and handles whichever one is ready.

### Multiplexing Use Case

Handling multiple input sources or timeout scenarios.

### Multiplexing Example

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from ch1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from ch2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

---

## 8. Context for Cancellation

### Cancellation Description

Manages goroutines with cancellation and deadlines using `context.Context`.

### Cancellation Use Case

Graceful shutdown or timeout handling.

### Cancellation Example

```go
import "context"

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker stopped")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go worker(ctx)

    time.Sleep(3 * time.Second)
    fmt.Println("Main done")
}
```

---

## 9. Actor Model

### Actor Description

Encapsulates state and communicates through messages.

### Actor Use Case

High-concurrency systems with minimal shared state.

### Actor Example

```go
func actor(ch chan int) {
    state := 0
    for msg := range ch {
        state += msg
        fmt.Println("Current state:", state)
    }
}

func main() {
    ch := make(chan int)
    go actor(ch)

    ch <- 1
    ch <- 2
    ch <- 3

    close(ch)
}
```

---

## 10. Deadlock Prevention

### Deadlock Prevention Description

Avoids cyclic dependencies between goroutines.

### Deadlock Prevention Use Case

Safeguarding systems from freezing.

### Note

Avoid shared locks or use proper lock ordering in your design.
