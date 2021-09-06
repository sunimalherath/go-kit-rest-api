package go_kit_rest_api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/sunimalherath/go-kit-rest-api/regionsvc"
	"net/url"
	"strings"
)

type Endpoints struct {
	PostRegionEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s regionsvc.Service) Endpoints {
	return Endpoints{
		PostRegionEndpoint: MakePostRegionEndpoint(s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}

	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		PostRegionEndpoint: httptransport.NewClient("POST", tgt, encodePostRegionRequest, decodePostRegionRequest, options...).Endpoint(),
	}, nil
}

func MakePostRegionEndpoint(s regionsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postRegionRequest)
		e := s.PostRegion(ctx, req.Region)
		return postRegionResponse{Err: e}, nil
	}
}

type postRegionRequest struct {
	Region regionsvc.Region
}

type postRegionResponse struct {
	Err error `json:"err,omitempty"`
}

func (e Endpoints) PostRegion(ctx context.Context, r regionsvc.Region) error {
	request := postRegionRequest{Region: r}
	response, err := e.PostRegionEndpoint(ctx, request)
	if err != nil {
		return err
	}
	respns := response.(postRegionResponse)
	return respns.Err
}
