package pom

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"path/filepath"
)

const (
	 datadirPrivateKey    = "nodekey"    //Path within the datadir to the node's private key
	 challengeInterval   = 120           //挑战周期120个区块间隔
     challengeDelay      = 20            //本次挑战从选择目标节点到提交POMReceipt有效时间为20个区块
	//chainId      = 2023                  //本次挑战从选择目标节点到提交POMReceipt有效时间为20个区块
	chainId      = 5                  //本次挑战从选择目标节点到提交POMReceipt有效时间为20个区块
    getTargetEventTopic = "0x53b0b52a85961621cf1f25fa9513d34315a2815b92e27c5d7c5603bd8ecb3385"
)
var(
	//pomContractAddress      = common.HexToAddress("0x0000000000000000000000000000000000004000")
	pomContractAddress      = common.HexToAddress("0x69C08f45330dAC6f4acbc8A29C96E50e04941741")

)
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
func GetPrivateKey() (prv *ecdsa.PrivateKey){
	dir,_:=os.Getwd()
	keyfile :=filepath.Join(dir, datadirPrivateKey)
	if FileExist(keyfile){
		prv,_=crypto.LoadECDSA(keyfile)
	}else{
		prv,_ =crypto.GenerateKey()
		crypto.SaveECDSA(keyfile, prv)
	}
	return
}
