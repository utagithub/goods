package file

import "fmt"

type OXS struct {
	// Endpoint 访问域名
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	// AccessKeyID AK
	AccessKeyID string `yaml:"accesskeyid" json:"accesskeyid"`
	// AccessKeySecret AKS
	AccessKeySecret string `yaml:"accesskeysecret" json:"accesskeysecret"`
	// BucketName 桶名称
	BucketName string `yaml:"bucketname" json:"bucketname"`
}

// Setup 配置文件存储driver
func (e *OXS) Setup(driver DriverType, options ...ClientOption) FileStoreType {
	fileStoreType := driver
	var fileStore FileStoreType
	switch fileStoreType {
	case AliYunOSS:
		fileStore = new(ALiYunOSS)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName)
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	case HuaweiOBS:
		fileStore = new(HuaWeiOBS)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName)
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	case QiNiuKodo:
		fileStore = new(QiNiuKODO)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName)
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	}

	return nil
}
