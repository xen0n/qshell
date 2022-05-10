package download

import (
	"github.com/qiniu/qshell/v2/iqshell/common/data"
	"github.com/qiniu/qshell/v2/iqshell/common/log"
	"github.com/qiniu/qshell/v2/iqshell/common/utils"
	"github.com/qiniu/qshell/v2/iqshell/storage/object"
	"os"
)

type LocalFileInfo struct {
	File                string // 被检测的文件 【必填】
	Bucket              string // 文件所在 bucket 用于检查 hash【选填】
	Key                 string // 文件被保存的 key  用于检查 hash【选填】
	FileHash            string // 文件 hash，有值则会检测 hash【选填】
	FileSize            int64  // 文件大小，有值则会检测文件大小【选填】
	RemoveFileWhenError bool   // 当遇到错误时是否该移除文件【选填】
}

func (l *LocalFileInfo) CheckDownloadFile() (err *data.CodeError) {
	defer func() {
		if err != nil && l.RemoveFileWhenError {
			e := os.Remove(l.File)
			if e != nil && !os.IsNotExist(e) {
				log.WarningF("download file check: remove file error:%v", e)
			}
		}
	}()

	err = l.CheckFileSizeOfDownloadFile()
	if err != nil {
		return err
	}

	err = l.CheckFileHashOfDownloadFile()
	return
}

func (l *LocalFileInfo) CheckFileSizeOfDownloadFile() *data.CodeError {
	if l.FileSize <= 0 {
		log.Debug("download file check size: needn't to check")
		return nil
	}

	tempFileStatus, err := os.Stat(l.File)
	if err != nil {
		return data.ConvertError(err)
	}

	if tempFileStatus == nil {
		return data.NewEmptyError().AppendDesc("download file check: can't get file status:" + l.File)
	}

	if l.FileSize != tempFileStatus.Size() {
		return data.NewEmptyError().AppendDesc("download file check: download file size is unexpected:" + l.File)
	}

	return nil
}

func (l *LocalFileInfo) CheckFileHashOfDownloadFile() *data.CodeError {
	if len(l.FileHash) == 0 {
		log.Debug("download file check hash: needn't to check")
		return nil
	}

	hashFile, oErr := os.Open(l.File)
	if oErr != nil {
		return data.NewEmptyError().AppendDesc("download file check: get temp file error when check hash:" + oErr.Error())
	}

	// 计算本地文件 hash
	var hash string
	if utils.IsSignByEtagV2(l.FileHash) {
		log.Debug("download file check hash: get etag by v2 for key:" + l.Key)
		if len(l.Bucket) == 0 || len(l.Key) == 0 {
			return data.NewEmptyError().AppendDesc("download file check hash: etag v2 check should provide bucket and key")
		}

		stat, sErr := object.Status(object.StatusApiInfo{
			Bucket: l.Bucket,
			Key:    l.Key,
		})
		if sErr != nil {
			return data.NewEmptyError().AppendDesc("download file check hash: etag v2 get file status error:" + sErr.Error())
		} else {
			if h, eErr := utils.EtagV2(hashFile, stat.Parts); eErr != nil {
				return data.NewEmptyError().AppendDesc("download file check: get file etag v2 error:" + eErr.Error())
			} else {
				hash = h
			}
		}
	} else {
		log.Debug("download file check hash: get etag by v1 for key:" + l.Key)
		if h, eErr := utils.EtagV1(hashFile); eErr != nil {
			return data.NewEmptyError().AppendDesc("download file check: get file etag v1 error:" + eErr.Error())
		} else {
			hash = h
		}
	}

	if hash != l.FileHash {
		return data.NewEmptyError().AppendDesc("download file check: file hash doesn't match for key:" + l.Key + "download file hash:" + hash + " excepted:" + l.FileHash)
	}

	return nil
}
