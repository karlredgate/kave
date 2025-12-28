package kvm

import (
    "fmt"
    "syscall"
    "unsafe"
)

type UserspaceMemoryRegion struct {
    Slot uint32
    Flags uint32
    GuestPhysAddr uint64
    MemorySize uint64
    UserspaceAddr uint64
}

func (self *VM) SetUserMemoryRegion(region *UserspaceMemoryRegion) (error) {
	arg := uintptr(unsafe.Pointer(region))

	_/*value*/, err := ioctl(uintptr(self.fd), KVM_SET_USER_MEMORY_REGION, arg)
	if err != nil {
		return err
	}

	return nil
}

func (self *VM) ConfigureMemory() {
	size := 1024*1024*1024
	addr, err := syscall.Mmap(-1, 0, size,
	    syscall.PROT_READ | syscall.PROT_WRITE,
	    syscall.MAP_SHARED | syscall.MAP_ANONYMOUS)

	if err != nil {
	    panic(err)
	}

	// is this really necessary?
	arg := uint64(uintptr(unsafe.Pointer(&addr[0])))
	d := &UserspaceMemoryRegion{
	    Slot: 0,
	    Flags: 0,
	    GuestPhysAddr: 0,
	    MemorySize: uint64(size),
	    UserspaceAddr: arg,
	}
	fmt.Printf("MEM 0x%08x\n", &addr[0] )

	err = self.SetUserMemoryRegion(d)
	if err != nil { panic(err) }
}
