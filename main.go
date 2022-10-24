package noise

import (
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type Noise struct {
	Elements [][]uint8
	Height   int
	Width    int
}

func MakeNoise2D(w int, h int) *Noise {
	n := make([][]uint8, h)
	for i := range n {
		n[i] = make([]uint8, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n[i][j] = uint8(rand.Intn(256))
		}
	}

	o := Noise{
		Elements: n,
		Height:   h,
		Width:    w,
	}

	return &o
}

func DrawNoise2D(n *Noise, c *gg.Context) {
	for i := 0; i < n.Height; i++ {
		for j := 0; j < n.Width; j++ {
			c.SetColor(color.Gray{n.Elements[i][j]})
			c.SetPixel(j, i)
		}
	}
}

func DrawNoise2D_RGB(r *Noise, g *Noise, b *Noise, c *gg.Context) {
	if r.Height != g.Height || g.Height != b.Height {
		panic("non-matching heights")
	}
	if r.Width != g.Width || g.Width != b.Width {
		panic("non-matching widths")
	}

	for i := 0; i < r.Height; i++ {
		for j := 0; j < r.Width; j++ {
			c.SetColor(color.RGBA{
				r.Elements[i][j],
				g.Elements[i][j],
				b.Elements[i][j],
				uint8(255),
			})
			c.SetPixel(j, i)
		}
	}
}
