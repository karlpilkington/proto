package protodecode

import (
	"testing"
)

func TestIPv6(t *testing.T) {
	b := []byte{
		96, 0, 0, 0, 0, 40, 6, 64, 38, 32, 1, 0,
		80, 7, 0, 6, 0, 0, 0, 0, 0, 1, 0, 3, 38,
		32, 1, 0, 80, 7, 0, 2, 0, 0, 0, 0, 0, 0,
		0, 2, 217, 103, 0, 80, 145, 114, 114, 15,
		0, 0, 0, 0, 160, 2, 22, 128, 91, 254, 0,
		0, 2, 4, 5, 160, 4, 2, 8, 10, 184, 68, 81,
		187, 0, 0, 0, 0, 1, 3, 3, 7,
	}

	packet := DecodeIPv6(b)

	if packet.DestinationAddress.String() != "2620:100:5007:2::2" {
		t.Errorf("expected destination address %v, got %v",
			"2620:100:5007:2::2", packet.DestinationAddress)
	}
}
