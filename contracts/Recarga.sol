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
    // Evento emitido ao recarregar energia
    event EnergiaRecarregada(uint256 id, address usuario, uint256 energiaTotal);

    // Função para iniciar uma nova sessão de recarga (reserva)
    function reservarSessao() public returns (uint256) {
        uint256 id = contador++;
        sessoes[id] = Sessao(msg.sender, 0, 0, false);
        return id;
    }

    // Função antiga para compatibilidade (opcional)
    function iniciarSessao() external returns (uint256) {
        return this.reservarSessao();
    }

    // Função para recarregar energia em uma sessão (pode ser chamada múltiplas vezes antes de finalizar)
    function recarregarSessao(uint256 id, uint256 energia) external {
        Sessao storage s = sessoes[id];
        require(s.usuario == msg.sender, unicode"Não autorizado");
        require(!s.finalizada, unicode"Sessão já finalizada");
        s.energiaConsumida += energia;
        emit EnergiaRecarregada(id, msg.sender, s.energiaConsumida);
    }

    // Função para finalizar uma sessão de recarga e efetuar o pagamento
    function finalizarSessao(uint256 id) external payable {
        Sessao storage s = sessoes[id];
        require(s.usuario == msg.sender, unicode"Não autorizado");
        require(!s.finalizada, unicode"Já finalizada");
        uint256 valor = s.energiaConsumida * precoPorKWh;
        require(msg.value >= valor, "Pagamento insuficiente");
        s.valorPago = msg.value;
        s.finalizada = true;
        emit RecargaFinalizada(id, msg.sender, s.energiaConsumida, msg.value);
    }
}