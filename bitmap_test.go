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

func TestSearch(t *testing.T) {
	cases := []struct{
		g Graph
		start Vertex
		searchFunc func(int) bool
		expected int
	}{
		{
			g: NewGraph(2, 2),
			start: Vertex{0, 0},
			searchFunc: func(n int) bool {
				return n != 0
			},
			expected: 0, // Nothing found
		},
		{
			g: func() Graph {
				g := NewGraph(2, 2)

				g[1][1] = white

				return g
			}(),
			start: Vertex{0, 0},
			searchFunc: func(n int) bool {
				return n == white
			},
			expected: 2,
		},
	}

	for _, c := range cases {
		got := Search(c.g, c.start, c.searchFunc)

		if got != c.expected {
			t.Errorf("failed, expected = %v, got = %v", c.expected, got)
		}
	}
}