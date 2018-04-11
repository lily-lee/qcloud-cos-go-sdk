# qcloud-cos-go-sdk
腾讯云 COS Golang SDK (XML API)

# Install
```bash

go get github.com/lily-lee/qcloud-cos-go-sdk/cos

```

# Usage

```go
package main

import (
    "github.com/lily-lee/qcloud-cos-go-sdk/cos"
    "log"
)

func main()  {
    client, err := cos.NewClient("AppID", "SecretID", "SecretKey", "https", 600)
    if err != nil {
        // Handle Error
    }
    // GetService()
    result, err := client.GetService()
    if err != nil {
        // Handle Error
    }
    log.Println("get service result:", result)
    
    bucket, err := client.NewBucket("BucketName", "Region")
    if err != nil {
        // Handle Error
    }
    
    err = bucket.PutObjectFromFile("license", "./LICENSE", nil)
    if err != nil {
        // Handle Error
    }
    
}
```

# API
```bash
godoc github.com/lily-lee/qcloud-cos-go-sdk/cos

# then you can see all the types and functions. 

# 实现了腾讯API提供的接口;
# 在腾讯提供的API接口基础上，
# 实现了PutObjectFromFile(),GetObjectToFile(), GetSTS()(获取临时凭证)等接口;
# 下面简单列一下主要的type 和 func，没加参数和返回值。
```
- Client
    - GetAuth() 
    - GetSTS()
    - NewBucket()
    - GetService()
    - DeleteBucket()
    - DeleteBucketCORS()
    - DeleteBucketLifecycle()
    - GetBucket()
    - GetBucketACL()
    - GetBucketCORS()
    - GetBucketLocation()
    - GetBucketLifecycle()
    - HeadBucket()
    - ListMultipartUploads()
    - PutBucket()
    - PutBucketACL()
    - PutBucketCORS() // TODO
    - PutBucketLifecycle() // TODO
    
- Bucket
    - AbortMultipartUpload()
    - CompleteMultipartUpload()
    - DeleteMultiObject()
    - DeleteObject()
    - GetObject()
    - GetObjectToFile()
    - GetObjectACL()
    - HeadObject()
    - InitiateMultipartUpload()
    - ListParts()
    - OptionsObject()
    - PutObject()
    - PutObjectFromFile()
    - PutObjectCopy()
    - PutObjectACL()
    - UploadPart()
    - UploadPartCopy()
    

# 参考资料
- [腾讯云官方API文档](https://cloud.tencent.com/document/product/436/7751)
- [腾讯云官方XML Node.js SDK](https://github.com/tencentyun/cos-nodejs-sdk-v5)
- [腾讯云官方XML JAVA SDK](https://github.com/tencentyun/cos-java-sdk-v5)
- [阿里云官方OSS Golang SDK](https://github.com/aliyun/aliyun-oss-go-sdk)

# LICENSE
[MIT License](https://github.com/lily-lee/qcloud-cos-go-sdk/blob/master/LICENSE)
