// get bash terminal size?
package main

import (
	"log"
	"os"
	"time"

	term "golang.org/x/term"
)

// firt get viewport size
func getViewportSize() (int, int) {
	fd := int(os.Stdin.Fd())
	w, h, err := term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	return w - 1, h
}

// declare display object
type Display struct {
	Height     int
	Widht      int
	TablePixel [][]Pixel
}
type Pixel struct {
	Pixel      byte
	ColorPixel Color
}
type Color struct {
	Color string
}

const (
	Red   = "\x1B[38;2;255;0;0m"
	Green = "\x1B[38;2;0;255;0"
	Blue  = "\x1B[38;2;0;0;255m"
)

var Clear = []byte("\033c")

var Screen Display

// make display
func buildScreen() {
	x, y := getViewportSize()
	Screen.Widht = x
	Screen.Height = y
	tablePixel := make([][]Pixel, Screen.Height)
	for i := 0; i <= Screen.Height-1; i++ {
		tablePixel[i] = make([]Pixel, Screen.Widht)
	}
	Screen.TablePixel = tablePixel

	//test_________display

	//test_________display
}
func main() {
	buildScreen()
	for _, each := range Screen.TablePixel {
		for range each {
			bf = append(bf, '_')
		}
		bf = append(bf, '\n')
	}
	display()
}

var bf []byte

func display() {
	// cmd := exec.Command("echo", "salut")
	// b, err := cmd.Output()
	os.Stdout.Write(Clear)
	os.Stdout.Write(bf)

	time.Sleep(time.Millisecond * 50)
	display()
}

//print

//print with stdout?
