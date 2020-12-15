package main

import (
	"container/list"
	"fmt"
	"log"
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

// FromRows returns a graph made of passed rows entries.
// If not rows passed an empty graph will be returned.
func FromRows(rows ...[]int) Graph {
	if len(rows) == 0 {
		return NewGraph(0, 0)
	}

	g := make([][]int, len(rows))
	copy(g, rows)

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
		// grid := NewGraph(n, m)

		whites := make([]Vertex, 0)
		// Read the bitmap description words
		for i := 0; i < n; i += 1 {
			w := ""
			if _, err := fmt.Scanf("%s", &w); err != nil {
				log.Fatalf("failed to read from STDOUT: %v", err)
			}

			for j, r := range []rune(w) {
				v, _ := strconv.ParseInt(string(r), 10, 64)
				bitmap[i][j] = int(v)

				if v == white {
					whites = append(whites, Vertex{i: i, j: j})
				}
			}
		}

		Show(Map(bitmap, whites))

		// Done with the test case
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

func Map(g Graph, whites []Vertex) Graph {
	n, m := len(g), len(g[0])
	grid := NewGraph(n, m)

	if len(whites) == 0 {
		return grid
	}

	q := list.New()

	for _, w := range whites {
		q.PushBack(w)
	}

	e := q.Front()

	// Change the current vertex to the next one in the queue.
	current := e.Value.(Vertex)
	q.Remove(e)

	d := 1

	nbrs := make([]Vertex, 0)

	for {
		fmt.Printf("distance = %d, current = [%d;%d]\n", d, current.i, current.j)

		for _, n := range FindNeighbors(g, current) {
			if g[n.i][n.j] != white && grid[n.i][n.j] == 0 {
				grid[n.i][n.j] = d
				nbrs = append(nbrs, n)
			}
		}

		// Prepare the next batch of nodes to visit.
		e := q.Front()
		if e == nil {
			// Nothing else to visit, exit the loop
			if len(nbrs) == 0 {
				break
			}

			d += 1
			for _, n := range nbrs {
				q.PushBack(n)
			}
			nbrs = make([]Vertex, 0)

			e = q.Front()
		}

		// Change the current vertex to the next one in the queue.
		current = e.Value.(Vertex)

		// Make sure to remove the vertex from the queue, so we don't visit
		// it more than once.
		q.Remove(e)
	}

	return grid
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
