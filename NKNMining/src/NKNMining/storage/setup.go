package storage

import (
	"io/ioutil"
	"NKNMining/common"
	"bytes"
	"encoding/json"
	// "github.com/pborman/uuid"
	"fmt"
	"os"
	"sync"
)

const (
	setupFile = "setup/info.json"
)

const (
	SETUP_STEP_GEN_WALLET = iota
	SETUP_STEP_SUCCESS
	SETUP_NODE_UPDATE
)

var IsRemote = false
var saveMutex = &sync.Mutex{}

func InitSetupInfo()  {
	NKNSetupInfo.Load()

	// if SETUP_STEP_GEN_WALLET == NKNSetupInfo.CurrentStep {
	// 	if "" == NKNSetupInfo.SerialNumber {
	// 		sn := "NKN-" + uuid.NewUUID().String()
	// 		NKNSetupInfo.SerialNumber = sn
	// 	}

	// 	if nil == NKNSetupInfo.Save() {
	// 		if IsRemote {
	// 			fmt.Println("serial number(sn): " + NKNSetupInfo.SerialNumber)
	// 		}
	// 	} else {
	// 		common.Log.Fatal("initialization NKN setup info failed!")
	// 	}
	// 	return
	// } else {
	// 	if common.IsWindowsOS() {
	// 		fmt.Println("NKNMining start up")
	// 	} else {
	// 		fmt.Println("NKNMining start up in background")
	// 	}
	// 	return
	// }

	fmt.Println("NKNMining start up in background")
}

type SetupInfo struct {
	SerialNumber string
	CurrentStep	int
	Key string
	WKey string
	BinVersion string
	SelfNode string
}

var NKNSetupInfo = &SetupInfo{
	CurrentStep: 0,
	BinVersion: "",
	SelfNode: "http://127.0.0.1:30003",
}

func (s *SetupInfo) Reset()  {
	NKNSetupInfo = &SetupInfo{SerialNumber: NKNSetupInfo.SerialNumber, CurrentStep: SETUP_STEP_GEN_WALLET}
	NKNSetupInfo.Save()
}

func (s *SetupInfo) GetRequestKey() string {
	return s.SerialNumber
}


func (s *SetupInfo) GetWalletKey() string {
	return s.SerialNumber + s.Key
}

func (s *SetupInfo) Load()  {
	setupInfo, err := ioutil.ReadFile(setupFile)
	if nil != err {
		s.Save()
		common.Log.Error("no setup info! use default setup param")
		return
	}

	setupInfo = bytes.TrimPrefix(setupInfo, []byte("\xef\xbb\xbf"))
	if 0 == len(setupInfo) {
		s.Save()
		common.Log.Error("blank setup info! use default setup param")
		return
	}

	err = json.Unmarshal(setupInfo, s)
	if nil != err {
		s.Save()
		common.Log.Error("setup file damaged! use default setup param")
		return
	}
}

func (s *SetupInfo) Save() error {
	saveMutex.Lock()
	defer func() {
		saveMutex.Unlock()
	}()
	setupInfo, _ := json.MarshalIndent(s, "", "    ")

	err := ioutil.WriteFile(setupFile, setupInfo, 0666)
	if nil != err {
		common.Log.Fatal("save setup file failed!")
		os.Exit(-3)
	}

	return err
}