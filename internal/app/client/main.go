package main

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/oklog/run"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var (
		localG  run.Group
		signals = []os.Signal{
			os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
			syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
			syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
		}
	)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	ctxAll, cancelAll := context.WithCancel(context.TODO())

	//监听系统信号
	{
		localG.Add(func() error {
			for {
				select {
				case <-ch:
					{
						glog.Error(ctxAll, "recv signal exist", <-ch)
						return gerror.New("recv signal exist")

					}
				case <-ctxAll.Done():
					{
						glog.Error(ctxAll, "recv other signal exist")
						return gerror.New("recv other signal exist")

					}
				}
			}
		}, func(err error) {
			cancelAll()
		})
	}
	{
		localG.Add(func() error {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:

				case <-ctxAll.Done():
					return gerror.New("Context Done")

				}
			}

			return nil
		}, func(err error) {
			cancelAll()
		})
	}
	localG.Run()
}
