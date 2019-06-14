package apiServerAction

import (
	"NKNMining/container"
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
)

var GetBalanceAPI IRestfulAPIAction = &getBalance{}

type getBalance struct {
	restfulAPIBase
}

func (g *getBalance) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/get/balance"
}

func (g *getBalance) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)
	balance := "0"

	balanceRes, err := container.CmdApp.SyncRun([]string{"-c", "cat ./wallet.pswd | ./nknc wallet -l balance"}, "")

	if nil != err {
		log.Println(err)
		balance = "0"
	}

	m := make(map[string]interface{})
	reg, _ := regexp.Compile(`\{.*\}`)
	nReg, _ := regexp.Compile(`\n`)

	err = json.Unmarshal([]byte(reg.FindString(nReg.ReplaceAllString(balanceRes, ""))), &m)
	if err != nil {
		log.Println("Umarshal failed:", err)
		balance = "0"
	}

	response.Success(map[string]interface{}{
		"Balance": balance,
	})

	return
}
