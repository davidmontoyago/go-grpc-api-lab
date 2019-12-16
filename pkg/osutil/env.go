package osutil

import "os"

func GetenvOrDefault(key string, def string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return def
	}
	return val
}
