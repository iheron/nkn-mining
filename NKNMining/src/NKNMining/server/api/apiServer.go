package apiServer

import (
	. "NKNMining/server/api/actions"
	"NKNMining/server/api/const"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	apiList = []IRestfulAPIAction{
		GetStatusAPI,
		GetVersionAPI,
		SetWalletAPI,
		SetWalletConfigAPI,
		GetWalletConfigAPI,
		GetWalletAPI,
		StartNodeAPI,
		StopNodeAPI,
		ResetShellAPI,
		LoginAPI,
		GetSNAPI,
	}
)

func InitRouters(router *gin.Engine, baseURI string) {
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
				c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
			}
			c.AbortWithStatus(http.StatusOK)
		}
	})

	router.Use(func(c *gin.Context) {
		if strings.Contains(c.Request.RemoteAddr, "127.0.0.1") {
			c.Params = append(c.Params, gin.Param{Key: apiServerConsts.PARAM_FROM_LOCAL, Value: "true"})
		}
	})

	setApiRouters(router, baseURI)
}

func setApiRouters(router *gin.Engine, baseURI string) {
	for _, v := range apiList {
		router.Any(v.URI(baseURI), v.Action)
	}
}
