// This function computes the correlation of two sequences. 
// Imagine a video is a sequence of values
// Also, imagine a feature vector is a sequence

package main

import "fmt"

func mean(x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func variance(x []float64) float64 {
	avg := mean(x)
	sum := 0.0
	for _, v := range x {
		sum += (v - avg) * (v - avg)
	}
	return sum / float64(len(x)-1)

}

func covariance(x []float64, y []float64) float64 {
	avgx := mean(x)
	avgy := mean(y)

	sum := 0.0

	for i := 0; i < len(x); i++ {
		sum += (x[i] - avgx) * (y[i] - avgy)
	}

	return sum / float64(len(x)-1)
}

func ssim(x []float64, y []float64) float64 {
	avgx := mean(x)
	avgy := mean(y)

	c1 := (0.01 * 255) * (0.01 * 255)
	c2 := (0.03 * 255) * (0.03 * 255)

	covxy := covariance(x, y)

	a := (2*avgx*avgy + c1) * (2*covxy + c2)
	b := (avgx*avgx + avgy*avgy + c1) * (variance(x) + variance(y) + c2)

	return a / b
}

func main() {
	x := []float64{10, 22, 33, 44, 55}
	y := []float64{10, 22, 33, 44, 55}
	fmt.Println(x, y, "correlation", ssim(x,y))

	x = []float64{10, 22, 33, 44, 55}
	y = []float64{80, 25, 123, 238, 205}
	fmt.Println(x, y, "correlation", ssim(x,y))

	x = []float64{20, 222, 213, 83, 83}
	y = []float64{183, 83, 1, 18, 2}
	fmt.Println(x, y, "correlation", ssim(x,y))

	x = []float64{0, 0, 0, 0, 0}
	y = []float64{255, 255, 255, 255, 255}
	fmt.Println(x, y, "correlation", ssim(x,y))

	y = []float64{11.20, 22.06, 33.98, 44.04, 54.98}
	fmt.Println(x, y, "correlation", ssim(x,y))

	x = []float64{192, 192, 192, 192, 192}
	y = []float64{191, 192, 189, 190, 191}
	fmt.Println(x, y, "correlation", ssim(x,y))
}
/*
$ go run ssim.go
[10 22 33 44 55] [10 22 33 44 55] correlation 1.00
[10 22 33 44 55] [80 25 123 238 205] correlation 0.15
[20 222 213 83 83] [183 83 1 18 2] correlation -0.31
[0 0 0 0 0] [255 255 255 255 255] correlation 0.00
[0 0 0 0 0] [11.2 22.06 33.98 44.04 54.98] correlation 0.00
[192 192 192 192 192] [191 192 189 190 191] correlation 0.97
*/
