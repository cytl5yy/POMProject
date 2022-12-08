
package main
import (
	"POMProject/src/pom"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	"fmt"
)

func main() {

	var prv *ecdsa.PrivateKey
	prv = pom.GetPrivateKey()
	//client, err := ethclient.Dial("ws://127.0.0.1:8546")
	client, err := ethclient.Dial("wss://eth-goerli.g.alchemy.com/v2/3gwNayGODX4zFCMUrh7ZlYydaVULpvia")
	//client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/3gwNayGODX4zFCMUrh7ZlYydaVULpvia")
	if err != nil {
		fmt.Println("Dial","err",err)
		return
	}
	pom.NewEvent(client,prv)
	select {

	}

}



