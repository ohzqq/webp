package webp

import (
	"bytes"
	"image"
	"io"

	"github.com/HugoSmits86/nativewebp"
)

// EncodeAll writes the animation to w.
func EncodeAll(w io.Writer, images []image.Image, o ...Options) error {
	anim := &nativewebp.Animation{
		Images:    make([]image.Image, len(images)),
		Durations: make([]uint, len(images)),
		Disposals: make([]uint, len(images)),
		LoopCount: 0,
	}
	for i, f := range images {
		d, err := EncodeImg(f, o...)
		if err != nil {
			return err
		}
		frame, err := Decode(bytes.NewReader(d))
		if err != nil {
			return err
		}
		anim.Images[i] = frame
		anim.Durations[i] = 80
		anim.Disposals[i] = 0
	}
	return nativewebp.EncodeAll(w, anim, &nativewebp.Options{UseExtendedFormat: true})
}

// encodeAll writes the animation to w.
func encodeAll(w io.Writer, webp *WEBP) error {
	anim := &nativewebp.Animation{
		Images:    webp.Image,
		Durations: make([]uint, len(webp.Delay)),
		Disposals: make([]uint, len(webp.Disposals)),
		LoopCount: uint16(webp.LoopCount),
	}
	for i, d := range webp.Delay {
		anim.Durations[i] = uint(d)
	}
	for i, d := range webp.Disposals {
		anim.Durations[i] = uint(d)
	}
	return nativewebp.EncodeAll(w, anim, &nativewebp.Options{UseExtendedFormat: true})
}
