// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tikvpb.proto

/*
	Package tikvpb is a generated protocol buffer package.

	It is generated from these files:
		tikvpb.proto

	It has these top-level messages:
*/
package tikvpb

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"

	coprocessor "github.com/pingcap/kvproto/pkg/coprocessor"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	raft_serverpb "github.com/pingcap/kvproto/pkg/raft_serverpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tikv service

type TikvClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvImport(ctx context.Context, in *kvrpcpb.ImportRequest, opts ...grpc.CallOption) (*kvrpcpb.ImportResponse, error)
	KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error)
	KvDeleteRange(ctx context.Context, in *kvrpcpb.DeleteRangeRequest, opts ...grpc.CallOption) (*kvrpcpb.DeleteRangeResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error)
	// SQL push down commands.
	Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error)
	CoprocessorStream(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (Tikv_CoprocessorStreamClient, error)
	// Raft commands (tikv <-> tikv).
	Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error)
	Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error)
	// Region commands.
	SplitRegion(ctx context.Context, in *kvrpcpb.SplitRegionRequest, opts ...grpc.CallOption) (*kvrpcpb.SplitRegionResponse, error)
	// transaction debugger commands.
	MvccGetByKey(ctx context.Context, in *kvrpcpb.MvccGetByKeyRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByKeyResponse, error)
	MvccGetByStartTs(ctx context.Context, in *kvrpcpb.MvccGetByStartTsRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByStartTsResponse, error)
}

type tikvClient struct {
	cc *grpc.ClientConn
}

func NewTikvClient(cc *grpc.ClientConn) TikvClient {
	return &tikvClient{cc}
}

func (c *tikvClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvScan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvPrewrite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvCommit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvImport(ctx context.Context, in *kvrpcpb.ImportRequest, opts ...grpc.CallOption) (*kvrpcpb.ImportResponse, error) {
	out := new(kvrpcpb.ImportResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvImport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error) {
	out := new(kvrpcpb.CleanupResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvCleanup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error) {
	out := new(kvrpcpb.BatchGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvBatchGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvBatchRollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error) {
	out := new(kvrpcpb.ScanLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvScanLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvResolveLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error) {
	out := new(kvrpcpb.GCResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvGC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvDeleteRange(ctx context.Context, in *kvrpcpb.DeleteRangeRequest, opts ...grpc.CallOption) (*kvrpcpb.DeleteRangeResponse, error) {
	out := new(kvrpcpb.DeleteRangeResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvDeleteRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawPut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error) {
	out := new(kvrpcpb.RawScanResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawScan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error) {
	out := new(coprocessor.Response)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/Coprocessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) CoprocessorStream(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (Tikv_CoprocessorStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tikv_serviceDesc.Streams[0], c.cc, "/tikvpb.Tikv/CoprocessorStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvCoprocessorStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tikv_CoprocessorStreamClient interface {
	Recv() (*coprocessor.Response, error)
	grpc.ClientStream
}

type tikvCoprocessorStreamClient struct {
	grpc.ClientStream
}

func (x *tikvCoprocessorStreamClient) Recv() (*coprocessor.Response, error) {
	m := new(coprocessor.Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tikv_serviceDesc.Streams[1], c.cc, "/tikvpb.Tikv/Raft", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvRaftClient{stream}
	return x, nil
}

type Tikv_RaftClient interface {
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvRaftClient struct {
	grpc.ClientStream
}

func (x *tikvRaftClient) Send(m *raft_serverpb.RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvRaftClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tikv_serviceDesc.Streams[2], c.cc, "/tikvpb.Tikv/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvSnapshotClient{stream}
	return x, nil
}

type Tikv_SnapshotClient interface {
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvSnapshotClient struct {
	grpc.ClientStream
}

func (x *tikvSnapshotClient) Send(m *raft_serverpb.SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvSnapshotClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) SplitRegion(ctx context.Context, in *kvrpcpb.SplitRegionRequest, opts ...grpc.CallOption) (*kvrpcpb.SplitRegionResponse, error) {
	out := new(kvrpcpb.SplitRegionResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/SplitRegion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) MvccGetByKey(ctx context.Context, in *kvrpcpb.MvccGetByKeyRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByKeyResponse, error) {
	out := new(kvrpcpb.MvccGetByKeyResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/MvccGetByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) MvccGetByStartTs(ctx context.Context, in *kvrpcpb.MvccGetByStartTsRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByStartTsResponse, error) {
	out := new(kvrpcpb.MvccGetByStartTsResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/MvccGetByStartTs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tikv service

type TikvServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvImport(context.Context, *kvrpcpb.ImportRequest) (*kvrpcpb.ImportResponse, error)
	KvCleanup(context.Context, *kvrpcpb.CleanupRequest) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(context.Context, *kvrpcpb.BatchGetRequest) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(context.Context, *kvrpcpb.ScanLockRequest) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(context.Context, *kvrpcpb.GCRequest) (*kvrpcpb.GCResponse, error)
	KvDeleteRange(context.Context, *kvrpcpb.DeleteRangeRequest) (*kvrpcpb.DeleteRangeResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(context.Context, *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error)
	// SQL push down commands.
	Coprocessor(context.Context, *coprocessor.Request) (*coprocessor.Response, error)
	CoprocessorStream(*coprocessor.Request, Tikv_CoprocessorStreamServer) error
	// Raft commands (tikv <-> tikv).
	Raft(Tikv_RaftServer) error
	Snapshot(Tikv_SnapshotServer) error
	// Region commands.
	SplitRegion(context.Context, *kvrpcpb.SplitRegionRequest) (*kvrpcpb.SplitRegionResponse, error)
	// transaction debugger commands.
	MvccGetByKey(context.Context, *kvrpcpb.MvccGetByKeyRequest) (*kvrpcpb.MvccGetByKeyResponse, error)
	MvccGetByStartTs(context.Context, *kvrpcpb.MvccGetByStartTsRequest) (*kvrpcpb.MvccGetByStartTsResponse, error)
}

func RegisterTikvServer(s *grpc.Server, srv TikvServer) {
	s.RegisterService(&_Tikv_serviceDesc, srv)
}

func _Tikv_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvImport(ctx, req.(*kvrpcpb.ImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCleanup(ctx, req.(*kvrpcpb.CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchGet(ctx, req.(*kvrpcpb.BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScanLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScanLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScanLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScanLock(ctx, req.(*kvrpcpb.ScanLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvGC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvGC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvGC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvGC(ctx, req.(*kvrpcpb.GCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvDeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.DeleteRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvDeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvDeleteRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvDeleteRange(ctx, req.(*kvrpcpb.DeleteRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawScan(ctx, req.(*kvrpcpb.RawScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_Coprocessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(coprocessor.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).Coprocessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/Coprocessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).Coprocessor(ctx, req.(*coprocessor.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_CoprocessorStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(coprocessor.Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TikvServer).CoprocessorStream(m, &tikvCoprocessorStreamServer{stream})
}

type Tikv_CoprocessorStreamServer interface {
	Send(*coprocessor.Response) error
	grpc.ServerStream
}

type tikvCoprocessorStreamServer struct {
	grpc.ServerStream
}

func (x *tikvCoprocessorStreamServer) Send(m *coprocessor.Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Tikv_Raft_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Raft(&tikvRaftServer{stream})
}

type Tikv_RaftServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.RaftMessage, error)
	grpc.ServerStream
}

type tikvRaftServer struct {
	grpc.ServerStream
}

func (x *tikvRaftServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvRaftServer) Recv() (*raft_serverpb.RaftMessage, error) {
	m := new(raft_serverpb.RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tikv_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Snapshot(&tikvSnapshotServer{stream})
}

type Tikv_SnapshotServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.SnapshotChunk, error)
	grpc.ServerStream
}

type tikvSnapshotServer struct {
	grpc.ServerStream
}

func (x *tikvSnapshotServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvSnapshotServer) Recv() (*raft_serverpb.SnapshotChunk, error) {
	m := new(raft_serverpb.SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tikv_SplitRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SplitRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).SplitRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/SplitRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).SplitRegion(ctx, req.(*kvrpcpb.SplitRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_MvccGetByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.MvccGetByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).MvccGetByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/MvccGetByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).MvccGetByKey(ctx, req.(*kvrpcpb.MvccGetByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_MvccGetByStartTs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.MvccGetByStartTsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).MvccGetByStartTs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/MvccGetByStartTs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).MvccGetByStartTs(ctx, req.(*kvrpcpb.MvccGetByStartTsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tikv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tikvpb.Tikv",
	HandlerType: (*TikvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _Tikv_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _Tikv_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _Tikv_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _Tikv_KvCommit_Handler,
		},
		{
			MethodName: "KvImport",
			Handler:    _Tikv_KvImport_Handler,
		},
		{
			MethodName: "KvCleanup",
			Handler:    _Tikv_KvCleanup_Handler,
		},
		{
			MethodName: "KvBatchGet",
			Handler:    _Tikv_KvBatchGet_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _Tikv_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvScanLock",
			Handler:    _Tikv_KvScanLock_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _Tikv_KvResolveLock_Handler,
		},
		{
			MethodName: "KvGC",
			Handler:    _Tikv_KvGC_Handler,
		},
		{
			MethodName: "KvDeleteRange",
			Handler:    _Tikv_KvDeleteRange_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _Tikv_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _Tikv_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _Tikv_RawDelete_Handler,
		},
		{
			MethodName: "RawScan",
			Handler:    _Tikv_RawScan_Handler,
		},
		{
			MethodName: "Coprocessor",
			Handler:    _Tikv_Coprocessor_Handler,
		},
		{
			MethodName: "SplitRegion",
			Handler:    _Tikv_SplitRegion_Handler,
		},
		{
			MethodName: "MvccGetByKey",
			Handler:    _Tikv_MvccGetByKey_Handler,
		},
		{
			MethodName: "MvccGetByStartTs",
			Handler:    _Tikv_MvccGetByStartTs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CoprocessorStream",
			Handler:       _Tikv_CoprocessorStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Raft",
			Handler:       _Tikv_Raft_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Snapshot",
			Handler:       _Tikv_Snapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tikvpb.proto",
}

func init() { proto.RegisterFile("tikvpb.proto", fileDescriptorTikvpb) }

var fileDescriptorTikvpb = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x57, 0x69, 0x94, 0xcd, 0xdb, 0x04, 0x73, 0x0b, 0x74, 0x61, 0xab, 0x60, 0x27, 0xc4,
	0x21, 0xfc, 0x95, 0x38, 0x0c, 0x10, 0x34, 0x95, 0x2a, 0xc8, 0x2a, 0x55, 0xe9, 0x2e, 0x9c, 0x90,
	0x1b, 0xbd, 0x4b, 0xa3, 0xa4, 0x71, 0x70, 0x5c, 0x57, 0xfb, 0x26, 0x5c, 0xf9, 0x36, 0x1c, 0xf9,
	0x08, 0xa8, 0x7c, 0x11, 0x94, 0xa4, 0x76, 0xec, 0xa6, 0x45, 0xe2, 0xd4, 0xe4, 0xf7, 0xbc, 0xcf,
	0x93, 0xb7, 0xce, 0x6b, 0x07, 0x1d, 0xf2, 0x30, 0x12, 0xe9, 0xc4, 0x4e, 0x19, 0xe5, 0x14, 0x37,
	0xcb, 0x3b, 0xeb, 0xd8, 0xa7, 0x29, 0xa3, 0x3e, 0x64, 0x19, 0x65, 0xa5, 0x64, 0x1d, 0x45, 0x82,
	0xa5, 0xbe, 0xac, 0xb4, 0x5a, 0x8c, 0x5c, 0xf3, 0xaf, 0x19, 0x30, 0x01, 0x4c, 0xc1, 0x76, 0x40,
	0x03, 0x5a, 0x5c, 0x3e, 0xcb, 0xaf, 0x4a, 0xfa, 0xf2, 0xc7, 0x21, 0xda, 0xbd, 0x0a, 0x23, 0x81,
	0x5f, 0xa3, 0x5b, 0xae, 0x18, 0x00, 0xc7, 0x2d, 0x5b, 0x86, 0x0d, 0x80, 0x7b, 0xf0, 0x6d, 0x0e,
	0x19, 0xb7, 0xda, 0x26, 0xcc, 0x52, 0x9a, 0x64, 0x70, 0xbe, 0x83, 0xdf, 0xa0, 0xa6, 0x2b, 0xc6,
	0x3e, 0x49, 0x70, 0x55, 0x91, 0xdf, 0x4a, 0xdf, 0xbd, 0x35, 0xaa, 0x8c, 0x0e, 0x42, 0xae, 0x18,
	0x31, 0x58, 0xb0, 0x90, 0x03, 0xee, 0xa8, 0x32, 0x89, 0x64, 0xc0, 0xc9, 0x06, 0x45, 0x85, 0xbc,
	0x43, 0x7b, 0xae, 0x70, 0xe8, 0x6c, 0x16, 0x72, 0x7c, 0x5f, 0x15, 0x96, 0x40, 0x06, 0x3c, 0xa8,
	0x71, 0xd3, 0xfe, 0x69, 0x96, 0x52, 0xa6, 0xdb, 0x4b, 0x50, 0xb7, 0x4b, 0xae, 0xec, 0x1f, 0xd0,
	0xbe, 0x2b, 0x9c, 0x18, 0x48, 0x32, 0x4f, 0xb1, 0xf6, 0x98, 0x92, 0xc8, 0x80, 0x4e, 0x5d, 0x30,
	0x17, 0xa1, 0x47, 0xb8, 0x3f, 0xcd, 0x17, 0xbe, 0xaa, 0x94, 0xa8, 0xbe, 0x08, 0x95, 0xa2, 0x42,
	0x3c, 0x74, 0x67, 0x15, 0xe2, 0xd1, 0x38, 0x9e, 0x10, 0x3f, 0xc2, 0x67, 0x66, 0xbd, 0xe4, 0x32,
	0xae, 0xbb, 0x4d, 0x36, 0x1b, 0xcb, 0xdf, 0xd8, 0x25, 0xf5, 0x23, 0xad, 0x31, 0x89, 0xea, 0x8d,
	0x55, 0x8a, 0x0a, 0xb9, 0x44, 0x47, 0xae, 0xf0, 0x20, 0xa3, 0xb1, 0x80, 0x22, 0xe7, 0xa1, 0xaa,
	0xd6, 0xa8, 0x8c, 0x3a, 0xdd, 0x2c, 0xaa, 0xb4, 0x17, 0x68, 0xd7, 0x15, 0x03, 0x07, 0xe3, 0x6a,
	0x12, 0x1d, 0xe9, 0x6d, 0x19, 0xcc, 0x6c, 0xa0, 0x0f, 0x31, 0x70, 0xf0, 0x48, 0x12, 0x80, 0xd6,
	0x80, 0x46, 0xeb, 0x0d, 0x18, 0xa2, 0x4a, 0xbb, 0x40, 0x4d, 0x8f, 0x2c, 0xf2, 0x17, 0x55, 0xcd,
	0x4a, 0x09, 0xea, 0xb3, 0x22, 0xf9, 0x9a, 0x79, 0x34, 0x5f, 0x33, 0x8f, 0xe6, 0x9b, 0xcd, 0x05,
	0x57, 0xe6, 0x3e, 0xda, 0xf7, 0xc8, 0xa2, 0xec, 0x0a, 0x9f, 0xe8, 0x75, 0xab, 0x4e, 0x57, 0x11,
	0xd6, 0x26, 0x49, 0xa5, 0xbc, 0x47, 0xb7, 0x3d, 0xb2, 0x28, 0xf6, 0xaa, 0xf1, 0x2c, 0x7d, 0xbb,
	0x76, 0xea, 0x82, 0xf2, 0xbf, 0x45, 0x07, 0x4e, 0x75, 0xf0, 0xe0, 0xb6, 0xad, 0x1f, 0x43, 0xd5,
	0x7e, 0x37, 0xa9, 0xf6, 0x1f, 0x8e, 0x35, 0xf7, 0x98, 0x33, 0x20, 0xb3, 0xff, 0xcc, 0x78, 0xde,
	0xc0, 0x17, 0x68, 0xd7, 0x23, 0xd7, 0x1c, 0x5b, 0xb6, 0x79, 0xc2, 0xe5, 0x70, 0x08, 0x59, 0x46,
	0x02, 0xb0, 0x5a, 0x6b, 0x5a, 0x9f, 0x26, 0x70, 0xbe, 0xf3, 0xa4, 0x81, 0x3f, 0xa2, 0xbd, 0x71,
	0x42, 0xd2, 0x6c, 0x4a, 0x39, 0x3e, 0x5d, 0x2b, 0x92, 0x82, 0x33, 0x9d, 0x27, 0xd1, 0xf6, 0x88,
	0xcf, 0xe8, 0x60, 0x9c, 0xc6, 0xf9, 0x21, 0x12, 0x84, 0x34, 0xd1, 0xe6, 0x49, 0xa3, 0xf5, 0x79,
	0x32, 0x44, 0xb5, 0x22, 0x43, 0x74, 0x38, 0x14, 0xbe, 0x3f, 0x00, 0xde, 0xbb, 0x71, 0xe1, 0x06,
	0x57, 0xf5, 0x3a, 0x96, 0x69, 0x67, 0x5b, 0x54, 0x15, 0xf7, 0x05, 0xdd, 0x55, 0xca, 0x98, 0x13,
	0xc6, 0xaf, 0x32, 0xfc, 0xa8, 0x6e, 0x5a, 0x49, 0x32, 0xf6, 0xf1, 0x3f, 0x2a, 0x64, 0x74, 0xef,
	0xe9, 0xcf, 0x65, 0xb7, 0xf1, 0x6b, 0xd9, 0x6d, 0xfc, 0x5e, 0x76, 0x1b, 0xdf, 0xff, 0x74, 0x77,
	0x50, 0xc7, 0xa7, 0x33, 0x3b, 0x0d, 0x93, 0xc0, 0x27, 0xa9, 0x9d, 0x7f, 0x96, 0xec, 0x48, 0x14,
	0xdf, 0x93, 0x49, 0xb3, 0xf8, 0x79, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xea, 0xa8, 0xe1, 0xc4,
	0xbb, 0x06, 0x00, 0x00,
}
