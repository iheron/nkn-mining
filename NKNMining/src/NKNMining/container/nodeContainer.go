package container

import (
	"NKNMining/common"
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var Node = &NodeContainer{}
var NodeStatus = &NodeContainer{}
var CmdApp = &NodeContainer{}

var manuallyStopped = true

func InitNodeContainers() {
	shellWorkPath := common.GetCurrentDirectory()
	nodeWorkPath := shellWorkPath + "/bin"
	nodeApp := nodeWorkPath + "/nknd"
	nodeVersionApp := nodeWorkPath + "/nknc"
	cmdApp := "sh"
	if common.IsWindowsOS() {
		nodeApp += ".exe"
		nodeVersionApp += ".exe"
	}

	if !Node.InitEnvironment(nodeApp, nodeWorkPath) {
		common.Log.Fatal("initialization of NKN node failed!")
	}

	if !NodeStatus.InitEnvironment(nodeVersionApp, nodeWorkPath) {
		common.Log.Fatal("initialization of NKN node status getter failed!")
	}

	if !CmdApp.InitEnvironment(cmdApp, nodeWorkPath) {
		common.Log.Fatal("initialization of cmd app failed!")
	}
}

type NodeContainer struct {
	controller *exec.Cmd
	path       string
	workPath   string
	mutex      *sync.Mutex
}

func (c *NodeContainer) InitEnvironment(path string, workPath string) bool {
	//if !common.FileExist(path) {
	//	return false
	//}

	c.path = path
	c.workPath = workPath
	c.mutex = new(sync.Mutex)

	return true
}

func (c *NodeContainer) buildCommand(args []string, input string) (cmd *exec.Cmd, err error) {
	if nil != c.controller {
		cmd = c.controller
		return
	}
	c.controller = exec.Command(c.path, args...)
	c.controller.Dir = c.workPath

	stdin, err := c.controller.StdinPipe()
	if nil != err {
		return
	}

	_, err = stdin.Write([]byte(input + "\r\n"))
	if nil != err {
		common.Log.Error(err)
		return
	}

	cmd = c.controller
	return
}

func (c *NodeContainer) SyncRun(args []string, input string) (output string, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cmd, err := c.buildCommand(args, input)
	if nil != err {
		return "", err
	}

	outBuf := bytes.NewBuffer(nil)
	cmd.Stdout = outBuf
	cmd.Run()

	c.controller = nil

	if 0 == outBuf.Len() {
		output = "UNKNOWN"
	} else {
		outBuf.Truncate(outBuf.Len() - 1)
		output = outBuf.String()
	}

	return
}

func (c *NodeContainer) AsyncRun(args []string, input string) (cmd *exec.Cmd, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	common.Log.Error("async run")
	manuallyStopped = false
	cmd, err = c.buildCommand(args, input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if nil != err {
		return
	}

	go func() {
		for {
			common.Log.Error("run start")
			cmd.Run()
			cmd.Wait()
			fmt.Println(string(stdout.Bytes()))
			fmt.Println(string(stderr.Bytes()))
			common.Log.Error("run end")

			c.mutex.Lock()
			if nil == c.controller {
				c.mutex.Unlock()
				common.Log.Error("clean upped")
				break
			}

			common.Log.Error("clean up")

			c.controller = nil
			cmd, _ = c.buildCommand(args, input)
			time.Sleep(time.Second)
			c.mutex.Unlock()
		}
	}()

	return
}

func (c *NodeContainer) Stop() {
	c.mutex.Lock()
	defer func() {
		c.controller = nil
		c.mutex.Unlock()
	}()

	manuallyStopped = true
	if nil == c.controller {
		return
	}

	if nil != c.controller.ProcessState {
		return
	}

	if nil == c.controller.Process {
		return
	}

	c.controller.Process.Kill()
	c.controller.Process.Wait()
}

func (c *NodeContainer) Status() int {
	if manuallyStopped {
		return common.NKN_NODE_EXITED
	} else {
		return common.NKN_NODE_RUNNING
	}
}
