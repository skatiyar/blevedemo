package blevedemo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Index(filePath string) error {
	file, fileError := os.Open(filePath)
	if fileError != nil {
		return fileError
	}
	defer file.Close()

	reader := csv.NewReader(file)
	count := 0
	for {
		data, dataErr := reader.Read()
		if dataErr != nil {
			if dataErr != io.EOF {
				return dataErr
			}

			return nil
		}
		if len(data) != 2 {
			continue
		}

		doc, docErr := goquery.NewDocument("http://" + data[0])
		if docErr != nil {
			fmt.Println("Error: ", data[0], docErr.Error())
			continue
		}

		if indexErr := indexer.Index(strconv.Itoa(count), Page{
			Title:   doc.Find("title").Text(),
			Content: doc.Find("body").Find("script").Remove().End().Text(),
			Tags:    data[1],
		}); indexErr != nil {
			return indexErr
		}

		count += 1
	}
}
