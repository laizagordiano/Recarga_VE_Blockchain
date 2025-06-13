<img width="100%" src="https://capsule-render.vercel.app/api?type=waving&color=00FFFF&height=120&section=header"/>

<div align="center">

# Recarga_VE_Blockchain

## Sistema distribuído para gestão de recarga de veículos elétricos com uso de Blockchain

</div>

Este projeto visa resolver o problema de "ansiedade de autonomia" em veículos elétricos (VEs), garantindo **transparência**, **segurança** e **interoperabilidade entre empresas** por meio da tecnologia **Blockchain**. O sistema permite registrar reservas de recarga, rotas compostas e sessões concluídas de forma auditável, utilizando contratos inteligentes.

---

## 🛠️ Tecnologias Utilizadas

- **Backend:** Go (Golang)
- **Contratos inteligentes:** Solidity + Hardhat
- **Comunicação:** API REST
- **Blockchain local:** Hardhat Node
- **Contêineres:** Docker & Docker Compose
- **Cliente:** Simulação de veículo em Go
- **Gerador de chaves:** Script Go

---


## 🚀 Como Executar o Projeto

### 1. Clone o repositório

```bash
git clone <repo>
cd Recarga_VE_Blockchain
```

### 2. Suba a blockchain 

```bash
docker-compose up --build blockchain
```
### 3. Suba as empresas
```bash
docker-compose up --build empresa_a empresa_b empresa_c
```
### 3. Suba o cliente
```bash
docker-compose up --build carro_cliente
```

Isso irá:

- Iniciar o nó Hardhat (blockchain local)
- Fazer deploy automático do contrato inteligente
- Iniciar os servidores Go das empresas A, B e C
- Subir o cliente simulador de veículo

### 4. Acesse as APIs

- **Empresa A:** http://localhost:8085
- **Empresa B:** http://localhost:8086
- **Empresa C:** http://localhost:8087

---

## 🔁 Fluxo de Funcionamento

1. Empresa A inicia uma reserva via `POST /blockchain/reservar`
2. A transação é registrada na blockchain
3. Empresa B (destino) confirma a sessão de recarga
4. A sessão pode ser finalizada via `POST /blockchain/finalizar`
5. Todo o histórico fica disponível para consulta

---


## 🔧 Variáveis de Ambiente

Exemplo para Empresa A (`.env`):

```ini
NOME_EMPRESA=empresa_a
PORTA=8085
RPC_URL=http://blockchain:8545
CONTRACT_ADDRESS=0x...
PRIVATE_KEY=...
CHAIN_ID=31337
```

---

## 🔗 Principais Endpoints REST

- `POST /blockchain/reservar` — Cria uma nova reserva
- `POST /blockchain/finalizar` — Finaliza uma sessão
- `GET /blockchain/historico` — Lista sessões por empresa
- `GET /blockchain/sessao/:id` — Detalhes de uma sessão específica

---


## 📝 Observações

- Cada empresa só pode operar com sua própria chave privada
- O contrato impede que empresas não autorizadas confirmem sessões
- É possível expandir o sistema para incluir:
  - Novas empresas
  - Novos pontos de recarga
  - Lógica de rotas avançadas

---

## 👩‍💻 Alunos(as)

<table align='center'>
  <tr>
    <td align="center">
      <a href="https://github.com/LuisMarioRC">
        <img src="https://avatars.githubusercontent.com/u/142133059?v=4" width="100px;" alt=""/>
        <br /><sub><b>Luis Mario</b></sub>
      </a>
      <br />👨💻
    </td>
    <td align="center">
      <a href="https://github.com/laizagordiano">
        <img src="https://avatars.githubusercontent.com/u/132793645?v=4" width="100px;" alt=""/>
        <br /><sub><b>Laiza Gordiano</b></sub>
      </a>
      <br />👩💻
    </td>
    <td align="center">
      <a href="https://github.com/GHenryssg">
        <img src="https://avatars.githubusercontent.com/u/142272107?v=4" width="100px;" alt=""/>
        <br /><sub><b>Gabriel Henry</b></sub>
      </a>
      <br />👨💻
    </td>
  </tr>
</table>