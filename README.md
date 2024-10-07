# eBPF LSM

eBPF (extended Berkeley Packet Filter) LSM (Linux Security Module) integrates the flexibility of eBPF with security policies, allowing developers to attach custom security measures to various kernel events dynamically. 

## Attaching to eBPF LSM Hook

**NOTE**: Minimum 5.7 kernel version is required.

Before we’re ready to write any LSM-BPF code, we need to find two things:
- Available LSM hook points
- Context provided to the hook point

Unfortunately, I don’t know of any tool that would print this information for us, other than looking directly into the [Linux source code](https://github.com/torvalds/linux/blob/master/security/security.c).

## How to Run

First build and run eBPF program:

```
go generate
go build
sudo ./lsm
```

By default eBPF program will block execution of the `ls` binary for the user with ID `1001`. 

If you don't have one such, you can create it using:

```
sudo useradd restricted-user --uid 1001
```

And then try to run the command `sudo -u restricted-user -- ls` to see the execution is in fact blocked.
