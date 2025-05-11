# Go Compiler Bootstrapping Process

## Which programming language was Go made with

The Go programming language (Golang) was initially implemented in **C**. The original Go compiler, known as **gc (Go compiler)**, was written in C.

However, later on, a new Go compiler called **"Go compiler in Go" (cmd/compile)** was created, which is now **written in Go itself**. This is known as **bootstrapping**.

## 2. **How is a new version of the Go compiler built?**

To build a new version of the Go compiler, you need to use an **existing version of the Go compiler** (from an earlier release) to compile the new version's source code.

### The Bootstrapping Process

- **Step 1:** Start with an existing Go compiler (e.g., Go 1.21).
- **Step 2:** Use this version to **compile the source code** of the new version (e.g., Go 1.22).
- **Step 3:** The result is the new `go` binary with the updated compiler and toolchain.

This process ensures that each new version of Go can be built using the previous one.

### What if there's no existing Go compiler?

If you are starting from scratch (e.g., on a new architecture):

- You can use a **Go bootstrap toolchain** provided by the Go team.
- Alternatively, you can use an older version of a **C-based compiler** like `gccgo`, though it's now deprecated.

---

This bootstrapping approach allows Go to evolve and be self-hosted over time.
