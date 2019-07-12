package hasher

import (
	"testing"
)

func Test_ShouldHashAByteArrayAndReturnTheStringEncodingOfTheHashedValue(t *testing.T) {
	byteArray := []byte("adjust")
	hasher := NewHash()

	expected := "3b7770f7743e8f01f0fd807f304a21d0"
	result := hasher.Hash(byteArray)

	if result != expected {
		t.Fail()
	}
}
