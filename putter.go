package putter

import (
	"bufio"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Write(path string, name string, fh *multipart.FileHeader) error {
	//Create all directories about the received path.
	//It it the same as `mkdir -p`
	uts := time.Now().Unix()
	path = filepath.Join(path, strconv.FormatInt(uts, 10))
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	//Auto generation filepath.
	fpath := filepath.Join(path, name)

	//Create empty file.
	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	f, err := fh.Open()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(out)
	reader := bufio.NewReader(f)
	bsize := 4 * 1024 * 1024
	buf := make([]byte, bsize)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}

		_, err = writer.Write(buf[:n])
		if err != nil {
			break
		}
	}
	writer.Flush()
	return nil
}

func Test() string {
	return "Hello KIPI(•ө•)♡"
}