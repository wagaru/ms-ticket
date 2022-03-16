package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
)

type RoutingService interface {
	GetCinemaHallSeats(cinemaID string) ([]*CinemaSeat, error)
}

type proxyService struct {
	ctx           context.Context
	fetchEndpoint endpoint.Endpoint
	svc           RoutingService
}

func (p *proxyService) GetCinemaHallSeats(cinemaID string) ([]*CinemaSeat, error) {
	resp, err := p.fetchEndpoint(p.ctx, fetchRequest{})
	if err != nil {
		return nil, err
	}

	response := resp.(fetchResponse)
	return response.Seats, nil
}

type RoutingServiceMiddleware func(RoutingService) RoutingService

func NewProxyMiddleware(ctx context.Context, proxyURL string) RoutingServiceMiddleware {
	return func(next RoutingService) RoutingService {
		var e endpoint.Endpoint
		e = makeFetchEndpoint(ctx, proxyURL)
		e = circuitbreaker.Hystrix("fetchCinemaHallSeats")(e)
		return &proxyService{
			ctx:           ctx,
			svc:           next,
			fetchEndpoint: e,
		}
	}
}

func makeFetchEndpoint(ctx context.Context, proxyURL string) endpoint.Endpoint {
	u, err := url.Parse(proxyURL)
	if err != nil {
		panic(err)
	}
	return kithttp.NewClient("GET", u, encodeRequest, decodeResponse).Endpoint()
}

type fetchRequest struct {
}

type fetchResponse struct {
	Seats []*CinemaSeat `json:"seats"`
	Err   error         `json:"error"`
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	return nil
}

func decodeResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response fetchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
