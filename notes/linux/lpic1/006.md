# Run levels

## What are Runlevels?

Runlevels are a concept from the traditional SysV init system used in many older Linux distributions (and still in some modern ones, though many have moved to systemd).

A runlevel defines a specific state of the machine, mainly what services or processes should be running.

Think of them as modes that determine what your system is doing (e.g., running as a desktop, server, or in maintenance mode).

| Runlevel | Description                                                               |
| -------- | ------------------------------------------------------------------------- |
| `0`      | **Halt** — Shuts down the system.                                         |
| `1`      | **Single-user mode** — Maintenance mode, minimal services, no networking. |
| `2`      | Multi-user mode **without** networking (varies by distro).                |
| `3`      | Full multi-user mode **with** networking (text mode, no GUI).             |
| `4`      | Unused (user-definable).                                                  |
| `5`      | Multi-user with **GUI** (X11/graphical login).                            |
| `6`      | **Reboot** the system.                                                    |

💡 Not all distros use all runlevels the same way—some (like Debian) treat 2–5 the same.

to check current runlevel:

```sh
runlevel
// output sample
// N 5

init 3    # Switch to runlevel 3
telinit 5 # newer versions
```

but now days linux use systemd target unit instead of runlevel

## Runlevel vs systemd (Modern Systems)

| SysV Runlevel | systemd Target      |
| ------------- | ------------------- |
| 0             | `poweroff.target`   |
| 1             | `rescue.target`     |
| 3             | `multi-user.target` |
| 5             | `graphical.target`  |
| 6             | `reboot.target`     |

Check current target:

```bash
systemctl get-default
```

Change target:

```bash
systemctl isolate multi-user.target
```

Set default:

```bash
systemctl set-default graphical.target
```

To see how a target is built (which units it depends on), use:

```bash
systemctl list-dependencies <target>
systemctl list-dependencies graphical.target
# Another way to examine a target's configuration is by checking its associated .target unit file, which is typically written in a TOML-like format.
```

Output:

```txt
graphical.target
● ├─multi-user.target
● │ ├─network.target
● │ └─...
● └─display-manager.service
```

### main systemd targets

#### 1. `default.target`

- **What it is**: The default target the system boots into.
- **Alias of**: Usually points to `graphical.target` or `multi-user.target`.
- **Use**: Determines the default system state after boot.
- **Location**: `/etc/systemd/system/default.target` (usually a symlink)

---

#### 2. `graphical.target`

- **What it is**: Multi-user system **with a GUI**.
- **Includes**: Everything in `multi-user.target`, plus graphical session services.
- **Use**: Desktop systems or servers with a graphical interface.

---

#### 3. `multi-user.target`

- **What it is**: Multi-user system **without a GUI**.
- **Use**: Servers or headless systems.
- **Includes**: Networking, login prompt, and most essential system services.

---

#### 4. `rescue.target`

- **What it is**: Minimal **single-user mode** with a root shell.
- **Use**: For emergency maintenance or recovery.
- **Includes**: Basic system services and file systems.

---

#### 5. `emergency.target`

- **What it is**: The most minimal environment; boots to a root shell with only the root filesystem mounted (read-only).
- **Use**: For the most critical recovery tasks when even `rescue.target` cannot be used.
- **Includes**: No services started, no filesystems mounted except root (read-only).

#### Shutdown targets

- `halt.target`: Stops the system without powering off (cpu halts)
- `poweroff.target`: Shuts down and powers off
- `reboot.target`: Reboots the machine

## What are /etc/rc1-6 and /etc/init?

### 1. `/etc/rc[1-6].d/` directories and `/etc/init.d/`

These are legacy directories and scripts used by the *SysV* init system (the older init system before systemd became popular).

**Purpose**: They define runlevels and contain startup and shutdown scripts.

For example, `rc1.d` corresponds to runlevel 1 (single-user mode), rc3.d to multi-user mode without GUI, and rc5.d to multi-user with GUI.

Inside these directories are symbolic links that point to scripts in `/etc/init.d/` which are run during boot, shutdown, or runlevel changes.

---

### 2. `/etc/init`

- This directory (and files inside) is related to Upstart, another init system used mostly in Ubuntu before switching to systemd.

- Upstart uses event-driven jobs described in `/etc/init/*.conf` files.

These jobs define how and when services start, stop, or respawn.

---

### So, how does this relate to systemd?

- Modern Linux distros mostly use *systemd* now instead of *SysV* init or *Upstart*.
- systemd uses `.service`, `.target`, and other unit files located mainly in /etc/systemd/system/ and /usr/lib/systemd/system/.
- Many distros still keep `/etc/init.d/` scripts for backwards compatibility — systemd can run these SysV init scripts via a compatibility layer.
- But `/etc/rc1-6.d/` and `/etc/init/` are considered legacy, mostly replaced by systemd targets and services.

| File/Dir          | Used by   | Purpose                             | Status in modern Linux     |
| ----------------- | --------- | ----------------------------------- | -------------------------- |
| `/etc/rc[1-6].d/` | SysV init | Runlevel-specific startup scripts   | Legacy, compatibility only |
| `/etc/init.d/`    | SysV init | Service scripts                     | Legacy, compatibility only |
| `/etc/init/`      | Upstart   | Job configuration files             | Mostly obsolete            |
| `systemd` units   | systemd   | Unit files for services and targets | Current standard           |
