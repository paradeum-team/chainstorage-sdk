package bfsutils

import (
	"time"
)

var local, _ = time.LoadLocation("Local")
var layoutTime = "2006-01-02 15:04:05"
var sysTimeFmt4compact  ="20060102150405"

/**
 * 获取当前 14 位字符长度的日期
 */
func GetCurrentDate()(dateLen14 string){
	currentDate :=time.Now().Format(sysTimeFmt4compact)[:14]
	return currentDate
}

/* 获取当前 8 位字符长度的日期
*/
func GetCurrentDate8()(dateLen8 string){
	currentDate :=time.Now().Format("20060102")[:8]
	return currentDate
}

func CurrentTime()string{
	timeStr:=time.Now().Format(layoutTime)  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
//	log.Println(timeStr)    //打印结果：2017-04-11 13:24:04
	return timeStr
}

//字符串转时间戳
func StrTimeConvertToTimeStamp(strTime string) int64 {
	formatTime,_:=time.ParseInLocation(sysTimeFmt4compact,strTime,local)
	return formatTime.Unix()
}


func ParseTime(fmtTime string) string {
	stamp, err := time.ParseInLocation(sysTimeFmt4compact, fmtTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	timeStr := stamp.String()[:19]
	if err == nil{
		return timeStr
	}else {
		return fmtTime
	}
}