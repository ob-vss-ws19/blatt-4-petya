package MovieService

import (
	"context"
	"fmt"
	MovieService "github.com/ob-vss-ws19/blatt-4-petya/MovieService/messages"
	ShowService "github.com/ob-vss-ws19/blatt-4-petya/ShowService/messages"
	"sync"
)

type Movie struct {
	title string
}

type MovieMicroService struct {
	MovieRepository map[int32]*Movie
	NextID          int32
	mu              *sync.RWMutex
	ShowService     func() ShowService.ShowService
}

func Spawn() *MovieMicroService {
	return &MovieMicroService{
		MovieRepository: make(map[int32]*Movie),
		NextID:          1,
		mu:              &sync.RWMutex{},
	}
}

func (msrv MovieMicroService) CreateMovie(context context.Context, req *MovieService.CreateMovieMessage, res *MovieService.CreateMovieResponse) error {
	msrv.mu.Lock()

	msrv.MovieRepository[msrv.NextID] = &Movie{title: req.Title}
	msrv.NextID++

	msrv.mu.Unlock()
	return nil
}

func (msrv MovieMicroService) DeleteMovie(context context.Context, req *MovieService.DeleteMovieMessage, res *MovieService.DeleteMovieResponse) error {
	msrv.mu.Lock()

	_, ok := msrv.MovieRepository[req.MovieID]

	if ok {
		s := msrv.ShowService()

		message := &ShowService.KillShowsMovieMessage{
			MovieID: req.MovieID,
		}

		s.KillShowsMovie(context, message)
	}

	msrv.mu.Unlock()
	return nil
}

func (msrv MovieMicroService) GetMovie(context context.Context, req *MovieService.GetMovieMessage, res *MovieService.GetMovieResponse) error {
	m, ok := msrv.MovieRepository[req.MovieID]

	if ok {
		res.MovieID = req.MovieID
		res.Title = m.title
		return nil
	}

	return fmt.Errorf("The movie could not be found.")
}
