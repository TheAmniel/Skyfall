package utils

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/klauspost/compress/gzip"
)

var (
	AppName    string
	ConfigFile string
	Version    string
	Commit     string
	Branch     string
	BuiltAt    string
)

const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	supportedImage = regexp.MustCompile(`^image\/(png|jpeg|gif|webp)$`)
	supportedVideo = regexp.MustCompile(`^video\/(mp4|webm|3gpp|quicktime)$`)
	isImage        = regexp.MustCompile(`(png|jpe?g|gif|webp)$`)
	isVideo        = regexp.MustCompile(`(mp4|webm|mov|3gpp?)$`)
	values         = regexp.MustCompile(`[#]\{([\w\.]+)\}`)
)

func ParseFilename(raw string) (string, string) {
	values := strings.Split(raw, ".")
	n := len(values) - 1
	return strings.Join(values[:n], ""), values[n]
}

func SupportedMediaType(mime string) bool {
	return supportedImage.MatchString(mime) || supportedVideo.MatchString(mime)
}

func IsVideo(raw string) bool {
	return isVideo.MatchString(raw)
}

func IsImage(raw string) bool {
	return isImage.MatchString(raw)
}

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

func GetPath(p string) (string, error) {
	dir, _, err := Executable()
	if err != nil {
		return "", err
	}
	dir += p
	if dir[:len(dir)-1] != "/" {
		dir += "/"
	}
	return dir, nil
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
