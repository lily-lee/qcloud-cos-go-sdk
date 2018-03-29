package cos

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Config ...
type Config struct {
	APPID       string
	SecretKey   string
	SecretID    string
	Region      string
	Scheme      string      // http or https
	AuthExpired int64       // Authorization 有效时间 second，默认600秒
	HTTPTimeout HTTPTimeout // HTTP的超时时间设置

	IsEnableMD5 bool
}

// HTTPTimeout http timeout
type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

func getDefaultConfig() *Config {
	config := Config{}
	config.Scheme = "https"
	config.AuthExpired = 600
	config.HTTPTimeout.ConnectTimeout = time.Second * 30   // 30s
	config.HTTPTimeout.ReadWriteTimeout = time.Second * 60 // 60s
	config.HTTPTimeout.HeaderTimeout = time.Second * 60    // 60s
	config.HTTPTimeout.LongTimeout = time.Second * 300     // 300s
	config.HTTPTimeout.IdleConnTimeout = time.Second * 50  // 50s

	return &config
}

func makeHost(bucketName, appID, region string) (host string) {
	if bucketName == "" {
		host = "service.cos.myqcloud.com"
	} else {
		if !strings.Contains(bucketName, appID) {
			bucketName = fmt.Sprintf("%s-%s", bucketName, appID)
		}
		host = fmt.Sprintf("%s.cos.%s.myqcloud.com", bucketName, region)
	}

	return host
}

func makeRequestUrl(scheme, bucketName, appID, region, uri string) (url string) {
	host := makeHost(bucketName, appID, region)
	url = fmt.Sprintf("%s://%s%s", scheme, host, uri)

	log.Println("request url:", url)

	return
}

func makePublicRequestHeader() {

}
