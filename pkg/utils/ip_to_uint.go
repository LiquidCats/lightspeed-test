package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// IpToUint32 converts an IPv4 address string to a uint32 representation.
func IpToUint32(ipStr string) (uint32, error) {
	// Split the IP address into its four octets.
	parts := strings.Split(ipStr, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid IP address: %s", ipStr)
	}

	var ip uint32
	for i := 0; i < 4; i++ {
		// Parse each octet to an integer.
		n, err := strconv.ParseUint(parts[i], 10, 8)
		if err != nil {
			return 0, fmt.Errorf("invalid IP address: %s", ipStr)
		}
		// Shift the previous bits by 8 and add the new octet.
		ip = ip<<8 + uint32(n)
	}
	return ip, nil
}
