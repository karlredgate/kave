package kvm

type Translation struct {
	LinearAddress uint64
	PhysicalAddress uint64
	Valid uint8
	Writable uint8
	UserMode uint8
	_  [5]uint8 // pad
}
