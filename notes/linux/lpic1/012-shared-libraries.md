# shared libraries

## commands and utilities to learn

- ldd
- ldconfig
- /etc/ld.so.conf
- LD_LIBRARY_PATH

## Linking

When you compile a program (e.g., with gcc or clang), two main steps happen:

Compilation – each source file (.c) becomes an object file (.o)

Linking – object files and libraries are combined into a final executable

Linking resolves the references your code makes to external functions—for example, functions from the C standard library (printf, malloc, etc.) or any other library.

## Types of Libraries

1. Static Libraries (.a)

- Library code is copied into the final executable
- Executable becomes self-contained
- Larger file size
- No need for library to be present on the system at runtime
- Example static library: libm.a (math)

1. Dynamic / Shared Libraries (.so)

- Not copied into the executable
- Loaded at runtime from the system
- Executable is smaller
- Library updates do not require recompilation
- Requires the library to be present on the system

Example shared library: libm.so

## How to Link Libraries

Assume you have a program main.c that uses the math library (-lm, because math functions are in libm).

Compile and link:

```sh
gcc main.c -o main -lm
```

How GCC interprets this:

- -l\<NAME> → search for lib\<NAME>.so or lib\<NAME>.a

- Search path is controlled by -L

Example:

```sh
gcc main.c -L/path/to/libs -lmylib
```

This looks for:

```sh
/path/to/libs/libmylib.so
/path/to/libs/libmylib.a
```

## Where the Linker Searches for Libraries (Compile Time)

The linker searches:

- `/lib` -> 32 bit, `/lib64` -> 64 bit
- `/usr/lib`
- `/usr/local/lib`
- Any directory passed with `-L`
- Directories from environment variable `LIBRARY_PATH`

## LDD

ldd is a Linux command-line tool that prints the shared libraries required by a given program or shared object.

Example:

```sh
ldd /bin/ls
```

You’ll see output like:

```sh
linux-vdso.so.1 (0x00007ffd2d5fe000)
libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x7f4e27834000)
ld-linux-x86-64.so.2 (0x7f4e27a42000)
```

## Short Explanation of `ldconfig`

`ldconfig` updates the system's list of where shared libraries (`.so` files) are located so programs can find them.

It performs three main tasks:

1. Scans library directories (like `/lib`, `/usr/lib`, etc.)
2. Updates the library cache (`/etc/ld.so.cache`)
3. Fixes/creates library version symlinks

You typically run it after installing a new library.

## Example: Adding a Library and Successful `ldconfig` Output

### Command

```bash
sudo cp libfoo.so.1.2.3 /usr/local/lib/
sudo ldconfig
```

### Example Successful Output

```txt
/usr/local/lib: libfoo.so.1 -> libfoo.so.1.2.3
/usr/local/lib: libfoo.so -> libfoo.so.1
ldconfig: /usr/local/lib/libfoo.so.1.2.3 is the most recent version
```
