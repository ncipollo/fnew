package merger

import (
	"github.com/stretchr/testify/mock"
	"github.com/ncipollo/fnew/manifest"
)

type MockMerger struct{
	mock.Mock
}

func NewMockMerger() *MockMerger {
	mockMerger := MockMerger{}
	mockMerger.On("MergedManifest").Return(manifest.FullManifest())
	return &mockMerger
}

func (merger *MockMerger) MergedManifest() *manifest.Manifest {
	args := merger.Called()

	return args.Get(0).(*manifest.Manifest)
}



