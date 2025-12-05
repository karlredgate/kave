package kvm

type Run struct {
	RequestInterruptWindow uint8
	ImmediateExit uint8
	Padding [6]uint8
	ExitReason uint32
	ReadyForInterruptInjection uint8
	IfFlag uint16 // ???
	Flags uint16
	CR8 uint64
	APICBase uint64
}

