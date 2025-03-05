package fuse2grpc_test

import (
	"testing"

	"github.com/hanwen/go-fuse/v2/fuse"
)

type mockRawFileSystem struct {
	getAttrFunc func(cancel <-chan struct{}, in *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status
	setAttrFunc func(cancel <-chan struct{}, in *fuse.SetAttrIn, out *fuse.AttrOut) fuse.Status
}

func (m *mockRawFileSystem) String() string                                                                { return "mock" }
func (m *mockRawFileSystem) Init(server *fuse.Server)                                                     {}
func (m *mockRawFileSystem) StatFs(cancel <-chan struct{}, in *fuse.InHeader, out *fuse.StatfsOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Lookup(cancel <-chan struct{}, header *fuse.InHeader, name string, out *fuse.EntryOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Forget(nodeID uint64, nlookup uint64)                                        {}
func (m *mockRawFileSystem) Access(cancel <-chan struct{}, input *fuse.AccessIn) fuse.Status             { return fuse.ENOSYS }
func (m *mockRawFileSystem) GetXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string, dest []byte) (uint32, fuse.Status) {
	return 0, fuse.ENOSYS
}
func (m *mockRawFileSystem) ListXAttr(cancel <-chan struct{}, header *fuse.InHeader, dest []byte) (uint32, fuse.Status) {
	return 0, fuse.ENOSYS
}
func (m *mockRawFileSystem) SetXAttr(cancel <-chan struct{}, input *fuse.SetXAttrIn, attr string, data []byte) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) RemoveXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Create(cancel <-chan struct{}, input *fuse.CreateIn, name string, out *fuse.CreateOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Mkdir(cancel <-chan struct{}, input *fuse.MkdirIn, name string, out *fuse.EntryOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Unlink(cancel <-chan struct{}, header *fuse.InHeader, name string) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Rmdir(cancel <-chan struct{}, header *fuse.InHeader, name string) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Rename(cancel <-chan struct{}, input *fuse.RenameIn, oldName string, newName string) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Link(cancel <-chan struct{}, input *fuse.LinkIn, filename string, out *fuse.EntryOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Symlink(cancel <-chan struct{}, header *fuse.InHeader, pointedTo string, linkName string, out *fuse.EntryOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Readlink(cancel <-chan struct{}, header *fuse.InHeader) ([]byte, fuse.Status) {
	return nil, fuse.ENOSYS
}
func (m *mockRawFileSystem) Mknod(cancel <-chan struct{}, input *fuse.MknodIn, name string, out *fuse.EntryOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) GetAttr(cancel <-chan struct{}, input *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status {
	if m.getAttrFunc != nil {
		return m.getAttrFunc(cancel, input, out)
	}
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) SetAttr(cancel <-chan struct{}, input *fuse.SetAttrIn, out *fuse.AttrOut) fuse.Status {
	if m.setAttrFunc != nil {
		return m.setAttrFunc(cancel, input, out)
	}
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Open(cancel <-chan struct{}, input *fuse.OpenIn, out *fuse.OpenOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Read(cancel <-chan struct{}, input *fuse.ReadIn, buf []byte) (fuse.ReadResult, fuse.Status) {
	return nil, fuse.ENOSYS
}
func (m *mockRawFileSystem) Lseek(cancel <-chan struct{}, in *fuse.LseekIn, out *fuse.LseekOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) GetLk(cancel <-chan struct{}, input *fuse.LkIn, out *fuse.LkOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) SetLk(cancel <-chan struct{}, input *fuse.LkIn) fuse.Status { return fuse.ENOSYS }
func (m *mockRawFileSystem) SetLkw(cancel <-chan struct{}, input *fuse.LkIn) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) Release(cancel <-chan struct{}, input *fuse.ReleaseIn)                       {}
func (m *mockRawFileSystem) Write(cancel <-chan struct{}, input *fuse.WriteIn, data []byte) (uint32, fuse.Status) {
	return 0, fuse.ENOSYS
}
func (m *mockRawFileSystem) CopyFileRange(cancel <-chan struct{}, input *fuse.CopyFileRangeIn) (uint32, fuse.Status) {
	return 0, fuse.ENOSYS
}
func (m *mockRawFileSystem) Flush(cancel <-chan struct{}, input *fuse.FlushIn) fuse.Status { return fuse.ENOSYS }
func (m *mockRawFileSystem) Fsync(cancel <-chan struct{}, input *fuse.FsyncIn) fuse.Status { return fuse.ENOSYS }
func (m *mockRawFileSystem) Fallocate(cancel <-chan struct{}, input *fuse.FallocateIn) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) OpenDir(cancel <-chan struct{}, input *fuse.OpenIn, out *fuse.OpenOut) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) ReadDir(cancel <-chan struct{}, input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) ReadDirPlus(cancel <-chan struct{}, input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	return fuse.ENOSYS
}
func (m *mockRawFileSystem) ReleaseDir(input *fuse.ReleaseIn)                   {}
func (m *mockRawFileSystem) FsyncDir(cancel <-chan struct{}, input *fuse.FsyncIn) fuse.Status          { return fuse.ENOSYS }

func TestMockFileSystem(t *testing.T) {
	mock := &mockRawFileSystem{}
	if mock.String() != "mock" {
		t.Errorf("Expected mock.String() to return 'mock'")
	}
}
