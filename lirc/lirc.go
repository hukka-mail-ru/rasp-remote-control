package lirc

import (
	"encoding/binary"
	"fmt"
)

func ConvertToArrayOfUint32(bytes []byte) []uint32 {

	var array []uint32

	for i := 0; i < len(bytes)-4; i += 4 {

		res := binary.LittleEndian.Uint32(bytes[i : i+4])

		array = append(array, res)
	}

	return array
}

func PrintPulseSpace(values []uint32) {

	const PULSE_MASK = 0x01000000
	const TIMEOUT_MASK = 0x03000000
	const VALUE_MASK = 0x00FFFFFF

	for i := 0; i < len(values); i++ {

		fmt.Printf("Uint32: %08x ", values[i])

		pulse := (values[i] & PULSE_MASK) > 0
		timeout := (values[i] & TIMEOUT_MASK) > 0
		val := values[i] & VALUE_MASK

		if pulse {

			fmt.Printf("pulse %d", val)
			if val > 8000 {

			}

		} else if timeout {

			fmt.Printf("timeout\n")

		} else {

			fmt.Printf("space %d\n", val)
		}
	}

}
