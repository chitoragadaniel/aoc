package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	d5()
}

func d1() {
	// Initialization
	data, _ := os.ReadFile("input/d1.txt")

	// Part 1
	var floor int // The floor where Santa is
	for _, v := range data {
		if v == 40 {
			floor++
		} else {
			floor--
		}
	}
	printP1(floor)

	// Part 2
	floor = 0
	for i := 0; i < len(data); i++ {
		if data[i] == 40 {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			printP2(i + 1)
			break
		}
	}
}
func d2() {
	// Initialization
	data, _ := os.ReadFile("input/d2.txt")
	gifts := strings.Split(string(data), "\r\n")

	// Part 1 & 2
	var paper int
	var ribbon int
	for i := 0; i < len(gifts); i++ {
		// Splitting each row on "x"
		row := strings.Split(gifts[i], "x")
		l, _ := strconv.Atoi(row[0])
		h, _ := strconv.Atoi(row[1])
		w, _ := strconv.Atoi(row[2])
		paper += min(l*h, h*w, w*l) + 2*(l*h+h*w+w*l)
		ribbon += min(l+h, h+w, w+l)*2 + l*h*w
	}
	printP1(paper)
	printP2(ribbon)
}
func d3() {
	// Initialization
	data, _ := os.ReadFile("input/d3.txt")

	type Vertex struct {
		x int
		y int
	}

	update := func(x, y *int, c byte) {
		switch string(c) {
		case ">":
			*x++
		case "<":
			*x--
		case "^":
			*y++
		case "v":
			*y--
		}
	}

	// Part 1
	var m = map[Vertex]bool{
		{0, 0}: true,
	}
	var x int
	var y int
	for _, c := range data {
		update(&x, &y, c)
		m[Vertex{x, y}] = true
	}
	printP1(len(m))

	// Part 2
	clear(m)
	m[Vertex{0, 0}] = true

	var x1 int
	var x2 int
	var y1 int
	var y2 int

	for i, c := range data {
		if i%2 == 1 {
			update(&x1, &y1, c)
			m[Vertex{x1, y1}] = true
		} else {
			update(&x2, &y2, c)
			m[Vertex{x2, y2}] = true
		}
	}
	printP2(len(m))
}
func d4() {
	// Part 1
	key := "yzbqklnj"
	hash := func(inp string) string {
		hasher := md5.New()
		hasher.Write([]byte(inp))
		return hex.EncodeToString(hasher.Sum(nil))
	}
	nr := 1
	for {
		res := hash(key + strconv.Itoa(nr))
		if strings.HasPrefix(res, "00000") {
			printP1(nr)
			break
		}
		nr++
	}

	// Part 2
	var l sync.Mutex
	var w sync.WaitGroup

	loop := func(i, t int) {
		for {
			res := hash(key + strconv.Itoa(i))
			if strings.HasPrefix(res, "000000") {
				l.Lock()
				if nr == 0 {
					nr = i
					w.Done()
				}
				l.Unlock()
				break
			}
			i += t
		}
	}

	nr = 0
	for i := 0; i < 8; i++ {
		go loop(i, 8)
	}
	w.Add(1)
	w.Wait()
	printP2(nr)
}
func d5() {
	// Initialization
	lines := scanLines("input/d5.txt")

	// Part 1
	count := 0
	v_list := []string{"a", "e", "i", "o", "u"}
	il_list := []string{"ab", "cd", "pq", "xy"}

	// Iterating through the list of strings
	for _, str := range lines {
		arr := strings.Split(str, "")
		vs := 0
		twice := false
		legal := true
		pr_c := arr[0]
		if slices.Contains(v_list, pr_c) {
			vs++
		}

		// Iterating through characters of the string
		for i := 1; i < len(arr); i++ {
			c := arr[i]
			if slices.Contains(v_list, c) {
				vs++
			}
			if twice || c == pr_c {
				twice = true
			}
			if slices.Contains(il_list, pr_c+c) {
				legal = false
				break
			}
			pr_c = c
		}

		// Check if the string is legal
		if vs > 2 && twice && legal {
			count++
		}
	}

	printP1(count)
}
