package bfsutils

import (
	"chainstorage-sdk/common/dict"
	"errors"
	"fmt"
	"time"
)

/**
* yyyyMMddHHmmss 转换成时间对象
 */
func ParseTime14Len(fmtTime string) (time.Time, error) {
	if fmtTime == "" {
		return time.Now(), fmt.Errorf("time is invalid . Can not be nil ")
	}
	fmtTemplate := dict.SysTimeFmt4compact
	return ParseTime2TimeType(fmtTime, fmtTemplate)
}

//
//func ParseTime2TimeType(fmtTime string, timeFormatTemplate string) (time.Time, error) {
//	stamp, err := time.ParseInLocation(timeFormatTemplate, fmtTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
//	//stamp, err := time.ParseInLocation("20060102150405", "20190108135030", time.Local)
//	//log.Println(stamp.Unix())  //输出：1546926630
//
//	return stamp, err
//}

/**
 * 解析时间至时间类型
 */
func ParseTime2TimeType(fmtTime ...string) (time.Time, error) {
	if len(fmtTime) == 0 {
		return time.Time{}, errors.New("parameter is invalid.")
	}

	timeValue := fmtTime[0]
	formatTemplate := ""
	if len(fmtTime) == 1 || len(fmtTime[1]) == 0 {
		formatTemplate = dict.SysTimeFmt
	} else {
		formatTemplate = fmtTime[1]
	}

	stamp, err := time.ParseInLocation(formatTemplate, timeValue, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	//stamp, err := time.ParseInLocation("20060102150405", "20190108135030", time.Local)
	//log.Println(stamp.Unix())  //输出：1546926630

	return stamp, err
}

func FmtTime14Len(t time.Time) string {
	return FmtTime(t, dict.SysTimeFmt4compact)
}
func FmtTime(t time.Time, format string) string {
	return t.Format(format)
}

func CurrentTime4compact() string {
	//timeStr:=time.Now().Format("2006-01-02 15:04:05")  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	timeStr := time.Now().Format(dict.SysTimeFmt4compact) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	//log.Println(timeStr)    //打印结果：2017-04-11 13:24:04
	return timeStr
}

func CurrentDate() string {
	timeStr := time.Now().Format(dict.SysTimeFmt4Date) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr
}

func Yesterday() string {
	timeStr := time.Now().AddDate(0, 0, -1).Format(dict.SysTimeFmt4Date) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr
}

func GetDateFromNow(days int) string {
	timeStr := time.Now().AddDate(0, 0, days).Format(dict.SysTimeFmt4Date) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return timeStr
}

//func DevTimeTrack(start time.Time, name string) {
//	elapsed := time.Since(start)
//	plogger.NewInstance().GetLogger().Debugf("[DEV]TIMETRACK,command: %s took %s\n", name, elapsed)
//}

func AddExpiredTime(days int) string {
	//expireTimeStr := time.Now().Local().Add(time.Second * time.Duration((days-1)*24*3600)).Format(dict.SysTimeFmt4compact)
	expireTimeStr := time.Now().Local().AddDate(0, 0, days).Format(dict.SysTimeFmt4compact)
	expireTimeStr = expireTimeStr[:8] + "235959"
	return expireTimeStr
}

func TodayExpiredTime() string {
	expireTimeStr := time.Now().Format(dict.SysTimeFmt4compact)
	expireTimeStr = expireTimeStr[:8] + "235959"
	return expireTimeStr
}

func netDay() string {
	expireTimeStr := time.Now().Local().AddDate(0, 0, 1).Format(dict.SysTimeFmt4compact)
	expireTimeStr = expireTimeStr[:8] + "000001"
	return expireTimeStr
}

/**
 * 特条件下用
 */
func AfterToday(dateTime14Len string) bool {
	nextDay := netDay()
	n1, err1 := ParseTime14Len(nextDay)
	t, err2 := ParseTime14Len(dateTime14Len)
	if err2 == nil && err1 == nil {
		if t.After(n1) {
			return true
		}
	}
	return false
}
