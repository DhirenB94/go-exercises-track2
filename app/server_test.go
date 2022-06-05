package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	t.Run("returns posts", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
		res := httptest.NewRecorder()

		PostServer(res, req)

		var got []Posts

		err := json.NewDecoder(res.Body).Decode(&got)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Posts, '%v'", res.Body, err)
		}

		assert.NotEmpty(t, res.Body, "empty response")
		assert.Equal(t, http.StatusOK, res.Code, "status should equal 200")
	})
}
