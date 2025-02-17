package operations

import (
	"fmt"
	"github.com/qiniu/qshell/v2/iqshell"
	"github.com/qiniu/qshell/v2/iqshell/common/alert"
	"github.com/qiniu/qshell/v2/iqshell/common/data"
	"github.com/qiniu/qshell/v2/iqshell/common/export"
	"github.com/qiniu/qshell/v2/iqshell/common/flow"
	"github.com/qiniu/qshell/v2/iqshell/common/log"
	"github.com/qiniu/qshell/v2/iqshell/common/utils"
	"github.com/qiniu/qshell/v2/iqshell/storage/object"
	"github.com/qiniu/qshell/v2/iqshell/storage/object/batch"
	"path/filepath"
	"time"
)

type StatusInfo object.StatusApiInfo

func (info *StatusInfo) Check() *data.CodeError {
	if len(info.Bucket) == 0 {
		return alert.CannotEmptyError("Bucket", "")
	}
	if len(info.Key) == 0 {
		return alert.CannotEmptyError("Key", "")
	}
	return nil
}

func Status(cfg *iqshell.Config, info StatusInfo) {
	if shouldContinue := iqshell.CheckAndLoad(cfg, iqshell.CheckAndLoadInfo{
		Checker: &info,
	}); !shouldContinue {
		return
	}

	result, err := object.Status(object.StatusApiInfo(info))
	if err != nil {
		data.SetCmdStatusError()
		log.ErrorF("Status Failed, [%s:%s], Error:%v",
			info.Bucket, info.Key, err)
		return
	}

	if result.IsSuccess() {
		log.Alert(getResultInfo(info.Bucket, info.Key, result))
	}
}

type BatchStatusInfo struct {
	BatchInfo batch.Info
	Bucket    string
}

func (info *BatchStatusInfo) Check() *data.CodeError {
	if err := info.BatchInfo.Check(); err != nil {
		return err
	}

	if len(info.Bucket) == 0 {
		return alert.CannotEmptyError("Bucket", "")
	}
	return nil
}

func BatchStatus(cfg *iqshell.Config, info BatchStatusInfo) {
	cfg.JobPathBuilder = func(cmdPath string) string {
		jobId := utils.Md5Hex(fmt.Sprintf("%s:%s:%s", cfg.CmdCfg.CmdId, info.Bucket, info.BatchInfo.InputFile))
		return filepath.Join(cmdPath, jobId)
	}
	if shouldContinue := iqshell.CheckAndLoad(cfg, iqshell.CheckAndLoadInfo{
		Checker: &info,
	}); !shouldContinue {
		return
	}

	exporter, err := export.NewFileExport(info.BatchInfo.FileExporterConfig)
	if err != nil {
		log.Error(err)
		data.SetCmdStatusError()
		return
	}

	batch.NewHandler(info.BatchInfo).
		EmptyOperation(func() flow.Work {
			return &object.StatusApiInfo{}
		}).
		SetFileExport(exporter).
		ItemsToOperation(func(items []string) (operation batch.Operation, err *data.CodeError) {
			key := items[0]
			if key != "" {
				return &object.StatusApiInfo{
					Bucket: info.Bucket,
					Key:    key,
				}, nil
			}
			return nil, alert.Error("key invalid", "")
		}).
		OnResult(func(operationInfo string, operation batch.Operation, result *batch.OperationResult) {
			apiInfo, ok := (operation).(*object.StatusApiInfo)
			if !ok {
				data.SetCmdStatusError()
				log.ErrorF("Status Failed, %s, Code: %d, Error: %s", operationInfo, result.Code, result.Error)
				return
			}
			in := (*StatusInfo)(apiInfo)
			if result.IsSuccess() {
				log.AlertF("%s\t%d\t%s\t%s\t%d\t%d",
					in.Key, result.FSize, result.Hash, result.MimeType, result.PutTime, result.Type)
			} else {
				data.SetCmdStatusError()
				log.ErrorF("Status Failed, [%s:%s], Code: %d, Error: %s", in.Bucket, in.Key, result.Code, result.Error)
			}
		}).
		OnError(func(err *data.CodeError) {
			data.SetCmdStatusError()
			log.ErrorF("Batch Status error:%v:", err)
		}).Start()
}

func getResultInfo(bucket, key string, status object.StatusResult) string {
	statInfo := fmt.Sprintf("%-20s%s\r\n", "Bucket:", bucket)
	statInfo += fmt.Sprintf("%-20s%s\r\n", "Key:", key)
	statInfo += fmt.Sprintf("%-20s%s\r\n", "FileHash:", status.Hash)
	statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "Fsize:", status.FSize, utils.FormatFileSize(status.FSize))

	putTime := time.Unix(0, status.PutTime*100)
	statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "PutTime:", status.PutTime, putTime.String())
	statInfo += fmt.Sprintf("%-20s%s\r\n", "MimeType:", status.MimeType)

	resoreStatus := ""
	if status.RestoreStatus > 0 {
		if status.RestoreStatus == 1 {
			resoreStatus = "解冻中"
		} else if status.RestoreStatus == 2 {
			resoreStatus = "解冻完成"
		}
	}
	if len(resoreStatus) > 0 {
		statInfo += fmt.Sprintf("%-20s%d(%s)\r\n", "RestoreStatus:", status.RestoreStatus, resoreStatus)
	}

	if status.Expiration > 0 {
		expiration := time.Unix(status.Expiration, 0)
		statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "Expiration:", status.Expiration, expiration.String())
	}

	if status.TransitionToIA > 0 {
		date := time.Unix(status.TransitionToIA, 0)
		statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "TransitionToIA:", status.TransitionToIA, date.String())
	}

	if status.TransitionToARCHIVE > 0 {
		date := time.Unix(status.TransitionToARCHIVE, 0)
		statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "TransitionToARCHIVE:", status.TransitionToARCHIVE, date.String())
	}

	if status.TransitionToDeepArchive > 0 {
		date := time.Unix(status.TransitionToDeepArchive, 0)
		statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "TransitionToDeepArchive:", status.TransitionToDeepArchive, date.String())
	}

	statInfo += fmt.Sprintf("%-20s%d -> %s\r\n", "FileType:", status.Type, getStorageTypeDescription(status.Type))

	return statInfo
}

var objectTypes = []string{"标准存储", "低频存储", "归档存储", "深度归档存储"}

func getStorageTypeDescription(storageType int) string {
	typeString := "未知类型"
	if storageType >= 0 && storageType < len(objectTypes) {
		typeString = objectTypes[storageType]
	}
	return typeString
}
