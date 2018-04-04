package cos

import (
	"io"
	"os"
)

// GetObjectToFile save object to local file
func (bucket Bucket) GetObjectToFile(objectName, filepath string) error {
	body, err := bucket.GetObject(objectName, nil, nil)
	defer body.Close()

	// 如果文件不存在则创建，存在则清空
	fd, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = io.Copy(fd, body)
	if err != nil {
		return err
	}

	return nil
}

// TODO DownloadFile 分片下载
func (bucket Bucket) DownloadFile() {

}
