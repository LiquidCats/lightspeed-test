package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strings"

	"github.com/LiquidCats/lightspeed-test/pkg/utils"
)

// Constants for bitmap calculations.
const totalBits = 1 << 32 // Total possible IPv4 addresses.
const bitsPerUint64 = 64  // Number of bits in uint64.
const bitmapSize = totalBits / bitsPerUint64

func main() {
	// Check if the filename argument is provided.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Open the file for reading.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func() {
		_ = file.Close()
	}()

	// Initialize the bitmap as a slice of uint64.
	bitmap := make([]uint64, bitmapSize)

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNumber++
		if line == "" {
			continue // Skip empty lines.
		}

		// Convert the IP address string to an uint32.
		ipUint, err := utils.IpToUint32(line)
		if err != nil {
			log.Printf("Error parsing IP address on line %d: %v", lineNumber, err)
			continue // Skip invalid IP addresses.
		}

		// Calculate the index and bit position in the bitmap.
		index := ipUint / bitsPerUint64
		bit := ipUint % bitsPerUint64

		// Set the corresponding bit in the bitmap.
		bitmap[index] |= 1 << bit
	}

	// Check for any errors encountered by the scanner.
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Count the number of set bits in the bitmap.
	var uniqueCount uint64
	for _, word := range bitmap {
		uniqueCount += uint64(bits.OnesCount64(word))
	}

	// Output the total number of unique IP addresses.
	fmt.Printf("Number of unique IP addresses: %d\n", uniqueCount)
}
