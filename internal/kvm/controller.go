package kvm

import (
	"os"
	"syscall"
)

type Controller struct {
	fd uintptr
}

func New() (*Controller, error) {
	f, err := os.OpenFile("/dev/kvm", os.O_RDWR, 0o666)
	if err != nil {
		return nil, err
	}

	controller := &Controller{
		fd: f.Fd(),
	}

	return controller, nil
}

func (self *Controller) GetVCPUMmapSize() (uintptr, error) {
	// result, err := self.ioctl(self.fd, KVM_CREATE_VCPU, uintptr(self.id))
	// result, err := ioctl(uintptr(self.fd), KVM_GET_VCPU_MMAP_SIZE, uintptr(self.ID))
	result, err := ioctl(uintptr(self.fd), KVM_GET_VCPU_MMAP_SIZE, uintptr(0))
	if err != nil {
		return uintptr(0), err
	}

	_ = result
	return result, nil
}

func (control *Controller) ioctl(command uint, value uintptr) (uintptr, error) {
	result, _, erno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(control.fd), uintptr(command), value)

	if erno != 0 {
		return result, erno
	}

	return result, nil
}
