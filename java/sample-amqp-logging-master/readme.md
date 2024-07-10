## How to ship logs with Logstash, Elasticsearch and RabbitMQ
## rabbitmq
docker run -d -it --name rabbit --hostname rabbit -p 30000:5672 -p 30001:15672 rabbitmq:management

## elasticsearch
error 1
Elasticsearch bootstrap check failure [1] of [2]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
fix error1
wsl -d docker-desktop -u root
sysctl -w vm.max_map_count=262144
error 2
bootstrap check failure [1] of [1]: the default discovery settings are unsuitable for production use; at least one of [discovery.seed_hosts, discovery.seed_providers, cluster.initial_master_nodes] must be configured
fix error2 : -e "discovery.type=single-node"
docker run -d -it --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.19

## logstash
docker run -d -it --name logstash logstash:7.17.19 -e 'input { rabbitmq {
host => "localhost" port => 30000 durable => true } }
output { elasticsearch { hosts => ["localhost"] } }'