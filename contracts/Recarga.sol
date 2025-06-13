// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Contrato para gerenciar sessões de recarga de veículos elétricos
contract RecargaVE {
    // Estrutura que representa uma sessão de recarga
    struct Sessao {
        address usuario;           // Endereço do usuário que iniciou a sessão
        uint256 energiaConsumida;  // Quantidade de energia consumida (em kWh)
        uint256 valorPago;         // Valor pago pela recarga
        bool finalizada;           // Indica se a sessão foi finalizada
    }

    uint256 public precoPorKWh = 5; // Preço fixo por kWh (exemplo: 5 wei por kWh)
    mapping(uint256 => Sessao) public sessoes; // Mapeamento de IDs para sessões
    uint256 public contador; // Contador para gerar novos IDs de sessão

    // Evento emitido ao finalizar uma recarga
    event RecargaFinalizada(uint256 id, address usuario, uint256 energia, uint256 valor);

    // Função para iniciar uma nova sessão de recarga
    function iniciarSessao() external returns (uint256) {
        uint256 id = contador++; // Gera um novo ID de sessão
        sessoes[id] = Sessao(msg.sender, 0, 0, false); // Cria a sessão para o usuário
        return id; // Retorna o ID da nova sessão
    }

    // Função para finalizar uma sessão de recarga e efetuar o pagamento
    function finalizarSessao(uint256 id, uint256 energiaConsumida) external payable {
        Sessao storage s = sessoes[id]; // Obtém a sessão pelo ID
        require(s.usuario == msg.sender, unicode"Não autorizado"); // Garante que só o dono pode finalizar
        require(!s.finalizada, unicode"Já finalizada"); // Garante que a sessão não foi finalizada antes

        uint256 valor = energiaConsumida * precoPorKWh; // Calcula o valor a ser pago
        require(msg.value >= valor, "Pagamento insuficiente"); // Verifica se o pagamento é suficiente

        s.energiaConsumida = energiaConsumida; // Atualiza a energia consumida
        s.valorPago = msg.value; // Registra o valor pago
        s.finalizada = true; // Marca a sessão como finalizada

        emit RecargaFinalizada(id, msg.sender, energiaConsumida, msg.value); // Emite o evento de finalização
    }
}