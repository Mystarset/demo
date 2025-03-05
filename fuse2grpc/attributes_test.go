package fuse2grpc

import (
	"context"
	"testing"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/stretchr/testify/assert"

	"github.com/Mystarset/demo/pb"
)

type mockRawFileSystem struct {
	fuse.RawFileSystem
	getAttrFn func(cancel <-chan struct{}, in *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status
}

func (m *mockRawFileSystem) GetAttr(cancel <-chan struct{}, in *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status {
	if m.getAttrFn != nil {
		return m.getAttrFn(cancel, in, out)
	}
	return fuse.ENOSYS
}

func TestServer_GetAttr(t *testing.T) {
	tests := []struct {
		name     string
		req      *pb.GetAttrRequest
		fsStatus fuse.Status
		want     *pb.GetAttrResponse
		wantErr  error
	}{
		{
			name: "success",
			req: &pb.GetAttrRequest{
				Header: &pb.InHeader{
					NodeId: 1,
					Caller: &pb.Caller{
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
					},
				},
			},
			fsStatus: fuse.OK,
			want: &pb.GetAttrResponse{
				AttrOut: &pb.AttrOut{
					Attr: &pb.Attr{
						Mode: 0644,
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
					},
				},
				Status: &pb.Status{Code: 0},
			},
		},
		{
			name: "error",
			req: &pb.GetAttrRequest{
				Header: &pb.InHeader{
					NodeId: 1,
					Caller: &pb.Caller{
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
					},
				},
			},
			fsStatus: fuse.ENOENT,
			want: &pb.GetAttrResponse{
				Status: &pb.Status{Code: int32(fuse.ENOENT)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &mockRawFileSystem{
				getAttrFn: func(cancel <-chan struct{}, in *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status {
					return tt.fsStatus
				},
			}
			s := NewServer(mockFS)

			got, err := s.GetAttr(context.Background(), tt.req)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.Status.Code, got.Status.Code)
		})
	}
}
