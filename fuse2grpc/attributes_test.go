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
	getAttrStatus fuse.Status
	setAttrStatus fuse.Status
	attrOut       fuse.AttrOut
}

func (m *mockRawFileSystem) String() string { return "mock" }

func (m *mockRawFileSystem) GetAttr(cancel <-chan struct{}, input *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status {
	if m.getAttrStatus != fuse.OK {
		return m.getAttrStatus
	}
	*out = m.attrOut
	return fuse.OK
}

func (m *mockRawFileSystem) SetAttr(cancel <-chan struct{}, in *fuse.SetAttrIn, out *fuse.AttrOut) fuse.Status {
	if m.setAttrStatus != fuse.OK {
		return m.setAttrStatus
	}
	*out = m.attrOut
	return fuse.OK
}

func TestGetAttr(t *testing.T) {
	tests := []struct {
		name           string
		getAttrStatus  fuse.Status
		expectedStatus int32
		attrOut       fuse.AttrOut
		shouldError    bool
	}{
		{
			name:           "success",
			getAttrStatus:  fuse.OK,
			expectedStatus: 0,
			attrOut: fuse.AttrOut{
				Attr: fuse.Attr{
					Ino:  123,
					Mode: 0644,
				},
				AttrValid:     1,
				AttrValidNsec: 2,
			},
		},
		{
			name:           "not implemented",
			getAttrStatus:  fuse.ENOSYS,
			expectedStatus: 0,
			shouldError:    true,
		},
		{
			name:           "error",
			getAttrStatus:  fuse.EACCES,
			expectedStatus: int32(fuse.EACCES),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &mockRawFileSystem{
				getAttrStatus: tt.getAttrStatus,
				attrOut:      tt.attrOut,
			}
			server := NewServer(mockFS)

			req := &pb.GetAttrRequest{
				Header: &pb.InHeader{
					NodeId: 1,
					Caller: &pb.Caller{
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
						Pid: 1234,
					},
				},
			}

			resp, err := server.GetAttr(context.Background(), req)

			if tt.shouldError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.Status.Code)

			if tt.getAttrStatus == fuse.OK {
				assert.Equal(t, tt.attrOut.AttrValid, resp.AttrOut.AttrValid)
				assert.Equal(t, tt.attrOut.AttrValidNsec, resp.AttrOut.AttrValidNsec)
				assert.Equal(t, tt.attrOut.Attr.Ino, resp.AttrOut.Attr.Ino)
				assert.Equal(t, tt.attrOut.Attr.Mode, resp.AttrOut.Attr.Mode)
			}
		})
	}
}

func TestSetAttr(t *testing.T) {
	tests := []struct {
		name           string
		setAttrStatus  fuse.Status
		expectedStatus int32
		attrOut       fuse.AttrOut
		shouldError    bool
	}{
		{
			name:           "success",
			setAttrStatus:  fuse.OK,
			expectedStatus: 0,
			attrOut: fuse.AttrOut{
				Attr: fuse.Attr{
					Ino:  123,
					Mode: 0644,
				},
			},
		},
		{
			name:           "not implemented",
			setAttrStatus:  fuse.ENOSYS,
			expectedStatus: 0,
			shouldError:    true,
		},
		{
			name:           "error",
			setAttrStatus:  fuse.EACCES,
			expectedStatus: int32(fuse.EACCES),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &mockRawFileSystem{
				setAttrStatus: tt.setAttrStatus,
				attrOut:      tt.attrOut,
			}
			server := NewServer(mockFS)

			req := &pb.SetAttrRequest{
				Header: &pb.InHeader{
					NodeId: 1,
					Caller: &pb.Caller{
						Owner: &pb.Owner{
							Uid: 1000,
							Gid: 1000,
						},
						Pid: 1234,
					},
				},
				Valid:   1,
				Mode:    0644,
				Owner:   &pb.Owner{Uid: 1000, Gid: 1000},
				Size:    1024,
				Atime:   123456,
				Mtime:   123456,
				Ctime:   123456,
				Padding: 0,
			}

			resp, err := server.SetAttr(context.Background(), req)

			if tt.shouldError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.Status.Code)

			if tt.setAttrStatus == fuse.OK {
				assert.Equal(t, tt.attrOut.Attr.Ino, resp.AttrOut.Attr.Ino)
				assert.Equal(t, tt.attrOut.Attr.Mode, resp.AttrOut.Attr.Mode)
			}
		})
	}
}
