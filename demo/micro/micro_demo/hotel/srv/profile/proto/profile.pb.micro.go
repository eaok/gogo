// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: profile.proto

package profile

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "context"

	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for Profile service

type ProfileService interface {
	GetProfiles(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error)
}

type profileService struct {
	c    client.Client
	name string
}

func NewProfileService(name string, c client.Client) ProfileService {
	return &profileService{
		c:    c,
		name: name,
	}
}

func (c *profileService) GetProfiles(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Profile.GetProfiles", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileHandler interface {
	GetProfiles(context.Context, *Request, *Result) error
}

func RegisterProfileHandler(s server.Server, hdlr ProfileHandler, opts ...server.HandlerOption) error {
	type profile interface {
		GetProfiles(ctx context.Context, in *Request, out *Result) error
	}
	type Profile struct {
		profile
	}
	h := &profileHandler{hdlr}
	return s.Handle(s.NewHandler(&Profile{h}, opts...))
}

type profileHandler struct {
	ProfileHandler
}

func (h *profileHandler) GetProfiles(ctx context.Context, in *Request, out *Result) error {
	return h.ProfileHandler.GetProfiles(ctx, in, out)
}
