# go-phpobj

[![tag](https://img.shields.io/github/tag/cyg-pd/go-phpobj.svg)](https://github.com/cyg-pd/go-phpobj/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-phpobj?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-phpobj)
![Build Status](https://github.com/cyg-pd/go-phpobj/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-phpobj)](https://goreportcard.com/report/github.com/cyg-pd/go-phpobj)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-phpobj)](https://codecov.io/gh/cyg-pd/go-phpobj)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-phpobj)](https://github.com/cyg-pd/go-phpobj/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-phpobj)](./LICENSE)

## 🚀 Install

```sh
go get github.com/cyg-pd/go-phpobj@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `go-phpobj` using:

```go
import (
	"log/slog"

	"github.com/cyg-pd/go-phpobj"
)

type PHPJSONData struct {
	Data *phpobj.W[struct{ Msg string }]
}

func main() {
	m := []byte(`{"Data":{"Msg":"hi"}}`)
	var w PHPJSONData

	if err := json.Unmarshal(m, &w); err != nil {
		panic(err)
	}

	slog.Info(w.Data.Data().Msg) // hi
}
```
