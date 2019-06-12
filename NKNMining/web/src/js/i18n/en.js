module.exports = {
  nav: {
    langSel: 'Language',
    langList: {
      en: 'English',
      zh: '中文',
    }
  },
  mark: 'NKN',
  nsNav: {
    clearCache: 'Clear cache',
    signOut: 'Sign out',
    resetNodeMining: 'Reset @:mark mining'
  },
  nsFooter: {
    copyright: 'Copyright {i} 2017-2018 @:mark    |    All rights reserved',
    attention: 'This software is in the early development stage and may not have all functions working properly. It should be used only for testing now. '
  },
  nsInput: {
    account: {
      title: 'Account',
      placeholder: 'Your @:mark mining account',
      errorInfo: 'Alphanumeric only!'
    },
    password: {
      title: 'Your wallet password',
      placeholder: '8-20 characters',
      errorInfo: 'Please input 8-20 characters'
    },
    rePassword: {
      title: 'Confirm password',
      placeholder: '8-20 characters',
      errorInfo: 'The password for the two input is inconsistent!'
    },
    sn: {
      title: 'System initialization serial number',
      placeholder: '40 characters',
      errorInfo: 'Invalid serial number'
    },
    wallet: {
      title: ' ',
      placeholder: 'Wallet file to load',
      errorInfo: 'please select a wallet file'
    },
    walletPassword: {
      title: 'Wallet password',
      placeholder: 'password ot the wallet to load',
      errorInfo: 'Please check if the password matches the wallet.'
    },
  },
  nsSignIn: {
    titleLabel: 'Welcome to',
    title: '@:mark NODE',
    signInbtn: 'Sign in',
    loginfail:'login failed!'
  },
  nsMain: {
    node: {
      title: 'Node connect',
      on: 'on',
      off: 'off',
      mining:{
        downloading: "you can't start mining when downloading NKN node program.",
        fail: 'start mining failed. please check log on the server.',
        error: "you can't mine when node updating.",
        stopfail: 'stop mining failed. please check log on the server.'
      },
      nodeStatus:{
        relayCount: 'Message relayed by node: {relayCount}',
        prefix: 'Node status: ',
        init: '@:nsMain.node.nodeStatus.prefix sync ...',
        chainDownloading: '@:nsMain.node.nodeStatus.prefix downloading NKN node program.',
        starting: 'starting',
        updating: '@:nsMain.node.nodeStatus.prefix updating',
        stopping: 'stopping',
        stoped: '@:nsMain.node.nodeStatus.prefix stopped',
        syncing: '@:nsMain.node.nodeStatus.prefix syncing',
        persisting: '@:nsMain.node.nodeStatus.prefix persisting',
        mining: '@:nsMain.node.nodeStatus.prefix mining'
      },
      version: {
        init: 'Querying version ...'
      },
      NKNBlockHeight: '@:mark network block height:',
      myBlockHeight: 'Node block height:'
    },
    wallet: {
      title: 'Wallet',
      downloadbtn: 'download',
      transferbtn: 'transfer',
      addressLabel: 'Address',
      publicKeyStatus: 'Public key:',
      beneficiaryStatus: 'Beneficiary address:',
      beneficiaryLabel: 'Beneficiary',
      autoTransferBtn: 'Auto transfer',
      autoTransferStopBtn: 'Stop',
      beneficiaryAddrPlaceholder: 'Beneficiary address then click auto transfer',
      autoTransferTip: 'Click Auto transfer to turn this feature on.',
      balance: {
        init: 'Querying balance ...'
      },
      balanceLabel: 'Balance',
      balanceUnit: '@:mark'
    },
    neighbors: {
      title: 'Neighbors',
      ID: 'ID',
      IP: 'IP',
      port: 'Port'
    },
    rewards: {
      title: 'Your latest mining rewards',
      transaction: 'Transaction',
      value: 'Value',
      height: 'Height',
      time: 'Time'
    },
    transfer: {
      title: 'Transfer @:mark',
      addressPlaceholder: 'target address',
      countPlaceholder: 'how much @:mark to transfer',
      passwordPlaceholder: 'wallet password',
      cancel: 'Cancel',
      confirm: 'Transfer',
      alertInfo: {
        default: 'transfer failed!'
      }
    },
    resetdlg: {
      title: 'Reset mining confirm',
      warning: 'Warning: This operation will reset all your @:mark data, include your Wallet!',
      passwordPlaceholder: 'Account password',
      cancel: 'Cancel',
      reset: 'Reset',
      alertInfo: {
        wrongPass: 'wrong password!',
        success: 'reset node mining success'
      }
    }
  },
  nsGenWallet: {
    title: 'Generate wallet',
    link: 'Load wallet',
    nextStepbtn: 'NEXT'
  },
  nsLoadWallet: {
    title: 'Load wallet',
    upload: 'Select file',
    link: 'Generate wallet',
    nextStepbtn: 'NEXT',
    netWorkError: 'Network error, please try again later.'
  },
  nsShowWallet: {
    title: 'Download wallet',
    privateKey: 'Private key',
    address: 'Wallet address',
    signInbtn: 'Sign in',
    downloadbtn: 'Download wallet'
  }
}