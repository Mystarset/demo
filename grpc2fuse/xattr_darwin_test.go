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

type MockRawFileSystemClient struct {
	mock.Mock
}

func (m *MockRawFileSystemClient) SetXAttr(ctx context.Context, in *pb.SetXAttrRequest, opts ...grpc.CallOption) (*pb.SetXAttrResponse, error) {
	args := m.Called(ctx, in, opts)
	if res := args.Get(0); res != nil {
		return res.(*pb.SetXAttrResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSetXAttr(t *testing.T) {
	mockClient := &MockRawFileSystemClient{}
	fs := &fileSystem{
		client: mockClient,
	}

	tests := []struct {
		name     string
		input    *fuse.SetXAttrIn
		attr     string
		data     []byte
		mockResp *pb.SetXAttrResponse
		mockErr  error
		want     fuse.Status
	}{
		{
			name: "success",
			input: &fuse.SetXAttrIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
				Size:     10,
				Flags:    1,
				Position: 0,
				Padding:  0,
			},
			attr: "user.test",
			data: []byte("test data"),
			mockResp: &pb.SetXAttrResponse{
				Status: &pb.Status{
					Code: 0,
				},
			},
			mockErr: nil,
			want:    fuse.OK,
		},
		{
			name: "error_response",
			input: &fuse.SetXAttrIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
			},
			attr: "user.test",
			data: []byte("test"),
			mockResp: &pb.SetXAttrResponse{
				Status: &pb.Status{
					Code: int32(fuse.EACCES),
				},
			},
			mockErr: nil,
			want:    fuse.EACCES,
		},
		{
			name: "grpc_error",
			input: &fuse.SetXAttrIn{
				InHeader: fuse.InHeader{
					NodeId: 1,
				},
			},
			attr:     "user.test",
			data:     []byte("test"),
			mockResp: nil,
			mockErr:  errors.New("grpc error"),
			want:     fuse.EIO,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient.On("SetXAttr", mock.Anything, &pb.SetXAttrRequest{
				Header:   toPbHeader(&tt.input.InHeader),
				Attr:     tt.attr,
				Data:     tt.data,
				Size:     tt.input.Size,
				Flags:    tt.input.Flags,
				Position: tt.input.Position,
				Padding:  tt.input.Padding,
			}, mock.Anything).Return(tt.mockResp, tt.mockErr).Once()

			got := fs.SetXAttr(make(chan struct{}), tt.input, tt.attr, tt.data)
			assert.Equal(t, tt.want, got)
			mockClient.AssertExpectations(t)
		})
	}
}
