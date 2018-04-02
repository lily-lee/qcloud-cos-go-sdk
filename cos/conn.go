package cos

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Conn ...
type Conn struct {
	config *Config
	client *http.Client
}

// init initialize Conn
func (conn *Conn) init(config *Config) error {
	// new Transport
	transport := newTransport(conn, config)

	conn.config = config
	conn.client = &http.Client{Transport: transport}

	return nil
}

func (conn Conn) Do(method, reqUrl string, headers, params map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(strings.ToUpper(method), reqUrl, body)
	if err != nil {
		log.Println("new request error:", err)
		return nil, err
	}

	if headers != nil {
		for k := range headers {
			req.Header.Set(k, headers[k])
		}
	}

	if body != nil {
		conn.handleBody(req, body)
	}

	am := &authMaker{
		method:      method,
		reqURL:      reqUrl,
		secretID:    conn.config.SecretID,
		secretKey:   conn.config.SecretKey,
		startTime:   time.Now(),
		authExpired: conn.config.AuthExpired,
		headers:     headers,
		params:      params,
	}

	auth, err := am.getAuth()
	if err != nil {
		log.Println("get auth error:", err)
		return nil, err
	}
	req.Header.Set(COS_AUTHORIZATION, auth)

	fmt.Println("request header:", req.Header)

	resp, err := conn.client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkResponse(resp)
}

// handle request body
func (conn Conn) handleBody(req *http.Request, body io.Reader) {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	req.Body = rc
	switch v := body.(type) {
	case *bytes.Buffer:
		req.ContentLength = int64(v.Len())
	case *bytes.Reader:
		req.ContentLength = int64(v.Len())
	case *strings.Reader:
		req.ContentLength = int64(v.Len())
	case *os.File:
		req.ContentLength = tryGetFileSize(v)
	}
	req.Header.Set(CONTENT_LENGTH, strconv.FormatInt(req.ContentLength, 10))

	//md5
	if req.Body != nil && conn.config.IsEnableMD5 {
		buf, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(buf))
		sum := md5.Sum(buf)
		b64 := base64.StdEncoding.EncodeToString(sum[:])
		req.Header.Set(CONTENT_MD5, b64)
	}
}

func tryGetFileSize(f *os.File) int64 {
	fInfo, _ := f.Stat()
	return fInfo.Size()
}
