package graph

import (
	"fmt"
	"math"

	"github.com/juan84n/estructurasdedatos/lists"
)

type Graph struct {
	value      int32
	vertices   int32
	mAdyacence [][]int32
	lAdyacence map[int32]*lists.ArrayList
}

func (g *Graph) AddEdge(vertice int32, edge int32) error {
	if g.mAdyacence != nil {
		return g.addEdgeMA(vertice, edge)
	} else if g.lAdyacence != nil {
		return g.addEdgeLA(vertice, edge)
	}
	return nil
}

func (g *Graph) GetNeighbors(vertice int32) *lists.ArrayList {
	if len(g.lAdyacence) > 0 {
		return g.getNeighborsLA(vertice)
	} else if len(g.mAdyacence) > 0 {
		return g.getNeighborsMA(vertice)
	}
	return nil
}

func (g *Graph) ShortestPath(from int32, to int32) (lists.ArrayList, error) {
	err := g.validateData(from, to)
	if err != nil {
		return lists.ArrayList{}, err
	}
	paths := *lists.NewArrayList()
	var shortest []lists.ArrayList
	if g.lAdyacence != nil {
		vert, ok := g.lAdyacence[from]
		if ok {
			index := vert.Search(to)
			if index >= 0 {
				paths.Push(from)
				paths.Push(to)
				return paths, nil
			}
			shortest, paths = g.shortestPathLA(from, to, paths, shortest)
		}
	} else if g.mAdyacence != nil {
		if g.mAdyacence[from][to] == 1 {
			paths.Push(from)
			paths.Push(to)
			return paths, nil
		}
		shortest, paths = g.shortestPathMA(from, to, paths, shortest)
	}
	temSize := math.MaxInt32
	var shortestNew lists.ArrayList
	for _, list := range shortest {
		if list.Size() < temSize {
			temSize = list.Size()
			shortestNew = list
		}
	}

	return shortestNew, nil
}

func (g *Graph) ShowGraph() {
	if g.mAdyacence != nil {
		for i := 0; i < len(g.mAdyacence); i++ {
			for j := 0; j < len(g.mAdyacence); j++ {
				fmt.Print(g.mAdyacence[i][j], " ")
			}
			fmt.Println("")
		}
	} else if g.lAdyacence != nil {
		for key, element := range g.lAdyacence {
			fmt.Println("Key:", key, "=>", "Elements:")
			element.PrintArrayList()
		}
	}
}

func (g *Graph) addEdgeMA(vertice int32, edge int32) error {
	err := g.validateData(vertice, edge)
	if err != nil {
		return err
	}
	g.mAdyacence[vertice][edge] = 1
	return nil
}

func (g *Graph) addEdgeLA(vertice int32, edge int32) error {
	err := g.validateData(vertice, edge)
	if err != nil {
		return err
	}
	vert, ok := g.lAdyacence[vertice]
	if !ok {
		arrayList := lists.NewArrayList()
		arrayList.Push(edge)
		g.lAdyacence[vertice] = arrayList
	} else {
		vert.Push(edge)
	}
	return nil
}

func (g *Graph) validateData(vertice int32, edge int32) error {
	if vertice < 0 || int(vertice) > int(g.vertices) {
		return fmt.Errorf("Vertice is out of range")
	}
	if edge < 0 || int(edge) > int(g.vertices) {
		return fmt.Errorf("Edge is out of range")
	}
	return nil
}

func (g *Graph) getNeighborsLA(vertice int32) *lists.ArrayList {
	return g.lAdyacence[vertice]
}
func (g *Graph) getNeighborsMA(vertice int32) *lists.ArrayList {
	arrayList := lists.NewArrayList()
	for i := 0; i < int(g.vertices); i++ {
		if g.mAdyacence[vertice][i] == 1 {
			arrayList.Push(int32(i))
		}
	}
	return arrayList
}
func (g *Graph) shortestPathMA(from int32, to int32, paths lists.ArrayList, shortest []lists.ArrayList) ([]lists.ArrayList, lists.ArrayList) {
	for i := 0; i < int(g.vertices); i++ {
		if g.mAdyacence[from][i] == 1 {
			paths.Push(from)
			if int32(i) == to {
				paths.Push(int32(i))
				shortest = append(shortest, paths)
				return shortest, *lists.NewArrayList()
			}
			shortest, paths = g.shortestPathMA(int32(i), to, paths, shortest)
		}
	}
	return shortest, paths
}
func (g *Graph) shortestPathLA(from int32, to int32, paths lists.ArrayList, shortest []lists.ArrayList) ([]lists.ArrayList, lists.ArrayList) {
	for i := 0; i < g.lAdyacence[from].Size(); i++ {
		value := g.lAdyacence[from].Get(i)
		if value == to {
			paths.Push(int32(i))
			shortest = append(shortest, paths)
			return shortest, *lists.NewArrayList()
		}
		shortest, paths = g.shortestPathLA(value, to, paths, shortest)
	}
	return shortest, paths
}

func NewGraphMA(vertices int32) *Graph {
	graph := &Graph{vertices: vertices, mAdyacence: make([][]int32, vertices)}
	for i := 0; i < int(vertices); i++ {
		graph.mAdyacence[i] = make([]int32, vertices)
		for j := 0; j < int(vertices); j++ {
			graph.mAdyacence[i][j] = 0
		}
	}
	return graph
}

func NewGraphLA(vertices int32) *Graph {
	graph := &Graph{vertices: vertices, lAdyacence: make(map[int32]*lists.ArrayList)}
	return graph
}
