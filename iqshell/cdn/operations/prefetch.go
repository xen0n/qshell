package operations

import (
	"github.com/qiniu/qshell/v2/iqshell"
	"github.com/qiniu/qshell/v2/iqshell/cdn"
	"github.com/qiniu/qshell/v2/iqshell/common/group"
	"github.com/qiniu/qshell/v2/iqshell/common/log"
	"strings"
)

type PrefetchInfo struct {
	UrlListFile string // url 信息文件
	SizeLimit   int    // 每次刷新最大 size 限制
	QpsLimit    int    // qps 限制
}

func (info *PrefetchInfo) Check() error {
	return nil
}

func Prefetch(cfg *iqshell.Config, info PrefetchInfo) {
	if shouldContinue := iqshell.CheckAndLoad(cfg, iqshell.CheckAndLoadInfo{
		Checker: &info,
	}); !shouldContinue {
		return
	}

	log.DebugF("qps limit: %d, max item-size: %d", info.QpsLimit, info.SizeLimit)

	handler, err := group.NewHandler(group.Info{
		InputFile:              info.UrlListFile,
		Force:                  true,
	})
	if err != nil {
		log.Error(err)
		return
	}

	createQpsLimitIfNeeded(info.QpsLimit)

	line := ""
	hasMore := false
	urlsToPrefetch := make([]string, 0, 50)
	for {
		line, hasMore = handler.Scanner().ScanLine()
		if !hasMore {
			break
		}

		url := strings.TrimSpace(line)
		if url == "" {
			continue
		}
		urlsToPrefetch = append(urlsToPrefetch, url)

		if len(urlsToPrefetch) == cdn.BatchPrefetchAllowMax ||
			(info.SizeLimit > 0 && len(urlsToPrefetch) >= info.SizeLimit) {
			prefetchWithQps(urlsToPrefetch)
			urlsToPrefetch = make([]string, 0, 50)
		}
	}

	if len(urlsToPrefetch) > 0 {
		prefetchWithQps(urlsToPrefetch)
	}
}

func prefetchWithQps(urlsToPrefetch []string) {

	waiterIfNeeded()

	log.Debug("cdnPrefetch, url size: %d", len(urlsToPrefetch))
	if len(urlsToPrefetch) > 0 {
		err := cdn.Prefetch(urlsToPrefetch)
		if err != nil {
			log.Error(err)
		}
	}
}
