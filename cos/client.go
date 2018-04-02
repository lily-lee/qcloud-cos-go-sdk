package cos

import (
	"encoding/xml"
	"log"
	"time"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	// Client ...
	Client struct {
		Config *Config
		Conn   *Conn
	}
	// ClientOption ...
	ClientOption func(*Client)
)

// NewClient 初始化一个Client
func NewClient(appID, secretID, secretKey, scheme string, authExpired int64, options ...ClientOption) (*Client, error) {
	// configuration
	config := getDefaultConfig()
	config.APPID = appID
	config.SecretID = secretID
	config.SecretKey = secretKey
	config.Scheme = scheme
	config.AuthExpired = authExpired

	// http connect
	conn := &Conn{config: config}

	// oss client
	client := &Client{config, conn}

	// client options parse
	for _, option := range options {
		option(client)
	}

	// create http connect
	err := conn.init(config)

	log.Println(client)

	return client, err
}

func (client Client) GetAuth(param AuthParam) (string, error) {
	conf := client.Config
	am := authMaker{
		method:      param.Method,
		reqURL:      makeRequestUrl(conf.Scheme, param.Bucket, conf.APPID, param.Region, param.URI),
		secretID:    conf.SecretID,
		secretKey:   conf.SecretKey,
		startTime:   time.Now(),
		authExpired: param.AuthExpired,
		headers:     param.Headers,
		params:      param.Params,
	}

	return am.getAuth()
}

// NewBucket 初始化一个Bucket
func (client Client) NewBucket(bucketName, region string) (*Bucket, error) {
	return &Bucket{client, bucketName, region}, nil
}

// GetService 列出该账户下所有 Client
func (client Client) GetService() (ListAllMyBucketsResult, error) {
	var result ListAllMyBucketsResult

	reqUrl := client.makeReqUrl("", "", "/")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// DeleteBucket 删除某个Bucket，
// 删除之前要求 Bucket 内的内容为空，
// 只有删除了 Bucket 内的信息，才能删除 Bucket 本身
// URI: /
func (client Client) DeleteBucket(bucketName, region string) error {
	reqUrl := client.makeReqUrl(bucketName, region, "/")
	resp, err := client.do("DELETE", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteBucketCORS 删除跨域访问配置信息
// URI: /?cors
func (client Client) DeleteBucketCORS(bucketName, region string) error {
	reqUrl := client.makeReqUrl(bucketName, region, "/?cors")
	resp, err := client.do("DELETE", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteBucketLifecycle 删除生命周期管理的配置
// Delete Bucket Lifecycle 用来删除 Bucket 的生命周期配置。
// 如果该 Bucket 没有配置生命周期规则会返回 NoSuchLifecycleConfiguration。
// URI: /?lifecycle
func (client Client) DeleteBucketLifecycle(bucketName, region string) error {
	reqUrl := client.makeReqUrl(bucketName, region, "/?lifecycle")
	resp, err := client.do("DELETE", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetBucket 列出指定 Client 下的部分或者全部 Bucket
// Get Bucket 请求等同于 List Object 请求，
// 可以列出该 Bucket 下的部分或者全部 Object。
// 此 API 调用者需要对 Bucket 有 Read 权限。
// 1. 每次默认返回的最大条目数为 1000 条，
// 如果无法一次返回所有的 list，则返回结果中的 IsTruncated 为 true，
// 同时会附加一个 NextMarker 字段，提示下一个条目的起点。
// 若一次请求，已经返回了整个 list，则不会有 NextMarker 这个字段，
// 同时 IsTruncated 为 false。
// 2. 若把 prefix 设置为某个文件夹的全路径名，
// 则可以列出以此 prefix 为开头的文件，即该文件夹下递归的所有文件和子文件夹。
// 如果再设置 delimiter 定界符为 “/”，则只列出该文件夹下的文件，
// 子文件夹下递归的文件和文件夹名将不被列出。
// 而子文件夹名将会以 CommonPrefix 的形式给出。
// URI: /
func (client Client) GetBucket(bucketName, region string) (ListBucketResult, error) {
	var result ListBucketResult

	reqUrl := client.makeReqUrl(bucketName, region, "/")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// GetBucketACL 获取 Client 的 ACL 表
// Get Bucket ACL 接口用来获取 Bucket 的 ACL(access control list)，
// 即用户空间（Bucket）的访问权限控制列表。 此 API 接口只有 Bucket 的持有者有权限操作。
// URI: /?acl
func (client Client) GetBucketACL(bucketName, region string) (AccessControlPolicy, error) {
	var result AccessControlPolicy

	reqUrl := client.makeReqUrl(bucketName, region, "/?acl")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// GetBucketCORS 获取 Client 的跨域访问配置信息
// Get Bucket CORS 接口实现 Bucket 持有者在 Bucket 上进行跨域资源共享的信息配置。
// （CORS 是一个 W3C 标准，全称是"跨域资源共享"（Cross-origin resource sharing））。
// 默认情况下，Bucket 的持有者直接有权限使用该 API 接口，Bucket 持有者也可以将权限授予其他用户。
// URI: /?cors
func (client Client) GetBucketCORS(bucketName, region string) (CORSConfiguration, error) {
	var result CORSConfiguration

	reqUrl := client.makeReqUrl(bucketName, region, "/?cors")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// GetBucketLocation 获取 Client 所在的地域信息
// 只有 Bucket 持有者才有该 API 接口的操作权限。
// URI: /?location
func (client Client) GetBucketLocation(bucketName, region string) (LocationConstraint, error) {
	var result LocationConstraint

	reqUrl := client.makeReqUrl(bucketName, region, "/?location")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// GetBucketLifecycle 读取生命周期管理的配置
// 如果该 Bucket 没有配置生命周期规则会返回 NoSuchLifecycleConfiguration。
// URI: /?lifecycle
func (client Client) GetBucketLifecycle(bucketName, region string) (LifecycleConfiguration, error) {
	var result LifecycleConfiguration

	reqUrl := client.makeReqUrl(bucketName, region, "/?lifecycle")
	resp, err := client.do("GET", reqUrl, nil, nil, nil)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	fmt.Println(string(body))
	printPretty(result)

	return result, err

}

// HeadBucket 确认指定账号下是否存在指定 Client
// Head 的权限与 Read 一致。
// 当该 Bucket 存在时，返回 HTTP 状态码 200；
// 当该 Bucket 无访问权限时，返回 HTTP 状态码 403；
// 当该 Bucket 不存在时，返回 HTTP 状态码 404
// URI: /
func (client Client) HeadBucket(bucketName, region string) error {
	reqUrl := client.makeReqUrl(bucketName, region, "/")
	resp, err := client.do("HEAD", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ListMultipartUploads 查询正在进行中的分块上传信息
// 单次请求操作最多列出 1000 个正在进行中的分块上传
//
// Params: delimiter,encoding-type,prefix,max-uploads,key-marker,upload-id-marker
//
// delimiter:	定界符为一个符号，对 Object 名字包含指定前缀且第一次出现 delimiter
// 				字符之间的 Object 作为一组元素：common prefix。
// 				如果没有 prefix，则从路径起点开始
// 				String	否
//
// encoding-type:	规定返回值的编码格式，合法值：url
// 					String	否
//
// prefix:	限定返回的 Object key 必须以 Prefix 作为前缀。
// 			注意使用 prefix 查询时，返回的 key 中仍会包含 Prefix
// 			String	否
//
// max-uploads:	设置最大返回的 multipart 数量，合法取值从1到1000，默认1000
// 				String	否
//
// key-marker:	与 upload-id-marker 一起使用
// 				当 upload-id-marker 未被指定时，ObjectName 字母顺序大于 key-marker 的条目将被列出
//				当upload-id-marker被指定时，ObjectName 字母顺序大于key-marker的条目被列出，
// 				ObjectName 字母顺序等于 key-marker 同时 UploadID 大于 upload-id-marker 的条目将被列出。
// 				String	否
//
// upload-id-marker:	与 key-marker 一起使用
// 						当 key-marker 未被指定时，upload-id-marker 将被忽略
//						当 key-marker 被指定时，ObjectName字母顺序大于 key-marker 的条目被列出，
// 						ObjectName 字母顺序等于 key-marker 同时 UploadID 大于 upload-id-marker 的条目将被列出。
// 						String	否
//
// URI: /?uploads
// TODO 问题，params用map传好还是struct好？？？
func (client Client) ListMultipartUploads(bucketName, region string, params *ListMultipartUploadsRequestParam) (ListMultipartUploadsResult, error) {
	var result ListMultipartUploadsResult

	p := structToMap(params)
	p = formatMapByKeys(Uploads_Param_Keys, p)
	paramStr := mapToStr(p, "&")

	reqUrl := client.makeReqUrl(bucketName, region, "/?uploads")
	if paramStr != "" {
		reqUrl = fmt.Sprintf("%s&%s", reqUrl, paramStr)
	}

	resp, err := client.do("GET", reqUrl, nil, p, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(body, &result)
	printPretty(result)

	return result, err
}

// PutBucket 在指定账号下创建一个 Client
// Put Bucket 接口请求可以在指定账号下创建一个 Bucket。
// 该 API 接口不支持匿名请求，您需要使用帯 Authorization 签名认证的请求才能创建新的 Bucket 。
// 创建 Bucket  的用户默认成为 Bucket 的持有者。
// 创建 Bucket 时，如果没有指定访问权限，则默认使用私有读写（private）权限。
//
// headers: x-cos-acl, x-cos-grant-read, x-cos-grant-write, x-cos-grant-full-control
//
// x-cos-acl:	定义 Object 的 ACL 属性。
// 			  	有效值：private，public-read-write，public-read；默认值：private。
// 			  	String 否

// x-cos-grant-read:	赋予被授权者读的权限。格式：x-cos-grant-read: id=" ",id=" "；
// 						当需要给子账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>"，
// 						当需要给根账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>"
// 						String 否

// x-cos-grant-write:	赋予被授权者写的权限。格式：x-cos-grant-write: id=" ",id=" "；
// 						当需要给子账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>"，
// 						当需要给根账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>"
// 						String	否

// x-cos-grant-full-control:	赋予被授权者读写权限。格式：x-cos-grant-full-control: id=" ",id=" "；
// 								当需要给子账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>"，
//								当需要给根账户授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>"
// 								String	否
func (client Client) PutBucket(bucketName, region string, headers map[string]string) error {
	headers = formatMapByKeys(Acl_Header_Keys, headers)
	reqUrl := client.makeReqUrl(bucketName, region, "/")
	resp, err := client.do("PUT", reqUrl, headers, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// PutBucketACL 写入 Client 的 ACL 表
// 通过 Header："x-cos-acl"，"x-cos-grant-read"，"x-cos-grant-write"，"x-cos-grant-full-control" 传入 ACL 信息
// 这是一个覆盖操作，传入新的 ACL 将覆盖原有 ACL。
// URI: /?acl
func (client Client) PutBucketACL(bucketName, region string, headers map[string]string) error {
	headers = formatMapByKeys(Acl_Header_Keys, headers)
	reqUrl := client.makeReqUrl(bucketName, region, "/?acl")
	resp, err := client.do("PUT", reqUrl, headers, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// PutBucketCORS 配置 Client 的跨域访问权限
// 通过传入 XML 格式的配置文件来实现配置，文件大小限制为64 KB。
// 默认情况下，Bucket 的持有者直接有权限使用该 API 接口，Bucket 持有者也可以将权限授予其他用户。
// TODO PutBucketCORS
func (client Client) PutBucketCORS() {

}

// PutBucketLifecycle 设置生命周期管理的功能
// COS 支持用户以生命周期配置的方式来管理 Bucket 中 Object 的生命周期。
// 生命周期配置包含一个或多个将应用于一组对象规则的规则集 (其中每个规则为 COS 定义一个操作)。
// 这些操作分为以下两种：
// 转换操作：
// 		定义对象转换为另一个存储类的时间。
// 		例如，您可以选择在对象创建 30 天后将其转换为低频存储（STANDARD_IA，适用于不常访问) 存储类别。
// 		同时也支持将数据沉降到归档存储（Archive，成本更低，目前支持国内园区）。具体参数参见请求示例说明中Transition项。
//
// 过期操作：指定 Object 的过期时间。COS 将会自动为用户删除过期的 Object。
//
// Put Bucket Lifecycle 用于为 Bucket 创建一个新的生命周期配置。
// 如果该 Bucket 已配置生命周期，使用该接口创建新的配置的同时则会覆盖原有的配置。
// URI: /?lifecycle
// TODO PutBucketLifecycle
func (client Client) PutBucketLifecycle() {

}

func (client Client) do(method, reqUrl string, headers, params map[string]string, body io.Reader) (*http.Response, error) {
	return client.Conn.Do(method, reqUrl, headers, params, body)
}

func (client Client) makeReqUrl(bucketName, region, uri string) string {
	return makeRequestUrl(client.Config.Scheme, bucketName, client.Config.APPID, region, uri)
}
