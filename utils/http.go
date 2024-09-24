package utils

import (
	"crypto/tls"
	"middleware-experience/constants"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/parnurzeal/gorequest"
)

// HTTPGet func
func HTTPGet(url string, header http.Header, bodyReq interface{}, retry bool) ([]byte, error) {
	if !NilInterface(bodyReq) {
		v, _ := query.Values(bodyReq)
		url = url + "?" + v.Encode()
	}
	request := gorequest.New()
	request.SetDebug(EnvBool("MainSetup.HttpClient.Debug", true))
	// timeout, _ := time.ParseDuration(EnvString("MainSetup.HttpClient.TimeOut", "60s"))
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	countRetry := 0
	if retry {
		countRetry = EnvInt("MainSetup.HttpClient.RetryBad", 0)
	}
	reqagent := request.Get(url)

	headers := make(map[string]string)
	for key, values := range header {
		headers[key] = strings.Join(values, ", ")
	}

	reqagent.Header = headers
	_, body, errs := reqagent.
		// Clone().Timeout(timeout).
		Retry(countRetry, time.Duration(EnvInt("MainSetup.HttpClient.RetryBadAttemp", 0))*time.Second, http.StatusInternalServerError, http.StatusRequestTimeout).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(EnvBool("MainSetup.HttpClient.Debug", true))
	timeout, _ := time.ParseDuration(EnvString("MainSetup.HttpClient.TimeOut", "60s"))
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	reqagent := request.Post(url)
	// reqagent.Header.Set("Content-Type", "application/json")
	reqagent.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(EnvInt("MainSetup.HttpClient.RetryBad", 0), time.Duration(EnvInt("MainSetup.HttpClient.RetryBadAttemp", 0))*time.Second, http.StatusInternalServerError, http.StatusRequestTimeout).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, header http.Header, retry bool) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(EnvBool("MainSetup.HttpClient.Debug", true))
	timeout, _ := time.ParseDuration(EnvString("MainSetup.HttpClient.TimeOut", "60s"))
	// _ := errors.New("Connection Problem")
	if strings.HasPrefix(url, "https") {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url).
		Send(jsondata).
		Timeout(timeout)
	if retry {
		reqagent = reqagent.Retry(EnvInt("MainSetup.HttpClient.RetryBad", 0), time.Duration(EnvInt("MainSetup.HttpClient.RetryBadAttemp", 0))*time.Second, http.StatusInternalServerError, http.StatusRequestTimeout)
	}

	headers := make(map[string]string)
	for key, values := range header {
		headers[key] = strings.Join(values, ", ")
	}

	reqagent.Header = headers
	_, body, errs := reqagent.End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPutWithHeader func
func HTTPPutWithHeader(url string, jsondata interface{}, header http.Header, retry bool) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(EnvBool("MainSetup.HttpClient.Debug", true))
	timeout, _ := time.ParseDuration(EnvString("MainSetup.HttpClient.TimeOut", "60s"))
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Put(url)
	countRetry := 0
	if retry {
		countRetry = EnvInt("MainSetup.HttpClient.RetryBad", 0)
	}

	headers := make(map[string]string)
	for key, values := range header {
		headers[key] = strings.Join(values, ", ")
	}

	reqagent.Header = headers
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(countRetry, time.Duration(EnvInt("MainSetup.HttpClient.RetryBadAttemp", 0))*time.Second, http.StatusInternalServerError, http.StatusRequestTimeout).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPDeleteWithHeader func
func HTTPDeleteWithHeader(url string, jsondata interface{}, header http.Header, retry bool) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(EnvBool("MainSetup.HttpClient.Debug", true))
	timeout, _ := time.ParseDuration(EnvString("MainSetup.HttpClient.TimeOut", "60s"))
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Delete(url)
	countRetry := 0
	if retry {
		countRetry = EnvInt("MainSetup.HttpClient.RetryBad", 0)
	}

	headers := make(map[string]string)
	for key, values := range header {
		headers[key] = strings.Join(values, ", ")
	}

	reqagent.Header = headers
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(countRetry, time.Duration(EnvInt("MainSetup.HttpClient.RetryBadAttemp", 0))*time.Second, http.StatusInternalServerError, http.StatusRequestTimeout).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// SendHttpRequest ..
func SendHttpRequest(method string, url string, header http.Header, body interface{}, retry bool) ([]byte, error) {
	var data []byte
	var err error
	switch method {

	case constants.HttpMethodGet:
		data, err = HTTPGet(url, header, body, retry)
		break
	case constants.HttpMethodPost:
		data, err = HTTPPostWithHeader(url, body, header, retry)
		break
	case constants.HttpMethodPut:
		data, err = HTTPPutWithHeader(url, body, header, retry)
		break
	case constants.HttpMethodDelete:
		data, err = HTTPDeleteWithHeader(url, body, header, retry)
		break
	}
	return data, err
}

func NilInterface(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
