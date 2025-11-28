# Logs

## kernel ring buffer

- Imagine a circular array.

- When it gets full, it wraps around and starts overwriting from the beginning.

This design:

- Minimizes memory use

- Avoids having to allocate more memory dynamically

- Ensures that the most recent messages are always available

### Related Tools & Commands

- `dmesg -w`: Follow kernel messages in real-time (like `tail -f` for logs).

- `journalctl -k`: On systemd-based systems, shows kernel messages from the journal.

- `loglevel`: Controls what messages are shown (e.g., KERN_INFO, KERN_ERR, etc.).

## dmesg

`dmesg` stands for "diagnostic message" and is a Linux command used to view the kernel ring buffer, which contains messages from the Linux kernel.

These messages include:

- Boot logs
- Hardware detection
- Driver initialization
- Kernel warnings and errors
- Messages from kernel modules (like printk() outputs)

## ‌ journalctl

`journalctl` is a command-line tool used on Linux systems (especially those using systemd) to view and query logs collected by the systemd journal service (`journald`). It replaces older logging mechanisms like `/var/log/messages` or `syslog`.

```terminal
# follows the logs
journalctl -f

# kernel logs
journalctl -k

# boot logs
journalctl -b
```
