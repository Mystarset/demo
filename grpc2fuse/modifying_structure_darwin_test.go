package grpc2fuse

import (
	"context"
	"testing"

	"github.com/Mystarset/demo/pb"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockRawFileSystemClient struct {
	mock.Mock
}

func (m *MockRawFileSystemClient) Mknod(ctx context.Context, in *pb.MknodRequest, opts ...grpc.CallOption) (*pb.MknodResponse, error) {
	args := m.Called(ctx, in, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.MknodResponse), args.Error(1)
}

// Mock other required methods
func (m *MockRawFileSystemClient) String(ctx context.Context, in *pb.StringRequest, opts ...grpc.CallOption) (*pb.StringResponse, error) {
	return nil, nil
}

func (m *MockRawFileSystemClient) Lookup(ctx context.Context, in *pb.LookupRequest, opts ...grpc.CallOption) (*pb.LookupResponse, error) {
	return nil, nil
}

func (m *MockRawFileSystemClient) Forget(ctx context.Context, in *pb.ForgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockRawFileSystemClient) GetAttr(ctx context.Context, in *pb.GetAttrRequest, opts ...grpc.CallOption) (*pb.GetAttrResponse, error) {
	return nil, nil
}

func (m *MockRawFileSystemClient) SetAttr(ctx context.Context, in *pb.SetAttrRequest, opts ...grpc.CallOption) (*pb.SetAttrResponse, error) {
	return nil, nil
}

func TestFileSystem_Mknod(t *testing.T) {
	mockClient := new(MockRawFileSystemClient)
	fs := &fileSystem{
		client: mockClient,
	}

	tests := []struct {
		name     string
		input    *fuse.MknodIn
		nodeName string
		mockResp *pb.MknodResponse
		mockErr  error
		want     fuse.Status
	}{
		{
			name: "successful mknod",
			input: &fuse.MknodIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
					Uid:    1000,
					Gid:    1000,
				},
				Mode: 0644,
				Rdev: 0,
			},
			nodeName: "testnode",
			mockResp: &pb.MknodResponse{
				Status: &pb.Status{
					Code: 0,
				},
				EntryOut: &pb.EntryOut{
					NodeId:     2,
					Generation: 1,
					EntryValid: 1000,
					AttrValid:  1000,
					Attr: &pb.Attr{
						Ino:  2,
						Mode: 0644,
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
					},
				},
			},
			mockErr: nil,
			want:    fuse.OK,
		},
		{
			name: "error response from server",
			input: &fuse.MknodIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Mode: 0644,
				Rdev: 0,
			},
			nodeName: "testnode",
			mockResp: &pb.MknodResponse{
				Status: &pb.Status{
					Code: int32(fuse.EACCES),
				},
			},
			mockErr: nil,
			want:    fuse.EACCES,
		},
		{
			name: "grpc error",
			input: &fuse.MknodIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Mode: 0644,
				Rdev: 0,
			},
			nodeName: "testnode",
			mockResp: nil,
			mockErr:  grpc.ErrServerStopped,
			want:     fuse.EIO,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient.On("Mknod", mock.Anything, &pb.MknodRequest{
				Header: toPbHeader(&tt.input.InHeader),
				Name:   tt.nodeName,
				Mode:   tt.input.Mode,
				Rdev:   tt.input.Rdev,
			}, mock.Anything).Return(tt.mockResp, tt.mockErr)

			out := &fuse.EntryOut{}
			got := fs.Mknod(make(chan struct{}), tt.input, tt.nodeName, out)

			assert.Equal(t, tt.want, got)
			if tt.want == fuse.OK {
				assert.Equal(t, uint64(2), out.NodeId)
				assert.Equal(t, uint64(1), out.Generation)
				assert.Equal(t, uint64(1000), out.EntryValid)
				assert.Equal(t, uint64(1000), out.AttrValid)
				assert.Equal(t, uint64(2), out.Attr.Ino)
				assert.Equal(t, uint32(0644), out.Attr.Mode)
				assert.Equal(t, uint32(1000), out.Attr.Uid)
				assert.Equal(t, uint32(1000), out.Attr.Gid)
			}
			mockClient.AssertExpectations(t)
		})
	}
}
