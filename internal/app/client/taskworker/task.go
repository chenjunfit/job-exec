package taskworker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/toolkits/pkg/sys"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

type Task struct {
	sync.Mutex

	JobId   int64
	Id      int64
	Clock   int64
	Action  string
	Status  string
	MetaDir string

	Alive  bool
	Cmd    *exec.Cmd
	StdOut bytes.Buffer
	StdErr bytes.Buffer

	Scripts string
	Args    []string
	Account string
}

func (t *Task) GetStdOut() string {
	t.Mutex.Lock()
	bufferStr := t.StdOut.String()
	t.Mutex.Unlock()
	return bufferStr
}

func (t *Task) GetStdErr() string {
	t.Mutex.Lock()
	bufferStr := t.StdErr.String()
	t.Mutex.Unlock()
	return bufferStr
}

func (t *Task) GetStatus() string {
	t.Mutex.Lock()
	status := t.Status
	t.Mutex.Unlock()
	return status
}
func (t *Task) SetStatus(status string) {
	t.Mutex.Lock()
	t.Status = status
	t.Mutex.Unlock()
}

func (t *Task) GetAlive() bool {
	t.Mutex.Lock()
	alive := t.Alive
	t.Mutex.Unlock()
	return alive
}

func (t *Task) SetAlive(Alive bool) {
	t.Mutex.Lock()
	t.Alive = Alive
	t.Mutex.Unlock()

}

func (t *Task) RestBuff() {
	t.Mutex.Lock()
	t.StdOut.Reset()
	t.StdErr.Reset()
	t.Mutex.Unlock()

}

func (t *Task) DoneBefore() bool {
	doneFlag := path.Join(t.MetaDir, fmt.Sprint(t.Id), fmt.Sprintf("%d.done", t.Clock))
	return fileutil.Exist(doneFlag)
}

func (t *Task) LoadResult() {
	metaDir := t.MetaDir

	doneFlag := path.Join(metaDir, fmt.Sprint(t.Id), fmt.Sprintf("%d.done", t.Clock))
	stdOutFile := path.Join(metaDir, fmt.Sprint(t.Id), "stdout")
	stdErrFile := path.Join(metaDir, fmt.Sprint(t.Id), "stderr")

	t.Status = gfile.GetContents(doneFlag)
	stdOut := gfile.GetContents(stdOutFile)
	stdErr := gfile.GetContents(stdErrFile)

	t.StdOut = *bytes.NewBufferString(stdOut)
	t.StdErr = *bytes.NewBufferString(stdErr)
}

func (t *Task) Prepare() error {
	IdDir := path.Join(t.MetaDir, fmt.Sprint(t.Id))
	err := gfile.Mkdir(IdDir)
	if err != nil {
		glog.Error(context.Background(), "mkdir -p %s", IdDir, "failed")
		return err
	}

	writeFlag := path.Join(t.MetaDir, ".write")

	if gfile.Exists(writeFlag) {
		argsFile := path.Join(IdDir, "args")
		argsByte := gfile.GetBytes(argsFile)
		args := make([]string, 0)
		err = gjson.Unmarshal(argsByte, args)
		if err != nil {
			glog.Error(context.TODO(), "err: ", err)
		}
		accountFile := path.Join(IdDir, "account")
		account := gfile.GetContents(accountFile)
		scriptFile := path.Join(IdDir, "script")
		scripts := gfile.GetContents(scriptFile)
		t.Args = args
		t.Account = account
		t.Scripts = scripts

	} else {
		args, account, script := t.Args, t.Account, t.Scripts
		scriptFile := path.Join(IdDir, "script")
		err := gfile.PutContents(scriptFile, script)
		if err != nil {
			glog.Error(context.Background(), "write script to %s", scriptFile, "failed")
			return err
		}
		err = gfile.Chmod(scriptFile, os.ModePerm)
		if err != nil {
			glog.Error(context.Background(), "chmod +x %s", scriptFile, "failed")
			return err
		}
		argsFile := path.Join(IdDir, "args")
		argsByte, err := gjson.Marshal(args)
		if err != nil {
			glog.Error(context.Background(), "write args to %s", argsFile, "failed")
			return err
		}
		gfile.PutBytes(argsFile, argsByte)
		if err != nil {
			glog.Error(context.Background(), "write args to %s", argsFile, "failed")
			return err
		}

		accountFile := path.Join(IdDir, "account")
		err = gfile.PutContents(accountFile, account)
		if err != nil {
			glog.Error(context.Background(), "write account to %s", accountFile, "failed")
			return err
		}
	}
	return nil
}

func (t *Task) Start() {
	if t.GetAlive() {
		return
	}
	err := t.Prepare()
	if err != nil {
		glog.Error(context.TODO(), "prepare faild err: ", err)
		return
	}

	nowPath, _ := os.Getwd()
	scriptPath := path.Join(nowPath, t.MetaDir, fmt.Sprint(t.Id), "script")

	sh := fmt.Sprintf("%s", scriptPath)
	if len(t.Args) > 0 {
		for _, value := range t.Args {
			sh = fmt.Sprintf("%s %s", sh, value)
		}
	}

	var cmd *exec.Cmd
	if t.Account == "root" {
		cmd = exec.Command("sh", "-c", sh)
		cmd.Dir = "/tmp"
	} else {
		cmd = exec.Command("su", "-c", sh, "-", t.Account)
	}

	cmd.Stdout = &t.StdOut
	cmd.Stderr = &t.StdErr
	t.Cmd = cmd
	err = cmd.Start()

	if err != nil {
		glog.Error(context.TODO(), "start cmd failed: ", err)
		return
	}
	go runProcess(t)
}
func (t *Task) Kill() {
	go killProcess(t)
}
func runProcess(t *Task) {
	t.SetAlive(true)
	defer t.SetAlive(false)
	err := t.Cmd.Wait()
	if err != nil {
		if strings.Contains(err.Error(), "signal: killed") {
			t.SetStatus("killed")
		} else {
			t.SetStatus("failed")
		}
	} else {
		t.SetStatus("successed")
	}
	persistResult(t)
}
func persistResult(t *Task) {
	metadir := t.MetaDir

	stdout := path.Join(metadir, fmt.Sprint(t.Id), "stdout")
	stderr := path.Join(metadir, fmt.Sprint(t.Id), "stderr")
	doneFlag := path.Join(metadir, fmt.Sprint(t.Id), fmt.Sprintf("%d.done", t.Clock))

	gfile.PutContents(stdout, t.GetStdOut())
	gfile.PutContents(stderr, t.GetStdErr())
	gfile.PutContents(doneFlag, t.GetStatus())
}
func killProcess(t *Task) {
	t.SetAlive(true)
	defer t.SetAlive(false)
	err := KillProcessByTaskId(t.Id, t.MetaDir)
	if err != nil {
		t.SetStatus("killedFailed")
		glog.Error(context.TODO(), "kill failed: ", err)
	} else {
		t.SetStatus("killed")
		glog.Info(context.TODO(), "kill %d: ", t.Id)
	}
	persistResult(t)
}
func KillProcessByTaskId(id int64, metaDir string) error {
	dir := strings.TrimRight(metaDir, "/")
	arr := strings.Split(dir, "/")
	lst := arr[len(arr)-1]
	return sys.KillProcessByCmdline(fmt.Sprintf("%s/%d/script", lst, id))
}
