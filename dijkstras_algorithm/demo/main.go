package main

import (
	"fmt"
)

const inf = int(^uint(0) >> 1)

type Point struct {
	name   string
	cost   int
	parent *Point
	child  *Point
}

func main() {
	graph := map[string]map[string]int{
		"a": {"b": 5, "c": 2},
		"b": {"a": 5, "c": 7, "d": 8},
		"c": {"a": 2, "b": 7, "d": 4, "e": 8},
		"d": {"b": 8, "c": 4, "e": 6, "f": 4},
		"e": {"c": 8, "d": 6, "f": 3},
		"f": {"e": 3, "d": 4},
	}
	start := "a"
	end := "f"
	mainRoute := make([]string, 0)

	// самые короткие пути к каждой точке
	shortestPath := make(map[string]int, len(graph))
	// непосещенные точки
	unvisited := make(map[string]struct{})

	// сначала у всех точек максимальный путь и все непосещенные
	for k, _ := range graph {
		shortestPath[k] = inf
		unvisited[k] = struct{}{}
	}

	// для начальной точки стоимость пути = 0
	shortestPath[start] = 0

	// пути [ребенок]родитель
	path := make(map[string]string)

	for len(unvisited) > 0 {
		minPoint := ""

		// находим минимальную точку из непосещенных
		for point := range unvisited {
			if minPoint == "" || shortestPath[point] < shortestPath[minPoint] {
				minPoint = point
			}
		}

		// проходимся по детям минимальной точки
		for child, cost := range graph[minPoint] {
			if shortestPath[child] > cost+shortestPath[minPoint] {
				shortestPath[child] = cost + shortestPath[minPoint]
				path[child] = minPoint
			}
		}

		delete(unvisited, minPoint)

		node := end
		route := []string{}
		for node != start {
			if parent, ok := path[node]; ok {
				route = append([]string{node}, route...)
				node = parent
			} else {
				break
			}
		}
		route = append([]string{start}, route...)

		mainRoute = route
	}

	fmt.Println("Shortest Path:", shortestPath)
	fmt.Println("mainRoute: ", mainRoute)
}
