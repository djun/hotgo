// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"io"
	"path"
)

// 文件归属分类
const (
	KindImg   = "images"   // 图片
	KindDoc   = "document" // 文档
	KindAudio = "audio"    // 音频
	KindVideo = "video"    // 视频
	KindOther = "other"    // 其他
)

var (
	// 图片类型
	imgType = g.MapStrStr{
		"jpeg": "image/jpeg",
		"jpg":  "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"webp": "image/webp",
		"cr2":  "image/x-canon-cr2",
		"tif":  "image/tiff",
		"bmp":  "image/bmp",
		"heif": "image/heif",
		"jxr":  "image/vnd.ms-photo",
		"psd":  "image/vnd.adobe.photoshop",
		"ico":  "image/vnd.microsoft.icon",
		"dwg":  "image/vnd.dwg",
	}

	// 文档类型
	docType = g.MapStrStr{
		"doc":  "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"ppt":  "application/vnd.ms-powerpoint",
		"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	}

	// 音频类型
	audioType = g.MapStrStr{
		"mid":  "audio/midi",
		"mp3":  "audio/mpeg",
		"m4a":  "audio/mp4",
		"ogg":  "audio/ogg",
		"flac": "audio/x-flac",
		"wav":  "audio/x-wav",
		"amr":  "audio/amr",
		"aac":  "audio/aac",
		"aiff": "audio/x-aiff",
	}

	// 视频类型
	videoType = g.MapStrStr{
		"mp4":  "video/mp4",
		"m4v":  "video/x-m4v",
		"mkv":  "video/x-matroska",
		"webm": "video/webm",
		"mov":  "video/quicktime",
		"avi":  "video/x-msvideo",
		"wmv":  "video/x-ms-wmv",
		"mpg":  "video/mpeg",
		"flv":  "video/x-flv",
		"3gp":  "video/3gpp",
	}
)

// IsImgType 判断是否为图片
func IsImgType(ext string) bool {
	_, ok := imgType[ext]
	return ok
}

// IsDocType 判断是否为文档
func IsDocType(ext string) bool {
	_, ok := docType[ext]
	return ok
}

// IsAudioType 判断是否为音频
func IsAudioType(ext string) bool {
	_, ok := audioType[ext]
	return ok
}

// IsVideoType 判断是否为视频
func IsVideoType(ext string) bool {
	_, ok := videoType[ext]
	return ok
}

// GetImgType 获取图片类型
func GetImgType(ext string) (string, error) {
	if mime, ok := imgType[ext]; ok {
		return mime, nil
	}
	return "", gerror.New("Invalid image type")
}

// GetFileType 获取文件类型
func GetFileType(ext string) (string, error) {
	if mime, ok := imgType[ext]; ok {
		return mime, nil
	}
	if mime, ok := docType[ext]; ok {
		return mime, nil
	}
	if mime, ok := audioType[ext]; ok {
		return mime, nil
	}
	if mime, ok := videoType[ext]; ok {
		return mime, nil
	}
	return "", gerror.Newf("Invalid file type:%v", ext)
}

// GetFileKind 获取文件所属分类
func GetFileKind(ext string) string {
	if _, ok := imgType[ext]; ok {
		return KindImg
	}
	if _, ok := docType[ext]; ok {
		return KindDoc
	}
	if _, ok := audioType[ext]; ok {
		return KindAudio
	}
	if _, ok := videoType[ext]; ok {
		return KindVideo
	}
	return KindOther
}

// Ext 获取文件后缀
func Ext(baseName string) string {
	return gstr.StrEx(path.Ext(baseName), ".")
}

// UploadFileByte 获取上传文件的byte
func UploadFileByte(file *ghttp.UploadFile) ([]byte, error) {
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(open)
}

// CalcFileMd5 计算文件md5值
func CalcFileMd5(file *ghttp.UploadFile) (string, error) {
	f, err := file.Open()
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, file.Filename)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
