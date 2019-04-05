package canvas

import (
	"fmt"
	"strings"

	"github.com/kieron-pivotal/rays/color"
)

type Canvas struct {
	Width  int
	Height int

	pixels [][]color.Color
}

func New(w, h int) *Canvas {
	c := Canvas{
		Width:  w,
		Height: h,
	}

	for i := 0; i < h; i++ {
		c.pixels = append(c.pixels, make([]color.Color, w))
	}
	return &c
}

func (c *Canvas) Pixel(x, y int) color.Color {
	return c.pixels[y][x]
}

func (c *Canvas) SetPixel(x, y int, pixelColor color.Color) {
	c.pixels[y][x] = pixelColor
}

func (c *Canvas) ToPPM() string {
	var sb strings.Builder
	sb.WriteString("P3\n")
	sb.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	sb.WriteString("255\n")

	for _, row := range c.pixels {
		sep := ""
		var rowSb strings.Builder
		for _, p := range row {
			rowSb.WriteString(
				fmt.Sprintf("%s%d %d %d", sep, to255(p.Red()), to255(p.Green()), to255(p.Blue())),
			)
			sep = " "
		}
		for _, line := range split(rowSb.String(), 70) {
			sb.WriteString(line + "\n")
		}
	}

	return sb.String()
}

func split(s string, lim int) []string {
	l := []string{}
	for len(s) > lim {
		idx := strings.LastIndex(s[:lim+1], " ")
		l = append(l, s[:idx])
		s = s[idx+1:]
	}
	l = append(l, s)
	return l
}

func to255(f float64) int {
	v := int(f * 255)
	if v < 0 {
		v = 0
	} else if v > 255 {
		v = 255
	}
	return v
}
