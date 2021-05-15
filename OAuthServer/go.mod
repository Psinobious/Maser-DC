module main

go 1.15

require (
	OAuth v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.5.0
	github.com/neo4j/neo4j-go-driver/v4 v4.2.2
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.37.0
)

replace OAuth => ../src/Infrastructure/OAuth
