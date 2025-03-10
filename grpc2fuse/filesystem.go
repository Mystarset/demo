/*
 * Copyright 2022 Han Xin, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc2fuse

import (
	"context"

	"github.com/hanwen/go-fuse/v2/fuse"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Mystarset/demo/pb"
)

const (
	defaultName = "grpcfuse"
)

type fileSystem struct {
	fuse.RawFileSystem

	client pb.RawFileSystemClient
	opts   []grpc.CallOption
}

// NewFileSystem creates a new file system.
func NewFileSystem(client pb.RawFileSystemClient, opts ...grpc.CallOption) *fileSystem {
	return &fileSystem{
		RawFileSystem: fuse.NewDefaultRawFileSystem(),
		client:        client,
		opts:          opts,
	}
}

func (fs *fileSystem) String() string {
	res, err := fs.client.String(context.TODO(), &pb.StringRequest{}, fs.opts...)
	if err != nil {
		log.Errorf("String: %v", err)
		return defaultName
	}
	return res.Value
}
