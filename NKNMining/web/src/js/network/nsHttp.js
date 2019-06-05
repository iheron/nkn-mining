import Axios from "axios"
import Is from "is_js"
import {RequestData} from "./nsRequestData";

function axiosRequest(scope, apiURI, param, success, fail) {
  Axios.post(nknShellWebConfig.baseURI + apiURI, param, {
    headers: { 'Content-Type': 'text/plain' }
  }).then(function (response) {
      if(Is.function(success)) {
        success.call(scope, response.data)
      }
  }).catch(function (response) {
    if(Is.function(fail)) {
      fail.call(scope, response)
    }
  })

  return true
}

function getVersion(scope, success = null, fail = null) {
  return axiosRequest(scope, "/version", {}, success, fail)
}

function getStatus(scope, success = null, fail = null) {
  return axiosRequest(scope, "/status", {}, success, fail)
}

function setWallet(scope, walletInfo, reqKey, success, fail) {
  let setWalletData = new RequestData()

  setWalletData.setData(walletInfo)
  return axiosRequest(scope, '/set/wallet', setWalletData.encryptedData(reqKey), success, fail)
}

function getWalletConfig(scope, success = null, fail = null) {
  return axiosRequest(scope, "/get/wallet/config", {}, success, fail)
}

function setWalletConfig(scope, walletInfo, reqKey, success, fail) {
  let setWalletData = new RequestData()

  setWalletData.setData(walletInfo)
  return axiosRequest(scope, '/set/wallet/config', setWalletData.encryptedData(reqKey), success, fail)
}

function startMining(scope, reqKey, success, fail) {
  let startMiningData = new RequestData()

  startMiningData.setData("start")
  return axiosRequest(scope, '/start/node', startMiningData.encryptedData(reqKey), success, fail)
}

function stopMining(scope, reqKey, success, fail) {
  let stopMiningData = new RequestData()

  stopMiningData.setData("stop")
  return axiosRequest(scope, '/stop/node', stopMiningData.encryptedData(reqKey), success, fail)
}

function resetShell(scope, reqKey, success, fail) {
  let resetShellData = new RequestData()

 resetShellData.setData("resetShell")
  return axiosRequest(scope, '/reset/shell', resetShellData.encryptedData(reqKey), success, fail)
}

function getRewardsList(scope, address, height, success, fail) {
  return Axios.get(nknShellWebConfig.nknDataCenterURL + '/mining/rewards/' + address + '/' + height, [], {
    headers: { 'Content-Type': 'text/plain' }
  }).then(function (response) {
    if(Is.function(success)) {
      success.call(scope, response.data)
    }
  }).catch(function (response) {
    if(Is.function(fail)) {
      fail.call(scope, response)
    }
  })

}

function login(scope, data, pwd, success, fail) {
  let loginData = new RequestData()

  loginData.setData({
    Nonce: data.nonce,
    RandomKey: data.randomKey,
  })
  loginData.setPlainData(data.nonce)

  return axiosRequest(scope, '/login', loginData.fullData(pwd), success, fail)
}


export default {
  getStatus,
  getVersion,
  setWallet,
  getWalletConfig,
  setWalletConfig,
  startMining,
  stopMining,
  resetShell,
  login,
  getRewardsList,
}