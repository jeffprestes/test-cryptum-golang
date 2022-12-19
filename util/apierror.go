package util

import (
	"log"
	"runtime"
)

type APIRequestError interface {
	Error()
	LogError()
	LogErrorWithData()
}

type APIRequestErrorData struct {
	Message       string
	RealMessage   string
	HttpErrorCode int
	Err           error
}

func (re *APIRequestErrorData) SetError(err error) {
	re.Err = err
	re.RealMessage = err.Error()
}

func (re *APIRequestErrorData) Error() string {
	return re.Message
}

func (re *APIRequestErrorData) LogErrorWithAditionalMsg(aditionalMsg string) {
	pc, _, _, _ := runtime.Caller(1)
	funcOrPackageName := runtime.FuncForPC(pc).Name()
	log.Println("Error at", "["+funcOrPackageName+"]", aditionalMsg, re.RealMessage)
}

func (re *APIRequestErrorData) LogError() {
	pc, _, _, _ := runtime.Caller(1)
	funcOrPackageName := runtime.FuncForPC(pc).Name()
	log.Println("Error at", "["+funcOrPackageName+"]", re.RealMessage, "|", re.Message)
}

func (re *APIRequestErrorData) LogErrorWithData(data any) {
	pc, _, _, _ := runtime.Caller(1)
	funcOrPackageName := runtime.FuncForPC(pc).Name()
	log.Printf("Error at [%s] %s | %s\n%#v\n", funcOrPackageName, re.RealMessage, re.Message, data)
}

func NewAPIPostError(err error) (restError APIRequestErrorData) {
	restError.Message = "Error connecting to server"
	restError.RealMessage = err.Error()
	restError.Err = err
	restError.HttpErrorCode = 500
	return restError
}

func NewAPIPostMarshellingError(err error) (restError APIRequestErrorData) {
	restError.Message = "Error connecting to server"
	restError.RealMessage = "Error marsheling data: " + err.Error()
	restError.Err = err
	restError.HttpErrorCode = 500
	return restError
}

func NewAPIGetError(err error) (restError APIRequestErrorData) {
	restError.Message = "Error connecting to server"
	restError.RealMessage = err.Error()
	restError.Err = err
	restError.HttpErrorCode = 500
	return restError
}

func NewAPIGetErrorEmpty() (restError APIRequestErrorData) {
	return restError
}

func NewAPIGetPartnerError(err error, partner string) (restError APIRequestErrorData) {
	restError.Message = "Error connecting to " + partner + " server"
	restError.RealMessage = err.Error()
	restError.Err = err
	restError.HttpErrorCode = 500
	return restError
}
