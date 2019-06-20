package task

import (
	"NKNMining/common"
	"NKNMining/network/nknReleaseQuery"
	"NKNMining/status"
	"github.com/mholt/archiver"
	"os"
	"time"
)

const (
	nkn_chain_data_path = "/bin"
)

var chainDataDownloadProgress = 0.0

const fullPreChainDataLen = 771130143.0

var downloadedChainData = 0.0

func GetChainDataDownloadProgress() float64 {
	return chainDataDownloadProgress
}

func updateChainDataDownloadProgress() {
	for {
		if status.IsInitFinished() {
			chainDataDownloadProgress = 1
			return
		}

		chainDataDownloadProgress = downloadedChainData / fullPreChainDataLen
		time.Sleep(time.Second)
	}
}

func downloadChainData() (err error) {
	basicPath := common.GetCurrentDirectory() + nkn_chain_data_path

	if common.FileExist(basicPath + "/Chain") {
		status.SetChainDataDownloaded()
		return
	}

	fileName := "Chain_182381.zip"
	fullName := basicPath + "/" + fileName
	err = nknReleaseQuery.DownloadNKN("https://storage.googleapis.com/nkn-testnet-snapshot/Chain_182381.zip",
		fullName, &downloadedChainData)
	if nil != err {
		return
	}

	err = archiver.Zip.Open(fullName, common.GetCurrentDirectory()+nkn_chain_data_path)

	if nil != err {
		common.Log.Error("unzip bin file failed: ", err)
		return
	}

	time.Sleep(time.Second)
	os.Remove(fullName)

	status.SetChainDataDownloaded()

	return

}

func ChainDataDownloading() {
	go updateChainDataDownloadProgress()

	for {
		downloaded := downloadChainData()
		if nil != downloaded {
			common.Log.Error(downloaded)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}
