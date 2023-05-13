package utils

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"regexp"

	"github.com/klauspost/compress/gzip"
)

var (
	AppName    string
	ConfigFile string
	Version    string
	Commit     string
	Branch     string
	BuiltAt    string

	values = regexp.MustCompile(`[#]\{([\w\.]+)\}`)
)

const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ReplaceValues(bsrc []byte) []byte {
	for _, items := range values.FindAllSubmatch(bsrc, -1) {
		env := os.Getenv(string(items[1]))
		if env != "" {
			bsrc = bytes.ReplaceAll(bsrc, items[0], []byte(env))
		}
	}
	return bsrc
}

func RandomString(lenght int) string {
	bigint := big.NewInt(int64(len(characters)))
	b := make([]byte, lenght)
	for i := range b {
		num, err := rand.Int(rand.Reader, bigint)
		if err != nil {
			panic(err)
		}
		b[i] = characters[num.Int64()]
	}
	return string(b)
}

func Executable() (string, string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", "", err
	}
	dir, file := filepath.Split(executable)
	n := len(file) - 4
	if file[n:] == ".exe" {
		file = file[:n]
	}
	return dir, file, nil
}

func Gzip(data []byte) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	w, err := gzip.NewWriterLevel(buff, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func Gunzip(data []byte) ([]byte, error) {
	zip, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer
	if _, err := zip.WriteTo(&buff); err != nil {
		return nil, err
	}
	if err := zip.Close(); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
