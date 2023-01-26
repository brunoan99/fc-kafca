# Apache Kafka

## Requisitos

* Docker e Compose

## Variáveis de Ambiente

Lembrete: antes de iniciar o projeto configure as variáveis de ambiente.
Conforme arquivo .env.example

## Iniciando Projeto

Para Iniciar o Projeto utilize o comando
```bash
docker compose up -d
```
Esse comando irá criar os containers necessários para o projeto.

Após a criação dos containers, assim que o container estiver pronto é preciso criar o Topic no Kafka, para criação acesse o container gokafka
```bash
docker exec -it gokafka bash
```

Após acessar o container utilize o go para criação do tópico com o seguinte comando:
```bash
go run cmd/topic/main.go
```

## Implementando Producer e Consumer

Uma implementação de Teste foi realizada tanto para Producer como para Consumer, e estão disponíveis nas pastas cmd/consumer e cmd/producer respectivamente.

Para novos testes modifique os arquivos main.go
