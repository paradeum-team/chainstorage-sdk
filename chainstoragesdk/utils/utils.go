package utils

import (
	"chainstoragesdk/consts"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var local, _ = time.LoadLocation("Local")
var layoutTime = "2006-01-02 15:04:05"
var sysTimeFmt4compact = "20060102150405"

const ledgerLogPrefix = "./llog_un_"
const downloadAnalyticsLogPrefix = "./download_summary_"

/**
 * 获取当前 14 位字符长度的日期
 */
func GetCurrentDate() (dateLen14 string) {
	currentDate := time.Now().Format(sysTimeFmt4compact)[:14]
	return currentDate
}

/* 获取当前 8 位字符长度的日期
 */
func GetCurrentDate8() (dateLen8 string) {
	currentDate := time.Now().Format("20060102")[:8]
	return currentDate
}

func CurrentTime() string {
	timeStr := time.Now().Format(layoutTime) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	//	log.Println(timeStr)    //打印结果：2017-04-11 13:24:04
	return timeStr
}

// 字符串转时间戳
func StrTimeConvertToTimeStamp(strTime string) int64 {
	formatTime, _ := time.ParseInLocation(sysTimeFmt4compact, strTime, local)
	return formatTime.Unix()
}

func ParseTime(fmtTime string) string {
	stamp, err := time.ParseInLocation(sysTimeFmt4compact, fmtTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	timeStr := stamp.String()[:19]
	if err == nil {
		return timeStr
	} else {
		return fmtTime
	}
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
func TrimQuotes(s string) string {
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")
	return s
}

// randStringRunes returns a random string of length n from an array of characters.
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function.
	// Always returns a valid content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// keyForValue returns the integer key in a map if the given value v exists in the map, or -1 otherwise.
func KeyForValue(m map[int]string, v string) int {
	for k, x := range m {
		if x == v {
			return k
		}
	}
	return -1
}

func LogCriticalError(op, msg string) {
	op = strings.Trim(op, "\n")
	msg = strings.Trim(msg, "\n")
	log.Printf("CRITICAL-ERROR: %s, while attempting %s\n", msg, op)
}

// devGetLLogPath returns a filename for the log file that changes hourly.
func DevGetLLogPath() string {
	return ledgerLogPrefix + time.Now().Format("20060102150405")[:8] + ".log"
}

// devGetDnAnaLogPath returns a filename for the analytics log file that changes hourly.
func DevGetDnAnaLogPath() string {
	return downloadAnalyticsLogPrefix + time.Now().Format("20060102150405")[:8] + ".log"
}

// devLogAnalytics write log in file at logFilePath in specific format for later analysis.
// start is the start time of the operation.
// end is the end time of the operation.
// op is the operation identifier.
func DevLogAnalytics(id int, logFilePath string, start, end time.Time, op, afid string, msg ...string) {
	if logFilePath != "" {
		f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			//DevLog2(id, "devAnalytics", fmt.Sprintf("%s append|create failed:%v", logFilePath, err))
			return
		}
		defer f.Close()

		elapsed := ""
		if !end.IsZero() && !start.IsZero() {
			elapsed = fmt.Sprintf("%f", float64(end.Sub(start))/float64(time.Millisecond))
		}

		msgStr := ""
		for _, m := range msg {
			msgStr = msgStr + " " + m
		}
		l := log.New(f, "", 0)
		l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " ")
		l.Println(fmt.Sprintf("%-8d,", id), fmt.Sprintf("%-8s,", op), fmt.Sprintf("%15s,", elapsed), afid+",", "...", msgStr)
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

/**
 * 获取本地的ip <br/>
 * return [] string
 * 可能有双网卡，或者多个ip 。数组最后一个是默认的ip：127.0.0.1
 */
func GetLocalHostIps() (ipHosts []string) {
	address, err := net.InterfaceAddrs()
	localhost := "127.0.0.1"
	ips := make([]string, 0)
	if err != nil {
		log.Printf("Error: get lochost ip is wrong ..")
		ips = append(ips, localhost)
		return ips
	}

	for _, addr := range address {
		IpDr := addr.String()
		match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, IpDr)
		if !match {
			continue
		}
		ip := strings.Split(IpDr, "/")[0]
		//fmt.Printf("ip[%d]=%s \n",idx,ip)
		if localhost != ip {
			ips = append(ips, ip)
		}
	}
	ips = append(ips, localhost)

	return ips
}

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"    //开发模式
	ModeTest Mode = "test"   //测试模式
	ModeProd Mode = "prod"   //生产模式
	Mysql         = "mysql"  //mysql数据库标识
	Sqlite        = "sqlite" //sqlite
)

func (e Mode) String() string {
	return string(e)
}

func CurrentDate() string {
	timeStr := time.Now().Format(consts.SysTimeFmt4Date) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr
}
