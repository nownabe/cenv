/*
	cenv provides functions to get configs
	from environment variables.
*/
package cenv

import (
	"os"
	"strconv"
	"strings"
)

func Bool(keys ...string) (bool, error) {
	return strconv.ParseBool(get(keys))
}

func Float32(keys ...string) (float32, error) {
	v, err := strconv.ParseFloat(get(keys), 32)
	return float32(v), err
}

func Float64(keys ...string) (float64, error) {
	return strconv.ParseFloat(get(keys), 64)
}

func Int(keys ...string) (int, error) {
	return strconv.Atoi(get(keys))
}

func Int32(keys ...string) (int32, error) {
	v, err := strconv.ParseInt(get(keys), 10, 32)
	return int32(v), err
}

func Int64(keys ...string) (int64, error) {
	return strconv.ParseInt(get(keys), 10, 64)
}

func String(keys ...string) string {
	return get(keys)
}

func Uint(keys ...string) (uint, error) {
	v, err := strconv.ParseUint(get(keys), 10, 0)
	return uint(v), err
}

func Uint32(keys ...string) (uint32, error) {
	v, err := strconv.ParseUint(get(keys), 10, 32)
	return uint32(v), err
}

func Uint64(keys ...string) (uint64, error) {
	return strconv.ParseUint(get(keys), 10, 64)
}

func MustBool(keys ...string) bool {
	v, err := strconv.ParseBool(must(keys))
	chkErr(keys, err)
	return v
}

func MustFloat32(keys ...string) float32 {
	v, err := strconv.ParseFloat(must(keys), 32)
	chkErr(keys, err)
	return float32(v)
}

func MustFloat64(keys ...string) float64 {
	v, err := strconv.ParseFloat(must(keys), 64)
	chkErr(keys, err)
	return v
}

func MustInt(keys ...string) int {
	v, err := strconv.Atoi(must(keys))
	chkErr(keys, err)
	return v
}

func MustInt32(keys ...string) int32 {
	v, err := strconv.ParseInt(must(keys), 10, 32)
	chkErr(keys, err)
	return int32(v)
}

func MustInt64(keys ...string) int64 {
	v, err := strconv.ParseInt(must(keys), 10, 64)
	chkErr(keys, err)
	return v
}

func MustString(keys ...string) string {
	return must(keys)
}

func MustUint(keys ...string) uint {
	v, err := strconv.ParseUint(must(keys), 10, 0)
	chkErr(keys, err)
	return uint(v)
}

func MustUint32(keys ...string) uint32 {
	v, err := strconv.ParseUint(must(keys), 10, 32)
	chkErr(keys, err)
	return uint32(v)
}

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
