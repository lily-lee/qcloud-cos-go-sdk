# qcloud-cos-go-sdk
腾讯云 COS Golang SDK (XML API)

# Install & Update
```bash
# Install
go get github.com/lily-lee/qcloud-cos-go-sdk/cos

# Update
go get -u github.com/lily-lee/qcloud-cos-go-sdk/cos

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

# RoadMap

### Instance
- [x] New Client (NewClient())
- [x] New Bucket (client.NewBucket())

### Service
- [x] GET Service (client.GetService())

### STS 临时密钥
- [x] CAM STS (client.GetSTS())

### Bucket
- [x] DELETE Bucket (client.DeleteBucket())
- [x] DELETE Bucket cors (client.DeleteBucketCORS())
- [x] DELETE Bucket lifecycle (client.DeleteBucketLifecycle())
- [x] GET Bucket (client.GetBucket())
- [x] GET Bucket acl (client.GetBucketACL())
- [x] GET Bucket cors (client.GetBucketCORS())
- [x] GET Bucket location (client.GetBucketLocation()) 
- [x] GET Bucket lifecycle (client.GetBucketLifecycle())
- [x] HEAD Bucket (client.HeadBucket())
- [x] List Multipart Uploads (client.ListMultipartUploads())
- [x] PUT Bucket (client.PutBucket())
- [x] PUT Bucket acl (client.PutBucketACL())
- [ ] PUT Bucket cors (client.PutBucketCORS())
- [ ] PUT Bucket lifecycle (client.PutBucketLifecycle())
    
### Object
- [x] Abort Multipart Upload (bucket.AbortMultipartUpload()
- [x] Complete Multipart Upload (bucket.CompleteMultipartUpload())
- [x] DELETE Multipart Objects (bucket.DeleteMultiObject())
- [x] DELETE Object (bucket.DeleteObject())
- [x] GET Object (bucket.GetObject())
- [x] GET Object to local file (bucket.GetObjectToFile())
- [x] GET Object acl (bucket.GetObjectACL())
- [x] HEAD Object (bucket.HeadObject())
- [x] Initiate Multipart Upload (bucket.InitiateMultipartUpload())
- [x] List Parts (bucket.ListParts())
- [x] OPTIONS Object (bucket.OptionsObject())
- [x] PUT Object (bucket.PutObject())
- [x] PUT Object from file (bucket.PutObjectFromFile())
- [x] PUT Object acl (bucket.PutObjectACL())
- [x] PUT Object - Copy (bucket.PutObjectCopy())
- [x] Upload Part (bucket.UploadPart())
- [x] Upload Part - Copy (bucket.UploadPartCopy())
    

# 参考资料
- [腾讯云官方API文档](https://cloud.tencent.com/document/product/436/7751)
- [腾讯云官方XML Node.js SDK](https://github.com/tencentyun/cos-nodejs-sdk-v5)
- [腾讯云官方XML JAVA SDK](https://github.com/tencentyun/cos-java-sdk-v5)
- [阿里云官方OSS Golang SDK](https://github.com/aliyun/aliyun-oss-go-sdk)

# LICENSE
[MIT License](https://github.com/lily-lee/qcloud-cos-go-sdk/blob/master/LICENSE)
