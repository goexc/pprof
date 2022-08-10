package pprof

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"sync"
)

var (
	once sync.Once
)

func StartAgent(host string, port int) {
	if len(host) == 0 {
		return
	}

	once.Do(func() {
		threading.GoSafe(func() {
			addr := fmt.Sprintf("%s:%d", host, port)
			logx.Infof("Starting pprof agent at %s", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				logx.Error(err)
			}
		})
	})
}
