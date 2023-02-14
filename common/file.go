package common

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func FileMd5(reader io.Reader) (string, error) {
	md5Handler := md5.New()
	if _, err := io.Copy(md5Handler, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5Handler.Sum(nil)), nil
}

func SaveFile(reader io.Reader, dstPath string) (bool, error) {
	writer, err := os.Create(dstPath)
	if err != nil {
		return false, err
	}
	defer writer.Close()
	buff := bufio.NewReader(reader)
	_, err = io.Copy(writer, buff)
	if err != nil {
		return false, err
	}
	return true, nil
}
