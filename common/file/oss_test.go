package file

import (
	"testing"
)

var (
	endpoint        string = "https://oss-cn-beijing.aliyuncs.com"
	accessKeyID     string = "LTAIMsginO6370pg"
	accessKeySecret string = "iC9KykZSpc2xSJOnZ3NLcr4NikVZvl"
	bucketName      string = "u-ta"
	imgPath         string = "image/"
	videogPath      string = "video/"
)

func TestOSSUpload(t *testing.T) {
	// 大括号内填写自己的测试信息即可
	e := OXS{endpoint, accessKeyID, accessKeySecret, bucketName}
	var oxs = e.Setup(AliYunOSS)

	signedURL, err := oxs.UpLoad(imgPath+"1.png", "/opt/homebrew/var/www/web_go/goods/storage/public/1.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Log("signedURL:" + signedURL)
}
