# installing bootloader

## review on how system boots

our system use BIOS or UEFI:

BIOS -> (start the firmware, firmware see the motherboard and do something else) -> firmware "start power on self test"  

Bios reads the first part of disk(0,0,1) 512 byte -> then send this part as a command to cpu. then computer loads the whole grub parts (512 byte was a small part for loading grub).  

then grub can find the disk, it can find init ram fs and ...

then kernel loads and starts init process.

UEFI uses stages:

- first is uefi security stage
- ...
- in last phase uefi looks for a EFI System Partition that is just a FAT32 partition with PE executables and runs them. (then grub loads and ...)

## grub legacy - very short

GRUB Legacy is an old Linux **bootloader** that shows a menu when the computer starts and lets you choose what to boot.

---

### Config file

Its settings are stored in:

/boot/grub/menu.lst

What’s inside menu.lst?

Two things:

#### 1. **General settings**

Example:  
default 0 # first entry boots automatically.  
timeout 5 # wait 5 seconds.  

#### 2. **Boot entries**

Each entry tells GRUB how to start an OS.

Example Linux entry:

```sh
title Linux
root (hd0,0)
kernel /boot/vmlinuz root=/dev/sda1
initrd /boot/initrd.img
```

Example Windows entry:

```sh
title Windows
rootnoverify (hd0,1)
chainloader +1
```

### Disk notation

GRUB counts from zero:

- `(hd0,0)` = first disk, first partition  
- `(hd0,1)` = first disk, second partition

## how to install grub legacy

Use the GRUB shell:

```bash
grub
root (hd0,0)
setup (hd0)
```

This installs GRUB to the first disk.

Use the simpler command:

```bash
grub-install /dev/sda
```

This installs GRUB to the MBR of the first drive.
