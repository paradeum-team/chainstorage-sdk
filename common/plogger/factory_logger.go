package plogger

import (
	"chainstorage-sdk/common/dict"
	"chainstorage-sdk/common/utils"
	"chainstorage-sdk/conf"
	"fmt"
	"github.com/kataras/golog"
	"github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

type PldLogger struct {
	logger *golog.Logger

	currentDate string //当前时间
}

var pldLoggerInstance *PldLogger

func NewInstance() *PldLogger {
	if pldLoggerInstance != nil {
		return pldLoggerInstance
	}
	currentDate := bfsutils.GetCurrentDate8() //当前的8位长度的日期
	pldLoggerInstance = &PldLogger{
		logger:      golog.Default,
		currentDate: currentDate,
	}
	pldLoggerInstance.logger.SetTimeFormat("2006-01-02 15:04:05")

	if conf.Logger.IsOutPutFile == false {
		return pldLoggerInstance
	}
	logInfoPath := CreateGinSysLogPath("pn")
	/*file, err := os.OpenFile(logInfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logInfoPath, err))
		return nil
	}*/
	//设置output
	logWriter := LogSplite(logInfoPath)
	pldLoggerInstance.logger.SetOutput(logWriter)
	//pldLoggerInstance.logger.AddOutput(logWriter)

	return pldLoggerInstance
}

func (lf *PldLogger) GetLogger() *golog.Logger {
	if pldLoggerInstance == nil {
		pldLoggerInstance = NewInstance()
		return pldLoggerInstance.logger

	} /*else {
		if lf.currentDate == bfsutils.GetCurrentDate8() {
			//同一天，说明日志不用切换文件，否则就新打开一个文件
		} else {
			NewInstance()
		}
	}*/
	return pldLoggerInstance.logger
}

/*
*
根据时间检测目录，不存在则创建
*/
func CreateDateDir(folderPath string) string {

	/*folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)*/
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

/*
*
创建系统日志的名字
*/
func CreateGinSysLogPath(filePrix string) string {
	baseLogPath := filepath.Join(conf.Config.DataPath, dict.LOG_FOLDER)
	writePath := CreateDateDir(baseLogPath) //根据时间检测是否存在目录，不存在创建
	//fileName := path.Join(writePath, filePrix + "_" + bfsutils.GetCurrentDate8() + ".log")
	fileName := path.Join(writePath, filePrix)
	return fileName
}

/*
*
使用io.WriteString()函数进行数据的写入，不存在则创建
*/
func WriteWithIo(filePath, content string) {
	fileObj, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	io.WriteString(fileObj, content)
}

/*
*
日志分割
*/
func LogSplite(logInfoPath string) io.Writer {
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logInfoPath+"_%Y%m%d.log",
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(logInfoPath),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(time.Duration(conf.Logger.MaxAgeDay*24)*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(conf.Logger.RotationTime*24)*time.Hour),
		//大小为这么多过期
		//rotatelogs.WithRotationCount(30),
	)
	return logWriter
}
