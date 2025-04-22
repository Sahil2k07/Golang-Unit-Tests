package services_test

import (
	"testing"

	"github.com/Sahil2k07/Golang-Unit-Tests/tests/mocks"
)

func TestGetUserPosts(t *testing.T) {
	testTable := []struct {
		name string
	}{}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.PostRepoMock)

			mockRepo.AssertExpectations(t)
		})
	}
}
