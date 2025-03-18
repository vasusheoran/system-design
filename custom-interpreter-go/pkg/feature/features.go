package feature

import (
	"os"
	"strconv"
)

const (
	EnableTracingKey = "ENABLE_TRACING"
)

var (
	EnableTracing = getBoolVal(EnableTracingKey, false)
)

func getBoolVal(name string, defaultValue bool) bool {
	val, ok := os.LookupEnv(name)

	if !ok {
		return defaultValue
	}

	booleanVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return booleanVal
}
