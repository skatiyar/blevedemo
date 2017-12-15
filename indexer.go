package blevedemo

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Index() error {
	file, fileError := os.Open("public/sites.csv")
	if fileError != nil {
		return fileError
	}

	reader := csv.NewReader(file)
	for {
		data, dataErr := reader.Read()
		if dataErr != nil {
			return dataErr
		}

		fmt.Println(data)
	}
}
