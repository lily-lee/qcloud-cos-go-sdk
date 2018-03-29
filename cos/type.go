package cos

import (
	"net/url"
	"time"
)

// ListAllMyBucketsResult is the result of client.GetService()
type ListAllMyBucketsResult struct {
	Owner   Owner        `xml:"Owner"`
	Buckets []BucketNode `xml:"Buckets>Bucket"`
}

// Owner 说明 Client 持有者的信息
type Owner struct {
	ID          string `xml:"ID"`
	DisplayName string `xml:"DisplayName"`
}

// BucketNode Container 节点 Client 的内容
type BucketNode struct {
	Name         string    `xml:"Name"`
	Location     string    `xml:"Location"`
	CreationDate time.Time `xml:"CreationDate"`
}

// ListBucketResult is the result of client.GetBucket()
type ListBucketResult struct {
	Name           string         `xml:"Name"`
	EncodingType   string         `xml:"Encoding-Type"`
	Prefix         string         `xml:"Prefix"`
	Marker         string         `xml:"Maxkeys"`
	MaxKeys        string         `xml:"MaxKeys"`
	IsTruncated    bool           `xml:"IsTruncated"`
	NextMarker     string         `xml:"NextMarker"`
	Contents       []Contents     `xml:"Contents"`
	CommonPrefixes CommonPrefixes `xml:"CommonPrefixes"`
}

type Contents struct {
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	ETag         string    `xml:"ETag"`
	Size         string    `xml:"Size"`
	Owner        Owner     `xml:"Owner"`
	StorageClass string    `xml:"StorageClass"`
}

type CommonPrefixes struct {
	Prefix string `xml:"Prefix"`
}

// AccessControlPolicy is the result of client.GetBucketACL()
type AccessControlPolicy struct {
	Owner             Owner   `xml:"Owner"`
	AccessControlList []Grant `xml:"AccessControlList>Grant"`
}

type Grant struct {
	Grantee    Grantee `xml:"Grantee"`
	Permission string  `xml:"Permission"`
}

type Grantee struct {
	ID          string `xml:"ID"`
	DisplayName string `xml:"DisplayName"`
	URI         string `xml:"URI"`
}

// CORSConfiguration is the result of client.GetBucketCORS()
type CORSConfiguration struct {
	CORSRule []CORSRule `xml:"CORSRule"`
}

type CORSRule struct {
	ID            string `xml:"ID"`
	AllowedOrigin string `xml:"AllowedOrigin"`
	AllowedMethod string `xml:"AllowedMethod"` // GET,PUT,HEAD,POST,DELETE
	AllowedHeader string `xml:"AllowedHeader"`
	MaxAgeSeconds int64  `xml:"MaxAgeSeconds"`
	ExposeHeader  string `xml:"ExposeHeader"`
}

// LocationConstraint is the result of client.GetBucketLocation()
type LocationConstraint string

// LifecycleConfiguration is the result of client.GetBucketLifecycle()
type LifecycleConfiguration struct {
	Rule []Rule `xml:"Rule"`
}

type Rule struct {
	ID                             string                         `xml:"ID"`
	Filter                         Filter                         `xml:"Filter"`
	Status                         string                         `xml:"Status"` // Enabled, Disabled
	Transition                     Transition                     `xml:"Transition"`
	Expiration                     Expiration                     `xml:"Expiration"`
	AbortIncompleteMultipartUpload AbortIncompleteMultipartUpload `xml:"AbortIncompleteMultipartUpload"`
	NoncurrentVersionExpiration    NoncurrentVersionExpiration    `xml:"NoncurrentVersionExpiration"`
	NoncurrentVersionTransition    NoncurrentVersionTransition    `xml:"NoncurrentVersionTransition"`
}

type AbortIncompleteMultipartUpload struct {
	DaysAfterInitiation int64 `xml:"DaysAfterInitiation"`
}

type NoncurrentVersionExpiration struct {
	NoncurrentDays int64 `xml:"NoncurrentDays"`
}

type NoncurrentVersionTransition struct {
	NoncurrentDays int64  `xml:"NoncurrentDays"`
	StorageClass   string `xml:"StorageClass"`
}

type Filter struct {
	And    And    `xml:"And"`
	Prefix string `xml:"Prefix"`
	Tag    Tag    `xml:"Tag"`
}

type And struct {
	Prefix string `xml:"Prefix"`
	Tag    Tag    `xml:"Tag"`
}

type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type Expiration struct {
	Days                      int64  `xml:"Days"`
	Date                      string `xml:"Date"`
	ExpiredObjectDeleteMarker string `xml:"ExpiredObjectDeleteMarker"`
}

type Transition struct {
	Days         int    `xml:"Days"`
	Date         string `xml:"Date"`
	StorageClass string `xml:"StorageClass"`
}

// ListMultipartUploadsResult is the result of client.ListMultipartUploads()
type ListMultipartUploadsResult struct {
	Bucket             string         `xml:"Bucket"`
	EncodingType       string         `xml:"Encoding-Type"`
	KeyMarker          string         `xml:"KeyMarker"`
	UploadIdMarker     string         `xml:"UploadIdMarker"`
	NextKeyMarker      string         `xml:"NextKeyMarker"`
	NextUploadIdMarker string         `xml:"NextUploadIdMarker"`
	MaxUploads         string         `xml:"MaxUploads"`
	IsTruncated        bool           `xml:"IsTruncated"`
	Prefix             string         `xml:"Prefix"`
	Delimiter          string         `xml:"Delimiter"`
	Upload             []Upload       `xml:"Upload"`
	CommonPrefixes     CommonPrefixes `xml:"CommonPrefixes"`
}

type Upload struct {
	Key          string    `xml:"Key"`
	UploadID     string    `xml:"UploadId"`
	StorageClass string    `xml:"StorageClass"`
	Initiator    Initiator `xml:"Initiator"`
	Owner        Owner     `xml:"Owner"`
	Initiated    time.Time `xml:"Initiated"`
}

type Initiator struct {
	UIN         string `xml:"UIN"`
	ID          string `xml:"ID"`
	DisplayName string `xml:"DisplayName"`
}

type GetBucketACLResult struct{}

// CompleteMultipartUpload is the request data type of bucket.CompleteMultipartUpload()
type CompleteMultipartUpload struct {
	Part Part `xml:"Part"`
}

type Part struct {
	PartNumber   string    `xml:"PartNumber"`
	LastModified time.Time `xml:"LastModified"`
	ETag         string    `xml:"ETag"`
	Size         string    `xml:"Size"`
}

type CompleteMultipartUploadResult struct {
	Location url.URL `xml:"Location"`
	Bucket   string  `xml:"Bucket"`
	Key      string  `xml:"Key"`
	ETag     string  `xml:"ETag"`
}

// DeleteObject is the request data of bucket.DeleteMultipleObjects()
type Delete struct {
	Quiet  bool     `xml:"Quiet"`
	Object []Object `xml:"Object"`
}

type Object struct {
	Key string `xml:"Key"`
}

// DeleteObject is the result of bucket.DeleteMultipleObjects()
type DeleteResult struct {
	Deleted []Deleted         `xml:"Deleted"`
	Error   DeleteResultError `xml:"Error"`
}

type Deleted struct {
	Key string `xml:"Key"`
}

type DeleteResultError struct {
	Key     string `xml:"Key"`
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}

// InitiateMultipartUploadResult InitiateMultipartUpload请求返回结果
type InitiateMultipartUploadResult struct {
	Bucket   string `xml:"Bucket"`   // Bucket名称
	Key      string `xml:"Key"`      // 上传Object名称
	UploadID string `xml:"UploadId"` // 生成的UploadId
}

type ListPartRequest struct {
	UploadID         string `json:"uploadId"`
	EncodingType     string `json:"encoding-type"`
	MaxParts         int    `json:"max-parts"`
	PartNumberMarker string `json:"part-number-marker"`
}

type ListPartsResult struct {
	Bucket               string    `xml:"Bucket"`
	EncodingType         string    `xml:"Encoding-type"`
	Key                  string    `xml:"Key"`
	UploadId             string    `xml:"UploadId"`
	Initiator            Initiator `xml:"Initiator"`
	Owner                Owner     `xml:"Owner"`
	StorageClass         string    `xml:"StorageClass"` // STANDARD, STANDARD_IA, NEARLINE
	PartNumberMarker     string    `xml:"PartNumberMarker"`
	NextPartNumberMarker string    `xml:"NextPartNumberMarker"`
	MaxParts             string    `xml:"MaxParts"`
	IsTruncated          bool      `xml:"IsTruncated"`
	Part                 Part      `xml:"Part"`
}

// PostResponse is the result of bucket.PostObject()
type PostResponse struct {
	Location string `xml:"Location"`
	Bucket   string `xml:"Bucket"`
	Key      string `xml:"key"`
	ETag     string `xml:"ETag"`
}

// RestoreRequest is the request data of bucket.PostObjectRestore()
type RestoreRequest struct {
	Days             int              `xml:"Days"`
	CASJobParameters CASJobParameters `xml:"CASJobParameters"`
}

type CASJobParameters struct {
	Tier string `xml:"Tier"` // Expedited,Standard,Bulk; 默认值:Standard
}

// CopyObjectResult is the result of bucket.PutObjectCopy()
type CopyObjectResult struct {
	ETag         string    `xml:"ETag"`
	LastModified time.Time `xml:"LastModified"`
}

type ListMultipartUploadsRequestParam struct {
	Delimiter      string `json:"delimiter,omitempty"`
	EncodingType   string `json:"encoding-type,omitempty"`
	Prefix         string `json:"prefix,omitempty"`
	MaxUploads     string `json:"max-uploads,omitempty"`
	KeyMarker      string `json:"key-marker,omitempty"`
	UploadIDMarker string `json:"upload-id-marker,omitempty"`
}
