package utils

import (
	"fmt"
	"math/rand"
)

// RandomIPAddress generates a random valid IPv4 address.
func RandomIPAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256))
}
