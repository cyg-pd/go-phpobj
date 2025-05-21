package gophpobj

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	Data *W[struct{ Msg string }]
}

func TestMarshal(t *testing.T) {
	var w mock
	w.Data = new(W[struct{ Msg string }])
	w.Data.SetData(&struct{ Msg string }{"mock"})

	b, err := json.Marshal(&w)

	assert.NoError(t, err)
	assert.JSONEq(t, `{"Data":{"Msg":"mock"}}`, string(b))
}

func TestUnmarshal(t *testing.T) {
	m := []byte(`{"Data":{"Msg":"mock"}}`)
	var w mock

	err := json.Unmarshal(m, &w)
	assert.NoError(t, err)
	assert.EqualValues(t, "mock", w.Data.Data().Msg)
}

func TestEmptyObject(t *testing.T) {
	m := []byte(`{"Data":{}}`)
	var w mock

	err := json.Unmarshal(m, &w)
	assert.NoError(t, err)
	assert.NotNil(t, w.Data.Data())
	assert.Empty(t, w.Data.Data().Msg)
}

func TestEmptyArray(t *testing.T) {
	m := []byte(`{"Data":[]}`)
	var w mock

	err := json.Unmarshal(m, &w)
	assert.NoError(t, err)
	assert.NotNil(t, w.Data.Data())
	assert.Empty(t, w.Data.Data().Msg)
}

func TestNull(t *testing.T) {
	m := []byte(`{"data":null}`)
	var w mock

	err := json.Unmarshal(m, &w)
	assert.NoError(t, err)
	assert.Nil(t, w.Data)
}
