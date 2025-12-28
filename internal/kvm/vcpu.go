package kvm

type VCPU struct {
	fd int
	vm *VM
	ID uint
	run *Run
}

func (vm *VM) CreateVCPU(id uint) (*VCPU, error) {
    fd, err := ioctl(uintptr(vm.fd), KVM_CREATE_VCPU, uintptr(id)) /// WRONG FD
    if err != nil {
	    return nil, err
    }

    vcpu := &VCPU{
	    fd: int(fd),
	    vm: vm,
	    ID: id,
    }

    return vcpu, nil
}

func (self *VCPU) Run() (error) {
	// result, err := self.ioctl(self.fd, KVM_CREATE_VCPU, uintptr(self.id))
	result, err := ioctl(uintptr(self.fd), KVM_RUN, uintptr(self.ID))
	if err != nil {
	    panic(err)
		return err
	}
	_ = result
	return nil
}
