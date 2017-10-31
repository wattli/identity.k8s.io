package util

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()

	stat, err := fsrc.Stat()
	if err != nil {
		return err
	}

	fdst, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, stat.Mode())
	if err != nil {
		return err
	}
	defer fdst.Close()

	_, err = io.Copy(fdst, fsrc)
	return err
}
