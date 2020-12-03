package main

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"strconv"
)

// Graph is a simple alias for multidimensional slice (dynamic array).
// We can treat elements of the slice as vertices of the graph and
// arithmetic relations between coordinates of the elements as edges.
type Graph [][]int

// NewGraph returns an initialized Graph with dimensions of underlying
// slice of (n, m).
func NewGraph(n, m int) Graph {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, m)
	}

	return g
}

// Vertex represents the node of a graph.
type Vertex struct {
	i, j int
}

// white color code, use to avoid magic numbers
const white = 1

func main() {
	// Number of the test cases to run
	t := 0

	fmt.Print("> ")
	if _, err := fmt.Scanf("%d", &t); err != nil {
		log.Fatalf("failed to read from STDOUT: %v", err)
	}

	for t > 0 {
		// Bitmap dimensions for a test case
		n, m := 0, 0
		if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
			log.Fatalf("failed to read from STDOUT: %v", err)
		}

		// A graph to store a bitmap for the testcase.
		bitmap := NewGraph(n, m)

		// Grid will contain a distance map,
		// it should have the same dimensions as the original bitmap.
		grid := NewGraph(n, m)

		// Read the bitmap description words
		for i := 0; i < n; i += 1 {
			w := ""
			if _, err := fmt.Scanf("%s", &w); err != nil {
				log.Fatalf("failed to read from STDOUT: %v", err)
			}

			for j, r := range []rune(w) {
				v, _ := strconv.ParseInt(string(r), 10, 64)
				bitmap[i][j] = int(v)
			}
		}

		// Path lookup
		for i := range bitmap {
			for j := range bitmap[i] {
				if bitmap[i][j] != white {
					grid[i][j] = Search(bitmap, Vertex{i, j}, func(n int) bool {
						return n == white
					})
				}
			}
		}

		Show(grid)

		t -= 1
	}
}

// Show pretty-prints the given slice to the stdout.
func Show(s [][]int) {
	for i := range s {
		for j := range s[i] {
			fmt.Printf("%d ", s[i][j])
		}
		fmt.Printf("\n")
	}
}

// Search traverses the graph 'g' using modified Breadth First Search
// algorithm - instead of traversing the whole graph, it returns as
// soon as 'searchFunc' returns 'true'.
// The nature of BFS guarantees that first node we find that satisfies
// the 'searchFunc' clause will be the closest to the 'start' vertex.
func Search(g Graph, start Vertex, searchFunc func(int) bool) int {
	visited := make(map[Vertex]bool)
	current := start

	// To avoid implementing our own type we will use
	// 'container/list' instead.
	// Using 'PushBack', 'Front' and 'Remove' we can make
	// behave like a queue.
	q := list.New()

	for {
		// Make sure we mark the current vertice as visited to avoid endless
		// loops.
		visited[current] = true

		// Found what we're looking for, return early.
		if searchFunc(g[current.i][current.j]) {
			return int(math.Abs(float64(start.i-current.i)) +
				math.Abs(float64(start.j-current.j)))
		}

		// Find all the neighbors of the current vertex.
		for _, n := range FindNeighbors(g, current) {
			if !visited[n] {
				q.PushBack(n)
			}
		}

		e := q.Front()
		if e == nil {
			// Nothing else to visit, exit the loop
			break
		}

		// Change the current vertex to the next one in the queue.
		current = e.Value.(Vertex)

		// Make sure to remove the vertice from the queue, so we don't visit
		// it more than once.
		q.Remove(e)
	}

	return 0
}

// FindNeighbors returns all the 'o' vertex
// neighbors based on the 'o's coordinates 'i', 'j' and the graph
// dimensions.
func FindNeighbors(g Graph, o Vertex) []Vertex {
	var nbrs []Vertex
	n, m := len(g), len(g[0])

	if o.i-1 >= 0 {
		nbrs = append(nbrs, Vertex{i: o.i - 1, j: o.j})
	}

	if o.i+1 < n {
		nbrs = append(nbrs, Vertex{i: o.i + 1, j: o.j})
	}

	if o.j-1 >= 0 {
		nbrs = append(nbrs, Vertex{i: o.i, j: o.j - 1})
	}

	if o.j+1 < m {
		nbrs = append(nbrs, Vertex{i: o.i, j: o.j + 1})
	}

	return nbrs
}
