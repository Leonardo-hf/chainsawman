// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: algo.proto

package algo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AlgoClient is the client API for Algo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlgoClient interface {
	CreateAlgo(ctx context.Context, in *CreateAlgoReq, opts ...grpc.CallOption) (*AlgoReply, error)
	QueryAlgo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AlgoReply, error)
	DropAlgo(ctx context.Context, in *DropAlgoReq, opts ...grpc.CallOption) (*AlgoReply, error)
	Degree(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*RankReply, error)
	Pagerank(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*RankReply, error)
	Louvain(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*ClusterReply, error)
	ShortestPath(ctx context.Context, in *ShortestPathReq, opts ...grpc.CallOption) (*ClusterReply, error)
	AvgShortestPath(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*MetricsReply, error)
	AvgClustering(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*MetricsReply, error)
	Custom(ctx context.Context, in *CustomAlgoReq, opts ...grpc.CallOption) (*CustomAlgoReply, error)
}

type algoClient struct {
	cc grpc.ClientConnInterface
}

func NewAlgoClient(cc grpc.ClientConnInterface) AlgoClient {
	return &algoClient{cc}
}

func (c *algoClient) CreateAlgo(ctx context.Context, in *CreateAlgoReq, opts ...grpc.CallOption) (*AlgoReply, error) {
	out := new(AlgoReply)
	err := c.cc.Invoke(ctx, "/services.algo/createAlgo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) QueryAlgo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AlgoReply, error) {
	out := new(AlgoReply)
	err := c.cc.Invoke(ctx, "/services.algo/queryAlgo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) DropAlgo(ctx context.Context, in *DropAlgoReq, opts ...grpc.CallOption) (*AlgoReply, error) {
	out := new(AlgoReply)
	err := c.cc.Invoke(ctx, "/services.algo/dropAlgo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) Degree(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*RankReply, error) {
	out := new(RankReply)
	err := c.cc.Invoke(ctx, "/services.algo/degree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) Pagerank(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*RankReply, error) {
	out := new(RankReply)
	err := c.cc.Invoke(ctx, "/services.algo/pagerank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) Louvain(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := c.cc.Invoke(ctx, "/services.algo/louvain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) ShortestPath(ctx context.Context, in *ShortestPathReq, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := c.cc.Invoke(ctx, "/services.algo/shortestPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) AvgShortestPath(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*MetricsReply, error) {
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, "/services.algo/avgShortestPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) AvgClustering(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*MetricsReply, error) {
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, "/services.algo/avgClustering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) Custom(ctx context.Context, in *CustomAlgoReq, opts ...grpc.CallOption) (*CustomAlgoReply, error) {
	out := new(CustomAlgoReply)
	err := c.cc.Invoke(ctx, "/services.algo/custom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgoServer is the server API for Algo service.
// All implementations must embed UnimplementedAlgoServer
// for forward compatibility
type AlgoServer interface {
	CreateAlgo(context.Context, *CreateAlgoReq) (*AlgoReply, error)
	QueryAlgo(context.Context, *Empty) (*AlgoReply, error)
	DropAlgo(context.Context, *DropAlgoReq) (*AlgoReply, error)
	Degree(context.Context, *BaseReq) (*RankReply, error)
	Pagerank(context.Context, *BaseReq) (*RankReply, error)
	Louvain(context.Context, *BaseReq) (*ClusterReply, error)
	ShortestPath(context.Context, *ShortestPathReq) (*ClusterReply, error)
	AvgShortestPath(context.Context, *BaseReq) (*MetricsReply, error)
	AvgClustering(context.Context, *BaseReq) (*MetricsReply, error)
	Custom(context.Context, *CustomAlgoReq) (*CustomAlgoReply, error)
	mustEmbedUnimplementedAlgoServer()
}

// UnimplementedAlgoServer must be embedded to have forward compatible implementations.
type UnimplementedAlgoServer struct {
}

func (UnimplementedAlgoServer) CreateAlgo(context.Context, *CreateAlgoReq) (*AlgoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlgo not implemented")
}
func (UnimplementedAlgoServer) QueryAlgo(context.Context, *Empty) (*AlgoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAlgo not implemented")
}
func (UnimplementedAlgoServer) DropAlgo(context.Context, *DropAlgoReq) (*AlgoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropAlgo not implemented")
}
func (UnimplementedAlgoServer) Degree(context.Context, *BaseReq) (*RankReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Degree not implemented")
}
func (UnimplementedAlgoServer) Pagerank(context.Context, *BaseReq) (*RankReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pagerank not implemented")
}
func (UnimplementedAlgoServer) Louvain(context.Context, *BaseReq) (*ClusterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Louvain not implemented")
}
func (UnimplementedAlgoServer) ShortestPath(context.Context, *ShortestPathReq) (*ClusterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortestPath not implemented")
}
func (UnimplementedAlgoServer) AvgShortestPath(context.Context, *BaseReq) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvgShortestPath not implemented")
}
func (UnimplementedAlgoServer) AvgClustering(context.Context, *BaseReq) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvgClustering not implemented")
}
func (UnimplementedAlgoServer) Custom(context.Context, *CustomAlgoReq) (*CustomAlgoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Custom not implemented")
}
func (UnimplementedAlgoServer) mustEmbedUnimplementedAlgoServer() {}

// UnsafeAlgoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlgoServer will
// result in compilation errors.
type UnsafeAlgoServer interface {
	mustEmbedUnimplementedAlgoServer()
}

func RegisterAlgoServer(s grpc.ServiceRegistrar, srv AlgoServer) {
	s.RegisterService(&Algo_ServiceDesc, srv)
}

func _Algo_CreateAlgo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlgoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).CreateAlgo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/createAlgo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).CreateAlgo(ctx, req.(*CreateAlgoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_QueryAlgo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).QueryAlgo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/queryAlgo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).QueryAlgo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_DropAlgo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropAlgoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).DropAlgo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/dropAlgo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).DropAlgo(ctx, req.(*DropAlgoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_Degree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).Degree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/degree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).Degree(ctx, req.(*BaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_Pagerank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).Pagerank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/pagerank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).Pagerank(ctx, req.(*BaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_Louvain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).Louvain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/louvain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).Louvain(ctx, req.(*BaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_ShortestPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortestPathReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).ShortestPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/shortestPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).ShortestPath(ctx, req.(*ShortestPathReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_AvgShortestPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).AvgShortestPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/avgShortestPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).AvgShortestPath(ctx, req.(*BaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_AvgClustering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).AvgClustering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/avgClustering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).AvgClustering(ctx, req.(*BaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_Custom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomAlgoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).Custom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.algo/custom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).Custom(ctx, req.(*CustomAlgoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Algo_ServiceDesc is the grpc.ServiceDesc for Algo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Algo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.algo",
	HandlerType: (*AlgoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAlgo",
			Handler:    _Algo_CreateAlgo_Handler,
		},
		{
			MethodName: "queryAlgo",
			Handler:    _Algo_QueryAlgo_Handler,
		},
		{
			MethodName: "dropAlgo",
			Handler:    _Algo_DropAlgo_Handler,
		},
		{
			MethodName: "degree",
			Handler:    _Algo_Degree_Handler,
		},
		{
			MethodName: "pagerank",
			Handler:    _Algo_Pagerank_Handler,
		},
		{
			MethodName: "louvain",
			Handler:    _Algo_Louvain_Handler,
		},
		{
			MethodName: "shortestPath",
			Handler:    _Algo_ShortestPath_Handler,
		},
		{
			MethodName: "avgShortestPath",
			Handler:    _Algo_AvgShortestPath_Handler,
		},
		{
			MethodName: "avgClustering",
			Handler:    _Algo_AvgClustering_Handler,
		},
		{
			MethodName: "custom",
			Handler:    _Algo_Custom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algo.proto",
}
