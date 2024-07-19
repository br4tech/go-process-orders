# Go Order Processing with RabbitMQ and Docker

Este projeto demonstra como criar um sistema de processamento de pedidos em Go, utilizando RabbitMQ como broker de mensagens e Docker para facilitar o desenvolvimento e a implantação.

## Recursos

### Processamento de Pedidos:
- Gera 10 pedidos fictícios.
- Publica os pedidos em uma fila RabbitMQ chamada "orders".
- Consome os pedidos da fila, simulando o processamento.

### RabbitMQ:
- Broker de mensagens robusto e escalável.
- Permite a comunicação assíncrona entre os componentes do sistema.

### Docker:
- Facilita a criação de um ambiente de desenvolvimento isolado e consistente.
- Simplifica a implantação da aplicação em diferentes ambientes.

### Docker Compose:
- Orquestra os containers Docker do RabbitMQ e da aplicação Go.
- Facilita a inicialização e o gerenciamento do ambiente de desenvolvimento.

## Pré-requisitos
- **Docker:** Certifique-se de ter o Docker instalado e em execução.
- **Go:** Certifique-se de ter o Go (versão 1.18 ou superior) instalado.

## Como Usar

### Clone o Repositório:
```bash
  git clone https://github.com/br4tech/go-process-orders.git
```

### Inicie os Containers:
```bash
  docker-compose up -d
```
Isso iniciará os containers do RabbitMQ e da aplicação Go.

### Execute o Produtor de Pedidos:
```bash
  go run main.go
```

### Verifique os Logs:
```bash
  docker-compose logs -f
```

## Estrutura do Projeto
```bash
  go-process-order/
  |____README.md
  |____.gitignore
  |____Dockerfile
  |____cmd
  | |____main.go
  |____go.mod
  |____go.sum
  |____docker-compose.yml
  |____internal
  | |____adapter
  | | |____rabbitmq_adapter.go
  | |____port
  | | |____message.go
  | |____domain
  | | |____services
  | | | |____order.go
  | | |____entities
  | | | |____order.go
```






