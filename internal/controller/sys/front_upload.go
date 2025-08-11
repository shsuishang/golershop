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
 * 上传-控制器
 * @author Xinze
 * @since 2021/11/18
 * @File : Upload
 */
package sys

import (
	"context"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// 用户控制器管理对象
var Upload = new(cUpload)

type cUpload struct{}

// Upload 系统上传
func (c *cUpload) Upload(ctx context.Context, req *sys.UploadReq) (out sys.UploadRes, err error) {
	// 上传上传
	var info model.FileInfo
	info.Type = ".jpg"

	switch req.MaterialType {
	case "image":
		info, err = service.Upload().UpdImg(ctx)
	case "video":
		info, err = service.Upload().UpdVideo(ctx)
	case "document":
		info, err = service.Upload().UpdFile(ctx)
	}

	if err != nil {
		return out, err
	}

	//存入记录
	service.MaterialBase().Add(ctx, &do.MaterialBase{
		GalleryId:        req.GalleryId,
		MaterialType:     req.MaterialType,
		MaterialSize:     info.FileSize,
		MaterialName:     info.FileName,
		MaterialAlt:      info.FileName,
		MaterialUrl:      info.FileUrl,
		MaterialPath:     info.FilePath,
		MaterialMimeType: info.MimeType,
	})

	out.FileInfo = info

	return
}

// UploadImage 图片上传
func (c *cUpload) UploadImage(ctx context.Context, req *sys.UploadImageReq) (out sys.UploadRes, err error) {
	// 上传上传
	var info model.FileInfo
	info, err = service.Upload().UpdImg(ctx)

	if err != nil {
		return out, err
	}

	//存入记录
	service.MaterialBase().Add(ctx, &do.MaterialBase{
		GalleryId:        req.GalleryId,
		MaterialType:     req.MaterialType,
		MaterialSize:     info.FileSize,
		MaterialName:     info.FileName,
		MaterialAlt:      info.FileName,
		MaterialUrl:      info.FileUrl,
		MaterialPath:     info.FilePath,
		MaterialMimeType: info.MimeType,
	})

	out.FileInfo = info

	return
}

// UploadFile 文件上传
func (c *cUpload) UploadFile(ctx context.Context, req *sys.UploadFileReq) (out sys.UploadRes, err error) {
	// 上传上传
	var info model.FileInfo
	info, err = service.Upload().UpdImg(ctx)

	if err != nil {
		return out, err
	}

	//存入记录
	service.MaterialBase().Add(ctx, &do.MaterialBase{
		GalleryId:        req.GalleryId,
		MaterialType:     req.MaterialType,
		MaterialSize:     info.FileSize,
		MaterialName:     info.FileName,
		MaterialAlt:      info.FileName,
		MaterialUrl:      info.FileUrl,
		MaterialPath:     info.FilePath,
		MaterialMimeType: info.MimeType,
	})

	out.FileInfo = info

	return
}

// UploadImage 视频上传
func (c *cUpload) UploadVideo(ctx context.Context, req *sys.UploadVideoReq) (out sys.UploadRes, err error) {
	// 上传上传
	var info model.FileInfo
	info, err = service.Upload().UpdImg(ctx)

	if err != nil {
		return out, err
	}

	//存入记录
	service.MaterialBase().Add(ctx, &do.MaterialBase{
		GalleryId:        req.GalleryId,
		MaterialType:     req.MaterialType,
		MaterialSize:     info.FileSize,
		MaterialName:     info.FileName,
		MaterialAlt:      info.FileName,
		MaterialUrl:      info.FileUrl,
		MaterialPath:     info.FilePath,
		MaterialMimeType: info.MimeType,
	})

	out.FileInfo = info

	return
}
