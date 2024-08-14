package taskreport

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "job-exec/api/taskreport/v1"
	"job-exec/internal/app/client/taskworker"
	"net"
)

//组装taskresult定时发送个服务端
//返回的的任务调用任务解析
//链接服务
//组装本地完成的taskresult，发送服务端
//返回的taskmeta,调用本地任务处理

type Report struct {
}

func (r *Report) GetIP() string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		glog.Error(context.TODO(), error)
	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr).IP.String()
	return ipAddress
}

func (r *Report) ReportManager(ctx context.Context, client v1.ServiceClient) error {

	res, err := r.DoReport(ctx, client)
	if err != nil {
		glog.Error(ctx, "report err: ", err)
		return err
	}
	if len(res.AssignTasks) > 0 {
		for _, task := range res.AssignTasks {
			taskworker.Locals.AssignTask(task)
		}
	}

	return nil
}
func (r *Report) DoReport(ctx context.Context, client v1.ServiceClient) (res *v1.TaskReportRes, err error) {
	req := &v1.TaskReportReq{
		Results: taskworker.Locals.ReportTasks(),
		AgentIp: r.GetIP(),
	}
	res, err = client.TaskReport(context.TODO(), req)
	if err != nil {
		glog.Error(ctx, "msg: ", err)
	}
	return res, err
}
