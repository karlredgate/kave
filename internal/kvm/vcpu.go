package kvm

type VCPU struct {
	fd uintptr
	vm *VM
}
