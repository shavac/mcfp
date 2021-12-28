package mcfp

import "runtime"

type Machiner interface {
	Model() string
	Arch() string
	OS() string
	NCPU() int
	MAC() string
	RootDevPath() string
}

type BASE struct {
}

func (BASE) Arch() string {
	return runtime.GOARCH
}

func (BASE) Model() string {
	return "Unknown"
}

func (BASE) OS() string {
	return runtime.GOOS
}

func (BASE) NCPU() int {
	return runtime.NumCPU()
}

func (BASE) MAC() string {
	return GetNicString()
}

func (BASE) RootDevPath() string {
	return GetRootDevPath()
}
