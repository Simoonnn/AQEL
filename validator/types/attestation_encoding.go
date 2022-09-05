// Code generated by fastssz. DO NOT EDIT.
// Hash: 39b5c876c111dcae278b58cd6e461ed838b575ae5abb0936301a1fe36caf989d
// Version: 0.1.2-dev
package types

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/raidoNetwork/RDO_v2/proto/prototype"
)

// MarshalSSZ ssz marshals the Attestation object
func (a *Attestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the Attestation object to a target array
func (a *Attestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(32)

	// Field (0) 'Validator'
	if size := len(a.Validator); size != 20 {
		err = ssz.ErrBytesLengthFn("Attestation.Validator", size, 20)
		return
	}
	dst = append(dst, a.Validator...)

	// Offset (1) 'Block'
	dst = ssz.WriteOffset(dst, offset)
	if a.Block == nil {
		a.Block = new(prototype.Block)
	}
	offset += a.Block.SizeSSZ()

	// Offset (2) 'Signature'
	dst = ssz.WriteOffset(dst, offset)
	if a.Signature == nil {
		a.Signature = new(prototype.Sign)
	}
	offset += a.Signature.SizeSSZ()

	// Field (3) 'Type'
	dst = ssz.MarshalUint32(dst, uint32(a.Type))

	// Field (1) 'Block'
	if dst, err = a.Block.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature'
	if dst, err = a.Signature.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Attestation object
func (a *Attestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 32 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o2 uint64

	// Field (0) 'Validator'
	if cap(a.Validator) == 0 {
		a.Validator = make([]byte, 0, len(buf[0:20]))
	}
	a.Validator = append(a.Validator, buf[0:20]...)

	// Offset (1) 'Block'
	if o1 = ssz.ReadOffset(buf[20:24]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 32 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (2) 'Signature'
	if o2 = ssz.ReadOffset(buf[24:28]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (3) 'Type'
	a.Type = AttestationType(ssz.UnmarshallUint32(buf[28:32]))

	// Field (1) 'Block'
	{
		buf = tail[o1:o2]
		if a.Block == nil {
			a.Block = new(prototype.Block)
		}
		if err = a.Block.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'Signature'
	{
		buf = tail[o2:]
		if a.Signature == nil {
			a.Signature = new(prototype.Sign)
		}
		if err = a.Signature.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Attestation object
func (a *Attestation) SizeSSZ() (size int) {
	size = 32

	// Field (1) 'Block'
	if a.Block == nil {
		a.Block = new(prototype.Block)
	}
	size += a.Block.SizeSSZ()

	// Field (2) 'Signature'
	if a.Signature == nil {
		a.Signature = new(prototype.Sign)
	}
	size += a.Signature.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Attestation object
func (a *Attestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the Attestation object with a hasher
func (a *Attestation) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Validator'
	if size := len(a.Validator); size != 20 {
		err = ssz.ErrBytesLengthFn("Attestation.Validator", size, 20)
		return
	}
	hh.PutBytes(a.Validator)

	// Field (1) 'Block'
	if err = a.Block.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	if err = a.Signature.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'Type'
	hh.PutUint32(uint32(a.Type))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Attestation object
func (a *Attestation) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}