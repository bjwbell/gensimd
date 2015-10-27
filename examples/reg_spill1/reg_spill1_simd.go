// +build !amd64,gc

package main

func regspill1(x, y int32) int32 {
	dist := int32(0)
	xi := x
	xj := x
	yi := y
	yj := y
	dx := xj + xi
	dy := yj + yi
	sqX := dx * dx
	sqY := dy * dy
	sqDist := sqX + sqY
	dx2 := dx - dy
	t := 2*dx2 + dy
	dist = sqDist + t
	return dist
}
