package action

import (
    "testing"
    "github.com/ncipollo/fnew/testrepo"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
)

const cleanupActionPath = "/test"

func TestCleanupRepoAction_Perform_FailsOnDeleteError(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    printer := testmessage.NewTestPrinter()
    action := NewCleanupRepoAction(cleanupActionPath, mockRepo)

    mockRepo.StubDelete(true)
    mockRepo.StubInit(false)
    err := action.Perform(printer)

    assert.Error(t, err)
}

func TestCleanupRepoAction_Perform_FailsOnInitError(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    printer := testmessage.NewTestPrinter()
    action := NewCleanupRepoAction(cleanupActionPath, mockRepo)

    mockRepo.StubDelete(false)
    mockRepo.StubInit(true)
    err := action.Perform(printer)

    assert.Error(t, err)
}

func TestCleanupRepoAction_Perform_Success(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    printer := testmessage.NewTestPrinter()
    action := NewCleanupRepoAction(cleanupActionPath, mockRepo)

    mockRepo.StubDelete(false)
    mockRepo.StubInit(false)
    err := action.Perform(printer)

    assert.NoError(t, err)
    mockRepo.AssertCalled(t, "Delete", cleanupActionPath)
    mockRepo.AssertCalled(t, "Init", cleanupActionPath)
}
