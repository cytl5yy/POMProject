package pom

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

)
type Data struct{
	ID string
	Address common.Address
}


func(self *Data) SigPOMData(prv *ecdsa.PrivateKey) []byte {
	bytes,err := rlp.EncodeToBytes([]interface{}{
		self.ID,
		self.Address,
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var b Data
	rlp.DecodeBytes(bytes,&b)
	hash := crypto.Keccak256(bytes)
	sig, err :=crypto.Sign(hash,prv)
	if err!=nil{
		return nil
	}
	return sig
}