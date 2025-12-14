# Virtualization in Linux

## What is Virtualization?

**Virtualization** is a technology that allows you to run **multiple operating systems (virtual machines or VMs)** on a **single physical computer** at the same time.

In Linux, virtualization uses software to **simulate hardware** (CPU, memory, storage, network), allowing each VM to behave like a separate physical machine while sharing the same hardware.

### Why Virtualization is Used

* Run multiple operating systems on one machine
* Better hardware utilization
* Isolation between systems
* Software testing and development
* Foundation of cloud computing

### Example

A single Linux server can run:

* Ubuntu VM for a web server
* CentOS VM for a database
* Windows Server VM for applications

All running simultaneously.

---

## What is a Hypervisor?

A **hypervisor** is software that:

* Creates and manages virtual machines
* Allocates CPU, memory, storage, and network resources
* Ensures isolation between VMs

Linux commonly uses hypervisors such as **KVM**, **Xen**, and **VirtualBox**.

---

## Types of Hypervisors

There are **two main types of hypervisors**:

---

## Type 1 Hypervisor (Bare-Metal Hypervisor)

### Definition of Type 1

A **Type 1 hypervisor** runs **directly on the physical hardware**, without a host operating system.

### Architecture of Type 1

```scss
Hardware
↓
Type 1 Hypervisor
↓
Virtual Machines
```

### Characteristics of Type 1

* High performance
* Strong security
* Used in servers and data centers
* No host OS overhead

### Linux Examples

* **KVM (Kernel-based Virtual Machine)** – built into the Linux kernel
* **Xen**

### Use Cases

* Cloud platforms
* Enterprise servers
* Production environments

---

## Type 2 Hypervisor (Hosted Hypervisor)

### Definition of Type 2

A **Type 2 hypervisor** runs **on top of a host operating system** like a regular application.

### Architecture of Type 2

```scss
Hardware
↓
Host Operating System (Linux / Windows / macOS)
↓
Type 2 Hypervisor
↓
Virtual Machines
```

### Characteristics of Type 2

* Easy to install and use
* Slightly lower performance
* Depends on the host OS
* Best for learning and testing

### Linux Examples of Type 2

* **VirtualBox**
* **VMware Workstation**

### Use Cases of Type 2

* Desktop virtualization
* Learning Linux
* Software testing

---

## Comparison Table

| Feature          | Type 1 Hypervisor | Type 2 Hypervisor |
| ---------------- | ----------------- | ----------------- |
| Runs on hardware | Yes               | No                |
| Requires host OS | No                | Yes               |
| Performance      | High              | Medium            |
| Typical use      | Servers, Cloud    | Desktop, Testing  |
| Linux examples   | KVM, Xen          | VirtualBox        |

---

## Summary

* Virtualization allows multiple operating systems to run on one physical machine
* A hypervisor is the core component that enables virtualization
* **Type 1 hypervisors** run directly on hardware and offer high performance
* **Type 2 hypervisors** run on top of a host OS and are easier to use
