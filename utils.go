package main

import (
	"os"
	"strconv"
)

func ParseEnvUint(env string, defaultNumber uint64) uint64 {
	data := os.Getenv(env)
	if len(data) == 0 {
		return defaultNumber
	}

	uintNumber, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return defaultNumber
	}

	return uintNumber
}
