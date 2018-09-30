package merger

import (
	"github.com/stretchr/testify/mock"
	"github.com/ncipollo/fnew/manifest"
)

type MockMerger struct{
	mock.Mock
}

func NewMockMerger(mergedManaifest manifest.Manifest) *MockMerger {
	mockMerger := MockMerger{}
	mockMerger.On("MergedManifest").Return(&mergedManaifest)
	return &mockMerger
}

func (merger *MockMerger) MergedManifest() *manifest.Manifest {
	args := merger.Called()

	return args.Get(0).(*manifest.Manifest)
}



