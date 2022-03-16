package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/log"
	"github.com/wagaru/ticket/pkg/common_error"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHttpHandler(endpts Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/cinemas").Handler(httptransport.NewServer(
		endpts.PostCinemaEndpoint,
		decodePostCinemaRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas/{id}").Handler(httptransport.NewServer(
		endpts.GetCinemaEndpoint,
		decodeGetCinemaRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas").Queries("city", "{city:[0-9]+}").Handler(httptransport.NewServer(
		endpts.GetCinemaByCityEndpoint,
		decodeGetCinemasByCityRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas").Handler(httptransport.NewServer(
		endpts.GetCinemasEndpoint,
		decodeGetCinemasRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodePostCinemaRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request PostCinemaRequest
	if err := json.NewDecoder(r.Body).Decode(&request.Cinema); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetCinemaRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, common_error.ErrBadRouting
	}
	return GetCinemaRequest{id}, nil
}

func decodeGetCinemasRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return GetCinemasRequest{}, nil
}

func decodeGetCinemasByCityRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	city := r.URL.Query().Get("city")
	cityInt, _ := strconv.Atoi(city)
	return GetCinemasByCityRequest{City(cityInt)}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if e, ok := resp.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	w.WriteHeader(getCodeFrom(err))
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func getCodeFrom(err error) int {
	switch err {
	case common_error.ErrAlreadyExists:
		fallthrough
	case common_error.ErrNotFound:
		fallthrough
	case common_error.ErrBadRouting:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
