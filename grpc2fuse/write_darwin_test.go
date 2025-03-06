package grpc2fuse

import (
	"context"
	"errors"
	"testing"

	"github.com/Mystarset/demo/pb"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type mockRawFileSystemClient struct {
	mock.Mock
}

func (m *mockRawFileSystemClient) Write(ctx context.Context, in *pb.WriteRequest, opts ...grpc.CallOption) (*pb.WriteResponse, error) {
	args := m.Called(ctx, in, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.WriteResponse), args.Error(1)
}

// Implement other required methods of RawFileSystemClient interface
func (m *mockRawFileSystemClient) String(ctx context.Context, in *pb.StringRequest, opts ...grpc.CallOption) (*pb.StringResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Lookup(ctx context.Context, in *pb.LookupRequest, opts ...grpc.CallOption) (*pb.LookupResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Forget(ctx context.Context, in *pb.ForgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) GetAttr(ctx context.Context, in *pb.GetAttrRequest, opts ...grpc.CallOption) (*pb.GetAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) SetAttr(ctx context.Context, in *pb.SetAttrRequest, opts ...grpc.CallOption) (*pb.SetAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Mknod(ctx context.Context, in *pb.MknodRequest, opts ...grpc.CallOption) (*pb.MknodResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Mkdir(ctx context.Context, in *pb.MkdirRequest, opts ...grpc.CallOption) (*pb.MkdirResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Unlink(ctx context.Context, in *pb.UnlinkRequest, opts ...grpc.CallOption) (*pb.UnlinkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Rmdir(ctx context.Context, in *pb.RmdirRequest, opts ...grpc.CallOption) (*pb.RmdirResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Rename(ctx context.Context, in *pb.RenameRequest, opts ...grpc.CallOption) (*pb.RenameResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Link(ctx context.Context, in *pb.LinkRequest, opts ...grpc.CallOption) (*pb.LinkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Symlink(ctx context.Context, in *pb.SymlinkRequest, opts ...grpc.CallOption) (*pb.SymlinkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Readlink(ctx context.Context, in *pb.ReadlinkRequest, opts ...grpc.CallOption) (*pb.ReadlinkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Access(ctx context.Context, in *pb.AccessRequest, opts ...grpc.CallOption) (*pb.AccessResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) GetXAttr(ctx context.Context, in *pb.GetXAttrRequest, opts ...grpc.CallOption) (*pb.GetXAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) ListXAttr(ctx context.Context, in *pb.ListXAttrRequest, opts ...grpc.CallOption) (*pb.ListXAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) SetXAttr(ctx context.Context, in *pb.SetXAttrRequest, opts ...grpc.CallOption) (*pb.SetXAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) RemoveXAttr(ctx context.Context, in *pb.RemoveXAttrRequest, opts ...grpc.CallOption) (*pb.RemoveXAttrResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Create(ctx context.Context, in *pb.CreateRequest, opts ...grpc.CallOption) (*pb.CreateResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Open(ctx context.Context, in *pb.OpenRequest, opts ...grpc.CallOption) (*pb.OpenResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Read(ctx context.Context, in *pb.ReadRequest, opts ...grpc.CallOption) (pb.RawFileSystem_ReadClient, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Lseek(ctx context.Context, in *pb.LseekRequest, opts ...grpc.CallOption) (*pb.LseekResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) GetLk(ctx context.Context, in *pb.LkRequest, opts ...grpc.CallOption) (*pb.GetLkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) SetLk(ctx context.Context, in *pb.LkRequest, opts ...grpc.CallOption) (*pb.SetLkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) SetLkw(ctx context.Context, in *pb.LkRequest, opts ...grpc.CallOption) (*pb.SetLkResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Release(ctx context.Context, in *pb.ReleaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) CopyFileRange(ctx context.Context, in *pb.CopyFileRangeRequest, opts ...grpc.CallOption) (*pb.CopyFileRangeResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Flush(ctx context.Context, in *pb.FlushRequest, opts ...grpc.CallOption) (*pb.FlushResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Fsync(ctx context.Context, in *pb.FsyncRequest, opts ...grpc.CallOption) (*pb.FsyncResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) Fallocate(ctx context.Context, in *pb.FallocateRequest, opts ...grpc.CallOption) (*pb.FallocateResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) OpenDir(ctx context.Context, in *pb.OpenDirRequest, opts ...grpc.CallOption) (*pb.OpenDirResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) ReadDir(ctx context.Context, in *pb.ReadDirRequest, opts ...grpc.CallOption) (pb.RawFileSystem_ReadDirClient, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) ReadDirPlus(ctx context.Context, in *pb.ReadDirRequest, opts ...grpc.CallOption) (pb.RawFileSystem_ReadDirPlusClient, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) ReleaseDir(ctx context.Context, in *pb.ReleaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) FsyncDir(ctx context.Context, in *pb.FsyncRequest, opts ...grpc.CallOption) (*pb.FsyncResponse, error) {
	return nil, nil
}

func (m *mockRawFileSystemClient) StatFs(ctx context.Context, in *pb.StatfsRequest, opts ...grpc.CallOption) (*pb.StatfsResponse, error) {
	return nil, nil
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name           string
		input          *fuse.WriteIn
		data           []byte
		mockResponse   *pb.WriteResponse
		mockError      error
		expectedStatus fuse.Status
		expectedBytes  uint32
	}{
		{
			name: "successful write",
			input: &fuse.WriteIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Fh:         2,
				Offset:     100,
				Size:       200,
				WriteFlags: 0,
			},
			data: []byte("test data"),
			mockResponse: &pb.WriteResponse{
				Status:  &pb.Status{Code: 0},
				Written: 9,
			},
			mockError:      nil,
			expectedStatus: fuse.OK,
			expectedBytes:  9,
		},
		{
			name: "write error",
			input: &fuse.WriteIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
			},
			data:           []byte("test"),
			mockResponse:   nil,
			mockError:      errors.New("write failed"),
			expectedStatus: fuse.EIO,
			expectedBytes:  0,
		},
		{
			name: "write with non-zero status",
			input: &fuse.WriteIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
			},
			data: []byte("test"),
			mockResponse: &pb.WriteResponse{
				Status:  &pb.Status{Code: int32(fuse.ENOSYS)},
				Written: 0,
			},
			mockError:      nil,
			expectedStatus: fuse.ENOSYS,
			expectedBytes:  0,
		},
		{
			name: "write with empty data",
			input: &fuse.WriteIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Size: 0,
			},
			data: []byte{},
			mockResponse: &pb.WriteResponse{
				Status:  &pb.Status{Code: 0},
				Written: 0,
			},
			mockError:      nil,
			expectedStatus: fuse.OK,
			expectedBytes:  0,
		},
		{
			name: "write with large data",
			input: &fuse.WriteIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Size: 1024 * 1024,
			},
			data: make([]byte, 1024*1024),
			mockResponse: &pb.WriteResponse{
				Status:  &pb.Status{Code: 0},
				Written: 1024 * 1024,
			},
			mockError:      nil,
			expectedStatus: fuse.OK,
			expectedBytes:  1024 * 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockRawFileSystemClient{}

			fs := &fileSystem{
				client: mockClient,
			}

			cancel := make(chan struct{})

			expectedReq := &pb.WriteRequest{
				Header:     toPbHeader(&tt.input.InHeader),
				Fh:        tt.input.Fh,
				Offset:    tt.input.Offset,
				Data:      tt.data,
				Size:      tt.input.Size,
				WriteFlags: tt.input.WriteFlags,
			}

			mockClient.On("Write", mock.Anything, expectedReq, mock.Anything).Return(tt.mockResponse, tt.mockError)

			written, status := fs.Write(cancel, tt.input, tt.data)

			assert.Equal(t, tt.expectedStatus, status)
			assert.Equal(t, tt.expectedBytes, written)

			mockClient.AssertExpectations(t)
		})
	}
}
