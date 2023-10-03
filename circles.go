// Mathematical art based on:
// https://blogs.scientificamerican.com/guest-blog/making-mathematical-art/
package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/fzipp/canvas"
)

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g., '127.0.0.1:8080' or just ':8080')")
	flag.Parse()

	fmt.Println("Go to " + httpLink(*http))
	err := canvas.ListenAndServe(*http, run, &canvas.Options{
		Title: "Mathematical Art Demo",
		Width: w, Height: h,
		ScaleToPageHeight: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

const (
	w     = 1000
	h     = 1000
	N     = 14_000
	scale = 500.0
)

func run(ctx *canvas.Context) {
	ctx.SetLineWidth(0.5)
	for k := 0; k < N; k++ {
		ctx.SetStrokeStyle(C(k))
		circle(ctx, X(k)*scale+(w/2), Y(k)*scale+(h/2), R(k)*scale)
	}
	ctx.Flush()
}

func X(k int) float64 {
	return math.Cos((10*math.Pi*float64(k))/N) * (1 - 0.5*math.Pow(math.Cos((16*math.Pi*float64(k))/N), 2))
}

func Y(k int) float64 {
	return math.Sin((10*math.Pi*float64(k))/N) * (1 - 0.5*math.Pow(math.Cos((16*math.Pi*float64(k))/N), 2))
}

func R(k int) float64 {
	return (1.0 / 200.0) + 0.1*math.Pow(math.Sin((52*math.Pi*float64(k))/N), 4)
}

func C(k int) color.Color {
	return color.RGBA{
		R: uint8(200+255*float64(k)/N) % 255,
		G: uint8(100+255*float64(k)/N) % 255,
		B: uint8(10+255*float64(k)/N) % 255,
		A: 255,
	}
}

func circle(ctx *canvas.Context, x, y, r float64) {
	ctx.BeginPath()
	ctx.Arc(x, y, r, 0.0, 2*math.Pi, false)
	ctx.Stroke()
}

func httpLink(addr string) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	return "http://" + addr
}
