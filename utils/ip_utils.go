package utils

import (
	"regexp"
	"strings"
)


/**
 * 正则验证：ip
 * 1.0.0.0~255.255.255.255
 */
func CheckIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

/**
 *
A类：10.0.0.0-10.255.255.255
B类：172.16.0.0-172.31.255.255
C类：192.168.0.0-192.168.255.255
--
10.x.x.x
172.16-31.x.x
192.168.x.x
 */
func CheckInnerIp(ip string)bool{
	addr := strings.Trim(ip, " ")

	regStrA := `^((10)\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	regStrB := `^((172)\.)((1[6-9]|2[0-9]|3[0-1])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	regStrC := `^((192)\.)((168)\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`

	if "127.0.0.1"==addr{
		return true
	}

	if match, _ := regexp.MatchString(regStrA, addr); match {
		return true
	}

	if match, _ := regexp.MatchString(regStrB, addr); match {
		return true
	}
	if match, _ := regexp.MatchString(regStrC, addr); match {
		return true
	}

	return false
}


/**
 * 验证pn 的url 地址是否正确
 *
 * http://192.168.1.129:5145/
 */
func CheckPNUrl(url string ) bool {
	addr :=strings.Trim(url," ")
	addr=strings.ToLower(addr)
	regStr := `^(https?://)(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(:5145/)$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}
