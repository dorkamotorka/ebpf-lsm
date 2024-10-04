# eBPF LSM

eBPF (extended Berkeley Packet Filter) LSM (Linux Security Module) integrates the flexibility of eBPF with security policies, allowing developers to attach custom security measures to various kernel events dynamically. 

## Attaching to eBPF LSM Hook

TODO: prerequisite

Before we’re ready to write any LSM-BPF code, we need to find two things:
- Available LSM hook points
- Context provided to the hook point

Unfortunately, I don’t know of any tool that would print this information for us, other than looking directly into the [Linux source code](https://github.com/torvalds/linux/blob/master/security/security.c).

## How to Run

TODO
