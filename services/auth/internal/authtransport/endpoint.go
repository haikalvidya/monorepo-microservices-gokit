package authtransport

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/haikalvidya/monorepo-microservices-gokit/services/auth/internal/authservice"
)

type AuthEndpoint struct {
	GenerateTokenEndpoint endpoint.Endpoint
	ParseTokenEndpoint    endpoint.Endpoint
	DeleteTokenEndpoint   endpoint.Endpoint
}

type GenerateTokenRequest struct {

}

func NewAuthEndpoint(svc authservice.AuthService) AuthEndpoint {
	var generateTokenEndpoint endpoint.Endpoint
	// {
	// 	generateTokenEndpoint = 
	// }
}

func MakeGenerateTokenEndpoint(s authservice.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.()
	}
}