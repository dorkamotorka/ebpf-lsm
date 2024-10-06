package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go lsm lsm.c

import (
    "log"
    "time"
    "github.com/cilium/ebpf/link"
    "github.com/cilium/ebpf/rlimit"
)

func main() {
    // Allow the current process to lock memory for eBPF resources.
    if err := rlimit.RemoveMemlock(); err != nil {
	log.Fatal(err)
    }

    // Load pre-compiled programs and maps into the kernel.
    var lsmObjs lsmObjects
    lsmObjs = lsmObjects{}
    if err := loadLsmObjects(&lsmObjs, nil); err != nil {
	log.Fatal(err)
    }

    // Attach LSM program.
    lsmLink, err := link.AttachLSM(link.LSMOptions{
	Program:   lsmObjs.PolicePerm,
    })
    if err != nil {
	log.Fatal("Attaching LSM:", err)
    }
    defer lsmLink.Close()

    time.Sleep(time.Second * 30)
}
