# Understanding `dpkg` in Linux

## What is `dpkg`?

`dpkg` (Debian Package) is the low-level package manager for Debian-based Linux distributions such as Debian, Ubuntu, Linux Mint, and others. It is responsible for installing, removing, querying, and managing `.deb` packages.

Unlike higher-level tools like `apt` or `apt-get`, which automatically handle dependencies, `dpkg` works directly with `.deb` files and **does not** resolve dependencies automatically.

---

## Common `dpkg` Commands

### Install a Package

```bash
dpkq -i package.deb
```

Installs the specified `.deb` file. Errors may occur if dependencies are missing.

---

### Remove a Package (Keep Configuration Files)

```bash
dpkg -r package_name
```

Removes the package binaries but preserves config files.

---

### Purge a Package (Remove Everything)

```bash
dpkg -P package_name
```

Removes both the package and its configuration files.

---

### List Installed Packages

```bash
dpkg -l
```

Shows all installed packages.

To search for a specific package:

```bash
dpkg -l | grep name
```

---

### Get Information About a Package

```bash
dpkg -s package_name
```

Displays details: version, status, maintainer, description, etc.

---

### List Files Installed by a Package

```bash
dpkg -L package_name
```

Shows all files installed by the package.

---

### Find Which Package Owns a File

```bash
dpkg -S /path/to/file
```

Useful to identify which installed package contains a given file.

---

### Unpack a `.deb` File (Without Installing)

```bash
dpkg -x package.deb directory/
```

Extracts package contents without setting up or installing.

---

## Important Flags and Switches

| Flag          | Description                           |
| ------------- | ------------------------------------- |
| `-i`          | Install a `.deb` package              |
| `-r`          | Remove a package (keep configuration) |
| `-P`          | Purge package (remove config)         |
| `-l`          | List installed packages               |
| `-L`          | List files from installed package     |
| `-s`          | Display package info                  |
| `-S`          | Search for which package owns a file  |
| `-x`          | Extract `.deb` contents               |
| `--configure` | Configure an unpacked package         |

Example:

```bash
dpkg --configure package_name
```

---

## `dpkg-reconfigure`

`dpkg-reconfigure` is used to re-run the configuration scripts for an already installed package.

This is especially useful for packages that prompt configuration during installation (e.g., timezone, keyboard layout, services).

### Usage

```bash
dpkg-reconfigure package_name
```

This will bring up the configuration dialog again.

### Useful Options

| Option | Description                                               |
| ------ | --------------------------------------------------------- |
| `-f`   | Choose frontend (e.g., `dialog`, `noninteractive`)        |
| `-p`   | Set priority of questions (e.g., `low`, `medium`, `high`) |

Example:

```bash
dpkg-reconfigure -f dialog tzdata
```

Reconfigures time zone settings.

---

## Summary

* `dpkg` is the core tool for managing `.deb` packages.
* It works at a low level and does not handle dependency resolution.
* Use higher tools like `apt` for dependency management.
* `dpkg-reconfigure` allows you to re-run configuration for installed packages.

This knowledge is essential for system administration on Debian-based Linux systems.
