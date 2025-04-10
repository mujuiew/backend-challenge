package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "", bytes.NewBuffer([]byte("")))
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	GetBeef(w, r)

	var responseData PieFireDire
	_ = json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.Equal(t, 200, w.Result().StatusCode)

	assert.NotEmpty(t, len(responseData.Beef))
}
