package kvm

import (
	"syscall"
)

const (
    KVMIO = 0xAE00
    KVM_GET_API_VERSION = KVMIO + 0x00
    KVM_CREATE_VM       = KVMIO + 0x01
)

func ioctl( fd uintptr, operation uintptr, arg uintptr ) (uintptr,error) {
    result,_,errno := syscall.Syscall( syscall.SYS_IOCTL, fd, operation, arg )
    if errno != 0 {
	return result,errno
    }
    return result,nil
}

func GetAPIVersion( fd uintptr ) (uintptr,error) {
	v, err := ioctl( fd, uintptr(KVM_GET_API_VERSION), uintptr(0) )
	if err != nil {
	     return uintptr(0), err
	}
	return v, nil
}
