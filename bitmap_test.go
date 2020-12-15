package main

import (
	"reflect"
	"testing"
)

func TestFindNeighbors(t *testing.T) {
	cases := []struct {
		g        Graph
		o        Vertex
		expected []Vertex
	}{
		{
			g: NewGraph(2, 2),
			o: Vertex{0, 0},
			expected: []Vertex{
				{i: 1, j: 0},
				{i: 0, j: 1},
			},
		},
		{
			g: NewGraph(3, 4),
			o: Vertex{1, 1},
			expected: []Vertex{
				{0, 1},
				{2, 1},
				{1, 0},
				{1, 2},
			},
		},
	}

	for _, c := range cases {
		got := FindNeighbors(c.g, c.o)

		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("failed, expected = %v, got = %v", c.expected, got)
		}
	}
}

func TestMap(t *testing.T) {
	cases := []struct {
		g        Graph
		whites   []Vertex
		expected Graph
	}{
		{
			g:        NewGraph(2, 2),
			expected: NewGraph(2, 2), // Nothing found
		},
		{
			g:        FromRows([]int{0, 0}, []int{0, 1}),
			whites:   []Vertex{{0, 0}},
			expected: FromRows([]int{2, 1}, []int{1, 0}),
		},
	}

	for _, c := range cases {
		got := Map(c.g, c.whites)

		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("failed, expected = %v, got = %v", c.expected, got)
		}
	}
}
