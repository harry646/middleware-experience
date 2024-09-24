package utils

import (
	"middleware-experience/constants"
	"strconv"

	"github.com/spf13/viper"
)

func EnvString(key string, def string) string {
	ctrlName := PkgName + "EnvString"
	value, ok := viper.Get(key).(string)
	if !ok {
		LogData(ctrlName, "viper.Get(key).(string)", constants.LEVEL_LOG_WARNING, "Invalid type assertion - from = "+key)
		return def
	}
	return value
}
func EnvBool(key string, def bool) bool {
	ctrlName := PkgName + "EnvString"
	// value, ok := viper.Get(key).(bool)
	value2, ok := viper.Get(key).(string)
	if !ok {
		LogData(ctrlName, "viper.Get(key).(bool)", constants.LEVEL_LOG_WARNING, "Invalid type assertion - from = "+key)
		return def
	}
	value, _ := strconv.ParseBool(value2)
	return value
}
func EnvInt(key string, def int) int {
	ctrlName := PkgName + "EnvInt"
	//value, ok := viper.Get(key).(int)
	value2, ok := viper.Get(key).(string)
	if !ok {
		LogData(ctrlName, "viper.Get(key).(int)", constants.LEVEL_LOG_WARNING, "Invalid type assertion - from = "+key)
		return def
	}
	value, _ := strconv.Atoi(value2)
	return value
}
func EnvInterface(key string, def interface{}) interface{} {
	value := viper.Get(key)
	return value
}
