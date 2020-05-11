package fakes

import (
	"sync"

	"github.com/cloudfoundry/packit"
)

type EntrypointParser struct {
	ParseCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			ProcessSlice []packit.Process
			Error        error
		}
		Stub func(string) ([]packit.Process, error)
	}
}

func (f *EntrypointParser) Parse(param1 string) ([]packit.Process, error) {
	f.ParseCall.Lock()
	defer f.ParseCall.Unlock()
	f.ParseCall.CallCount++
	f.ParseCall.Receives.Path = param1
	if f.ParseCall.Stub != nil {
		return f.ParseCall.Stub(param1)
	}
	return f.ParseCall.Returns.ProcessSlice, f.ParseCall.Returns.Error
}
