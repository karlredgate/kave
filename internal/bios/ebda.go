package bios

import (
    "bytes"
    "encoding/binary"
)

const (
    mpfSig = (('_' << 24)|('P' << 16)|('M' << 8)|('_'))
    mpcSig = (('P' << 24)|('M' << 16)|('C' << 8)|('P'))
)

type mpcCPU struct {
    typ uint8
    apicID uint8
    apicVer uint8
    cpuFlag uint8
    cpuFeature uint32
    featureFlag uint32
    reserved1 uint32
    reserved2 uint32
}

type IntelMPF struct {
    signature uint32
    physPtr uint32
    length uint8
    specification uint8
    checksum uint8
    feature1 uint8
    feature2 uint8
    feature3 uint8
    feature4 uint8
    feature5 uint8
}

type MPC struct {
    signature uint32
    length uint16
    spec uint8
    checkSum uint8
    oem [8]uint8
    product [12]uint8
    oemPtr uint32
    oemSize uint16
    oemCount uint16
    lapic uint32
    reserved uint32
    // cpu
    cpulist  [64]mpcCPU
}

type EBDA struct {
    _  [16 * 3]uint8
    mpf IntelMPF
    mpc MPC
}

func NewEBDA() (*EBDA, error) {
    ebda := &EBDA{
	mpf: IntelMPF{
	    signature: mpfSig,
	    length: 1,
	    specification: 4,
	    // physPtr: 
	},
	mpc: MPC{
	},
    }

    ebda.setChecksum()

    return ebda, nil
}

func (self *EBDA) setChecksum() {
    buf := new( bytes.Buffer )
    err := binary.Write( buf, binary.LittleEndian, self.mpc )
    if err != nil { panic(err) }
    self.mpc.checkSum = 0
}
