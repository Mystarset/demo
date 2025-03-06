package grpc2fuse_test

import (
	"errors"
	"testing"

	"github.com/Mystarset/demo/grpc2fuse"
	"github.com/Mystarset/demo/pb"
	"github.com/stretchr/testify/assert"
)

type mockReadDirClient struct {
	responses []*pb.ReadDirResponse
	index     int
	err       error
}

func (m *mockReadDirClient) Recv() (*pb.ReadDirResponse, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.index >= len(m.responses) {
		return nil, nil
	}
	resp := m.responses[m.index]
	m.index++
	return resp, nil
}

func TestRawFileSystem_ReadDirClient_Recv(t *testing.T) {
	tests := []struct {
		name      string
		responses []*pb.ReadDirResponse
		err       error
		want      *pb.ReadDirResponse
		wantErr   error
	}{
		{
			name: "successful receive",
			responses: []*pb.ReadDirResponse{
				{
					Status: &pb.Status{Code: 0},
					Entries: []*pb.DirEntry{
						{
							Ino:  1,
							Name: []byte("test.txt"),
							Mode: 0644,
						},
					},
				},
			},
			want: &pb.ReadDirResponse{
				Status: &pb.Status{Code: 0},
				Entries: []*pb.DirEntry{
					{
						Ino:  1,
						Name: []byte("test.txt"),
						Mode: 0644,
					},
				},
			},
			wantErr: nil,
		},
		{
			name:      "error on receive",
			responses: nil,
			err:       errors.New("receive error"),
			want:      nil,
			wantErr:   errors.New("receive error"),
		},
		{
			name:      "empty response",
			responses: []*pb.ReadDirResponse{},
			want:      nil,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockReadDirClient{
				responses: tt.responses,
				err:       tt.err,
			}

			var readDirClient grpc2fuse.RawFileSystem_ReadDirClient = client

			got, err := readDirClient.Recv()

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
