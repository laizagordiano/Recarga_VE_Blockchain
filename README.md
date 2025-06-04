# Recarga_VE_Blockchain

## Como rodar o projeto

1. Instale as dependências:
```powershell
npm install
```

2. Compile o contrato:
```powershell
npx hardhat compile
```

3. Inicie a blockchain local do Hardhat (em um novo terminal):
```powershell
npx hardhat node
```

4. Em outro terminal, faça o deploy do contrato:
```powershell
npx hardhat run scripits/deploy.js --network localhost
```

O endereço do contrato será exibido no terminal após o deploy.