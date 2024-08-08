package common

import "strings"

const (
	RdsRoot = "game"
)

func SplitRdsKey(key string) []string {
	return strings.Split(key, ":")
}

func GetRdsKey(keys ...string) string {
	return RdsRoot + ":" + strings.Join(keys, ":")
}
