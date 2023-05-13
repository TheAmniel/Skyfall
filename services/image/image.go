package image

import (
	"bytes"
	"image"
	"sync"

	"github.com/disintegration/imaging"
)

type Image struct {
	sync.RWMutex
	raw     image.Image
	format  imaging.Format
	quality int
}

func New(buff []byte) *Image {
	img, format, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		panic(err)
	}
	formatImaging, err := imaging.FormatFromExtension(format)
	if err != nil {
		formatImaging = imaging.JPEG
	}
	return &Image{
		raw:     img,
		format:  formatImaging,
		quality: 95,
	}
}

func (img *Image) Thumbnail() *Image {
	img.Lock()
	img.quality = 75
	img.Unlock()
	return img.Resize(400, 280)
}

func (img *Image) Resize(w, h int) *Image {
	img.Lock()
	img.raw = imaging.Resize(img.raw, w, h, imaging.Lanczos)
	img.Unlock()
	return img
}

func (img *Image) Size(size int) *Image {
	// Min: 16px - Max: 4096px
	if size >= 16 && size <= 4096 && size%2 == 0 {
		x, y := img.raw.Bounds().Dx(), img.raw.Bounds().Dy()
		tx, ty := x, y

		if x > size {
			tx = size
			ty = (size / x) * y
		}

		if ty > size {
			tx = (size / ty) * tx
			ty = size
		}
		return img.Resize(tx, ty)
	}
	return img
}

func (img *Image) Process() ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	// TODO: better handler for configs
	if img.quality > 0 {
		if err := imaging.Encode(buff, img.raw, img.format, imaging.JPEGQuality(img.quality)); err != nil {
			return nil, err
		}
	} else {
		if err := imaging.Encode(buff, img.raw, img.format); err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}
