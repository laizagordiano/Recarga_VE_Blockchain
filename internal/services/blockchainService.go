package services

import (
	"context"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/laizagordiano/Recarga_VE_Blockchain/recarga"
)

// BlockchainService encapsula a lógica de interação com o contrato inteligente
// Fornece métodos para iniciar/finalizar sessões e consultar histórico

type BlockchainService struct {
	client      *ethclient.Client
	contract    *recarga.Recarga
	fromAddress common.Address
	auth        *bind.TransactOpts
}

// NewBlockchainService inicializa a conexão com o nó e o contrato
func NewBlockchainService() (*BlockchainService, error) {
	rpcURL := os.Getenv("RPC_URL")
	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	chainIDStr := os.Getenv("CHAIN_ID")

	if rpcURL == "" || contractAddr == "" || privateKeyHex == "" || chainIDStr == "" {
		return nil, errors.New("variáveis de ambiente RPC_URL, CONTRACT_ADDRESS, PRIVATE_KEY, CHAIN_ID são obrigatórias")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID := new(big.Int)
	chainID.SetString(chainIDStr, 10)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	contract, err := recarga.NewRecarga(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, err
	}
	return &BlockchainService{client, contract, fromAddress, auth}, nil
}

// IniciarSessao chama o método iniciarSessao do contrato
func (b *BlockchainService) IniciarSessao() (*big.Int, error) {
	tx, err := b.contract.IniciarSessao(b.auth)
	if err != nil {
		return nil, err
	}
	receipt, err := bind.WaitMined(context.Background(), b.client, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != 1 {
		return nil, errors.New("transação de iniciarSessao falhou")
	}
	// O ID da sessão é retornado pelo evento ou pelo contador
	contador, err := b.contract.Contador(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return new(big.Int).Sub(contador, big.NewInt(1)), nil // O último ID criado
}

// FinalizarSessao chama o método finalizarSessao do contrato
func (b *BlockchainService) FinalizarSessao(id *big.Int, energiaConsumida *big.Int) error {
	preco, err := b.contract.PrecoPorKWh(&bind.CallOpts{})
	if err != nil {
		return err
	}
	valor := new(big.Int).Mul(energiaConsumida, preco)
	b.auth.Value = valor
	tx, err := b.contract.FinalizarSessao(b.auth, id, energiaConsumida)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), b.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("transação de finalizarSessao falhou")
	}
	return nil
}

// Recarga chama o método recarga do contrato para adicionar energia a uma sessão aberta
func (b *BlockchainService) Recarga(id *big.Int, energiaAdicional *big.Int) error {
	preco, err := b.contract.PrecoPorKWh(&bind.CallOpts{})
	if err != nil {
		return err
	}
	valor := new(big.Int).Mul(energiaAdicional, preco)
	b.auth.Value = valor
	// Substitua "AdicionarEnergia" pelo nome correto do método gerado pelo abigen para adicionar energia.
	tx, err := b.contract.Recarga(b.auth, id, energiaAdicional)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), b.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("transação de recarga falhou")
	}
	return nil
}

// GetSessao retorna os dados de uma sessão pelo ID
func (b *BlockchainService) GetSessao(id *big.Int) (usuario common.Address, energiaConsumida, valorPago *big.Int, finalizada bool, err error) {
	sessao, err := b.contract.Sessoes(&bind.CallOpts{}, id)
	if err != nil {
		return common.Address{}, nil, nil, false, err
	}
	return sessao.Usuario, sessao.EnergiaConsumida, sessao.ValorPago, sessao.Finalizada, nil
}

// GetHistorico retorna o histórico de sessões (varre todos os IDs)
func (b *BlockchainService) GetHistorico() ([]map[string]interface{}, error) {
	contador, err := b.contract.Contador(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	historico := []map[string]interface{}{}
	for i := big.NewInt(0); i.Cmp(contador) < 0; i.Add(i, big.NewInt(1)) {
		sessao, err := b.contract.Sessoes(&bind.CallOpts{}, new(big.Int).Set(i))
		if err != nil {
			continue // ignora sessões inválidas
		}
		historico = append(historico, map[string]interface{}{
			"id":               new(big.Int).Set(i),
			"usuario":          sessao.Usuario.Hex(),
			"energiaConsumida": sessao.EnergiaConsumida.String(),
			"valorPago":        sessao.ValorPago.String(),
			"finalizada":       sessao.Finalizada,
		})
	}
	return historico, nil
}
