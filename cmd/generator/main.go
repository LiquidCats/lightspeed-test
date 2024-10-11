package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/LiquidCats/lightspeed-test/pkg/utils"
)

func main() {
	filename := "test/assets/ips.txt"
	totalLines := 10000
	duplicateRatio := 0.3

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	uniqueCount := int(float64(totalLines) * (1 - duplicateRatio))
	duplicatesCount := totalLines - uniqueCount

	// Generate and write unique IPs directly to file.
	ipSet := make(map[string]struct{})
	uniqueIPs := make([]string, 0, uniqueCount)
	for len(uniqueIPs) < uniqueCount {
		ip := utils.RandomIPAddress()
		if _, exists := ipSet[ip]; !exists {
			ipSet[ip] = struct{}{}
			uniqueIPs = append(uniqueIPs, ip)
			fmt.Fprintln(writer, ip)
		}
	}

	// Generate and write duplicates.
	for i := 0; i < duplicatesCount; i++ {
		idx := rand.Intn(len(uniqueIPs))
		fmt.Fprintln(writer, uniqueIPs[idx])
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Printf("Generated %d IP addresses in %s\n", totalLines, filename)
	}
}
