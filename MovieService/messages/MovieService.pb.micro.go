// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: MovieService.proto

package MovieService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MovieService service

type MovieService interface {
	CreateMovie(ctx context.Context, in *CreateMovieMessage, opts ...client.CallOption) (*CreateMovieResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieMessage, opts ...client.CallOption) (*DeleteMovieResponse, error)
	GetMovie(ctx context.Context, in *GetMovieMessage, opts ...client.CallOption) (*GetMovieResponse, error)
}

type movieService struct {
	c    client.Client
	name string
}

func NewMovieService(name string, c client.Client) MovieService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "movieservice"
	}
	return &movieService{
		c:    c,
		name: name,
	}
}

func (c *movieService) CreateMovie(ctx context.Context, in *CreateMovieMessage, opts ...client.CallOption) (*CreateMovieResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.CreateMovie", in)
	out := new(CreateMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) DeleteMovie(ctx context.Context, in *DeleteMovieMessage, opts ...client.CallOption) (*DeleteMovieResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.DeleteMovie", in)
	out := new(DeleteMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) GetMovie(ctx context.Context, in *GetMovieMessage, opts ...client.CallOption) (*GetMovieResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.GetMovie", in)
	out := new(GetMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceHandler interface {
	CreateMovie(context.Context, *CreateMovieMessage, *CreateMovieResponse) error
	DeleteMovie(context.Context, *DeleteMovieMessage, *DeleteMovieResponse) error
	GetMovie(context.Context, *GetMovieMessage, *GetMovieResponse) error
}

func RegisterMovieServiceHandler(s server.Server, hdlr MovieServiceHandler, opts ...server.HandlerOption) error {
	type movieService interface {
		CreateMovie(ctx context.Context, in *CreateMovieMessage, out *CreateMovieResponse) error
		DeleteMovie(ctx context.Context, in *DeleteMovieMessage, out *DeleteMovieResponse) error
		GetMovie(ctx context.Context, in *GetMovieMessage, out *GetMovieResponse) error
	}
	type MovieService struct {
		movieService
	}
	h := &movieServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MovieService{h}, opts...))
}

type movieServiceHandler struct {
	MovieServiceHandler
}

func (h *movieServiceHandler) CreateMovie(ctx context.Context, in *CreateMovieMessage, out *CreateMovieResponse) error {
	return h.MovieServiceHandler.CreateMovie(ctx, in, out)
}

func (h *movieServiceHandler) DeleteMovie(ctx context.Context, in *DeleteMovieMessage, out *DeleteMovieResponse) error {
	return h.MovieServiceHandler.DeleteMovie(ctx, in, out)
}

func (h *movieServiceHandler) GetMovie(ctx context.Context, in *GetMovieMessage, out *GetMovieResponse) error {
	return h.MovieServiceHandler.GetMovie(ctx, in, out)
}
