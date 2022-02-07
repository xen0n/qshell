package operations

import (
	"github.com/qiniu/qshell/v2/iqshell/common/group"
	"github.com/qiniu/qshell/v2/iqshell/common/log"
	"github.com/qiniu/qshell/v2/iqshell/common/utils"
	"github.com/qiniu/qshell/v2/iqshell/storage/object"
	"github.com/qiniu/qshell/v2/iqshell/storage/object/batch"
)

type ChangeMimeInfo object.ChangeMimeApiInfo

func ChangeMime(info ChangeMimeInfo) {
	result, err := object.ChangeMimeType(object.ChangeMimeApiInfo(info))
	if err != nil {
		log.ErrorF("Change Mime error:%v", err)
		return
	}

	if len(result.Error) != 0 {
		log.ErrorF("Change Mime:%v", result.Error)
		return
	}
}

type BatchChangeMimeInfo struct {
	BatchInfo batch.Info
	Bucket    string
}

func BatchChangeMime(info BatchChangeMimeInfo) {
	handler, err := group.NewHandler(info.BatchInfo.Info)
	if err != nil {
		log.Error(err)
		return
	}
	batch.NewFlow(info.BatchInfo).ReadOperation(func() (operation batch.Operation, hasMore bool) {
		line, success := handler.Scanner().ScanLine()
		if !success {
			return nil, false
		}

		items := utils.SplitString(line, info.BatchInfo.ItemSeparate)
		if len(items) > 1 {
			key, mime := items[0], items[1]
			if key != "" && mime != "" {
				return object.ChangeMimeApiInfo{
					Bucket: info.Bucket,
					Key:    key,
					Mime:   mime,
				}, true
			}
		}
		return nil, true
	}).OnResult(func(operation batch.Operation, result batch.OperationResult) {
		apiInfo, ok := (operation).(object.ChangeMimeApiInfo)
		if !ok {
			return
		}
		in := ChangeMimeInfo(apiInfo)
		if result.Code != 200 || result.Error != "" {
			handler.Export().Fail().ExportF("%s\t%s\t%d\t%s\n", in.Key, in.Mime, result.Code, result.Error)
			log.ErrorF("Chgm '%s' => '%s' Failed, Code: %d, Error: %s\n",
				in.Key, in.Mime, result.Code, result.Error)
		} else {
			handler.Export().Success().ExportF("%s\t%s\n", in.Key, in.Mime)
			log.ErrorF("Chgm '%s' => '%s' success\n", in.Key, in.Mime)
		}
	}).OnError(func(err error) {
		log.ErrorF("batch chgm error:%v:", err)
	}).Start()
}
