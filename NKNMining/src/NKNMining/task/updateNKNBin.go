package task

import (
	"NKNMining/common"
	"NKNMining/container"
	"NKNMining/crypto"
	"NKNMining/network/nknReleaseQuery"
	"NKNMining/status"
	"NKNMining/storage"
	"github.com/mholt/archiver"
	"os"
	"runtime"
	"time"
)

const (
	nkn_bin_file_path = "/bin-src"
)

var nknBinFirstUpdate = true

func nknBinNeedUpdate(version string) bool {
	if version == "" && common.NknBinExists() {
		return true
	}

	return false
}

func mvNKNBin(from string, to string) error {
	return os.Rename(from, to)
}

func doBinUpdate(toVersion string, url string) (err error) {
	needRestart := true
	initialization := nknBinFirstUpdate
	currentStep := storage.NKNSetupInfo.CurrentStep

	if common.NknBinExists() {
		status.SetBinDownloaded()
	}

	if currentStep == storage.SETUP_NODE_UPDATE && !nknBinFirstUpdate {
		if !nknBinNeedUpdate(toVersion) {
			storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
			storage.NKNSetupInfo.Save()
		}
		return
	}

	if common.NknBinExists() {
		if nknBinFirstUpdate && storage.SETUP_NODE_UPDATE == storage.NKNSetupInfo.CurrentStep {
			currentStep = storage.SETUP_STEP_SUCCESS
		}

		binRunStatus, errInfo := status.GetServerStatus()

		if !nknBinNeedUpdate(toVersion) {
			if storage.NKNSetupInfo.CurrentStep == storage.SETUP_STEP_GEN_WALLET {
				return
			}
			currentStep = storage.SETUP_STEP_SUCCESS
			return
		}

		if common.NS_STATUS_GEN_WALLET == binRunStatus ||
			"" != errInfo {
			return
		}

		if common.NS_STATUS_NODE_RUNNING == binRunStatus {
			needRestart = true
			container.Node.Stop()
		}
	}

	if storage.SETUP_STEP_GEN_WALLET != currentStep {
		storage.NKNSetupInfo.CurrentStep = storage.SETUP_NODE_UPDATE
		storage.NKNSetupInfo.Save()

		defer func() {
			storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
			storage.NKNSetupInfo.Save()
		}()
	}

	basicPath := common.GetCurrentDirectory() + nkn_bin_file_path
	unzippedBin := basicPath + "/" + runtime.GOOS + "-" + runtime.GOARCH + "/nknd"
	unzippedNKNCBin := basicPath + "/" + runtime.GOOS + "-" + runtime.GOARCH + "/nknc"
	fileName := runtime.GOOS + "-" + runtime.GOARCH + "." + toVersion + ".zip"
	fullName := basicPath + "/" + fileName
	err = nknReleaseQuery.DownloadNKN(url, fullName, nil)
	if nil != err {
		return
	}
	err = archiver.Zip.Open(fullName, common.GetCurrentDirectory()+nkn_bin_file_path)

	if nil != err {
		common.Log.Error("unzip bin file failed: ", err)
		return
	}

	if common.IsWindowsOS() {
		err = mvNKNBin(unzippedBin, basicPath+"/nknd.exe")
		err = mvNKNBin(unzippedNKNCBin, basicPath+"/nknc.exe")
	} else {
		err = mvNKNBin(unzippedBin, basicPath+"/nknd")
		err = mvNKNBin(unzippedNKNCBin, basicPath+"/nknc")
	}
	if nil != err {
		common.Log.Error("move bin file failed: ", err)
		return
	}
	nknBinFirstUpdate = false

	if initialization {
		status.SetBinDownloaded()
	}

	storage.NKNSetupInfo.BinVersion = toVersion
	storage.NKNSetupInfo.Save()
	if needRestart {
		wKey, err := crypto.AesDecrypt(storage.NKNSetupInfo.WKey, storage.NKNSetupInfo.GetWalletKey())
		if nil != err {
			common.Log.Error(err)
			return err
		}

		_, err = container.Node.AsyncRun([]string{"--no-check-port"}, wKey)
		if nil != err {
			common.Log.Error(err)
			return err
		}
	}

	return

}

func UpdateNKNBin() {
	for {
		version, url, err := nknReleaseQuery.LastVersion()
		if nil != err {
			common.Log.Error(err)
		} else {
			err = doBinUpdate(version, url)
			if nil != err {
				time.Sleep(5 * time.Second)
				continue
			}
		}

		time.Sleep(120 * time.Second)
	}
}
