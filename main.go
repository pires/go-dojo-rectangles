package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"image"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/pires/go-dojo-rectangles/geometry"
)

var (
	filepath = flag.String("json", "rectangles.json", "Path to the JSON file containing rectangles definition.")

	rectangles     map[string]geometry.Rectangle
	rectanglesKeys []string

	// errors
	errCantReadFile        = errors.New("Can't read JSON file.")
	errCantDecodeJSON      = errors.New("Can't decode JSON.")
	errNotEnoughRectangles = errors.New("There less than two (2) rectangles defined in the JSON file.")
)

// loadRectanglesFromJSON loads rectangles specified in a JSON file
func loadRectanglesFromJSON(r io.Reader) (map[string]geometry.Rectangle, error) {
	// Load JSON
	jsonParser := json.NewDecoder(r)

	var jsonRectangles []geometry.Rectangle
	if err := jsonParser.Decode(&jsonRectangles); err != nil {
		return nil, errCantDecodeJSON
	}

	// How many rectangles were loaded from JSON file?
	total := len(jsonRectangles)

	if total < 2 {
		return nil, errNotEnoughRectangles
	}

	// Convert []rectangles to []Rectangle
	processedRectangles := make(map[string]geometry.Rectangle, total)
	for _, jsonRectangle := range jsonRectangles {
		r := image.Rectangle{
			Min: image.Point{X: jsonRectangle.P1.X, Y: jsonRectangle.P1.Y},
			Max: image.Point{X: jsonRectangle.P2.X, Y: jsonRectangle.P2.Y},
		}
		processedRectangles[jsonRectangle.Name] = geometry.Rectangle{Name: jsonRectangle.Name, Inner: r}
	}

	return processedRectangles, nil
}

// loadRectanglesFromJsonFile loads rectangles specified in a JSON file
func loadRectanglesFromJSONFile(filepath string) (map[string]geometry.Rectangle, error) {
	// Open JSON file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errCantReadFile
	}

	return loadRectanglesFromJSON(file)
}

func main() {
	// Read flags
	flag.Parse()

	var err error

	// Load from JSON file
	if rectangles, err = loadRectanglesFromJSONFile(*filepath); err != nil {
		panic(err)
	}

	// Go map iteration with range isn't sorted but we want order!
	// See https://blog.golang.org/go-maps-in-action#TOC_7.
	for k := range rectangles {
		rectanglesKeys = append(rectanglesKeys, k)
	}
	sort.Strings(rectanglesKeys)

	// Compute and process all rectangles
	for _, name1 := range rectanglesKeys {
		// Print rectangle name
		fmt.Println(fmt.Sprintf("=> [%s]", name1))

		// Iteratively build operation results
		var intersects, contains, adjacents []string

		// Perform calculations and append results
		for _, name2 := range rectanglesKeys {
			if name1 != name2 {
				rectangle1 := rectangles[name1]
				rectangle2 := rectangles[name2]

				// Intersection
				if rectangle1.Intersects(rectangle2) {
					intersects = append(intersects, fmt.Sprintf("%s (intersection points: %+v)", rectangle2.Name, rectangle1.IntersectionPoints(rectangle2)))
				}

				// Containment
				if rectangle1.Contains(rectangle2) {
					contains = append(contains, rectangle2.Name)
				}

				// Adjacency
				if rectangle1.IsAdjacent(rectangle2) {
					adjacents = append(adjacents, rectangle2.Name)
				}
			}
		}

		// If any operations didn't return results, pretty print.
		if len(intersects) == 0 {
			intersects = append(intersects, "None")
		}
		if len(contains) == 0 {
			contains = append(contains, "None")
		}
		if len(adjacents) == 0 {
			adjacents = append(adjacents, "None")
		}

		// Print results
		fmt.Println(fmt.Sprintf("%s %s", "  => Intersects: ", strings.Join(intersects[:], ", ")))
		fmt.Println(fmt.Sprintf("%s %s", "  => Contains: ", strings.Join(contains[:], ", ")))
		fmt.Println(fmt.Sprintf("%s %s", "  => Is adjacent to: ", strings.Join(adjacents[:], ", ")))
	}
}
