package gophpobj

import (
	"bytes"
	"encoding/json"
)

type W[T any] struct {
	data *T
}

func (w *W[T]) Data() *T                     { return w.data }
func (w *W[T]) SetData(d *T)                 { w.data = d }
func (w *W[T]) MarshalJSON() ([]byte, error) { return json.Marshal(w.data) }
func (w *W[T]) UnmarshalJSON(b []byte) error {
	*w = W[T]{data: new(T)}
	if bytes.Equal(b, []byte(`[]`)) {
		return nil
	}
	return json.Unmarshal(b, w.data)
}
