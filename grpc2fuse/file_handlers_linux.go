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
	"github.com/Mystarset/demo/pb"

	"github.com/hanwen/go-fuse/v2/fuse"
)

func (fs *fileSystem) Create(cancel <-chan struct{}, input *fuse.CreateIn, name string, out *fuse.CreateOut) (code fuse.Status) {
	ctx := newContext(cancel)

	res, err := fs.client.Create(ctx, &pb.CreateRequest{
		Header:  toPbHeader(&input.InHeader),
		Name:    name,
		Flags:   input.Flags,
		Mode:    input.Mode,
		Umask:   input.Umask,
		Padding: input.Padding,
	}, fs.opts...)

	if st := dealGrpcError("Create", err); st != fuse.OK {
		return st
	}
	if res.Status.GetCode() != 0 {
		return fuse.Status(res.Status.GetCode())
	}
	toFuseEntryOut(&out.EntryOut, res.EntryOut)
	toFuseOpenOut(&out.OpenOut, res.OpenOut)
	return fuse.Status(res.Status.GetCode())
}
