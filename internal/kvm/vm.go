package kvm

import (
	//"golang/x/unix"
)

type VM struct {
	fd int
}

func (control *Controller) CreateVM() (*VM, error) {
    // err := unix.IoctlSetInt(control.fd, KVM_CREATE_VM, 0)
    fd, err := control.ioctl(KVM_CREATE_VM, 0)
    if err != nil {
	    return nil, err
    }

    vm := &VM{
	    fd: int(fd),
    }

    return vm, nil
}
