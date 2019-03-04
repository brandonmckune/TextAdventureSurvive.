package engine

import (
	"fmt"
)

type LevelDetails struct {
	ConnectingLevels map[string]*LevelDetails
	Details          [DISPLAYHEIGHT][DISPLAYWIDTH]string
}

var ()

func init() {

}

/* LoadLevelRow allows the loader to read a single line from a file and insert that
 * data into the Level Details 2D array.
 */
func (d LevelDetails) LoadRow(row string, rowIndex int) {
	if rowIndex < 0 || rowIndex > 79 {
		fmt.Println("Invalid row index passed in LoadRow(", rowIndex, ")")
		panic(-1)
	}

	for idx, char := range row {
		if idx >= DISPLAYWIDTH {
			break
		}

		d.Details[rowIndex][idx] = string(char)
	}
}

func (d LevelDetails) AddConnectingLevel(key string, level *LevelDetails) {
	if _, found := d.ConnectingLevels[key]; !found {
		d.ConnectingLevels[key] = level
	}

	//TODO What do I do if it's there? Replace it?
}

func (d LevelDetails) GetLevel(key string) *LevelDetails {
	if detail, found := d.ConnectingLevels[key]; found {
		return detail
	}

	panic("Level Not Found Exception was thrown!")
}

func (ld LevelDetails) Print() {
	for idy := 0; idy < len(ld.Details); idy++ {
		for idx := 0; idx < len(ld.Details[idy]); idx++ {
			fmt.Print(ld.Details[idy][idx])
		}
		fmt.Print("\n")
	}
}
