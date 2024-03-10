package test

import (
	util "AnimeImageDownloader/util"
	"bytes"
	"github.com/google/uuid"
	"io"
	"os"
	"testing"
)

func TestDown(t *testing.T) {
	get, err := util.NewHttpDos("https://www.dmoe.cc/random.php", nil, nil, nil).Get()
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(uuid.New().String() + "png")
	if err != nil {
		t.Fatal(err)
	}

	written, err := io.Copy(file, bytes.NewReader(get))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(written)

}
