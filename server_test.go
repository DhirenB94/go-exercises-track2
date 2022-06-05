package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	server := NewPostServer()
	t.Run("returns posts", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := PostLists

		err := json.NewDecoder(res.Body).Decode(&got)
		assert.NoError(t, err)

		assert.NotEmpty(t, res.Body, "empty response")
		assert.Equal(t, http.StatusOK, res.Code, "status should equal 200")
		assert.Len(t, PostLists, 3, "should return 3 posts")
	})
}
