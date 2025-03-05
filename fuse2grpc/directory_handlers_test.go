package fuse2grpc_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/stretchr/testify/require"

	"github.com/Mystarset/demo/pb"
)

func TestReadDdir(t *testing.T) {
	server, fs := startTestServices(t, 0)
	defer server.Stop()

	client, conn := newRawFileSystemClient(t, serverSocketPath)
	defer conn.Close()

	ctx, cancel := Context()
	defer cancel()

	req := &pb.ReadDirRequest{
		ReadIn: &pb.ReadIn{
			Header: TestInHeader,
			Fh:     1,
			Offset: 0,
			Size:   8192,
		},
	}

	fs.EXPECT().ReadDir(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(cancel <-chan struct{}, input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
			out.AddDirEntry(fuse.DirEntry{Name: "foo", Mode: 0100000, Ino: 1})
			out.AddDirEntry(fuse.DirEntry{Name: "foo2", Mode: 0100000, Ino: 2})
			out.AddDirEntry(fuse.DirEntry{Name: "foo3foo3", Mode: 0100000, Ino: 3})
			out.AddDirEntry(fuse.DirEntry{Name: "foo4", Mode: 0100000, Ino: 4})
			return fuse.OK
		})

	stream, err := client.ReadDir(ctx, req)
	require.NoError(t, err)

	res, err := stream.Recv()
	require.NoError(t, err)

	require.Equal(t, []*pb.DirEntry{
		{Name: []byte("foo"), Mode: 0100000, Ino: 1},
		{Name: []byte("foo2"), Mode: 0100000, Ino: 2},
		{Name: []byte("foo3foo3"), Mode: 0100000, Ino: 3},
		{Name: []byte("foo4"), Mode: 0100000, Ino: 4},
	}, res.Entries)
}
