package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"middleware-experience/constants"

	"github.com/sirupsen/logrus"
)

func LogFmtTemp(v ...string) {
	fmt.Printf("\n >>>>>>>>>  %s  <<<<<<<<< \n", v)
}
func LogData(pkgName string, actName string, level int, data string) {
	formatDate := fmt.Sprintf("logs/system/log_%d_%s_%d.log", day, month, year)
	file, err := os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
	// item := logrus.Fields{
	// 	"action": actName,
	// 	"data":   data,
	// }
	// item := logrus.Fields{
	// 	"version":       "1.1",
	// 	"host":          "dina-proc-transfer",
	// 	"short_message": actName,
	// 	"full_message":  data,
	// 	"level":         1,
	// 	"_some_info":    pkgName,
	// }

	item := logrus.Fields{
		"full_message": data,
		"some_info":    pkgName,
	}

	fmt.Print("\n")
	itemPrint := logrus.Fields{
		"services": pkgName,
		"action":   actName,
	}
	//log.AddHook(LogHook)
	// test := fmt.Sprintf("echo '%s' | nc localhost 5555", item)
	// fmt.Print(test)
	// log.Info(test)
	if level == constants.LEVEL_LOG_INFO {
		log.WithFields(item).Infof(actName)
		logrus.WithFields(itemPrint).Info(data)
	}
	if level == constants.LEVEL_LOG_WARNING {
		log.WithFields(item).Warning(pkgName)
		logrus.WithFields(itemPrint).Warning(data)
	}
	if level == constants.LEVEL_LOG_ERROR {
		log.WithFields(item).Error(pkgName)
		logrus.WithFields(itemPrint).Error(data)
	}
	if level == constants.LEVEL_LOG_FATAL {
		log.WithFields(item).Fatal(pkgName)
		logrus.WithFields(itemPrint).Fatal(data)
	}
}
func LogDBerrInsert(data interface{}, info string) {
	formatDate := fmt.Sprintf("logs/db/log_%d_%s_%d.log", day, month, year)
	file, err := os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logDb.Out = file
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
	dataByte, _ := json.Marshal(data)
	item := logrus.Fields{
		"data": string(dataByte),
	}
	logDb.WithFields(item).Warning(info)
}
