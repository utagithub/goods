/*
 * @Author: 尤_Ta
 * @Date:  23:29
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:29
 */

package admin

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"goods/common/api"
	"goods/common/file"
	"goods/common/tools"
	"goods/common/utils"
	"goods/config"
	"io/ioutil"
	"strings"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const path = "storage/public/"

type FileUploadController struct {
	api.Api
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/file/upload [post]
// @Security Bearer
func (e FileUploadController) UploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1": // 单图
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2": // 多图
		multipartFile := e.multipleFile(c, urlPrefix)
		if len(multipartFile) == 0 {
			e.Error(200, errors.New(""), "图片不能为空")
			return
		}
		e.OK(multipartFile, "上传成功")
		return
	case "3": // base64
		fileResponse = e.base64Img(c, fileResponse, urlPrefix)
		e.OK(fileResponse, "上传成功")
	default:
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	}

}

// singleFile 单图
func (e FileUploadController) singleFile(c *gin.Context, fileResponse FileResponse, urlPerfix string) (FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		e.Error(200, err, "图片不能为空")
		return FileResponse{}, true
	}
	// 上传文件至指定目录
	guid := uuid.New().String()

	fileName := guid + utils.GetExt(files.Filename)

	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, err, "初始化文件路径失败")
	}
	singleFile := path + fileName
	err = c.SaveUploadedFile(files, singleFile)
	if err != nil {
		e.Error(500, err, "图片图片上传失败")
		return FileResponse{}, true
	}
	fileType, _ := utils.GetType(singleFile)
	fileResponse = FileResponse{
		Size:     tools.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPerfix + singleFile,
		Name:     files.Filename,
		Type:     fileType,
	}
	//source, _ := c.GetPostForm("source")
	//err = thirdUpload(source, fileName, singleFile)
	//if err != nil {
	//	e.Error(200, errors.New(""), "上传第三方失败")
	//	return FileResponse{}, true
	//}
	//fileResponse.Path = "/" + path + fileName
	//fileResponse.FullPath = "/" + path + fileName
	return fileResponse, false
}

// multipleFile 多图
func (e FileUploadController) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	//source, _ := c.GetPostForm("source")
	driver := config.ExtConfig.UpLoad.Driver
	var multipartFile []FileResponse
	if len(files) == 0 {
		return multipartFile
	}
	for _, f := range files {
		guid := uuid.New().String()
		fileName := guid + utils.GetExt(f.Filename)

		err := utils.IsNotExistMkDir(path)
		if err != nil {
			e.Error(500, err, "初始化文件路径失败")
		}
		multipartFileName := path + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := utils.GetType(multipartFileName)
		if err1 == nil {
			singURL, err := thirdUpload(driver, fileName, multipartFileName)
			fmt.Printf("singURL:%v\n", singURL)
			if err != nil {
				e.Error(500, errors.New(""), "上传第三方失败")
			} else {
				fileResponse := FileResponse{
					Size:     tools.GetFileSize(multipartFileName),
					Path:     multipartFileName,
					FullPath: urlPerfix + multipartFileName,
					Name:     f.Filename,
					Type:     fileType,
				}
				if driver != "Local" {
					fileResponse.Path = singURL
					fileResponse.FullPath = singURL
				}
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	return multipartFile
}

// base64Img
func (e FileUploadController) base64Img(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"
	err := utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	base64File := path + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size:     tools.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	driver, _ := c.GetPostForm("source")
	//driver := config.Extend{}.UpLoad.Driver
	_, err = thirdUpload(driver, fileName, base64File)
	if err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}
	if driver != "Local" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	}
	return fileResponse
}

func thirdUpload(driver string, name string, path string) (string, error) {
	xos := file.OXS{config.ExtConfig.UpLoad.Secret.Endpoint, config.ExtConfig.UpLoad.Secret.AccessKeyID, config.ExtConfig.UpLoad.Secret.AccessKeySecret, config.ExtConfig.UpLoad.Secret.BucketName}
	switch driver {
	case "AliYunOSS":
		oss := xos.Setup(file.AliYunOSS)
		return oss.UpLoad(imgPath+name, path)
	case "QiNiuKodo":
		kodo := xos.Setup(file.QiNiuKodo)
		return kodo.UpLoad(imgPath+name, path)
	case "HuaweiOBS":
		obs := xos.Setup(file.HuaweiOBS)
		return obs.UpLoad(imgPath+name, path)
	}
	return "", nil
}

var (
	imgPath string = "image/"
	//videogPath      string = "video/"
	//uploadConfig = config.ExtendConfig
)
