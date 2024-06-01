// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

/**
 * 文件上传-服务类
 * @author Xinze风雨
 * @since 2021/7/23
 * @File : upload
 */
package logic

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"log"
	"net/http"
	"os"
	"strings"
)

type sUpload struct{}

// 初始化
func init() {
	service.RegisterUpload(New())
	// 文件上传路径
	fmt.Println("初始化配置")
}

func New() *sUpload {
	return &sUpload{}
}

func (s *sUpload) upload(ctx context.Context, fileExt string, fileSize string, materialType string) (model.FileInfo, error) {
	// 上传文件
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("upfile")

	// 允许上传文件后缀
	//fileExt := "jpg,gif,png,bmp,jpeg,JPG"
	// 检查上传文件后缀
	fileType, err := checkFileExt(file.Filename, fileExt)
	if err != nil {
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.WriteHeader(http.StatusInternalServerError)
		r.Response.WriteExit(fmt.Sprintf("上传文件格式不正确，文件后缀只允许为：%s 的文件", fileExt))

		//return model.FileInfo{}, gerror.New(fmt.Sprintf("上传文件格式不正确，文件后缀只允许为：%s 的文件", fileExt))
	}
	// 允许文件上传最大值
	//fileSize := "5M"
	// 检查上传文件大小
	isvalid, err := checkFileSize(file.Size, fileSize)
	if err != nil {
		return model.FileInfo{}, err
	}
	if !isvalid {
		return model.FileInfo{}, gerror.New("上传文件大小不得超过：" + fileSize)
	}

	// 临时存储目录
	upPath := utility.TempPath(ctx)

	//判读单登录用户
	loginUser := service.BizCtx().GetUser(ctx)
	var userId uint = 0

	if loginUser != nil {
		userId = loginUser.UserId
	} else {
		// 临时存储目录
	}

	switch materialType {
	case "image":
		upPath = fmt.Sprintf("%s/%d", utility.ImagePath(ctx), userId)
	case "video":
		upPath = fmt.Sprintf("%s/%d", utility.VideoPath(ctx), userId)
	case "document":
		upPath = fmt.Sprintf("%s/%d", utility.FilePath(ctx), userId)
	}

	savePath := upPath + "/" + gtime.Now().Format("Ymd")

	// 创建文件夹
	ok := utility.CreateDir(savePath)
	if !ok {
		return model.FileInfo{}, gerror.New("存储路径创建失败")
	}

	// 上传文件
	fileName, err := file.Save(savePath, true)
	if err != nil {
		return model.FileInfo{}, err
	}

	filePath := savePath + "/" + fileName
	fileUrl := gstr.Replace(savePath, utility.UploadPath(ctx), "/uploads") + "/" + fileName

	//mine信息
	fileTmp, err := os.Open(filePath)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)
	_, err = fileTmp.Read(buff)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	// 返回结果
	result := model.FileInfo{
		FileName: file.Filename,
		FileSize: file.Size,
		FileType: fileType,
		FilePath: filePath,
		FileUrl:  utility.GetWebUrl(ctx, fileUrl),
		MimeType: filetype,
		UserId:   userId,
	}

	return result, nil
}

func (s *sUpload) UpdImg(ctx context.Context) (model.FileInfo, error) {
	// 允许上传文件后缀
	fileExt := "jpg,gif,png,bmp,jpeg"
	// 允许文件上传最大值
	fileSize := "50M"

	// 返回结果
	return s.upload(ctx, fileExt, fileSize, "video")
}

func (s *sUpload) UpdVideo(ctx context.Context) (model.FileInfo, error) {
	// 允许上传文件后缀
	fileExt := "flv,swf,mkv,avi,rm,rmvb,mpeg,mpg,ogg,ogv,mov,wmv,mp4,webm,mp3,wav,mid"
	// 允许文件上传最大值
	fileSize := "4M"

	// 返回结果
	return s.upload(ctx, fileExt, fileSize, "image")
}

func (s *sUpload) UpdFile(ctx context.Context) (model.FileInfo, error) {
	// 允许上传文件后缀
	fileExt := "rar,zip,tar,gz,7z,bz2,cab,iso,doc,docx,xls,xlsx,ppt,pptx,pdf,txt,md,xml,csv" // 允许上传文件后缀
	// 允许文件上传最大值
	fileSize := "50M"

	// 返回结果
	return s.upload(ctx, fileExt, fileSize, "document")
}

// 检查文件格式是否合法
func checkFileExt(fileName string, typeString string) (suffix string, err error) {
	// 上传文件后缀
	suffix = gstr.SubStrRune(fileName, gstr.PosRRune(fileName, ".")+1, gstr.LenRune(fileName)-1)
	// 允许上传文件后缀
	exts := gstr.Split(typeString, ",")
	// 是否验证通过
	isValid := false
	for _, v := range exts {
		// 对比文件后缀
		if gstr.Equal(strings.ToLower(suffix), v) {
			isValid = true
			break
		}
	}

	if !isValid {
		err = gerror.New("文件格式不合法")
	}

	return suffix, err
}

// 检查上传文件大小
func checkFileSize(fileSize int64, maxSize string) (bool, error) {
	// 匹配上传文件最大值
	match, err := gregex.MatchString(`^([0-9]+)(?i:([a-z]*))$`, maxSize)
	if err != nil {
		return false, err
	}
	if len(match) == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB）")
		return false, err
	}
	var cfSize int64
	switch gstr.ToUpper(match[2]) {
	case "MB", "M":
		cfSize = gconv.Int64(match[1]) * 1024 * 1024
	case "KB", "K":
		cfSize = gconv.Int64(match[1]) * 1024
	case "":
		cfSize = gconv.Int64(match[1])
	}
	if cfSize == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB），最大单位为MB")
		return false, err
	}
	return cfSize >= fileSize, nil
}
