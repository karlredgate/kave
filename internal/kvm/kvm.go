package kvm

import (
	"syscall"
)

const (
    KVMIO = 0xAE00
    KVM_GET_API_VERSION     = KVMIO + 0x00
    KVM_CREATE_VM           = KVMIO + 0x01
    // KVM_GET_MSR_INDEX_LIST  = KVMIO + 0x02
    KVM_CHECK_EXTENSION     = KVMIO + 0x03
    KVM_GET_VCPU_MMAP_SIZE  = KVMIO + 0x04
    // KVM_GET_SUPPORTED_CPUID = KVMIO + 0x05
    KVM_TRACE_ENABLE        = KVMIO + 0x06 // deprecated
    KVM_TRACE_PAUSEE        = KVMIO + 0x07 // deprecated
    KVM_TRACE_DISABLE       = KVMIO + 0x08 // deprecated
    // KVM_GET_EMULATED_CPUID  = KVMIO + 0x09
    // KVM_GET_MSR_FEATURE_INDEX_LIST = KVMIO + 0x0A
    KVM_GET_MSR_FEATURE_INDEX_LIST = KVMIO + 0x0A

    KVM_CREATE_VCPU                = KVMIO + 0x41

    // need better way
    KVM_GET_REGS = 0x8090ae81
    KVM_SET_REGS = 0x4138AE84
)
//      NNN 41
//      MMM 38
//    KVMIO AE
// SET_REGS 84

func OLDioctl( fd uintptr, operation uintptr, arg uintptr ) (uintptr,error) {
    result,_,errno := syscall.Syscall( syscall.SYS_IOCTL, fd, operation, arg )
    if errno != 0 {
	return result,errno
    }
    return result,nil
}

func GetAPIVersion( fd uintptr ) (uintptr,error) {
	v, err := ioctl( fd, KVM_GET_API_VERSION, uintptr(0) )
	if err != nil {
	     return uintptr(0), err
	}
	return v, nil
}
