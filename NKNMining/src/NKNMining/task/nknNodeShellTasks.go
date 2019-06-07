package task

func StartAllTask() {

	go UpdateNodeState()
	go UpdateNodeNeighbor()
	//go UpdateNKNBin()
	go UpdateCurrentHeight()
	go UpdateNetworkHeight()
	//go ChainDataDownloading()

}
