// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/repository/repository.proto

/*
	Package repository is a generated protocol buffer package.

	Repository Service

	Repository Service API performs CRUD actions against repository resources

	It is generated from these files:
		server/repository/repository.proto

	It has these top-level messages:
		RepoQuery
		RepoResponse
		RepoUpdateRequest
*/
package repository

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "k8s.io/api/core/v1"
import github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// RepoQuery is a query for Repository resources
type RepoQuery struct {
	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (m *RepoQuery) Reset()                    { *m = RepoQuery{} }
func (m *RepoQuery) String() string            { return proto.CompactTextString(m) }
func (*RepoQuery) ProtoMessage()               {}
func (*RepoQuery) Descriptor() ([]byte, []int) { return fileDescriptorRepository, []int{0} }

func (m *RepoQuery) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

type RepoResponse struct {
}

func (m *RepoResponse) Reset()                    { *m = RepoResponse{} }
func (m *RepoResponse) String() string            { return proto.CompactTextString(m) }
func (*RepoResponse) ProtoMessage()               {}
func (*RepoResponse) Descriptor() ([]byte, []int) { return fileDescriptorRepository, []int{1} }

type RepoUpdateRequest struct {
	Url  string                                                                `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Repo *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
}

func (m *RepoUpdateRequest) Reset()                    { *m = RepoUpdateRequest{} }
func (m *RepoUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*RepoUpdateRequest) ProtoMessage()               {}
func (*RepoUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptorRepository, []int{2} }

func (m *RepoUpdateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RepoUpdateRequest) GetRepo() *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func init() {
	proto.RegisterType((*RepoQuery)(nil), "repository.RepoQuery")
	proto.RegisterType((*RepoResponse)(nil), "repository.RepoResponse")
	proto.RegisterType((*RepoUpdateRequest)(nil), "repository.RepoUpdateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepositoryService service

type RepositoryServiceClient interface {
	// List returns list of repos
	List(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.RepositoryList, error)
	// Create creates a repo
	Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Get returns a repo by name
	Get(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Update updates a repo
	Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Update updates a repo (special handler intended to be used only by the gRPC gateway)
	UpdateREST(ctx context.Context, in *RepoUpdateRequest, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Delete updates a repo
	Delete(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*RepoResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) List(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.RepositoryList, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.RepositoryList)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) Get(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) UpdateREST(ctx context.Context, in *RepoUpdateRequest, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/UpdateREST", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) Delete(ctx context.Context, in *RepoQuery, opts ...grpc.CallOption) (*RepoResponse, error) {
	out := new(RepoResponse)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepositoryService service

type RepositoryServiceServer interface {
	// List returns list of repos
	List(context.Context, *RepoQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.RepositoryList, error)
	// Create creates a repo
	Create(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Get returns a repo by name
	Get(context.Context, *RepoQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Update updates a repo
	Update(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Update updates a repo (special handler intended to be used only by the gRPC gateway)
	UpdateREST(context.Context, *RepoUpdateRequest) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository, error)
	// Delete updates a repo
	Delete(context.Context, *RepoQuery) (*RepoResponse, error)
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).List(ctx, req.(*RepoQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Create(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Get(ctx, req.(*RepoQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Update(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_UpdateREST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).UpdateREST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/UpdateREST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).UpdateREST(ctx, req.(*RepoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Delete(ctx, req.(*RepoQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repository.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RepositoryService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RepositoryService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RepositoryService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RepositoryService_Update_Handler,
		},
		{
			MethodName: "UpdateREST",
			Handler:    _RepositoryService_UpdateREST_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RepositoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/repository/repository.proto",
}

func (m *RepoQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Repo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRepository(dAtA, i, uint64(len(m.Repo)))
		i += copy(dAtA[i:], m.Repo)
	}
	return i, nil
}

func (m *RepoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RepoUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRepository(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.Repo != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRepository(dAtA, i, uint64(m.Repo.Size()))
		n1, err := m.Repo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintRepository(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RepoQuery) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovRepository(uint64(l))
	}
	return n
}

func (m *RepoResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RepoUpdateRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovRepository(uint64(l))
	}
	if m.Repo != nil {
		l = m.Repo.Size()
		n += 1 + l + sovRepository(uint64(l))
	}
	return n
}

func sovRepository(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRepository(x uint64) (n int) {
	return sovRepository(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RepoQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepository
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RepoQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepository
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRepository
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepository(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepository
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RepoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepository
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RepoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRepository(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepository
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RepoUpdateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepository
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RepoUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepository
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRepository
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepository
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRepository
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repo == nil {
				m.Repo = &github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository{}
			}
			if err := m.Repo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepository(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepository
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRepository(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRepository
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRepository
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRepository
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRepository
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRepository
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRepository(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRepository = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRepository   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/repository/repository.proto", fileDescriptorRepository) }

var fileDescriptorRepository = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x4d, 0x5b, 0x06, 0x1a, 0x45, 0x34, 0x54, 0xa9, 0x63, 0xbb, 0x2d, 0xf1, 0xb2, 0x14,
	0x9a, 0xb0, 0xf5, 0x22, 0xf5, 0xa6, 0x16, 0x29, 0x78, 0x71, 0xaa, 0x07, 0x3d, 0x28, 0xe9, 0xcc,
	0x63, 0x1a, 0x77, 0x9c, 0xc4, 0x24, 0x33, 0x50, 0xa4, 0x20, 0x1e, 0xc4, 0xbb, 0x17, 0x0f, 0x7e,
	0x0f, 0xbf, 0x82, 0x47, 0xc1, 0x2f, 0x20, 0x8b, 0xdf, 0xc2, 0x8b, 0x24, 0x33, 0xdd, 0x5d, 0xdb,
	0x6d, 0x4f, 0x73, 0xe8, 0xed, 0x3f, 0x6f, 0x92, 0x97, 0x5f, 0xfe, 0x2f, 0xef, 0x61, 0x6a, 0xc1,
	0xd4, 0x60, 0xb8, 0x01, 0xad, 0xac, 0x74, 0xca, 0x1c, 0x4e, 0x49, 0xa6, 0x8d, 0x72, 0x8a, 0xe0,
	0x49, 0x24, 0x5e, 0xca, 0x55, 0xae, 0x42, 0x98, 0x7b, 0xd5, 0xac, 0x88, 0x57, 0x72, 0xa5, 0xf2,
	0x02, 0xb8, 0xd0, 0x92, 0x8b, 0xb2, 0x54, 0x4e, 0x38, 0xa9, 0x4a, 0xdb, 0xfe, 0xa5, 0xc3, 0x7b,
	0x96, 0x49, 0x15, 0xfe, 0xa6, 0xca, 0x00, 0xaf, 0x07, 0x3c, 0x87, 0x12, 0x8c, 0x70, 0x90, 0xb5,
	0x6b, 0x76, 0x73, 0xe9, 0x0e, 0xaa, 0x7d, 0x96, 0xaa, 0xb7, 0x5c, 0x98, 0x70, 0xc4, 0x9b, 0x20,
	0x36, 0xd3, 0x8c, 0xeb, 0x61, 0xee, 0x37, 0x5b, 0x2e, 0xb4, 0x2e, 0x64, 0x1a, 0x92, 0xf3, 0x7a,
	0x20, 0x0a, 0x7d, 0x20, 0x4e, 0xa5, 0xa2, 0x6b, 0x78, 0x31, 0x01, 0xad, 0x9e, 0x56, 0x60, 0x0e,
	0x09, 0xc1, 0x0b, 0x9e, 0x7e, 0x19, 0xad, 0xa3, 0xfe, 0x62, 0x12, 0x34, 0xbd, 0x8a, 0xaf, 0xf8,
	0x05, 0x09, 0x58, 0xad, 0x4a, 0x0b, 0xf4, 0x03, 0xc2, 0xd7, 0x7d, 0xe0, 0xb9, 0xce, 0x84, 0x83,
	0x04, 0xde, 0x55, 0x60, 0x1d, 0xb9, 0x86, 0xe7, 0x2b, 0x53, 0xb4, 0x1b, 0xbd, 0x24, 0x2f, 0xda,
	0x5c, 0x73, 0xeb, 0xa8, 0x7f, 0x79, 0x6b, 0x87, 0x4d, 0x90, 0xd9, 0x31, 0x72, 0x10, 0xaf, 0xd3,
	0x8c, 0xe9, 0x61, 0xce, 0x3c, 0x32, 0x9b, 0x42, 0x66, 0xc7, 0xc8, 0x2c, 0x19, 0x1b, 0xda, 0x20,
	0x6d, 0xfd, 0x8d, 0x1a, 0x84, 0x26, 0xb8, 0x07, 0xa6, 0x96, 0x29, 0x90, 0x4f, 0x08, 0x2f, 0x3c,
	0x91, 0xd6, 0x91, 0x1b, 0x6c, 0xaa, 0x28, 0xe3, 0xcb, 0xc5, 0xbb, 0x9d, 0x20, 0xf8, 0x13, 0xe8,
	0xca, 0xc7, 0x5f, 0x7f, 0xbe, 0xcc, 0xdd, 0x24, 0x4b, 0xa1, 0x4a, 0xf5, 0x60, 0xf2, 0x0a, 0x24,
	0x58, 0xf2, 0x1d, 0xe1, 0xe8, 0xa1, 0x01, 0xe1, 0x80, 0x74, 0x73, 0xed, 0xb8, 0x9b, 0x34, 0x74,
	0x2d, 0x60, 0xdf, 0xa2, 0x33, 0xb1, 0xb7, 0xd1, 0x06, 0xf9, 0x8c, 0xf0, 0xfc, 0x63, 0x38, 0xd3,
	0xc1, 0x8e, 0x30, 0xee, 0x04, 0x8c, 0x55, 0x72, 0x7b, 0x16, 0x06, 0x7f, 0xef, 0xbf, 0x8e, 0xc8,
	0x57, 0x84, 0xa3, 0xe6, 0x89, 0x5d, 0x30, 0x13, 0x2f, 0x91, 0x6f, 0x08, 0xe3, 0xf6, 0xf5, 0xef,
	0xec, 0x3d, 0x23, 0xab, 0x27, 0xcd, 0xfa, 0xaf, 0x33, 0xba, 0x3a, 0xb6, 0x1f, 0x4c, 0xa3, 0x71,
	0x3c, 0xdb, 0xb4, 0xca, 0x14, 0x47, 0xdb, 0xa1, 0x3b, 0xc8, 0x2b, 0x1c, 0x3d, 0x82, 0x02, 0x1c,
	0x9c, 0x55, 0xc6, 0xe5, 0x93, 0xe1, 0x71, 0x6f, 0xb7, 0x95, 0xd9, 0x38, 0xaf, 0x32, 0x0f, 0xee,
	0xff, 0x18, 0xf5, 0xd0, 0xcf, 0x51, 0x0f, 0xfd, 0x1e, 0xf5, 0xd0, 0xcb, 0xcd, 0xf3, 0x46, 0xd1,
	0xa9, 0x71, 0xb9, 0x1f, 0x85, 0xa9, 0x73, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa0,
	0x05, 0x80, 0x4a, 0x05, 0x00, 0x00,
}
