package main

import (
	"fmt"
	"math/rand"
)

func GenerateRandomOTP() string {
	otp := rand.Intn(10000)         // generates a number in the range [0, 9999]
	return fmt.Sprintf("%04d", otp) // formats the number to ensure it is 4 digits
}
