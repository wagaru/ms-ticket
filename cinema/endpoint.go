package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PostCinemaEndpoint      endpoint.Endpoint
	GetCinemasEndpoint      endpoint.Endpoint
	GetCinemaEndpoint       endpoint.Endpoint
	GetCinemaByCityEndpoint endpoint.Endpoint
}

type PostCinemaRequest struct {
	Cinema *Cinema
}

type PostCinemaResponse struct {
	Err error `json:"error,omitempty"`
}

func (p PostCinemaResponse) error() error { return p.Err }

type GetCinemasByCityRequest struct {
	City City
}

type GetCinemasByCityResponse struct {
	Cinemas []*Cinema `json:"cinemas"`
	Err     error     `json:"error,omitempty"`
}

func (p GetCinemasByCityResponse) error() error { return p.Err }

type GetCinemasRequest struct {
}

type GetCinemasResponse struct {
	Cinemas []*Cinema `json:"cinemas"`
	Err     error     `json:"error,omitempty"`
}

func (p GetCinemasResponse) error() error { return p.Err }

type GetCinemaRequest struct {
	ID string
}

type GetCinemaResponse struct {
	Cinema *Cinema `json:"cinema"`
	Err    error   `json:"error,omitempty"`
}

func (p GetCinemaResponse) error() error { return p.Err }

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		PostCinemaEndpoint:      makePostCinemaEndpoint(svc),
		GetCinemasEndpoint:      makeGetCinemas(svc),
		GetCinemaEndpoint:       makeGetCinema(svc),
		GetCinemaByCityEndpoint: makeGetCinemasByCityEndpoint(svc),
	}
}

func makePostCinemaEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostCinemaRequest)
		err = svc.PostCinema(req.Cinema)
		return PostCinemaResponse{Err: err}, nil
	}
}

func makeGetCinemasByCityEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCinemasByCityRequest)
		cinemas, err := svc.GetCinemasByCity(req.City)
		return GetCinemasByCityResponse{cinemas, err}, nil
	}
}

func makeGetCinemas(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		cinemas, err := svc.GetCinemas()
		return GetCinemasResponse{cinemas, err}, nil
	}
}

func makeGetCinema(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCinemaRequest)
		cinema, err := svc.GetCinema(req.ID)
		return GetCinemaResponse{cinema, err}, nil
	}
}
