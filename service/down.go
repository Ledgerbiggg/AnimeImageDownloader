package service

import (
	util "AnimeImageDownloader/util"
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
)

func Down(p string) {
	get, err := util.NewHttpDos("https://www.dmoe.cc/random.php", nil, nil, nil).Get()
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile(p+uuid.New().String()+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	written, err := io.Copy(file, bytes.NewReader(get))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(written)
}
