package client

import (
	"bytes"
	"github.com/go-resty/resty"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func RestyGet(url string) (httpStatus int,body []byte,err error) {

	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().Get(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout,nil, err
		}else{
			return http.StatusServiceUnavailable,nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(),body, nil
}

func RestyPut(url string,data interface{} ) (httpStatus int,body []byte,err error) {
	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().SetBody(data).Put(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout,nil, err
		}else{
			return http.StatusServiceUnavailable,nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(),body, nil
}

func RestyPost(url string,data interface{} ) (httpStatus int,body []byte, err error) {
	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().SetBody(data).Post(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout,nil, err
		}else{
			return http.StatusServiceUnavailable,nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(),body, nil
}

func RestyPostForm(filename string ,filePath string,formData map[string]string,url string  )(httpStatus int,body []byte,err error){
	fileb, _ := ioutil.ReadFile(filePath)
	resp, err := resty.SetTimeout(time.Duration(60)*time.Second).R().SetFileReader("file", filename, bytes.NewReader(fileb)).
		SetFormData(formData).
		Post(url)

	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout,nil, err
		}else{
			return http.StatusServiceUnavailable,nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(),body ,err

}