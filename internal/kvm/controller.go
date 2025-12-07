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

func (control *Controller) ioctl(command uint, value uintptr) (uintptr, error) {
	result, _, erno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(control.fd), uintptr(command), value)

	if erno != 0 {
		return result, erno
	}

	return result, nil
}
