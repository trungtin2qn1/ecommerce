// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.12
// source: v1/buyer/buyer.proto

package buyer

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBuyerSelf = "/api.v1.auth.buyer.Buyer/Self"

type BuyerHTTPServer interface {
	Self(context.Context, *SelfRequest) (*SelfReply, error)
}

func RegisterBuyerHTTPServer(s *http.Server, srv BuyerHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/auth/buyers/self", _Buyer_Self1_HTTP_Handler(srv))
}

func _Buyer_Self1_HTTP_Handler(srv BuyerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SelfRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBuyerSelf)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Self(ctx, req.(*SelfRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SelfReply)
		return ctx.Result(200, reply)
	}
}

type BuyerHTTPClient interface {
	Self(ctx context.Context, req *SelfRequest, opts ...http.CallOption) (rsp *SelfReply, err error)
}

type BuyerHTTPClientImpl struct {
	cc *http.Client
}

func NewBuyerHTTPClient(client *http.Client) BuyerHTTPClient {
	return &BuyerHTTPClientImpl{client}
}

func (c *BuyerHTTPClientImpl) Self(ctx context.Context, in *SelfRequest, opts ...http.CallOption) (*SelfReply, error) {
	var out SelfReply
	pattern := "/v1/auth/buyers/self"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBuyerSelf))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
