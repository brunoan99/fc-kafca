# Apache Kafka

## Requisitos

* Docker e Compose

## Iniciando Projeto

Para Iniciar o Projeto utilize o comando
```bash
docker compose up -d
```
Esse comando irá criar os containers necessários para o projeto.

Após a criação dos containers, algumas configurações são realizadas de forma manual no Kafka.
Para realiza-las acesse o container do kafka usando o nome do container:
```bash
docker exec -it **nome-do-container** bash
```

Após acessar o container crie um tópico no kafka através do comando:
```bash
kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3
```

## Implementando Producer e Consumer

