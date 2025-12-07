package kvm

type VCPU struct {
	fd uintptr
	vm *VM
}

func (self *VCPU) SetRegisters(registers Registers) (error) {
	// regs := uintptr(unsafe.Pointer(&registers))
	// ioctl(self.fd, uintptr(KVM_SET_REGS), uintptr(unsafe.Pointer(&registers)
	return nil
}
