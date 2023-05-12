package chainstoragesdk

import (
	"chainstoragesdk/utils"
	"fmt"
	"github.com/kataras/golog"
	"github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type PldLogger struct {
	logger      *golog.Logger
	config      *LoggerConf
	currentDate string //当前时间
}

var onceLogger sync.Once
var pldLoggerInstance *PldLogger

func newLogger(cfg *LoggerConf) *PldLogger {
	if pldLoggerInstance != nil {
		return pldLoggerInstance
	}

	onceLogger.Do(func() {
		currentDate := utils.GetCurrentDate8() //当前的8位长度的日期
		pldLoggerInstance = &PldLogger{
			logger:      golog.Default,
			config:      cfg,
			currentDate: currentDate,
		}
		pldLoggerInstance.logger.SetTimeFormat("2006-01-02 15:04:05")

		if cfg.IsOutPutFile == false {
			//return pldLoggerInstance
			return
		}

		logInfoPath := pldLoggerInstance.createGinSysLogPath("chainstorage")
		/*file, err := os.OpenFile(logInfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logInfoPath, err))
			return nil
		}*/

		//设置output
		logWriter := pldLoggerInstance.logSplite(logInfoPath)
		pldLoggerInstance.logger.SetOutput(logWriter)
		//pldLoggerInstance.logger.AddOutput(logWriter)
	})

	return pldLoggerInstance
}

//func NewInstance() *PldLogger {
//	if pldLoggerInstance != nil {
//		return pldLoggerInstance
//	}
//	currentDate := utils.GetCurrentDate8() //当前的8位长度的日期
//	pldLoggerInstance = &PldLogger{
//		logger:      golog.Default,
//		currentDate: currentDate,
//	}
//	pldLoggerInstance.logger.SetTimeFormat("2006-01-02 15:04:05")
//
//	if conf.Logger.IsOutPutFile == false {
//		return pldLoggerInstance
//	}
//
//	logInfoPath := createGinSysLogPath("pn")
//	/*file, err := os.OpenFile(logInfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
//	if err != nil {
//		log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logInfoPath, err))
//		return nil
//	}*/
//	//设置output
//	logWriter := logSplite(logInfoPath)
//	pldLoggerInstance.logger.SetOutput(logWriter)
//	//pldLoggerInstance.logger.AddOutput(logWriter)
//
//	return pldLoggerInstance
//}

func GetLogger(cfg *LoggerConf) *PldLogger {
	if pldLoggerInstance == nil {
		pldLoggerInstance = newLogger(cfg)
		//return pldLoggerInstance.logger
	} /*else {
		if lf.currentDate == bfsutils.GetCurrentDate8() {
			//同一天，说明日志不用切换文件，否则就新打开一个文件
		} else {
			NewInstance()
		}
	}*/
	return pldLoggerInstance
}

/*
*
创建系统日志的名字
*/
func (pl *PldLogger) createGinSysLogPath(filePrix string) string {
	baseLogPath := pl.config.LogPath
	writePath := utils.CreateDateDir(baseLogPath) //根据时间检测是否存在目录，不存在创建
	//fileName := path.Join(writePath, filePrix + "_" + bfsutils.GetCurrentDate8() + ".log")
	fileName := path.Join(writePath, filePrix)
	return fileName
}

/*
*
使用io.WriteString()函数进行数据的写入，不存在则创建
*/
func (pl *PldLogger) writeWithIo(filePath, content string) {
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
func (pl *PldLogger) logSplite(logInfoPath string) io.Writer {
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logInfoPath+"_%Y%m%d.log",
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(logInfoPath),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(time.Duration(pl.config.MaxAgeDay*24)*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(pl.config.RotationTime*24)*time.Hour),
		//大小为这么多过期
		//rotatelogs.WithRotationCount(30),
	)
	return logWriter
}
