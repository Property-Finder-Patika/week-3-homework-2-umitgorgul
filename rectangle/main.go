package main

import ( // too much import for a simple rectangle functions tho
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func showAreaAndCircumference(choice1 int, choice2 int) {
	// simple math here no need to describe
	area := choice1 * choice2          // for find area
	circumference := choice1 + choice2 // for find circumference
	fmt.Println("area of rectangle : ", area)
	fmt.Println("circumference of rectangle : ", circumference)
}

func letsDrawIt(choice1 int, choice2 int) {
	rectangle := "rectangle.png" //for creating rectangle as image
	// I can't find other alternative than using image library
	rectImage := image.NewRGBA(image.Rect(0, 0, choice1, choice2))
	//image library used for basic 2d creations for more information you can check it on https://pkg.go.dev/image
	// green := color.RGBA{0, 100, 0, 255} // other colors for reference if you wanna change it
	// red := color.RGBA{100, 0, 0, 255}
	blue := color.RGBA{0, 0, 100, 255} // I use blue because I wanna

	// simple draw function for draw it
	draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// if there is error in creation of image it will alert us
	file, err := os.Create(rectangle)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	//this will encode rectangle as png format I can also use another formats such as jpeg, gif vb vb vb
	png.Encode(file, rectImage)
}

func main() {
	// lets get user inputs first
	var choice1 int
	var choice2 int
	// can't remember which used for kenar in english
	fmt.Print("enter the sides or edges for your rectangle :) : ")
	_, err := fmt.Scan(&choice1, &choice2)
	if err != nil {
		fmt.Println("Invalid inputs try again: err:", err)
	}

	showAreaAndCircumference(choice1, choice2)
	letsDrawIt(choice1, choice2)

}
