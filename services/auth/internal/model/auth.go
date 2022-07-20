package model

// this model file just for data that want to store on redis/database
// if want to see data representation that used for request and response
// see protofile or on the service/transport file
type TokenSession struct {
	Token    string
	Username string
	Email    string
	Role     string
}

// TODO create two port, one for tls server side, another one for mutual tls
// server side tls going to check the jwt
// mutual tls arent going to check the jwt
