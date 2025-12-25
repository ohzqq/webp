package webp

import (
	"bytes"
	"image"
	"io"

	"github.com/HugoSmits86/nativewebp"
)

// EncodeAll writes the animation to w.
func EncodeAll(w io.Writer, images []image.Image, o ...Options) error {
	var opt Options
	if o != nil {
		opt = o[0]
	}
	anim := newAnim(len(images), opt)
	for i, f := range images {
		d, err := encodeImg(f, opt)
		if err != nil {
			return err
		}
		frame, err := Decode(bytes.NewReader(d))
		if err != nil {
			return err
		}
		anim.Images[i] = frame
	}
	return nativewebp.EncodeAll(w, anim, &nativewebp.Options{UseExtendedFormat: true})
}

func newAnim(total int, o Options) *nativewebp.Animation {
	anim := &nativewebp.Animation{
		Images:    make([]image.Image, total),
		Durations: make([]uint, total),
		Disposals: make([]uint, total),
		LoopCount: uint16(o.LoopCount),
	}
	for i := 0; i < total; i++ {
		anim.Durations[i] = uint(DefaultDuration)
		if i < len(o.Durations) {
			anim.Durations[i] = uint(o.Durations[i])
		}
		anim.Disposals[i] = 0
		if i < len(o.Disposals) {
			anim.Disposals[i] = uint(o.Disposals[i])
		}
	}
	return anim
}
