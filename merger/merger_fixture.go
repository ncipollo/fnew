package merger

import (
    "github.com/ncipollo/fnew/manifest"
    "github.com/stretchr/testify/mock"
)

type MockMerger struct {
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
