module Infrastructure

go 1.15

require (
	OAuth v0.0.0-00010101000000-000000000000
	Publication v0.0.0-00010101000000-000000000000
	Test v0.0.0-00010101000000-000000000000
	User v0.0.0-00010101000000-000000000000
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210311100734-5d6b0c808457
	github.com/johnnadratowski/golang-neo4j-bolt-driver v0.0.0-20200323142034-807201386efa // indirect
)

replace OAuth => ./OAuth

replace Password => ./Password

replace User => ../Application/Users

replace Publication => ../Application/Publications

replace Test => ./Test
