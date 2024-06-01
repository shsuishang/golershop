package utility

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"log"
	"os"
	"regexp"
	"strings"
)

// 附件目录
func UploadPath(ctx context.Context) string {
	// 附件存储路径
	uploadDir, _ := g.Cfg().Get(ctx, "shopSuite.uploadDir")
	if uploadDir.String() != "" {
		return uploadDir.String()
	} else {
		// 获取项目根目录
		curDir, _ := os.Getwd()
		return curDir + "/public/uploads"
	}
}

// 临时目录
func TempPath(ctx context.Context) string {
	return UploadPath(ctx) + "/temp"
}

// 图片存放目录
func ImagePath(ctx context.Context) string {
	return UploadPath(ctx) + "/images"
}

// 视频存放目录
func VideoPath(ctx context.Context) string {
	return UploadPath(ctx) + "/videos"
}

// 文件目录(非图片/视频目录)
func FilePath(ctx context.Context) string {
	return UploadPath(ctx) + "/files"
}

// 创建文件夹并设置权限
func CreateDir(path string) bool {
	// 判断文件夹是否存在
	if IsExist(path) {
		return true
	}
	// 创建文件夹
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// 判断文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		if os.IsExist(err) {
			// 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

func ImgUrl(ctx context.Context) string {
	// 附件存储路径
	imageUrl, _ := g.Cfg().Get(ctx, "shopSuite.urlBase")
	if imageUrl.String() != "" {
		return imageUrl.String()
	} else {
	}

	return ""
}

// 获取文件地址
func GetWebUrl(ctx context.Context, path string) string {
	return ImgUrl(ctx) + path
}

func SaveImage(ctx context.Context, url string, dirname string) (string, error) {
	// 判断文件地址是否为空
	if gstr.Equal(url, "") {
		return "", gerror.New("文件地址不能为空")
	}

	// 判断是否本站图片
	if gstr.Contains(url, ImgUrl(ctx)) {
		// 本站图片

		// 是否临时图片
		if gstr.Contains(url, "temp") {
			// 临时图片

			// 创建目录
			dirPath := ImagePath(ctx) + "/" + dirname + "/" + gtime.Now().Format("Ymd")
			if !CreateDir(dirPath) {
				return "", gerror.New("文件目录创建失败")
			}
			// 原始图片地址
			oldPath := gstr.Replace(url, ImgUrl(ctx), UploadPath(ctx))
			// 目标目录地址
			newPath := ImagePath(ctx) + "/" + dirname + gstr.Replace(url, ImgUrl(ctx)+"/temp", "")
			// 移动文件
			os.Rename(oldPath, newPath)
			return gstr.Replace(newPath, UploadPath(ctx), ""), nil
		} else {
			// 非临时图片
			path := gstr.Replace(url, ImgUrl(ctx), "")
			return path, nil
		}
	} else {
		// 远程图片
		// TODO...
	}
	return "", gerror.New("保存文件异常")
}

// 处理富文本
func SaveImageContent(ctx context.Context, content string, title string, dirname string) string {
	str := `<img src="(?s:(.*?))"`
	//解析、编译正则
	ret := regexp.MustCompile(str)
	// 提取图片信息
	alls := ret.FindAllStringSubmatch(content, -1)
	// 遍历图片数据
	for _, v := range alls {
		// 获取图片地址
		item := v[1]
		if item == "" {
			continue
		}
		// 保存图片至正式目录
		image, _ := SaveImage(ctx, item, dirname)
		if image != "" {
			content = strings.ReplaceAll(content, item, "[IMG_URL]"+image)
		}
	}
	// 设置ALT标题
	if strings.Contains(content, "alt=\"\"") && title != "" {
		content = strings.ReplaceAll(content, "alt=\"\"", "alt=\""+title+"\"")
	}
	return content
}
