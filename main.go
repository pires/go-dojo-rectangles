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

	"github.com/pires/go-dojo-rectangles/geometry"

	"github.com/olekukonko/tablewriter"
)

const (
	yes           = "X"
	no            = ""
	notApplicable = "-"
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

// computeIntersectionsTable computes intersections and returns results formatted as an ASCII table.
func computeIntersectionsTable() (table *tablewriter.Table) {
	table = newTable("Intersects?")
	// Matrix values
	for _, name1 := range rectanglesKeys {
		// Print rectangle name
		row := make([]string, 0, len(rectanglesKeys)+1)
		row = append(row, name1)
		for _, name2 := range rectanglesKeys {
			if name1 == name2 {
				row = append(row, notApplicable)
			} else {
				rectangle1 := rectangles[name1]
				rectangle2 := rectangles[name2]

				if rectangle1.Intersects(rectangle2) {
					row = append(row, yes)
				} else {
					row = append(row, no)
				}
			}
		}
		table.Append(row)
	}

	return
}

// computeContainsTable computes contains and returns results formatted an ASCII table.
func computeContainsTable() (table *tablewriter.Table) {
	table = newTable("Contains?")
	// Matrix values
	for _, name1 := range rectanglesKeys {
		// Print rectangle name
		row := make([]string, 0, len(rectanglesKeys)+1)
		row = append(row, name1)
		for _, name2 := range rectanglesKeys {
			if name1 == name2 {
				row = append(row, notApplicable)
			} else {
				rectangle1 := rectangles[name1]
				rectangle2 := rectangles[name2]

				if rectangle1.Contains(rectangle2) {
					row = append(row, yes)
				} else {
					row = append(row, no)
				}
			}
		}
		table.Append(row)
	}

	return
}

// computeIsAdjacentTable computes adjacency and returns results formatted an ASCII table.
func computeIsAdjacentTable() (table *tablewriter.Table) {
	table = newTable("Is Adjacent?")
	// Matrix values
	for _, name1 := range rectanglesKeys {
		// Print rectangle name
		row := make([]string, 0, len(rectanglesKeys)+1)
		row = append(row, name1)
		for _, name2 := range rectanglesKeys {
			if name1 == name2 {
				row = append(row, notApplicable)
			} else {
				rectangle1 := rectangles[name1]
				rectangle2 := rectangles[name2]

				if rectangle1.IsAdjacent(rectangle2) {
					row = append(row, yes)
				} else {
					row = append(row, no)
				}
			}
		}
		table.Append(row)
	}

	return
}

// newTable returns a table ready to be populated.
func newTable(op string) (table *tablewriter.Table) {
	table = tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	// Matrix header
	header := make([]string, 0, len(rectanglesKeys))
	header = append(header, op)
	for _, rectangleName := range rectanglesKeys {
		header = append(header, rectangleName)
	}
	table.SetHeader(header)

	return
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

	// Render
	println()
	computeIntersectionsTable().Render()

	for _, name1 := range rectanglesKeys {
		// Print rectangle name
		for _, name2 := range rectanglesKeys {
			if name1 != name2 {
				rectangle1 := rectangles[name1]
				rectangle2 := rectangles[name2]
				if rectangle1.Intersects(rectangle2) {
					fmt.Printf("%+v \n", rectangle1.IntersectionPoints(rectangle2))
				}
			}
		}
	}

	println()
	computeContainsTable().Render()
	println()
	computeIsAdjacentTable().Render()
	println()
}
