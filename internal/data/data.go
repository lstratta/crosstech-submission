package data

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lstratta/crosstech-submission/internal/models"
)

func ParseJsonData() ([]models.Track, error) {
	tArr := []models.Track{}

	jsonData, err := os.ReadFile("./data/data.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var newStr string

	strSplit := strings.Split(string(jsonData), "\n")
	for _, s := range strSplit {

		s2 := strings.Replace(s, "NaN", "0.0", -1)
		newS := strings.Replace(s2, "null", "\"null\"", -1)

		newStr = newStr + newS
	}

	if err = json.Unmarshal([]byte(newStr), &tArr); err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %v", err)
	}

	return tArr, nil
}
