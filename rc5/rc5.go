package rc5

import (
	"encoding/binary"
	"math/bits"
)

const (
	w   = 32
	r   = 12
	b   = 16
	t   = 2 * (r + 1)
	P_w = 0xb7e15163
	Q_w = 0x9e3779b9
)

type cipher32 struct {
	S [t]uint32
}

func RC5_SETUP(key []byte) {
	if len(key) != b {
		return
	}

	S := make([]uint32, t)
	S[0] = P_w
	for i := uint(1); i < t; i++ {
		S[i] = S[i-1] + Q_w
	}

	var L [w / 8]uint32
	for i := 0; i < w/8; i++ {
		L[i] = binary.LittleEndian.Uint32(key[:w/8])
		key = key[w/8:]
	}

	var A uint32
	var B uint32
	var i, j int
	for k := 0; k < 3*t; k++ {
		S[i] = bits.RotateLeft32(S[i]+(A+B), 3)
		A = S[i]
		L[j] = bits.RotateLeft32(L[j]+(A+B), int(A+B))
		B = L[j]
		i = (i + 1) % t
		j = (j + 1) % (w / 8)
	}
}

func (C *cipher32) RC5_ENCRYPT(pt, ct []byte) {
	A := binary.LittleEndian.Uint32(pt[:w/8]) + C.S[0]
	B := binary.LittleEndian.Uint32(pt[w/8:]) + C.S[1]
	for i := 0; i < r; i++ {
		A = bits.RotateLeft32(A^B, int(B)) + C.S[2*i]
		B = bits.RotateLeft32(B^A, int(A)) + C.S[2*i+1]
	}
	binary.LittleEndian.PutUint32(ct[:w/8], A)
	binary.LittleEndian.PutUint32(ct[w/8:], B)
}

func (C *cipher32) RC5_DECRYPT(ct, pt []byte) {
	A := binary.LittleEndian.Uint32(ct[:w/8])
	B := binary.LittleEndian.Uint32(ct[w/8:])
	for i := r; i > 0; i-- {
		B = bits.RotateLeft32(B-C.S[2*i+1], -int(A)) ^ A
		A = bits.RotateLeft32(A-C.S[2*i], -int(B)) ^ B
	}
	binary.LittleEndian.PutUint32(pt[w/8:], B-C.S[1])
	binary.LittleEndian.PutUint32(pt[:w/8], A-C.S[0])
}
