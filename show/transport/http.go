package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/wagaru/ticket/pkg/common_error"
	"github.com/wagaru/ticket/show/endpoint"
)

func MakeHttpHandler(endpts endpoint.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/movies").Handler(httptransport.NewServer(
		endpts.AddMovieEndpoint,
		decodeAddMovieRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/cinemas").Handler(httptransport.NewServer(
		endpts.AddCinemaEndpoint,
		decodeAddCinemaRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas").Handler(httptransport.NewServer(
		endpts.GetCinemasEndpoint,
		decodeGetCinemasRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas/{id}/movies").Handler(httptransport.NewServer(
		endpts.GetCinemaMoviesEndpoint,
		decodeGetCinemaMoviesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/shows").Handler(httptransport.NewServer(
		endpts.AddShowEndpoint,
		decodeAddShowRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/cinemas/{cinema_id}/movies/{movie_id}/shows").Handler(httptransport.NewServer(
		endpts.GetShowsEndpoint,
		decodeGetShowsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/shows/{id}/seats").Handler(httptransport.NewServer(
		endpts.GetShowSeatsEndpoint,
		decodeGetShowSeatsRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeAddMovieRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&request.Movie); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAddCinemaRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddCinemaRequest
	if err := json.NewDecoder(r.Body).Decode(&request.Cinema); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetCinemasRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.GetCinemasRequest{}, nil
}

func decodeGetCinemaMoviesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	cinemaID, ok := vars["id"]
	if !ok {
		return nil, common_error.ErrBadRouting
	}
	ID, err := strconv.Atoi(cinemaID)
	if err != nil {
		return nil, err
	}
	return &endpoint.GetCinemaMoviesRequest{CinemaID: uint(ID)}, nil
}

func decodeAddShowRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddShowRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, common_error.ErrInvalidInput
	}
	return request, nil
}

func decodeGetShowsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	cinemaID, ok := vars["cinema_id"]
	if !ok {
		return nil, common_error.ErrInvalidInput
	}
	_cinemaID, _ := strconv.Atoi(cinemaID)
	movieID, ok := vars["movie_id"]
	if !ok {
		return nil, common_error.ErrInvalidInput
	}
	_movieID, _ := strconv.Atoi(movieID)
	return endpoint.GetShowsRequest{MovieID: uint(_movieID), CinemaID: uint(_cinemaID)}, nil
}

func decodeGetShowSeatsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	showID, ok := vars["id"]
	if !ok {
		return nil, common_error.ErrInvalidInput
	}
	_showId, _ := strconv.Atoi(showID)
	return endpoint.GetShowSeatsRequest{ShowID: uint(_showId)}, nil
}

type Error interface {
	GetError() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if e, ok := resp.(Error); ok && e.GetError() != nil {
		encodeError(ctx, e.GetError(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; chartset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; chartset=utf-8")
	w.WriteHeader(getCodeFrom(err))
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func getCodeFrom(err error) int {
	switch err {
	case common_error.ErrNotFound:
		fallthrough
	case common_error.ErrAlreadyExists:
		fallthrough
	case common_error.ErrBadRouting:
		fallthrough
	case common_error.ErrInvalidInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
