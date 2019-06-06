<template>
    <div class="ns-create-account nkn-card-shadow nkn-after-clear">
        <div class="nkn-sign-in-panel">
            <label class="nkn-page-title-label">{{$t("nsSignIn.titleLabel")}}</label>
            <h1 class="nkn-page-title text-main-blue">{{$t("nsSignIn.title")}}</h1>
            <div v-for="(inputItem, idx) in inputs" :key="idx">
                <ns-input-item v-if="showInput(idx)" :config="inputItem" />
            </div>
            <div class="nkn-sign-in">
                <button class="nkn-normal-btn sign-in-button" type="button" @click="signIn">{{$t("nsSignIn.signInbtn")}}</button>
            </div>
        </div>
        <div class="nkn-setup-page-wallpaper">
            <img class="nkn-wall-background" :src="'./static/img/wallpaper.png'"/>
            <img class="nkn-wall-pad" :src="'./static/img/wallpad.png'"/>
        </div>
        <div style="display: none">{{lang}}</div>
        <ns-loading v-if="!this.$store.state.global.pageLoaded"/>
    </div>
</template>

<script>
  import NsInputItem from "../components/nsInputItem"
  import NsLoading from "../components/nsLoading"
  import NSLocalStorage from "../js/nsLocalStorage"
  import Crypto from "../js/crypto/algorithm"
  import Mathjs from "mathjs"
  import Http from "../js/network/nsHttp"
  import LangMix from "../js/mixin/lang.js"
  import {inputIdPrefix} from "../js/nsConst"
  import {nsNamespace} from "../js/nsNamespace"
  import {loadPage} from "../js/nsLoading"

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    mixins:[LangMix],
    name: "ns-sign-in",
    mounted() {
      if(NSLocalStorage.checkLogin()) {
        this.$nextTick(function () {
          this.$router.push({name: nsNamespace.MAIN})
        })
      } else {
        loadPage.call(this)
      }
    },
    methods: {
      showInput(idx) {
        if('sn' === idx) {
          return this.showSN
        }

        return true
      },
      array2HexString(bytes) {
        return Array.from(bytes, function(byte) {
          return ('0' + (byte & 0xFF).toString(16)).slice(-2)
        }).join('')
      },

      getAccountData() {
        let $this = $(this.$el)

        return {
          account: $this.find('.' + inputIdPrefix() + 'account').val(),
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
        }
      },

      accountDataCollect() {
        let verifyFailed = false
        let accountInfo = this.getAccountData()
        this.inputs.password.errorInfo = ''

        if(accountInfo.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = this.$t('nsInput.password.errorInfo')
        }

        return verifyFailed ? null : accountInfo
      },

      getSNInput() {
        let $this = $(this.$el)

        return $this.find('.' + inputIdPrefix() + 'serialNumber').val()
      },

      signIn() {
        let accountInfo = this.accountDataCollect()
        if(null === accountInfo) {
          return
        }

        let encKey = NSLocalStorage.getReqKey() || '' //this.getSNInput()
        //if(!encKey || encKey.length !== 40) {
        //  this.inputs.sn.errorInfo = this.$t('nsInput.sn.errorInfo')
        //  return
        //}

        let loginKey = Crypto.AESEnc(accountInfo.password, encKey)
        loginKey = Crypto.HmacSHA256(loginKey, loginKey)

        let loginData = {
          nonce : this.array2HexString(Mathjs.random([16], 0, 255)),
          randomKey : this.array2HexString(Mathjs.random([32], 0, 255)),
        }

        Http.login(
          this,
          loginData,
          loginKey,
          function (data) {
            let key = loginData.nonce + loginData.randomKey
            let responseData = Crypto.AESDec(data.Data, key)

            responseData = JSON.parse(responseData)

            if(responseData.Nonce !== loginData.nonce) {
              return
            }

            let walletJsonObj = JSON.parse(responseData.Wallet)

            NSLocalStorage.setReqKey(encKey)
            NSLocalStorage.setLogin(
              responseData.Wallet,
              walletJsonObj.Address,
            )

            this.$router.push({name: nsNamespace.MAIN})
          },
          function () {
            alert(this.$t('nsSignIn.loginfail'))
          }
        )
      }
    },
    data: function () {
      return {
        showSN: false,
        inputs: {
          password: {
            inputId: inputIdPrefix() + "password",
            title: this.$t('nsInput.password.title'),
            placeholder: this.$t('nsInput.password.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          //sn: {
          //  inputId: inputIdPrefix() + "serialNumber",
          //  title: this.$t('nsInput.sn.title'),
          //  placeholder: this.$t('nsInput.sn.placeholder'),
          //  hasAppend: true,
          //  inputType: 'text',
          //  maxSize: 40,
          //  errorInfo: '',
          //}
        }
      }
    },
    computed: {
      lang() {
        this.inputs = {
          password: {
            inputId: inputIdPrefix() + "password",
            title: this.$t('nsInput.password.title'),
            placeholder: this.$t('nsInput.password.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          //sn: {
          //  inputId: inputIdPrefix() + "serialNumber",
          //  title: this.$t('nsInput.sn.title'),
          //  placeholder: this.$t('nsInput.sn.placeholder'),
          //  hasAppend: true,
          //  inputType: 'text',
          //  maxSize: 40,
          //  errorInfo: '',
          //}
        }

        return this.$i18n.locale
      }
    },
  }
</script>

<style scoped>
    .ns-create-account {
        padding: 60px 50px;
        background-color: #ffffff;
    }

    .ns-create-account > div {
        float: left;
        width: 50%;
        height: 875px;
    }

    .nkn-sign-in {
        width: 100%;
        padding-right: 94px;
        margin-top: 40px;
    }

    .sign-in-button {
        margin-top: 50px;
        width: 100%;
    }
</style>
