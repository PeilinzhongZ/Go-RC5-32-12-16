package rc5

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestCipher32(t *testing.T) {
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
