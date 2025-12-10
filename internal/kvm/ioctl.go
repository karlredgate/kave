package kvm

import (
	"syscall"
)

func ioctl(fd uintptr, command uint, value uintptr) (uintptr, error) {
	result, _, erno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(command), value)

	if erno != 0 {
		return result, erno
	}

	return result, nil
}
