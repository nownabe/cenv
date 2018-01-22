// Package cenv provides simple functions to get configs
// from environment variables.
//
// Key Conversions
//
// Functions provided by cenv get environment variables from
// a list of string keys.
// Keys are converted to uppercase and
// dots are converted to underscores.
// If a list of keys is given, they are joined by underscores.
//
// All following codes get the environment variable
// which is associated to APP_KEY as an integer value.
//
//	cenv.Int("app_key")
//	cenv.Int("app.key")
//	cenv.Int("app", "key")
package cenv

import (
	"os"
	"strconv"
	"strings"
)

// Bool returns the boolean value from environment variable.
// It accepts boolean string values from strconv.ParseBool.
// Any other value returns an error.
func Bool(keys ...string) (bool, error) {
	return strconv.ParseBool(get(keys))
}

// Float32 returns the float32 value from environment variable.
// When it couldn't get the value as float32, it returns an error.
// It is based on strconv.ParseFloat.
func Float32(keys ...string) (float32, error) {
	v, err := strconv.ParseFloat(get(keys), 32)
	return float32(v), err
}

// Float64 returns the float64 value from environment variable.
// When it couldn't get the value as float64, it returns an error.
// It is based on strconv.ParseFloat.
func Float64(keys ...string) (float64, error) {
	return strconv.ParseFloat(get(keys), 64)
}

// Int returns the int value from environment variable.
// When it couldn't get the value as int, it returns an error.
// It is based on strconv.Atoi.
func Int(keys ...string) (int, error) {
	return strconv.Atoi(get(keys))
}

// Int32 returns the int32 value from environment variable.
// When it couldn't get the value as int32, it returns an error.
// It is based on strconv.ParseInt.
func Int32(keys ...string) (int32, error) {
	v, err := strconv.ParseInt(get(keys), 10, 32)
	return int32(v), err
}

// Int64 returns the int64 value from environment variable.
// When it couldn't get the value as int64, it returns an error.
// It is based on strconv.ParseInt.
func Int64(keys ...string) (int64, error) {
	return strconv.ParseInt(get(keys), 10, 64)
}

// String returns the string value from environment variable.
func String(keys ...string) string {
	return get(keys)
}

// Uint returns the uint value from environment variable.
// When it couldn't get the value as uint, it returns an error.
// It is based on strconv.ParseUint.
func Uint(keys ...string) (uint, error) {
	v, err := strconv.ParseUint(get(keys), 10, 0)
	return uint(v), err
}

// Uint32 returns the uint32 value from environment variable.
// When it couldn't get the value as uint32, it returns an error.
// It is based on strconv.ParseUint.
func Uint32(keys ...string) (uint32, error) {
	v, err := strconv.ParseUint(get(keys), 10, 32)
	return uint32(v), err
}

// Uint64 returns the uint64 value from environment variable.
// When it couldn't get the value as uint64, it returns an error.
// It is based on strconv.ParseUint.
func Uint64(keys ...string) (uint64, error) {
	return strconv.ParseUint(get(keys), 10, 64)
}

// MustBool returns the boolean value from environment variable.
// It accepts boolean string values from strconv.ParseBool.
// If any other value is given or the variable is not present, it panics.
func MustBool(keys ...string) bool {
	v, err := strconv.ParseBool(must(keys))
	chkErr(keys, err)
	return v
}

// MustFloat32 returns the float32 value from environment variable.
// When it couldn't get the value as float32 or the variable is not present,
// it panics.
// It is based on strconv.ParseFloat.
func MustFloat32(keys ...string) float32 {
	v, err := strconv.ParseFloat(must(keys), 32)
	chkErr(keys, err)
	return float32(v)
}

// MustFloat64 returns the float64 value from environment variable.
// When it couldn't get the value as float64 or the variable is not present,
// it panics.
// It is based on strconv.ParseFloat.
func MustFloat64(keys ...string) float64 {
	v, err := strconv.ParseFloat(must(keys), 64)
	chkErr(keys, err)
	return v
}

// MustInt returns the int value from environment variable.
// When it couldn't get the value as int or the variable is not present,
// it panics.
// It is based on strconv.Atoi.
func MustInt(keys ...string) int {
	v, err := strconv.Atoi(must(keys))
	chkErr(keys, err)
	return v
}

// MustInt32 returns the int32 value from environment variable.
// When it couldn't get the value as int32 or the variable is not present,
// it panics.
// It is based on strconv.ParseInt.
func MustInt32(keys ...string) int32 {
	v, err := strconv.ParseInt(must(keys), 10, 32)
	chkErr(keys, err)
	return int32(v)
}

// MustInt64 returns the int64 value from environment variable.
// When it couldn't get the value as int64 or the variable is not present,
// it panics.
// It is based on strconv.ParseInt.
func MustInt64(keys ...string) int64 {
	v, err := strconv.ParseInt(must(keys), 10, 64)
	chkErr(keys, err)
	return v
}

// MustString returns the string value from environment variable.
// When the variable is not present, it panics.
func MustString(keys ...string) string {
	return must(keys)
}

// MustUint returns the uint value from environment variable.
// When it couldn't get the value as uint or the variable is not present,
// it panics.
// It is based on strconv.ParseUint.
func MustUint(keys ...string) uint {
	v, err := strconv.ParseUint(must(keys), 10, 0)
	chkErr(keys, err)
	return uint(v)
}

// MustUint32 returns the uint32 value from environment variable.
// When it couldn't get the value as uint32 or the variable is not present,
// it panics.
// It is based on strconv.ParseUint.
func MustUint32(keys ...string) uint32 {
	v, err := strconv.ParseUint(must(keys), 10, 32)
	chkErr(keys, err)
	return uint32(v)
}

// MustUint64 returns the uint64 value from environment variable.
// When it couldn't get the value as uint64 or the variable is not present,
// it panics.
// It is based on strconv.ParseUint.
func MustUint64(keys ...string) uint64 {
	v, err := strconv.ParseUint(must(keys), 10, 64)
	chkErr(keys, err)
	return v
}

func chkErr(keys []string, err error) {
	if err != nil {
		panic(convertKeys(keys) + " can't be got by the error: " + err.Error())
	}
}

func convertKeys(keys []string) string {
	key := strings.Join(keys, "_")
	key = strings.ToUpper(key)
	key = strings.Replace(key, ".", "_", -1)
	return key
}

func get(keys []string) string {
	return os.Getenv(convertKeys(keys))
}

func must(keys []string) string {
	k := convertKeys(keys)
	v, ok := os.LookupEnv(k)
	if !ok {
		panic(k + " must be set")
	}
	return v
}
