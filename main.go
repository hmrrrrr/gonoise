package noise

import (
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type Noise struct {
	elements [][]uint8
	height   int
	width    int
}

func MakeNoise2D(w int, h int) *Noise {
	n := make([][]uint8, h)
	for i := range n {
		n[i] = make([]uint8, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// n[i][j] = uint8(math.Pow(float64(rand.Intn(256)), float64(i+j)/float64(h+w)))
			n[i][j] = uint8(rand.Intn(256))
		}
	}

	o := Noise{
		elements: n,
		height:   h,
		width:    w,
	}

	return &o
}

func DrawNoise2D(n *Noise, c *gg.Context) {
	for i := 0; i < n.height; i++ {
		for j := 0; j < n.width; j++ {
			c.SetColor(color.Gray{n.elements[i][j]})
			c.SetPixel(j, i)
		}
	}
}

func DrawNoise2D_RGB(r *Noise, g *Noise, b *Noise, c *gg.Context) {
	if r.height != g.height || g.height != b.height {
		panic("non-matching heights")
	}
	if r.width != g.width || g.width != b.width {
		panic("non-matching widths")
	}

	for i := 0; i < r.height; i++ {
		for j := 0; j < r.width; j++ {
			c.SetColor(color.RGBA{
				r.elements[i][j],
				g.elements[i][j],
				b.elements[i][j],
				uint8(255),
			})
			c.SetPixel(j, i)
		}
	}
}
