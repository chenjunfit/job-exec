package main

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/oklog/run"
	v1 "job-exec/api/taskreport/v1"
	"job-exec/internal/app/client/taskreport"
	"job-exec/internal/app/client/taskworker"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	taskworker.InitLocals("./meta")
	var (
		Client v1.ServiceClient
	)
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	conn := grpcx.Client.MustNewGrpcClientConn("job-exec")
	Client = v1.NewServiceClient(conn)
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
	//报告任务
	{
		localG.Add(func() error {
			ticker := time.NewTicker(20 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					report := taskreport.Report{}
					err := report.ReportManager(ctxAll, Client)
					if err != nil {
						glog.Error(ctxAll, err)
					}
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
