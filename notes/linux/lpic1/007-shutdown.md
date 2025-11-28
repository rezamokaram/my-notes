# shutdown and notifying users

## ACPI (Advanced Configuration and Power interface)

ACPI (Advanced Configuration and Power Interface) in Linux is a standard that the operating system uses to:

- Manage power (sleep, suspend, hibernate, shutdown).
- Control hardware states (CPU frequency scaling, fan speed, thermal zones, battery).
- Handle events (like pressing the power button or closing a laptop lid).

Linux uses the acpid daemon (or systemd-logind) to listen for ACPI events and act on them.

**$In$ $short$: ACPI is how Linux talks to your computer’s firmware to manage power and hardware events.**

## mesg

`mesg` command in Linux controls whether other users are allowed to send you messages using the `write` or `wall` commands.

- `mesg y` → allow others to send you messages.
- `mesg n` → disallow messages (default in many  systems for privacy).
- `mesg` (with no args) → shows current status (is y or is n).

## shutdown

```bash
shutdown [OPTION] [TIME] [MESSAGE]
```

### Common options

- `shutdown now` → shut down immediately.
- `shutdown +10 "System will go down"` → shut down after 10 minutes with a warning message.
- `shutdown -h now` → halt and power off now.
- `shutdown -r now` → reboot immediately.
- `shutdown -c` → cancel a scheduled shutdown.

| Command      | Action                            | Notes                                                      |
| ------------ | --------------------------------- | ---------------------------------------------------------- |
| **shutdown** | Safely halt, power off, or reboot | Can schedule (e.g., `+10`), notifies users, safest option. |
| **poweroff** | Powers off the system immediately | Equivalent to `shutdown -h now` (on systemd).              |
| **halt**     | Halts (stops) the system          | CPUs stop, may not fully power off depending on hardware.  |
| **reboot**   | Restarts the system immediately   | Equivalent to `shutdown -r now`.                           |

## reboot

## wall

The wall command in Linux is used to send a message to all logged-in users.

- It writes your message to the terminals of all users currently logged in.
- Usually requires root to broadcast to everyone, but regular users can send to users in the same group/terminal.

## write

The write command in Linux lets you send a message directly to another logged-in user’s terminal.

```bash
$ write username
Hello! Are you there?
```

## type

The `type` command in Linux is used to identify how a given command name would be interpreted by the shell. It tells you whether the command is a shell built-in, an alias, a function, or an external executable.

```bash
type [options] command_name
```

```bash
$ type ls
ls is aliased to `ls --color=auto`

$ type cd
cd is a shell builtin

$ type python3
python3 is /usr/bin/python3
```

### What it tells you

- **Alias** → a shortcut defined in your shell (`alias ll='ls -l'`).
- **Builtin** → a command built into the shell (`cd`, `echo`).
- **Function** → a shell function you or scripts defined.
- **File/Executable** → an external program in your PATH (`/usr/bin/ls`).

### Useful options

- `type -a <command>` → shows all occurrences (aliases, builtins, and binaries) of the command.
- `type -t <command>` → shows only the type (alias, builtin, file, function).
