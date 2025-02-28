package util

import (
	"os"
	"strconv"
	"strings"
)

func GetEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	strVal := GetEnv(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsInt64(key string, defaultVal int64) int64 {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseInt(strVal, 10, 64); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsUint64(key string, defaultVal uint64) uint64 {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseUint(strVal, 10, 64); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsUint32(key string, defaultVal uint32) uint32 {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseUint(strVal, 10, 32); err == nil {
		return uint32(val)
	}

	return defaultVal
}

func GetEnvAsUint8(key string, defaultVal uint8) uint8 {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseUint(strVal, 10, 8); err == nil {
		return uint8(val)
	}

	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseBool(strVal); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsStringArr reads ENV and returns the values split by separator.
func GetEnvAsStringArr(key string, defaultVal []string, separator ...string) []string {
	strVal := GetEnv(key, "")

	if len(strVal) == 0 {
		return defaultVal
	}

	sep := ","
	if len(separator) >= 1 {
		sep = separator[0]
	}

	return strings.Split(strVal, sep)
}

// GetEnvAsStringArrTrimmed reads ENV and returns the whitespace trimmed values split by separator.
func GetEnvAsStringArrTrimmed(key string, defaultVal []string, separator ...string) []string {
	slc := GetEnvAsStringArr(key, defaultVal, separator...)

	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}

	return slc
}

func GetEnvAsFloat64(key string, defaultVal float64) float64 {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseFloat(strVal, 64); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsRune(key string, defaultVal rune) rune {
	strVal := GetEnv(key, "")

	if strVal != "" {
		// return first rune
		for _, r := range strVal {
			return r
		}
	}

	return defaultVal
}
