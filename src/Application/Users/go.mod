module User

go 1.15

require (
	OAuth v0.0.0-00010101000000-000000000000
	Password v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/neo4j/neo4j-go-driver/v4 v4.2.4
	google.golang.org/grpc v1.37.0
)

replace Password => ../../Infrastructure/Password

replace OAuth => ../../Infrastructure/OAuth
