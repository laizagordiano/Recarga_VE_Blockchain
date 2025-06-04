package main

import (
    "crypto/ecdsa"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)
    fmt.Printf("Chave privada (hex): %x\n", privateKeyBytes)

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("erro ao converter chave pública")
    }

    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Printf("Endereço Ethereum: %s\n", address)
}
