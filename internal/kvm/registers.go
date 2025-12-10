package kvm

import (
	"unsafe"
)

type Registers struct {
	RAX uint64
	RBX uint64
	RCX uint64
	RDX uint64
	RSI uint64
	RDI uint64
	RSP uint64
	RBP uint64
	R08 uint64
	R09 uint64
	R10 uint64
	R11 uint64
	R12 uint64
	R13 uint64
	R14 uint64
	R15 uint64
	RIP uint64
	RFLAGS uint64
}

type Segments struct {
	CS Segment
	DS Segment
	ES Segment
	FS Segment
	GS Segment
	SS Segment
	TR Segment
	LDT Descriptor // why is the other uy using Segment here?
	GDT Descriptor
	IDT Descriptor
	CR0 uint64
	CR1 uint64
	CR2 uint64
	CR3 uint64
	CR4 uint64
	CR5 uint64
	CR6 uint64
	CR7 uint64
	CR8 uint64
	EFER uint64
	// APICBase uint64
	// IntBitmap [(numInterrupts + 63) / 64]uint64
}

type Segment struct {
	Base     uint64
	Limit    uint32
	Selector uint16
	Typ      uint8
	Present  uint8
	DPL      uint8
	DB       uint8
	S        uint8
	L        uint8
	G        uint8
	AVL      uint8
	Unusable uint8
	_        uint8
}

type Descriptor struct {
	Base uint64
	Limit uint64
	_  [3]uint16
}

func (self *VCPU) SetRegisters(registers *Registers) (error) {
	// regs := uintptr(unsafe.Pointer(&registers))
	regs := uintptr(unsafe.Pointer(registers))
	// ioctl(self.fd, uintptr(KVM_SET_REGS), uintptr(unsafe.Pointer(&registers)
	_/*value*/, err := ioctl(uintptr(self.fd), KVM_SET_REGS, regs)
	return err
}

func (self *VCPU) GetRegisters() (*Registers, error) {
	registers := &Registers{}

	// regs := uintptr(unsafe.Pointer(&registers))
	regs := uintptr(unsafe.Pointer(registers))
	// ioctl(self.fd, uintptr(KVM_SET_REGS), uintptr(unsafe.Pointer(&registers)
	_/*value*/, err := ioctl(uintptr(self.fd), KVM_GET_REGS, regs)
	if err != nil {
		return nil, err
	}
	return registers, nil
}

func (self *VCPU) SetSegments(segments *Segments) (error) {
	// regs := uintptr(unsafe.Pointer(&registers))
	// ioctl(self.fd, uintptr(KVM_SET_REGS), uintptr(unsafe.Pointer(&registers)
	return nil
}

func (self *VCPU) GetSegments() (*Segments, error) {
	// regs := uintptr(unsafe.Pointer(&registers))
	// ioctl(self.fd, uintptr(KVM_SET_REGS), uintptr(unsafe.Pointer(&registers)
	return nil, nil
}
