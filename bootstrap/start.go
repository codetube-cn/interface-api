package bootstrap

import (
	"codetube.cn/interface-api/components"
	"codetube.cn/interface-api/routes/v1"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var BootErrChan chan error

func Start() {
	BootErrChan = make(chan error)
	go func() {
		//加载各版本的 API 路由
		v1.ApiRouter.Load(v1.LoadRoutes...)

		components.RouterEngine.Run("0.0.0.0:8080")
	}()
	//监听事件
	go func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		BootErrChan <- fmt.Errorf("%", <-sigC)
	}()
	getErr := <-BootErrChan
	log.Println(getErr)
}
