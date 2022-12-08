package pom

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)
type POMReceipt struct {
	target common.Address
	v []uint8
	r [][32]byte
	s [][32]byte

}
type Event struct {
	client        *ethclient.Client
	prv           *ecdsa.PrivateKey
	newHeader     chan *types.Header
	newChallenge  chan common.Address
	newPOMReceipt chan POMReceipt
}


func NewEvent(client *ethclient.Client,prv *ecdsa.PrivateKey)*Event{
	header := make(chan *types.Header)
	challenge := make(chan common.Address)
	pomReceipt := make(chan POMReceipt)
	event := &Event{client:client,prv:prv,newHeader:header,newChallenge:challenge,newPOMReceipt:pomReceipt}
	go event.ListenBlockEvent()
	go event.ListenPOMContractEvent()
	go event.update()
	return event
}
//用于监听mesh网络区块同步信息
func (self *Event)ListenBlockEvent(){
	headers := make(chan *types.Header)
	sub, err :=self.client.SubscribeNewHead(context.Background(),headers)
	if err != nil {
		fmt.Println("SubscribeNewHead", "err", err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("sub.Err", "err", err)
		case header := <-headers:
			self.newHeader <-header
		}
	}
}
//尝试发起挑战
func (self *Event)TryChallenge(header *types.Header){
	pom,_ :=NewPom(pomContractAddress,self.client)
	lastChallengeBlock, err := pom.GetLastChallengeBlock(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	},crypto.PubkeyToAddress(self.prv.PublicKey))
	if err != nil {
		fmt.Println("GetLastChallengeBlock", "err", err)
		return
	}
	fmt.Println("TryChallenge",header.Number.Uint64(),lastChallengeBlock.Uint64())
	//符合挑战间隔区块，获取目标节点
	if header.Number.Uint64()-lastChallengeBlock.Uint64()>= challengeInterval{
		self.GetPOMTarget()
	}
}
func (self *Event)GetPOMTarget(){
	auth, err := bind.NewKeyedTransactorWithChainID(self.prv, new(big.Int).SetInt64(chainId))
    pom,_ :=NewPom(pomContractAddress,self.client)
	tx, err := pom.GetTarget(
		&bind.TransactOpts{
		From:        crypto.PubkeyToAddress(self.prv.PublicKey),
		Signer: auth.Signer,
		})

	fmt.Println("GetPOMTarget","tx",tx.Hash().String(),"err",err)

}

//用于监听pom合约获取Target方法返回目标节点信息
func (self *Event)ListenPOMContractEvent() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{pomContractAddress},
	}
	logs := make(chan types.Log)
	sub, err := self.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println("SubscribeFilterLogs","err",err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("sub.Err","err",err)
		case vLog := <-logs:
			//判断是GetTarget Event
			if len(vLog.Topics)==2 && strings.EqualFold(vLog.Topics[0].Hex(),getTargetEventTopic) {
				from := vLog.Topics[1].Hex()
				if len(from)==66{
					from = "0x"+ from[26:]
					data :=hexutil.Encode(vLog.Data)
					//判断是本地址发送的判断是本节点钱包调用的getTarget方法
					if strings.EqualFold(from,crypto.PubkeyToAddress(self.prv.PublicKey).String()) && len(data)==66{
						target := "0x"+ data[26:]
						self.newChallenge <- common.HexToAddress(target)
					}
				}
			}
		}
	}
}
//发送见证人返回的签名数据报文到pom合约
func (self *Event)SendPOMReceipt(target common.Address,v []uint8, r [][32]byte, s [][32]byte){
	auth, _ := bind.NewKeyedTransactorWithChainID(self.prv, new(big.Int).SetInt64(chainId))
	pom,_ :=NewPom(pomContractAddress,self.client)
	tx, err := pom.SendPOMReceipt(
		&bind.TransactOpts{
			From:        crypto.PubkeyToAddress(self.prv.PublicKey),
			Signer: auth.Signer,
		},
		target,
		v,
		r,
		s)
	fmt.Println("SendPOMReceipt", "tx", tx.Hash().String(),"err",err)
}
//向目标节点发起挑战
func (self *Event)SendMsgtoTarget(target common.Address){
	fmt.Println("SendMsgtoTarget", "target", target.String())
	//TO DO

}
//监听部分事件
func (self *Event) update() {
	for {
		select {
		  case header :=<- self.newHeader:
		  	self.TryChallenge(header)
		  case target :=<- self.newChallenge:
			self.SendMsgtoTarget(target)
          case pomReceipt :=<-self.newPOMReceipt:
          	self.SendPOMReceipt(pomReceipt.target,pomReceipt.v,pomReceipt.r,pomReceipt.s)
		}
	}
}