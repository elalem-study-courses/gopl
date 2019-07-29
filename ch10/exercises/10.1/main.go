package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outputFormat = flag.String("output-format", "", "Image output format")

func main() {
	flag.Parse()
	if err := toImg(os.Stdin, os.Stdout, *outputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toImg(in io.Reader, out io.Writer, outputFormat string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	switch outputFormat {
	case "png":
		err = png.Encode(out, img)
	case "jpg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	case "gif":
		err = gif.Encode(out, img, nil)
	}

	return err
}
