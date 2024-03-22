package tools

import (
	"math/rand"
	"strconv"
	"time"
)

type GenerateOtp struct{}

func (c *GenerateOtp) GenerateOTP() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Define the range for 6-digit random integers
	min := 100000
	max := 999999

	// Generate a random integer in the specified range
	randomInt := rand.Intn(max-min+1) + min

	return strconv.Itoa(randomInt)
}
