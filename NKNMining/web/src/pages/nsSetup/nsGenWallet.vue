<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <h1 class="nkn-page-title text-main-blue">{{$t("nsGenWallet.title")}}</h1>
            <div v-for="(inputItem, idx) in inputs" :key="idx">
                <ns-input-item v-if="showInput(idx)" :config="inputItem" />
            </div>
            <div class="setup-button nkn-after-clear nkn-wallet-setup-button-panel">
                <a class="nkn-link-load-wallet" @click="toLoadWallet">{{$t("nsGenWallet.link")}}</a>
                <button class="nkn-normal-btn" type="button" @click="nextStep">{{$t("nsGenWallet.nextStepbtn")}}</button>
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
  import NsInputItem from "../../components/nsInputItem"
  import Http from "../../js/network/nsHttp"
  import Crypto from "../../js/crypto/algorithm"
  import NSLocalStorage from "../../js/nsLocalStorage"
  import {nsNamespace} from "../../js/nsNamespace";
  import NsLoading from "../../components/nsLoading";
  import {inputIdPrefix} from "../../js/nsConst"
  import {loadPage} from "../../js/nsLoading";
  import LangMix from "../../js/mixin/lang.js"

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    name: "ns-gen-wallet",
    mounted() {
      loadPage.call(this)
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

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: this.$t('nsInput.rePassword.title'),
            placeholder: this.$t('nsInput.rePassword.placeholder'),
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

    methods: {
      showInput(idx) {
        if('sn' === idx) {
          return this.showSN
        }

        return true
      },
      getWalletPassword() {
        let $this = $(this.$el)

        return {
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
          rePassword: $this.find('.' + inputIdPrefix() + 'rePassword').val(),
        }
      },

      getSNInput() {
        let $this = $(this.$el)

        return $this.find('.' + inputIdPrefix() + 'serialNumber').val()
      },

      walletDataGen(encKey) {
        let verifyFailed = false
        let walletPassword = this.getWalletPassword()
        this.inputs.password.errorInfo = ''
        this.inputs.rePassword.errorInfo = ''

        if(walletPassword.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = this.$t('nsInput.password.errorInfo')
        }

        if(walletPassword.password !== walletPassword.rePassword) {
          verifyFailed = true
          this.inputs.password.errorInfo = ''
          this.inputs.rePassword.errorInfo = this.$t('nsInput.rePassword.errorInfo')
        }

        if(verifyFailed) {
          return null
        }

        let wallet = nknWallet.newWallet(walletPassword.password)
        if(!wallet) {
          return null
        }

        let key = Crypto.AESEnc(walletPassword.password, encKey)

        return {
          reqData: {
            Wallet: wallet.toJSON(),
            Key: key,
          },

          wallet: wallet
        }
      },

      toLoadWallet() {
        this.$router.push({name: nsNamespace.SETUP.LOAD_WALLET})
      },

      nextStep() {
        let encKey = NSLocalStorage.getReqKey() || '' //this.getSNInput()

        //if(!encKey || encKey.length !== 40) {
        //  this.inputs.sn.errorInfo = this.$t('nsInput.sn.errorInfo')
        //  return
        //}

        let walletData = this.walletDataGen(encKey)
        if(null === walletData) {
          return
        }

        Http.setWallet(this, walletData.reqData, encKey, function () {
          NSLocalStorage.setReqKey(encKey)
          NSLocalStorage.setWallet(walletData.reqData.Wallet, walletData.wallet.address)
          this.$store.commit(nsNamespace.GLOBAL + "/updateWallet", walletData.wallet)
          this.$router.push({name: nsNamespace.SETUP.SHOW_WALLET})
        }, function (err) {
          console.error(err)
        })
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

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: this.$t('nsInput.rePassword.title'),
            placeholder: this.$t('nsInput.rePassword.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          //sn: {
          //  inputId: inputIdPrefix() + "serialNumber",
          //  title: this.$t('nsInput.sn.title'),
          //  placeholder: this.$t('nsInput.sn.placeholder'),
          //  hasAppend: false,
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
    .nkn-wallet-setup-button-panel {
        text-align: right;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-load-wallet {
        display: inline-block;
        line-height: 110px;
        margin-top: 25px;
        cursor: pointer;
        color: #498ae5;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-load-wallet:hover {
        text-decoration: underline;
    }
</style>