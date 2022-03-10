package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wagaru/ticket/movie/domain"
	"github.com/wagaru/ticket/movie/endpoint"
)

var ErrBadRouting = errors.New("invalid routing")

func MakeHttpHandler(endpts endpoint.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/movies").Handler(httptransport.NewServer(
		endpts.PostMovieEndpoint,
		decodePostMovieRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/movies").Handler(httptransport.NewServer(
		endpts.GetMoviesEndpoint,
		decodeGetMoviesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/movies/{id}").Handler(httptransport.NewServer(
		endpts.GetMovieEndpoint,
		decodeGetMovieRequest,
		encodeResponse,
		options...,
	))

	r.Path("/metrics").Handler(promhttp.Handler())

	return r
}

func decodePostMovieRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.PostMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Movie); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetMoviesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGetMovieRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoint.GetMovieRequest{ID: id}, nil
}

type myError interface {
	GetError() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(myError); ok && e.GetError() != nil {
		encodeError(ctx, e.GetError(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrBadRouting:
		fallthrough
	case domain.ErrAlreadyExists:
		fallthrough
	case domain.ErrNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
