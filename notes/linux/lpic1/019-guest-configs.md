# Guest-Specific Configuration in Virtual Machines

In virtualization, **guest-specific configuration** refers to settings that uniquely define and control the behavior, identity, and resources of an individual virtual machine (VM) (the *guest*) as seen by both the hypervisor and the guest operating system.

Below is a structured overview of the most common and important guest-specific configs, with emphasis on **UUIDs, port mappings**, and related concepts.

---

## 1. Guest Identity

These settings uniquely identify a VM and are often relied on by operating systems, licensing systems, and management tools.

### 1.1 UUID (Universal Unique Identifier)

* A **UUID** uniquely identifies a VM instance
* Used by:

  * Hypervisors (inventory, snapshots, migration)
  * Guest OS (systemd, udev, cloud-init)
  * Licensing / DRM software

Examples:

* VMware: `uuid.bios`, `uuid.location`
* libvirt/KVM: `<uuid>` in domain XML
* Hyper-V: VM GUID

Why it matters:

* Cloning a VM without regenerating UUIDs can cause:

  * Network conflicts
  * Licensing issues
  * Cluster membership collisions

---

### 1.2 Hostname & Machine ID

* Hostname: OS-level identity
* Machine ID:

  * Linux: `/etc/machine-id`
  * Windows: SID

Used by:

* Configuration management (Ansible, Puppet)
* Authentication systems
* Logging and monitoring

---

## 2. Virtual Hardware Configuration

Defines what *hardware* the guest believes it is running on.

### 2.1 CPU Configuration

* Number of vCPUs
* CPU topology:

  * Sockets
  * Cores
  * Threads
* CPU features / flags (SSE, AVX, virtualization extensions)

Impacts:

* Performance
* NUMA awareness
* Software compatibility

---

### 2.2 Memory Configuration

* Assigned RAM
* Memory ballooning
* Huge pages
* NUMA node binding

Guest-visible effects:

* Available memory
* Memory latency
* OOM behavior

---

## 3. Storage Configuration

### 3.1 Virtual Disks

Per-guest disk settings include:

* Disk image path
* Disk format (qcow2, raw, vmdk)
* Bus type (IDE, SATA, SCSI, VirtIO)
* Disk UUID / serial number

Why it matters:

* Boot order
* Persistent device naming
* Performance tuning

---

### 3.2 Boot Configuration

* Firmware type:

  * BIOS
  * UEFI
* Boot order
* Secure Boot state

---

## 4. Networking Configuration

This is where **port mappings** come into play.

### 4.1 Virtual NICs

Guest-specific network attributes:

* MAC address (must be unique)
* NIC model (e1000, virtio-net, vmxnet3)
* Network backend:

  * Bridge
  * NAT
  * Overlay network

MAC address importance:

* DHCP leases
* Network identity
* Licensing systems

---

### 4.2 Port Mappings (NAT / Forwarding)

Port mappings expose guest services to external networks.

Example (host → guest):

```scss
Host TCP 2222 → Guest TCP 22
Host TCP 8080 → Guest TCP 80
```

Used in:

* NAT-based networking
* Containers-in-VM setups
* Development environments

Characteristics:

* Guest-specific
* Defined at hypervisor or virtual switch level
* Often stateful and protocol-specific

Limitations:

* Port collisions
* Performance overhead
* Not suitable for high-throughput workloads

---

## 5. Device Passthrough & Special Hardware

### 5.1 PCI / USB Passthrough

* Assigns physical devices directly to a guest
* Examples:

  * GPUs
  * NICs
  * USB devices

Guest-specific considerations:

* Device ownership is exclusive
* Migration may be disabled
* Strong coupling to host hardware

---

### 5.2 Virtual TPM (vTPM)

* Provides a Trusted Platform Mod
