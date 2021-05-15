module main

go 1.15

require (
	Infrastructure v0.0.0-00010101000000-000000000000
	OAuth v0.0.0-00010101000000-000000000000
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210311100734-5d6b0c808457
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/johnnadratowski/golang-neo4j-bolt-driver v0.0.0-20200323142034-807201386efa
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/neo4j/neo4j-go-driver/v4 v4.2.4
	github.com/rs/cors v1.7.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace Infrastructure => ../src/Infrastructure

replace Authentication => ../src/Infrastructure/OAuth

replace User => ../src/Application/Users

replace Publication => ../src/Application/Publications

replace OAuth => ../src/Infrastructure/OAuth

replace Test => ../src/Infrastructure/Test

replace Password => ../src/Infrastructure/Password
