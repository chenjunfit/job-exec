package main

import (
	"context"
	_ "job-exec/internal/app/server/logic/jobexec"
	"job-exec/internal/app/server/logic/tasksync"
	_ "job-exec/internal/boot"
	"job-exec/internal/cmd"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/oklog/run"
)

func main() {
	var (
		localG  run.Group
		t       = new(tasksync.TaskSync)
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
	//http服务
	{
		localG.Add(func() error {
			for {
				err := cmd.Http.RunWithError(gctx.New())
				if err != nil {
					return err
				}
				select {
				case <-ctxAll.Done():
					{
						glog.Error(ctxAll, "context done ")
						return gerror.New("context done")
					}

				}

			}

		}, func(err error) {
			cancelAll()
		})
	}
	//rpc服务
	{
		localG.Add(func() error {
			for {
				err := cmd.Rpc.RunWithError(ctxAll)
				if err != nil {
					return err
				}
				select {
				case <-ctxAll.Done():
					{
						glog.Error(ctxAll, "Context Done")
						return gerror.New("Context Done")
					}
				}
			}
		}, func(err error) {
			cancelAll()
		})
	}

	//读取任务到内存
	{
		localG.Add(func() error {
			err := t.SyncManager(ctxAll)
			if err != nil {
				glog.Error(ctxAll, "msg: ", err)
			}
			return err
		}, func(err error) {
			cancelAll()
		})
	}
	localG.Run()

}
