package tokenclient

import (
	"bytes"
	"github.com/go-resty/resty"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	//token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlLZXkiOiJaT0dKNVI0QTZPU1JIWUlOTUlMQSIsImV4cCI6MTY4NDgzNDUxNSwiaWRlbnRpdHkiOnsiSWQiOjYsIk5pY2tOYW1lIjoidXNlcjEiLCJQcm9maWxlSW1hZ2UiOiIiLCJXYWxsZXRBZGRyIjoidXNlcjEiLCJ0ZXJtaW5hbCI6InBjL3NkayJ9LCJvcmlnX2lhdCI6MTY4MjI0MjUyNH0.w0lQ-RGufy5j3yMfbnL94hOvLeeCSccZbSIwd6kXvbI"
	token     = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODMyNzg2MTMsImlkZW50aXR5Ijp7IklkIjo2LCJOaWNrTmFtZSI6InVzZXIxIiwiUHJvZmlsZUltYWdlIjoiIiwiV2FsbGV0QWRkciI6InVzZXIxIiwidGVybWluYWwiOiJwYy9zZGsifSwib3JpZ19pYXQiOjE2ODI0MTQ2MTN9.UjhYK96QT_EiIZi2TeHih62CoWI33S1eiwU0YN0wxXo"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
	headers   = map[string]string{
		"Accept":        "application/json",
		"Authorization": token,
		"User-Agent":    userAgent,
	}
)

func RestyGet(url string) (httpStatus int, body []byte, err error) {
	resp, err := resty.
		SetTimeout(time.Duration(60) * time.Second).
		SetHeaders(headers).
		R().Get(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout, nil, err
		} else {
			return http.StatusServiceUnavailable, nil, err
		}
	}

	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(), body, nil
}

func RestyPut(url string, data interface{}) (httpStatus int, body []byte, err error) {
	resp, err := resty.
		SetTimeout(time.Duration(60) * time.Second).
		SetHeaders(headers).
		R().SetBody(data).Put(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout, nil, err
		} else {
			return http.StatusServiceUnavailable, nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(), body, nil
}

func RestyPost(url string, data interface{}) (httpStatus int, body []byte, err error) {
	resp, err := resty.
		SetTimeout(time.Duration(60) * time.Second).
		SetHeaders(headers).
		R().SetBody(data).Post(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout, nil, err
		} else {
			return http.StatusServiceUnavailable, nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()
	return resp.StatusCode(), body, nil
}

func RestyPostForm(filename string, filePath string, formData map[string]string, url string) (httpStatus int, body []byte, err error) {
	fileb, err := os.ReadFile(filePath)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	resp, err := resty.
		SetTimeout(time.Duration(60)*time.Second).
		SetHeaders(headers).
		R().SetFileReader("file", filename, bytes.NewReader(fileb)).
		SetFormData(formData).
		Post(url)

	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout, nil, err
		} else {
			return http.StatusServiceUnavailable, nil, err

		}
	}
	defer resp.RawResponse.Body.Close()
	body = resp.Body()

	return resp.StatusCode(), body, err
}

func RestyDelete(url string, data interface{}) (httpStatus int, body []byte, err error) {
	resp, err := resty.
		SetTimeout(time.Duration(60) * time.Second).
		SetHeaders(headers).
		R().
		SetBody(data).
		Delete(url)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			return http.StatusGatewayTimeout, nil, err
		} else {
			return http.StatusServiceUnavailable, nil, err

		}
	}

	defer resp.RawResponse.Body.Close()

	body = resp.Body()

	return resp.StatusCode(), body, nil
}
