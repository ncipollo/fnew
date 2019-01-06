package action

import (
    "testing"
    "github.com/ncipollo/fnew/testrepo"
    "github.com/ncipollo/fnew/workspace"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
)

func TestUpdateAction_Perform_FailsIfOpenFails(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    mockWorkSpace := workspace.CreateMockWorkSpace(
        workspace.CreateMockDirectoryChecker(true),
        workspace.CreateMockDirectoryCreator(false))
    printer := testmessage.NewTestPrinter()
    action := NewUpdateAction(mockRepo, mockWorkSpace)

    mockRepo.StubOpen(true)
    mockRepo.StubPull(false)
    err := action.Perform(printer)

    assert.Error(t, err)
}

func TestUpdateAction_Perform_FailsIfPullFails(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    mockWorkSpace := workspace.CreateMockWorkSpace(
        workspace.CreateMockDirectoryChecker(true),
        workspace.CreateMockDirectoryCreator(false))
    printer := testmessage.NewTestPrinter()
    action := NewUpdateAction(mockRepo, mockWorkSpace)

    mockRepo.StubOpen(false)
    mockRepo.StubPull(true)
    err := action.Perform(printer)

    assert.Error(t, err)
}

func TestUpdateAction_Perform_Success(t *testing.T) {
    mockRepo := testrepo.NewMockRepo()
    mockWorkSpace := workspace.CreateMockWorkSpace(
        workspace.CreateMockDirectoryChecker(true),
        workspace.CreateMockDirectoryCreator(false))
    printer := testmessage.NewTestPrinter()
    action := NewUpdateAction(mockRepo, mockWorkSpace)

    mockRepo.StubOpen(false)
    mockRepo.StubPull(false)
    err := action.Perform(printer)

    assert.NoError(t, err)
    mockRepo.AssertNumberOfCalls(t, "Pull", 2)
}
