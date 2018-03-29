# qcloud-cos-go-sdk
腾讯云 COS Golang SDK (XML API)

# 安装

# 使用

```go

c, err := cos.NewClient("AppID", "SecretID", "SecretKey", "https", 600)
r, err := c.GetService()
if err != nil {
    HandleError...
}
fmt.Println(r)

```

# TODO
- [ ] PutBucketsCORS
- [ ] PutBucketLifecycle
- [ ] PostObject
- [ ] PostObjectCopy
- [ ] GetObjectToFile
- [ ] DownloadFile
- [ ] UploadFile
    - [ ] SliceFile
    - [ ] Concurrent Uploading File Slices
    - [ ] ...

- [ ] Documents and Comments
- [ ] Find Bugs and Fix them
- [ ] ...

# 参考资料
- [腾讯云官方API文档](https://cloud.tencent.com/document/product/436/7751)
- [腾讯云官方XML Node.js SDK](https://github.com/tencentyun/cos-nodejs-sdk-v5)
- [腾讯云官方XML JAVA SDK](https://github.com/tencentyun/cos-java-sdk-v5)
- [阿里云官方OSS Golang SDK](https://github.com/aliyun/aliyun-oss-go-sdk)

# LICENSE
[MIT License](https://github.com/lily-lee/qcloud-cos-go-sdk/blob/master/LICENSE)
