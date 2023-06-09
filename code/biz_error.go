package code

import (
	"fmt"
	"github.com/paradeum-team/chainstorage-sdk/utils"
	"strings"
)

type bizError struct {
	messages map[string]string
	code     int
	err      error
}

func NewBizError(code int, cnMsg string, enMsg string) *bizError {
	return &bizError{
		messages: map[string]string{
			"__default__": enMsg,
			"zh-cn":       cnMsg,
			"en":          enMsg,
		},
		code: code,
	}
}

func (e *bizError) AddMessage(lang, msg string) {
	e.messages[lang] = msg
}

func (e *bizError) Code() int {
	return e.code
}

func (e *bizError) Message(mode string, lang string) string {
	msg, ok := e.messages[strings.ToLower(lang)]
	if !ok {
		msg = e.messages["__default__"]
	}
	if mode != utils.ModeProd.String() {
		msg = fmt.Sprintf("%s : %d", msg, e.code)
		if e.err != nil {
			msg += ": " + e.err.Error()
		}
	}
	return msg
}

func (e *bizError) Error() string {
	return e.Message(utils.ModeDev.String(), "__default__")
}

func (e bizError) Wrap(err error) *bizError {
	e.err = err
	return &e
}

func (e bizError) Wrapf(format string, a ...interface{}) *bizError {
	e.err = fmt.Errorf(format, a...)
	return &e
}

func (e bizError) Is(err error) bool {
	if x, ok := err.(*bizError); ok && x.code == e.code {
		return true
	}
	return false
}
