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
	"github.com/hanwen/go-fuse/v2/fuse"

	"github.com/Mystarset/demo/pb"
)

func getUmask(in *fuse.MknodIn) uint16 {
	return uint16(in.Umask)
}

func setFlags(out *fuse.Attr, flags uint32) {
}

func setBlksize(out *fuse.Attr, size uint32) {
	out.Blksize = size
}

func setPadding(out *fuse.Attr, padding uint32) {
	out.Padding = padding
}

func toPbReadIn(in *fuse.ReadIn) *pb.ReadIn {
	return &pb.ReadIn{
		Header:    toPbHeader(&in.InHeader),
		Fh:        in.Fh,
		ReadFlags: in.ReadFlags,
		Offset:    in.Offset,
		Size:      in.Size,
		LockOwner: in.LockOwner,
		Flags:     in.Flags,
		Padding:   in.Padding,
	}
}
