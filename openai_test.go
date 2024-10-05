package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock for APIKeyReader
type MockAPIKeyReader struct {
	APIKey string
	Err    error
}

func (m *MockAPIKeyReader) ReadAPIKey() (string, error) {
	return m.APIKey, m.Err
}

// Mock for CommitMsgGenerator
type MockCommitMsgGenerator struct {
	Message string
	Err     error
}

func (m *MockCommitMsgGenerator) GenerateCommitMsg(apiKey, content string) (string, error) {
	if apiKey == "" {
		return "", errors.New("missing API key")
	}
	if content == "" {
		return "", errors.New("empty prompt")
	}
	return m.Message, m.Err
}

func TestGenerateCommitMsg_EmptyAPIKey(t *testing.T) {
	mockGenerator := &MockCommitMsgGenerator{Err: errors.New("missing API key")}
	_, err := mockGenerator.GenerateCommitMsg("", "some content")

	assert.Error(t, err, "missing API key")
}

func TestGenerateCommitMsg_EmptyRequestBody(t *testing.T) {
	mockAPIKeyReader := &MockAPIKeyReader{APIKey: "mock-api-key"}
	mockGenerator := &MockCommitMsgGenerator{Err: errors.New("empty prompt")}

	apiKey, err := mockAPIKeyReader.ReadAPIKey()
	assert.NoError(t, err, "Expected no error when reading API key")

	_, err = mockGenerator.GenerateCommitMsg(apiKey, "")

	assert.Error(t, err, "empty prompt")
}

func TestGenerateCommitMsg_Success(t *testing.T) {
	mockAPIKeyReader := &MockAPIKeyReader{APIKey: "mock-api-key"}
	mockGenerator := &MockCommitMsgGenerator{Message: "Mock commit message"}

	apiKey, err := mockAPIKeyReader.ReadAPIKey()
	assert.NoError(t, err, "Expected no error when reading API key")

	content := `+ // Added JWT token validation for API endpoints...`
	message, err := mockGenerator.GenerateCommitMsg(apiKey, content)

	assert.NoError(t, err, "Expected no error from the mock API call")
	assert.Equal(t, "Mock commit message", message, "Expected the mock commit message")
}
