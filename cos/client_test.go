package cos

import (
	"encoding/xml"
	"log"
	"testing"
	"time"
)

func TestGetService(t *testing.T) {
	//c, _ := NewClient("1234567890", "AKIDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "f4fyXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "https", 600)

	//c.GetService()

	//fmt.Println(c.DeleteBucket("lily", "ap-chengdu"))

	//fmt.Println(c.DeleteBucketCORS("lily", "ap-chengdu"))

	//fmt.Println(c.DeleteBucketLifecycle("lily", "ap-chengdu"))

	//c.GetBucket("lilytest-1234567890", "ap-chengdu")

	//c.GetBucketACL("lilytest-1234567890", "ap-chengdu")
	//c.GetBucketACL("lily-audio2-1234567890", "ap-shanghai")

	//c.GetBucketCORS("lilytest-1234567890", "ap-chengdu")
	//c.GetBucketCORS("lily-audio2-1234567890", "ap-shanghai")

	//c.GetBucketLocation("lilytest-1234567890", "ap-chengdu")

	//c.GetBucketLifecycle("lilytest-1234567890", "ap-chengdu")
	//c.GetBucketLifecycle("lily-audio2-1234567890", "ap-shanghai")

	//fmt.Println(c.HeadBucket("lilytest-1234567890", "ap-chengdu"))

	//c.ListMultipartUploads("lilytest-1234567890", "ap-chengdu", &ListMultipartUploadsRequestParam{EncodingType: "url"})

	//c.PutBucket("lilytest-1234567890", "ap-shanghai", map[string]string{
	//	"x-cos-acl": "public-read-write",
	//})
	//c.DeleteBucket("lilytest-1234567890", "ap-chengdu")

	//c.PutBucket("lilytest-1234567890", "ap-chengdu",nil)
	//c.ListMultipartUploads("lily-audio2-1234567890", "ap-shanghai", nil)
	//c.GetBucketACL("lilytest-1234567890", "ap-chengdu")
	//c.PutBucketACL("lilytest-1234567890", "ap-chengdu", map[string]string{
	//	"x-cos-acl": "public-read-write",
	//})
	//c.GetBucketACL("lilytest-1234567890", "ap-chengdu")

	//b, _ := c.NewBucket("lilytest-1234567890", "ap-chengdu")
	//b.PutObjectFromFile("test.txt", "../files/test.txt", nil)
	////filename := "莫娣.Maudie.2017.BD720P.x264.官方中文字幕.btrenren.torrent"
	////filename := "a.txt"
	//
	//filepath := "../files/阳光美男.tiff"
	//fd, err := os.Open(filepath)
	//stat, _ := fd.Stat()
	//filename := stat.Name()
	//log.Println(filename)
	//defer fd.Close()
	////stats, err := fd.Stat()
	//
	//if err != nil {
	//	log.Println("open file ", err)
	//}
	//
	//b.PutObject(filename, fd, nil)
	//h := map[string]string{
	//	"Content-Type": "image/tiff	",
	//}
	//b.GetObject("阳光美男.tiff", nil, nil)

	//b.InitiateMultipartUpload("test.txt")
	//b.ListParts("test.txt", "1522347178f07b2f9f11a2b43ac6ffc0eac6401f7ae7fd52aa510de76ea901e47cf607310c", nil)
	//fmt.Println(b.OptionsObject("test.js", nil))
	//b.HeadObject("test.txt")

	//b.GetObjectACL("test.txt")

	//filename := "a.txt"

	//filepath := "../files/test.txt"
	//fd, err := os.Open(filepath)
	//stat, _ := fd.Stat()
	//filename := stat.Name()
	//log.Println(filename)
	//defer fd.Close()
	////stats, err := fd.Stat()
	//
	//if err != nil {
	//	log.Println("open file ", err)
	//}
	//
	//b.UploadPart("test.txt", "1522246687d66fcfee8a8895d12c071841c4d277fef8e3c7747c86bc9298f495e7acbc9915", "1", fd)

	//b.PutObjectACL("a.txt", map[string]string{"x-cos-acl": "private"})

	//b.PutObjectCopy("a.copy.txt", "lilytest-1234567890.cos.ap-chengdu.myqcloud.com/test.txt", nil)

	//b.AbortMultipartUpload("test.txt", "15222476692365fb4f6bd40057d242cdffa431528c53860b577ef08d22f08d90ee9a564522")

	//b.CompleteMultipartUpload("test.txt", "1522246687d66fcfee8a8895d12c071841c4d277fef8e3c7747c86bc9298f495e7acbc9915", nil)

	//b.DeleteObject("test.txt")

	a := `<LocationConstraint>Lily</LocationConstraint>`
	//var f LocationConstraint
	var f string
	xml.Unmarshal([]byte(a), &f)
	log.Println(f)

	d := "Wed, 18 Jan 2017 16:17:03 GMT"
	layout := "Mon, 2 Jan 2006 15:04:05 GMT"
	fo, _ := time.Parse(layout, d)
	log.Println(fo)

}
