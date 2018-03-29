package cos

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RespError struct {
	Code         string `xml:"Code"`
	Message      string `xml:"Message"`
	StringToSign string `xml:"StringToSign"`
	Resource     string `xml:"Resource"`
	RequestID    string `xml:"RequestId"`
	TraceID      string `xml:"TraceId"`
}

func checkResponse(resp *http.Response) (*http.Response, error) {
	var (
		err     error
		respErr RespError
	)
	if resp.StatusCode < 300 {
		return resp, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = xml.Unmarshal(body, &respErr); err != nil {
		return nil, err
	}
	err = errors.New(fmt.Sprintf("qcloud cos error: Code=[%s], Message=[%s], RequestID=[%s], TraceID=[%s]", respErr.Code, respErr.Message, respErr.RequestID, respErr.TraceID))

	return nil, err
}
