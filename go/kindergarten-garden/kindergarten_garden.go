package kindergarten

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	childPlantsMap map[string][]*ChildPlant
}

const (
	Grass  = "G"
	Clover = "C"
	Radish = "R"
	Violet = "V"
)

const plantsPerChild = 4
const childPlantsPerRow = 2

type ChildPlant struct {
	Id          string
	Description string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

var validDiagramRegexp = regexp.MustCompile(`^\n(V|C|R|G)+\n(V|C|R|G)+$`)

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !validDiagramRegexp.Match([]byte(diagram)) {
		return nil, errors.New("invalid diagram")
	}

	d := strings.Trim(diagram, "\n")
	rows := strings.Split(d, "\n")
	if len(rows) != 2 {
		return nil, errors.New("invalid diagram")
	}

	garden := make([][]string, childPlantsPerRow)
	for i, row := range rows {
		garden[i] = strings.Split(row, "")
		if len(garden[i]) != len(children)*2 {
			return nil, errors.New("invalid diagram")
		}
	}

	// sort children alphabetically in ascending order
	slices.Sort(children)

	// map used to store the plants for each child indexed by the childName
	childMap := make(map[string][]*ChildPlant)

	// Cursor used to map each plants of each children in the garden
	childIndexCursor := 0
	for _, child := range children {
		if _, exists := childMap[child]; exists {
			return nil, errors.New("duplicate child")
		}

		childPlants := make([]*ChildPlant, plantsPerChild)

		// cursor used to index each row of the garden
		gardenRowCursor := 0
		for childPlantCursor := 0; childPlantCursor < len(childPlants); childPlantCursor += childPlantsPerRow {

			// offset from the childIndexCursor
			for offset := 0; offset < childPlantsPerRow; offset++ {
				cp, err := NewChildPlant(garden[gardenRowCursor][childIndexCursor+offset])
				if err != nil {
					return nil, err
				}

				childPlants[childPlantCursor+offset] = cp
			}

			gardenRowCursor++
		}

		childMap[child] = childPlants
		childIndexCursor += childPlantsPerRow
	}

	g := &Garden{
		childMap,
	}

	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, err := g.getPlantsForChild(child)
	return plants, err == nil
}

func (g *Garden) getPlantsForChild(child string) ([]string, error) {
	plants, ok := g.childPlantsMap[child]
	if !ok {
		return nil, errors.New("child not found")
	}

	translatedPlants := make([]string, len(plants))
	for i, p := range plants {
		translatedPlants[i] = p.Description
	}

	return translatedPlants, nil
}

func NewChildPlant(id string) (*ChildPlant, error) {
	description, err := translate(id)
	if err != nil {
		return nil, err
	}

	return &ChildPlant{Id: id, Description: description}, nil
}

func translate(c string) (string, error) {
	translationMap := map[string]string{
		Grass:  "grass",
		Clover: "clover",
		Radish: "radishes",
		Violet: "violets",
	}

	t, ok := translationMap[c]
	if !ok {
		return "", errors.New("invalid plant")
	}

	return t, nil
}
