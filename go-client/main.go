package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/laizagordiano/Recarga_VE_Blockchain/recarga"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	rpcURL := os.Getenv("RPC_URL")
	chainIDStr := os.Getenv("CHAIN_ID")

	chainID := new(big.Int)
	chainID.SetString(chainIDStr, 10)

	// Conecta com o nó RPC
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, instance, err := recarga.DeployRecarga(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contrato implantado! Endereço:", address.Hex())
	fmt.Println("Hash da transação:", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contrato minerado no bloco %d\n", receipt.BlockNumber.Uint64())

	// Pode usar `instance` para chamar funções do contrato
	_ = instance
}
