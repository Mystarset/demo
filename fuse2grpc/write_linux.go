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

package fuse2grpc

import (
	"context"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/hanwen/go-fuse/v2/fuse"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Mystarset/demo/pb"
)

func (s *server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	var (
		header fuse.InHeader
	)
	grpc_logrus.Extract(ctx).WithFields(log.Fields{
		"nodeId":     req.Header.NodeId,
		"fh":         req.Fh,
		"offset":     req.Offset,
		"size":       req.Size,
		"writeFlags": req.WriteFlags,
		"lockOwner":  req.LockOwner,
	}).Debug("Write")
	toFuseInHeader(req.Header, &header)

	writen, st := s.fs.Write(ctx.Done(), &fuse.WriteIn{InHeader: header, Fh: req.Fh, Offset: req.Offset, Size: req.Size, WriteFlags: req.WriteFlags, LockOwner: req.LockOwner}, req.Data)
	if st == fuse.ENOSYS {
		return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
	}
	return &pb.WriteResponse{Written: writen, Status: &pb.Status{Code: int32(st)}}, nil
}
