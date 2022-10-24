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

func make_noise2d(w int, h int) *Noise {
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

func draw_noise2d(n *Noise, c *gg.Context) {
	for i := 0; i < n.height; i++ {
		for j := 0; j < n.width; j++ {
			c.SetColor(color.Gray{n.elements[i][j]})
			c.SetPixel(j, i)
		}
	}
}

func draw_noise2d_rgb(r *Noise, g *Noise, b *Noise, c *gg.Context) {
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

// func main() {
// 	os.Mkdir("out", os.ModeDir)

// 	rand.Seed(time.Now().UnixMilli())
// 	n := make_noise2d(100, 200)
// 	dc := gg.NewContext(n.width, n.height)
// 	draw_noise2d(n, dc)
// 	dc.SavePNG("out/test.png")

// 	r := make_noise2d(1000, 1000)
// 	g := make_noise2d(1000, 1000)
// 	b := make_noise2d(1000, 1000)
// 	dc2 := gg.NewContext(r.width, r.height)
// 	draw_noise2d_rgb(r, g, b, dc2)
// 	dc2.SavePNG("out/test_rgb.png")
// }
