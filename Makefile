test:
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

#GET cluster ID
docker exec -ti broker ./opt/kafka/bin/kafka-cluster.sh cluster-id --bootstrap-server :9092

#CREATE a new TOPIC
docker exec -ti broker ./opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic orders --partitions 3

#DESCRIBE a TOPIC
docker exec -ti broker ./opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic orders

#DELETE a TOPIC
docker exec -ti broker ./opt/kafka/bin/kafka-topics.sh --bootstrap-server <host:port> --delete --topic <topic-name>

#PUBLISHING MESSAGES
docker exec -ti broker ./opt/kafka/bin/kafka-console-producer.sh --bootstrap-server :9092 --topic orders

#CONSUMING MESSAGES
docker exec -ti broker ./opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server :9092 --topic orders --from-beginning
--from-beginning search for all topics message
