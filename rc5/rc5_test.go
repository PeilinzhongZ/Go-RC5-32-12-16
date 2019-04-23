package rc5

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"
)

func TestRC5_Random(t *testing.T) {
	random := rand.New(rand.NewSource(99))
	max := 5000

	var encrypted [8]byte
	var decrypted [8]byte

	for i := 0; i < max; i++ {
		key := make([]byte, 16)
		random.Read(key)
		value := make([]byte, 8)
		random.Read(value)

		cipher, ok := RC5_SETUP(key)
		if ok {
			cipher.RC5_ENCRYPT(value[:], encrypted[:])
			cipher.RC5_DECRYPT(encrypted[:], decrypted[:])

			if !bytes.Equal(decrypted[:], value[:]) {
				t.Errorf("encryption/decryption failed: % 02x != % 02x\n", decrypted, value)
			}
		}
	}
}

func TestRC5_TESTVECTORS(t *testing.T) {
	var encrypted [8]byte
	var decrypted [8]byte

	key, _ := hex.DecodeString("5269F149D41BA0152497574D7F153125")
	pt, _ := hex.DecodeString("65C178B284D197CC")
	ct, _ := hex.DecodeString("EB44E415DA319824")

	cipher, ok := RC5_SETUP(key)
	if ok {
		cipher.RC5_ENCRYPT(pt[:], encrypted[:])
		if !bytes.Equal(ct[:], encrypted[:]) {
			t.Errorf("encryption/decryption failed: % 02x != % 02x\n", encrypted, ct)
		}

		cipher.RC5_DECRYPT(encrypted[:], decrypted[:])
		if !bytes.Equal(decrypted[:], pt[:]) {
			t.Errorf("encryption/decryption failed: % 02x != % 02x\n", decrypted, pt)
		}
	}
}
