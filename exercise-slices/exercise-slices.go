package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
pic := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        pic[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
            pic[i][j] = uint8((i*j)*(i*j))
        }
    }
    return pic
}

// NOTE: returns a string; in order to view the image, 
// find a base64 to image converter online and paste the string in
func main() {
    pic.Show(Pic)
}
