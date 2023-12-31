// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: party.proto

package protocolconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	gen_go "github.com/mrbanja/watchparty-proto/gen-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PartyServiceName is the fully-qualified name of the PartyService service.
	PartyServiceName = "com.mrbanja.watchparty.PartyService"
)

// PartyServiceClient is a client for the com.mrbanja.watchparty.PartyService service.
type PartyServiceClient interface {
	GetMagnet(context.Context, *connect_go.Request[gen_go.GetMagnetRequest]) (*connect_go.Response[gen_go.GetMagnetResponse], error)
	JoinRoom(context.Context) *connect_go.BidiStreamForClient[gen_go.RoomRequest, gen_go.RoomResponse]
}

// NewPartyServiceClient constructs a client for the com.mrbanja.watchparty.PartyService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPartyServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PartyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &partyServiceClient{
		getMagnet: connect_go.NewClient[gen_go.GetMagnetRequest, gen_go.GetMagnetResponse](
			httpClient,
			baseURL+"/com.mrbanja.watchparty.PartyService/GetMagnet",
			opts...,
		),
		joinRoom: connect_go.NewClient[gen_go.RoomRequest, gen_go.RoomResponse](
			httpClient,
			baseURL+"/com.mrbanja.watchparty.PartyService/JoinRoom",
			opts...,
		),
	}
}

// partyServiceClient implements PartyServiceClient.
type partyServiceClient struct {
	getMagnet *connect_go.Client[gen_go.GetMagnetRequest, gen_go.GetMagnetResponse]
	joinRoom  *connect_go.Client[gen_go.RoomRequest, gen_go.RoomResponse]
}

// GetMagnet calls com.mrbanja.watchparty.PartyService.GetMagnet.
func (c *partyServiceClient) GetMagnet(ctx context.Context, req *connect_go.Request[gen_go.GetMagnetRequest]) (*connect_go.Response[gen_go.GetMagnetResponse], error) {
	return c.getMagnet.CallUnary(ctx, req)
}

// JoinRoom calls com.mrbanja.watchparty.PartyService.JoinRoom.
func (c *partyServiceClient) JoinRoom(ctx context.Context) *connect_go.BidiStreamForClient[gen_go.RoomRequest, gen_go.RoomResponse] {
	return c.joinRoom.CallBidiStream(ctx)
}

// PartyServiceHandler is an implementation of the com.mrbanja.watchparty.PartyService service.
type PartyServiceHandler interface {
	GetMagnet(context.Context, *connect_go.Request[gen_go.GetMagnetRequest]) (*connect_go.Response[gen_go.GetMagnetResponse], error)
	JoinRoom(context.Context, *connect_go.BidiStream[gen_go.RoomRequest, gen_go.RoomResponse]) error
}

// NewPartyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPartyServiceHandler(svc PartyServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/com.mrbanja.watchparty.PartyService/GetMagnet", connect_go.NewUnaryHandler(
		"/com.mrbanja.watchparty.PartyService/GetMagnet",
		svc.GetMagnet,
		opts...,
	))
	mux.Handle("/com.mrbanja.watchparty.PartyService/JoinRoom", connect_go.NewBidiStreamHandler(
		"/com.mrbanja.watchparty.PartyService/JoinRoom",
		svc.JoinRoom,
		opts...,
	))
	return "/com.mrbanja.watchparty.PartyService/", mux
}

// UnimplementedPartyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPartyServiceHandler struct{}

func (UnimplementedPartyServiceHandler) GetMagnet(context.Context, *connect_go.Request[gen_go.GetMagnetRequest]) (*connect_go.Response[gen_go.GetMagnetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("com.mrbanja.watchparty.PartyService.GetMagnet is not implemented"))
}

func (UnimplementedPartyServiceHandler) JoinRoom(context.Context, *connect_go.BidiStream[gen_go.RoomRequest, gen_go.RoomResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("com.mrbanja.watchparty.PartyService.JoinRoom is not implemented"))
}
