package upload

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 重新命名文件/**/
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 不存在为true/**/
func CheckSavePath(dst string) bool {
	//利用oserror.ErrNotExist进行判断
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀名是否允许/**/
func CheckContainExt(t FileType, name string) bool{
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

		}
	return false

}

func CheckMaxSize(t FileType, f *multipart.FileHeader) bool {
	switch t {
	case TypeImage:
		if f.Size >= global.AppSetting.UploadImageMaxSize << 20 {
			return true
		}

	}

	return false
}

func CheckPermission(dst string) bool {
	//利用oserror.ErrPermission进行判断
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	//os.FileMode文件模式,和系统相关联
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return nil
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return nil
	}
	defer out.Close()

	//根据实现不通,可能通过读或者写进行拷贝
	_, err = io.Copy(out, src)

	return err
}