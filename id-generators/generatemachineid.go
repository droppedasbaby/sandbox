package generators

import (
	"encoding/binary"
	"encoding/hex"
	"net"
)

// GenerateMachineID generates a numeric machine id based on the MAC address of the first network interface.
// The max value of the machine id determined by machineBits.
func GenerateMachineID(machineBits uint8) uint64 {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Failed to get network interfaces")
	}

	for _, inter := range interfaces {
		if inter.HardwareAddr != nil {
			mac, decodeErr := hex.DecodeString(inter.HardwareAddr.String())
			if decodeErr != nil {
				panic("Failed to convert MAC address to a number")
			}
			machineID := binary.BigEndian.Uint64(mac)
			machineID &= (1 << machineBits) - 1
			return machineID
		}
	}

	panic("No network interface with a MAC address found")
}
