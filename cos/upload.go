package cos

import "os"

// PutObjectFromFile 将一个Object（文件/对象）上传至指定 Client
// 从本地文件上传
// URI: /<ObjectName>
func (bucket Bucket) PutObjectFromFile(objectName string, filepath string, headers map[string]string) error {
	fd, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	return bucket.PutObject(objectName, fd, headers)
}

// UploadFile 分片上传
// 将文件分片
// 将分片并发上传
// TODO UploadFile
func (bucket Bucket) UploadFile() {

}
