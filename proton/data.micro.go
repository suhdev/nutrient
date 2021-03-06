// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: data.proto

package proton

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Client API for Business service

type BusinessService interface {
	Search(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*Result, error)
}

type businessService struct {
	c    client.Client
	name string
}

func NewBusinessService(name string, c client.Client) BusinessService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proton"
	}
	return &businessService{
		c:    c,
		name: name,
	}
}

func (c *businessService) Search(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Business.Search", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Business service

type BusinessHandler interface {
	Search(context.Context, *QueryRequest, *Result) error
}

func RegisterBusinessHandler(s server.Server, hdlr BusinessHandler, opts ...server.HandlerOption) error {
	type business interface {
		Search(ctx context.Context, in *QueryRequest, out *Result) error
	}
	type Business struct {
		business
	}
	h := &businessHandler{hdlr}
	return s.Handle(s.NewHandler(&Business{h}, opts...))
}

type businessHandler struct {
	BusinessHandler
}

func (h *businessHandler) Search(ctx context.Context, in *QueryRequest, out *Result) error {
	return h.BusinessHandler.Search(ctx, in, out)
}

// Client API for DataStore service

type DataStoreService interface {
	Search(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*Result, error)
	Add(ctx context.Context, in *Result, opts ...client.CallOption) (*empty.Empty, error)
}

type dataStoreService struct {
	c    client.Client
	name string
}

func NewDataStoreService(name string, c client.Client) DataStoreService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proton"
	}
	return &dataStoreService{
		c:    c,
		name: name,
	}
}

func (c *dataStoreService) Search(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "DataStore.Search", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) Add(ctx context.Context, in *Result, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "DataStore.Add", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataStore service

type DataStoreHandler interface {
	Search(context.Context, *QueryRequest, *Result) error
	Add(context.Context, *Result, *empty.Empty) error
}

func RegisterDataStoreHandler(s server.Server, hdlr DataStoreHandler, opts ...server.HandlerOption) error {
	type dataStore interface {
		Search(ctx context.Context, in *QueryRequest, out *Result) error
		Add(ctx context.Context, in *Result, out *empty.Empty) error
	}
	type DataStore struct {
		dataStore
	}
	h := &dataStoreHandler{hdlr}
	return s.Handle(s.NewHandler(&DataStore{h}, opts...))
}

type dataStoreHandler struct {
	DataStoreHandler
}

func (h *dataStoreHandler) Search(ctx context.Context, in *QueryRequest, out *Result) error {
	return h.DataStoreHandler.Search(ctx, in, out)
}

func (h *dataStoreHandler) Add(ctx context.Context, in *Result, out *empty.Empty) error {
	return h.DataStoreHandler.Add(ctx, in, out)
}
