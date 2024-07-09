## How to ship logs with Logstash, Elasticsearch and RabbitMQ
docker run -d -it --name rabbit --hostname rabbit -p 30000:5672 -p 30001:15672 rabbitmq:management