# maxwell-mysql-kafka-example

https://maxwells-daemon.io/ prototyping 🔧

## producer=stdout

```bash
docker run -it --rm --network maxwell-mysql-kafka-example_default zendesk/maxwell bin/maxwell --user=root --password=root --host=mariadb --producer=stdout
```

## producer=kafka

```bash
# exec shell in kafka container
docker-compose exec -it kafka /bin/bash
# create topic with 20 partitions
kafka-topics.sh --bootstrap-server localhost:9092 --create --topic jobs --partitions 20 --replication-factor 1
# delete topic
kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic jobs
```

```bash
docker run -it --rm --network maxwell-mysql-kafka-example_default zendesk/maxwell bin/maxwell --user=root --password=root --host=mariadb --producer=kafka --kafka.bootstrap.servers=kafka:9092 --kafka_topic=jobs --producer_partition_by=primary_key
```

## php producer

```bash
docker run -it --rm --network maxwell-mysql-kafka-example_default -v $PWD/producer/producer.php:/app/producer.php jonczek/php-cli php /app/producer.php
```
