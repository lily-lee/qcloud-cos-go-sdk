package cos

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

type authMaker struct {
	method          string
	reqURL          string
	secretID        string
	secretKey       string
	startTime       time.Time
	authExpired     int64
	headers, params map[string]string
}

// GetAuth returns Authorization
func (am *authMaker) GetAuth() (string, error) {
	authKeys := []string{"q-sign-algorithm=sha1", "q-ak=%s", "q-sign-time=%s", "q-key-time=%s", "q-header-list=%s", "q-url-param-list=%s", "q-signature=%s"}

	signTime := getSignTime(am.startTime, am.authExpired)
	headerList := strings.ToLower(strings.Join(getParamKeys(am.headers), ";"))
	paramList := strings.ToLower(strings.Join(getParamKeys(am.params), ";"))
	headers := paramsToString(am.headers)
	parameters := paramsToString(am.params)

	URL, err := url.ParseRequestURI(am.reqURL)
	if err != nil {
		log.Println("get auth error:", err)
		return "", err
	}

	if err != nil {
		log.Println("getURLParams error:", err)
		return "", err
	}

	signature := getSignature(signTime, am.secretKey, am.method, URL.Path, parameters, headers)
	auth := fmt.Sprintf(strings.Join(authKeys, "&"), am.secretID, signTime, signTime, headerList, paramList, signature)
	log.Println("auth:", auth)

	return auth, err
}

// Sha1 ...
func Sha1(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSha1 ...
func HmacSha1(s, secretKey string) string {

	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(s))

	a := mac.Sum(nil)

	return fmt.Sprintf("%x", a)
}

//$SignKey = hash_hmac($SecretKey,"[q-key-time]")
//$HttpString = [HttpMethod]\n[HttpURI]\n[HttpParameters]\n[HttpHeaders]\n
//$StringToSign = [q-sign-algorithm]\n[q-sign-time]\nsha1($HttpString)\n
//$Signature = hash_hmac($SignKey,$StringToSign)
func getSignature(signTime, secretKey, httpMethod, httpURI, httpParameters, httpHeaders string) string {
	signKey := HmacSha1(signTime, secretKey)
	fmt.Println("signKey:", signKey)

	httpString := getHTTPString(httpMethod, httpURI, httpParameters, httpHeaders)
	fmt.Println("httpString:", httpString)

	sha1edHTTPString := Sha1(httpString)

	stringToSign := getStringToSign(signTime, sha1edHTTPString)
	fmt.Println("stringToSign:", stringToSign)

	signature := HmacSha1(stringToSign, signKey)
	log.Println(signature, sha1edHTTPString)

	return signature
}

func getSignTime(startTime time.Time, authExpired int64) string {
	endTime := startTime.Add(time.Duration(authExpired) * time.Second)

	return fmt.Sprintf("%d;%d", startTime.Unix(), endTime.Unix())
}

func getHTTPString(httpMethod, httpURI, httpParameters, httpHeaders string) string {

	return strings.Join([]string{strings.ToLower(httpMethod), httpURI, httpParameters, httpHeaders, ""}, "\n")
	//return fmt.Sprintf("%s\n%s\n%s\n%s\n", strings.ToLower(httpMethod), httpURI, httpParameters, httpHeaders)
}

func getStringToSign(signTime, sha1edHTTPString string) string {

	return strings.Join([]string{"sha1", signTime, sha1edHTTPString, ""}, "\n")
	//return fmt.Sprintf("sha1\n%s\n%s\n", signTime, sha1edHTTPString)
}

// key 转为小写，字典排序，value进行URLEncode
func paramsToString(params map[string]string) string {
	keys := getParamKeys(params)
	paramList := make([]string, len(keys))
	for i, k := range keys {
		val := params[k]
		paramList[i] = fmt.Sprintf("%s=%s", strings.ToLower(k), url.QueryEscape(val))
	}

	return strings.Join(paramList, "&")
}

func getParamKeys(params map[string]string) []string {
	keys := []string{}
	if params == nil {
		return keys
	}
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// make request str
func getReqStr(method, reqHost, reqPath string, params map[string]interface{}) string {
	str := []string{}
	for k, v := range params {
		s := fmt.Sprintf("%s=%v", k, v)
		str = append(str, s)
	}
	sort.Strings(str)

	reqStr := strings.Join(str, "&")
	return fmt.Sprintf("%s%s%s?%s", method, reqHost, reqPath, reqStr)
}

func sign(s, secretKey string) string {
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(s))
	s = base64.StdEncoding.EncodeToString([]byte(mac.Sum(nil)))
	fmt.Println(url.QueryEscape("0EEm/HtGRr/VJXTAD9tYMth1Bzm3lLHz5RCDv1GdM8s="))
	return url.QueryEscape(s)
}
