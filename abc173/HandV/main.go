package main

import "fmt"

func main() {
	var h, w, k int
	fmt.Scanf("%d %d %d", &h, &w, &k)
	c := make([][]string, h, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; i++ {
			fmt.Scan(&c[i][j])
		}
	}
	ans := 0

	for maskR := 0; maskR < (1<<h)-1; maskR++ {
		for maskC := 0; maskC < w; maskC++ {
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {

				}
			}
		}
	}
	fmt.Println(ans)
}
