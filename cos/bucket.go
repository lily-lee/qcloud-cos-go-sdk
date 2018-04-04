package cos

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"strconv"

	"time"

	"crypto/md5"
	"encoding/base64"

	"github.com/pkg/errors"
)

// Bucket ...
type Bucket struct {
	Client     Client
	BucketName string
	Region     string
}

// TODO 把每个方法中传headers map[string]string 改成 Header struct ？？？

// AbortMultipartUpload 实现舍弃一个分块上传操作并删除已上传的块
// 当您调用 Abort Multipart Upload 时，如果有正在使用这个 Upload Parts 上传块的请求，
// 则 Upload Parts 会返回失败。当该 UploadId 不存在时，会返回 404 NoSuchUpload。
// 注意：
// 建议您及时完成分块上传或者舍弃分块上传，因为已上传但是未终止的块会占用存储空间进而产生存储费用。
// URI: /<ObjectName>
func (bucket Bucket) AbortMultipartUpload(imur InitiateMultipartUploadResult) error {
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?uploadId=%s", imur.Key, imur.UploadID))
	params := map[string]string{
		"uploadId": imur.UploadID,
	}
	resp, err := bucket.do("DELETE", reqUrl, nil, params, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// CompleteMultipartUpload 完成整个文件的分块上传
// Complete Multipart Upload 接口请求用来实现完成整个分块上传。
// 当使用 Upload Parts 上传完所有块以后，必须调用该 API 来完成整个文件的分块上传。
// 在使用该 API 时，您必须在请求 Body 中给出每一个块的 PartNumber 和 ETag，用来校验块的准确性。
// 由于分块上传完后需要合并，而合并需要数分钟时间，因而当合并分块开始的时候，
// COS 就立即返回 200 的状态码，在合并的过程中，COS 会周期性的返回空格信息来保持连接活跃，
// 直到合并完成，COS会在 Body 中返回合并后块的内容。
// 当上传块小于 1 MB 的时候，在调用该 API 时，会返回 400 EntityTooSmall；
// 当上传块编号不连续的时候，在调用该 API 时，会返回 400 InvalidPart；
// 当请求 Body 中的块信息没有按序号从小到大排列的时候，在调用该 API 时，会返回 400 InvalidPartOrder；
// 当 UploadId 不存在的时候，在调用该 API 时，会返回 404 NoSuchUpload。
//
// 注意：
// 建议您及时完成分块上传或者舍弃分块上传，因为已上传但是未终止的块会占用存储空间进而产生存储费用
// URI: /<ObjectName>
func (bucket Bucket) CompleteMultipartUpload(imur InitiateMultipartUploadResult, parts []Part) (CompleteMultipartUploadResult, error) {
	var result CompleteMultipartUploadResult

	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?uploadId=%s", imur.Key, imur.UploadID))
	params := map[string]string{
		"uploadId": imur.UploadID,
	}

	data, err := xml.Marshal(CompleteMultipartUpload{
		Part: parts,
	})
	if err != nil {
		return result, errors.New("param error")
	}

	buffer := new(bytes.Buffer)
	buffer.Write(data)
	resp, err := bucket.do("POST", reqUrl, nil, params, buffer)
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

// DeleteMultiObject 在 Client 中批量删除 Bucket （文件/对象）
// Delete Multiple Object 接口请求实现在指定 Bucket 中批量删除 Object，
// 单次请求最大支持批量删除 1000 个 Object。
// 对于响应结果，COS 提供 Verbose 和 Quiet 两种模式：
// Verbose 模式将返回每个 Object 的删除结果；Quiet 模式只返回报错的 Object 信息。
// 注意：
// 此请求必须携带 Content-MD5 用来校验 Body 的完整性。
// 1. 每一个批量删除请求，最多只能包含 1000个 需要删除的对象；
// 2. 批量删除支持二种模式的放回，verbose 模式和 quiet 模式， 默认为 verbose 模式。
//    verbose 模式返回每个 key 的删除情况，quiet 模式只返回删除失败的 key 的情况；
// 3. 批量删除需要携带 Content-MD5 头部，用以校验请求 body 没有被修改；
// 4. 批量删除请求允许删除一个不存在的 key，仍然认为成功；
// URI: /?delete
func (bucket Bucket) DeleteMultiObject(headers map[string]string, param Delete) (DeleteResult, error) {
	var result DeleteResult

	reqUrl := bucket.makeReqUrl("/?delete")

	data, err := xml.Marshal(param)
	if err != nil {
		return result, errors.New("param error")
	}

	if headers == nil {
		headers = map[string]string{}
	}
	headers[CONTENT_TYPE] = "application/xml"
	headers[CONTENT_MD5] = base64.StdEncoding.EncodeToString(md5.Sum(data)[:])

	buffer := new(bytes.Buffer)
	buffer.Write(data)

	resp, err := bucket.do("POST", reqUrl, headers, nil, buffer)
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

// DeleteObject 在 Client 中删除指定 Bucket （文件/对象）
// Delete Object 接口请求可以在 COS 的 Bucket 中将一个文件（Object）删除。该操作需要请求者对 Bucket 有 WRITE 权限。
// 在 Delete Object 请求中删除一个不存在的 Object，仍然认为是成功的，返回 204 No Content。
// Delete Object 要求用户对该 Object 要有写权限。
// URI: /<ObjectName>
func (bucket Bucket) DeleteObject(objectName string) error {
	reqUrl := bucket.makeReqUrl("/" + objectName)
	resp, err := bucket.do("DELETE", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetObject 将一个Object（文件/对象）下载至本地
// URI: /<ObjectName>
func (bucket Bucket) GetObject(objectName string, params, headers map[string]string) (io.ReadCloser, error) {
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s", objectName))
	resp, err := bucket.do("GET", reqUrl, headers, params, nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetObjectACL 获取 Bucket（文件/对象）的ACL表
// URI: /<ObjectName>?acl
func (bucket Bucket) GetObjectACL(objectName string) (AccessControlPolicy, error) {
	var result AccessControlPolicy

	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?acl", objectName))
	resp, err := bucket.do("GET", reqUrl, nil, nil, nil)
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

// HeadObject 获得 Bucket 的 meta 信息
// URI: /<ObjectName>
func (bucket Bucket) HeadObject(objectName string) error {
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s", objectName))
	resp, err := bucket.do("HEAD", reqUrl, nil, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// InitiateMultipartUpload 初始化 Multipart Upload 上传操作
// URI: /<ObjectName>?uploads
func (bucket Bucket) InitiateMultipartUpload(objectName string) (InitiateMultipartUploadResult, error) {
	var result InitiateMultipartUploadResult

	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?uploads", objectName))
	resp, err := bucket.do("POST", reqUrl, nil, nil, nil)
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

// ListParts 查询特定分块上传操作中的已上传的块
// URI: /<ObjectName>?uploadId=UploadId
// TODO 改参数格式，extra --> ListPartRequest
func (bucket Bucket) ListParts(objectName, uploadId string, extra map[string]string) (ListPartsResult, error) {
	var result ListPartsResult

	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?uploadId=%s", objectName, uploadId))
	params := map[string]string{}
	if extra != nil {
		params = extra
	}
	params["uploadId"] = uploadId
	paramStr := mapToStr(params, "&")
	if paramStr != "" {
		reqUrl = fmt.Sprintf("%s&%s", reqUrl, paramStr)
	}

	resp, err := bucket.do("GET", reqUrl, nil, params, nil)
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

// OptionsObject 实现 Object 跨域访问配置的预请求
// 在发送跨域请求之前会发送一个 OPTIONS 请求并带上特定的来源域，
// HTTP 方法和 HEADER 信息等给 COS，以决定是否可以发送真正的跨域请求。
// 当 CORS 配置不存在时，请求返回 403 Forbidden。
// 可以通过 Put Bucket CORS 接口来开启 Bucket 的 CORS 支持。
// URI: /<ObjectName>
func (bucket Bucket) OptionsObject(objectName string, headers map[string]string) error {
	reqUrl := bucket.makeReqUrl("/coss3/" + objectName)
	resp, err := bucket.do("OPTIONS", reqUrl, headers, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// TODO PostObject
func (bucket Bucket) PostObject() {

}

// TODO PostObjectRestore
func (bucket Bucket) PostObjectRestore() {

}

// PutObject 将一个Object（文件/对象）上传至指定 Client
// 简单上传接口，可以将本地的小于 5 GB 的文件（Object）上传至指定 Bucket 中，
// 大于 5 GB 的文件请使用分片接口上传（Upload Part）。
// 该操作需要请求者对 Bucket 有 WRITE 权限。
// URI: /<ObjectName>
func (bucket Bucket) PutObject(objectName string, data io.Reader, headers map[string]string) error {
	reqUrl := bucket.makeReqUrl("/" + objectName)
	resp, err := bucket.do("PUT", reqUrl, headers, nil, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// PutObjectCopy 实现将一个文件从源路径复制到目标路径
// 建议文件大小 1M 到 5G，超过 5G 的文件请使用分块上传 Upload - Copy。
// 在拷贝的过程中，文件元属性和 ACL 可以被修改。
// 用户可以通过该接口实现文件移动，文件重命名，修改文件属性和创建副本。
// sourceURI: eg: <BucketName-APPID>.cos.<Region>.myqcloud.com/<sourceObjectName>
// URI /<destinationObject>
func (bucket Bucket) PutObjectCopy(destObjectName, sourceURI string, headers map[string]string) (CopyObjectResult, error) {
	var result CopyObjectResult

	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s", destObjectName))
	if headers == nil {
		headers = map[string]string{}
	}
	headers[X_COS_COPY_SOURCE] = sourceURI

	resp, err := bucket.do("PUT", reqUrl, headers, nil, nil)
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

// PutObjectACL 写入 Bucket （文件/对象）的 ACL 表
// 可以通过 Header："x-cos-acl"，"x-cos-grant-read"，"x-cos-grant-write"，"x-cos-grant-full-control" 传入 ACL 信息
// 是一个覆盖操作，传入新的 ACL 将覆盖原有 ACL
// URI: /<ObjectName>?acl
func (bucket Bucket) PutObjectACL(objectName string, headers map[string]string, param AccessControlPolicy) error {
	headers = formatMapByKeys(Acl_Header_Keys, headers)
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?acl", objectName))

	data, err := xml.Marshal(param)
	if err != nil {
		return errors.New("param error")
	}

	buffer := new(bytes.Buffer)
	buffer.Write(data)

	resp, err := bucket.do("PUT", reqUrl, headers, nil, buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// UploadPart 分块上传文件
// Upload Part 接口请求实现将对象按照分块的方式上传到 COS。
// 最多支持 10000 分块，每个分块大小为 1 MB 到 5 GB ，最后一个分块可以小于 1 MB。
// 1. 分块上传首先需要进行初始化，使用 Initiate Multipart Upload 接口实现， 初始化后会得到一个 uploadId ，唯一标识本次上传；
// 2. 在每次请求 Upload Part 时，需要携带 partNumber 和 uploadId，partNumber 为块的编号，支持乱序上传；
// 3. 当传入 uploadId 和 partNumber 都相同的时候，后传入的块将覆盖之前传入的块。当 uploadId 不存在时会返回 404 错误，NoSuchUpload。
// URI: /<ObjectName>
//imur InitiateMultipartUploadResult, parts []Part
func (bucket Bucket) UploadPart(imur InitiateMultipartUploadResult, data io.Reader, size int64, partNumber int) (Part, error) {
	var result Part
	params := map[string]string{
		"uploadId":   imur.UploadID,
		"partNumber": strconv.Itoa(partNumber),
	}
	paramStr := mapToStr(params, "&")
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?%s", imur.Key, paramStr))

	resp, err := bucket.do("PUT", reqUrl, nil, params, data)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	result.ETag = resp.Header.Get(ETAG)
	result.PartNumber = partNumber
	result.Size = size
	date := resp.Header.Get(DATE)
	//Wed，18 Jan 2017 16:17:03 GMT
	layout := "Mon, 2 Jan 2006 15:04:05 GMT"
	t, _ := time.Parse(layout, date)
	result.LastModified = t

	return result, nil
}

// UploadPartCopy 分块上传文件
// Upload Part 接口请求实现将对象按照分块的方式上传到 COS。
// 最多支持 10000 分块，每个分块大小为 1 MB 到 5 GB ，最后一个分块可以小于 1 MB。
// 1. 分块上传首先需要进行初始化，使用 Initiate Multipart Upload 接口实现， 初始化后会得到一个 uploadId ，唯一标识本次上传；
// 2. 在每次请求 Upload Part 时，需要携带 partNumber 和 uploadId，partNumber 为块的编号，支持乱序上传；
// 3. 当传入 uploadId 和 partNumber 都相同的时候，后传入的块将覆盖之前传入的块。当 uploadId 不存在时会返回 404 错误，NoSuchUpload。
// URI: /<destinationObject>
func (bucket Bucket) UploadPartCopy(destinationObject, uploadId, partNumber, sourceURI string, headers map[string]string) (CopyObjectResult, error) {
	var result CopyObjectResult

	params := map[string]string{
		"uploadId":   uploadId,
		"partNumber": partNumber,
	}
	paramStr := mapToStr(params, "&")
	reqUrl := bucket.makeReqUrl(fmt.Sprintf("/%s?%s", destinationObject, paramStr))

	if headers == nil {
		headers = map[string]string{}
	}

	headers[X_COS_COPY_SOURCE] = sourceURI

	resp, err := bucket.do("PUT", reqUrl, headers, params, nil)
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

func (bucket Bucket) makeReqUrl(uri string) string {
	return makeRequestUrl(bucket.Client.Config.Scheme, bucket.BucketName, bucket.Client.Config.APPID, bucket.Region, uri)
}

func (bucket Bucket) do(method, reqUrl string, headers, params map[string]string, body io.Reader) (*http.Response, error) {
	return bucket.Client.Conn.Do(method, reqUrl, headers, params, body)
}
