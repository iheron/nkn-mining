import {nsNamespace} from "./nsNamespace"
import {serverStatus} from "./nsConst"
import NSLocalStorage from "./nsLocalStorage"
import Http from "./network/nsHttp"

let goToPage = function(currentPage, targetPage, $router) {
  if(currentPage !== targetPage) {
    $router.push({name: targetPage})
  }
}

let loadPage = function () {
  let currentPage = this.$route.name
  let $router = this.$router

  Http.getStatus(this, function (data) {
    this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)

    if(data.Data.sn) {
      NSLocalStorage.setReqKey(data.Data.sn)
      this.showSN = false
    } else {
      if(!NSLocalStorage.checkInit()) {
        this.showSN = true
      }
    }

    switch(data.Data.shellStatus) {
      case serverStatus.NS_STATUS_GEN_WALLET():
        switch (currentPage) {
          case nsNamespace.SETUP.GEN_WALLET:
            goToPage(currentPage, nsNamespace.SIGN_IN, $router)
            break

          case nsNamespace.SETUP.LOAD_WALLET:
            goToPage(currentPage, nsNamespace.SETUP.LOAD_WALLET, $router)
            break

          default:
            goToPage(currentPage, nsNamespace.SIGN_IN, $router)
            break
        }
        break

      case serverStatus.NS_STATUS_NODE_RUNNING():
        if(!NSLocalStorage.checkLogin()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", true)
          goToPage(currentPage, nsNamespace.MAIN, $router)
        }
        break

      case serverStatus.NS_STATUS_NODE_EXITED():
      case serverStatus.NS_STATUS_INITIALIZATION():
      case serverStatus.NS_STATUS_UPDATE_BIN():
        if(!NSLocalStorage.checkLogin()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", false)
          goToPage(currentPage, nsNamespace.MAIN, $router)
        }
        break

      default:
        alert("invalid server!")
        break
    }
  }, function (e) {
    console.log(e)
    alert("connect to server failed!")
  })
}

export {
  loadPage
}