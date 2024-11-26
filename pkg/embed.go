package pkg

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
)

var (
	//go:embed template
	Templates embed.FS
)

const (
	Main = "main.go"
)

func ReadFsFile(filename string) ([]byte, error) {
	tpl, err := Templates.Open(fmt.Sprintf("template/%s", filename))
	if err != nil {
		return nil, err
	}
	buf, err := io.ReadAll(tpl)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
