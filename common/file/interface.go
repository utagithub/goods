package file

// DriverType 驱动类型
type DriverType string

const (
	// Local 本地存储
	Local DriverType = "Local"
	// AliYunOSS 阿里云OSS
	AliYunOSS DriverType = "AliYunOSS"
	// QiNiuKodo 七牛云kodo
	QiNiuKodo DriverType = "QiNiuKodo"
	// HuaweiOBS 华为云OBS
	HuaweiOBS DriverType = "HuaweiOBS"
)

type ClientOption map[string]interface{}

// TODO: FileStoreType名称待定

// FileStoreType OXS
type FileStoreType interface {
	// Setup 装载 endpoint sss
	Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error
	// UpLoad 上传
	UpLoad(yourObjectName string, localFile interface{}) (string, error)
	// GetTempToken 获取临时Token
	GetTempToken() (string, error)
}
