package services_test

import (
	"errors"
	"testing"

	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/Sahil2k07/Golang-Unit-Tests/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetUserPosts(t *testing.T) {
	testTable := []struct {
		name        string
		email       string
		mockPosts   []models.Post
		mockError   error
		expectedLen int
		expectedErr bool
	}{
		{
			name:        "successfully returns posts",
			email:       "test@example.com",
			mockPosts:   []models.Post{{Title: "Post 1"}, {Title: "Post 2"}},
			mockError:   nil,
			expectedLen: 2,
			expectedErr: false,
		},
		{
			name:        "user not found",
			email:       "missing@example.com",
			mockPosts:   nil,
			mockError:   errors.New("record not found"),
			expectedLen: 0,
			expectedErr: true,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.PostRepoMock)

			mockRepo.On("GetUserPosts", tt.email).Return(tt.mockPosts, tt.mockError)

			posts, err := mockRepo.GetUserPosts(tt.email)

			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, posts, tt.expectedLen)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
