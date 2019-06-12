package main

import (
	"NKNMining/cli"
	"NKNMining/common"
	"NKNMining/config"
	"NKNMining/container"
	"NKNMining/network/rpcRequest"
	"NKNMining/server"
	"NKNMining/storage"
	"NKNMining/task"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func startDaemon() (hasNewProcess bool) {
	defer func() {
		if !hasNewProcess {
			storage.Temp.SaveParentPid(-1)
		}
	}()

	ppid := os.Getppid()
	needPPidCheck := true

	if nil != storage.Temp.Load() {
		needPPidCheck = false
	}

	if ppid != 1 {
		if needPPidCheck && ppid == storage.Temp.LastPPid {
			return false
		}
		storage.Temp.SaveParentPid(ppid)

		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

		return true
	}

	return false
}

func getCliApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "NKN Mining"
	app.Version = "0.2.0"
	app.HelpName = "NKNMining"
	app.Usage = "NKN Mining application"
	app.UsageText = "NKNMining [options] [args]"
	app.HideHelp = false
	app.HideVersion = false

	return
}

func main() {
	//get app
	shell := getCliApp()

	//set some flags
	commands.SetFlags(shell)
	commands.SetAction(shell, &storage.IsRemote)

	//run
	err := shell.Run(os.Args)
	if nil != err {
		os.Exit(0)
	}

	//init logs
	common.InitLog(config.ShellConfig.LogFile)

	//start daemon mode in the os other then windows
	if !common.IsWindowsOS() {
		if startDaemon() {
			return
		}
	}
	container.InitNodeContainers()
	storage.InitSetupInfo()
	storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
	if common.ShellMutexCheck() {
		common.Log.Error("NKNMining is already running！")
		return
	}

	rpcRequest.Api.Build()

	task.StartAllTask()

	//start server
	server.Start()

	for {
		time.Sleep(time.Second)
	}
}
