package main

import (
	"fmt"
)

func main() {
	type point struct {
		x int
		y int
	}

	fmt.Println("******** points ********")

	points := []point{{0, 0}, {1, 1}}

	fmt.Printf("addr: point[0] %p\n", &points[0])
	fmt.Printf("addr: point[1] %p\n", &points[1])

	for i, p := range points {
		fmt.Printf("addr: p[%d] %p\n", i, &p)

		p.x += 1
		p.y += 1
	}

	fmt.Printf("points: %v\n", points)

	fmt.Println("******** segments ********")

	type segment struct {
		position int
		start    *point
		finish   *point
	}

	segs := []segment{
		{
			position: 0,
			start:    &point{x: 0, y: 0},
			finish:   &point{x: 0, y: 0},
		},
		{
			position: 1,
			start:    &point{x: 1, y: 1},
			finish:   &point{x: 1, y: 1},
		},
	}

	fmt.Printf("addr: segs[0] %p\n", &segs[0])
	fmt.Printf("addr: segs[0].start %p\n", segs[0].start)
	fmt.Printf("addr: segs[0].finish %p\n", segs[0].finish)

	fmt.Printf("addr: segs[1] %p\n", &segs[1])
	fmt.Printf("addr: segs[1].start %p\n", segs[1].start)
	fmt.Printf("addr: segs[1].finish %p\n", segs[1].finish)

	for i, s := range segs {
		fmt.Printf("addr: loop s %p\n", &s)

		fmt.Printf("addr: s[%d].start %p\n", i, s.start)
		fmt.Printf("addr: s[%d].finish %p\n", i, s.finish)

		s.position += 1
		s.start.x += 1
		s.finish.x += 1
	}

	for _, s := range segs {
		fmt.Printf("%d %v %v\n", s.position, *s.start, *s.finish)
	}

	psegs := []*segment{
		{
			position: 0,
			start:    &point{x: 0, y: 0},
			finish:   &point{x: 0, y: 0},
		},
		{
			position: 1,
			start:    &point{x: 1, y: 1},
			finish:   &point{x: 1, y: 1},
		},
	}

	for _, s := range psegs {
		fmt.Printf("%p\n", &s)

		s.position += 1
		s.start.x += 1
		s.finish.x += 1
	}

	for _, s := range psegs {
		fmt.Printf("%d %v %v\n", s.position, *s.start, *s.finish)
	}

	fmt.Println("******** collecting ********")

	points = []point{{0, 0}, {1, -1}, {3, -3}}

	negativeYPoints := []*point{}

	for _, p := range points {
		if p.y < 0 {
			negativeYPoints = append(negativeYPoints, &p)
		}
	}

	fmt.Printf("negativeY: %v\n", negativeYPoints)

	for _, p := range negativeYPoints {
		fmt.Printf("negativeY: p %v\n", p)
	}
}
