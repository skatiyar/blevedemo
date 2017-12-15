package blevedemo

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

const indexPath = "sites.index"

var indexer bleve.Index

func createMapping() (*mapping.IndexMappingImpl, error) {
	return bleve.NewIndexMapping(), nil
}

// initialise
func init() {
	index, indexErr := bleve.Open(indexPath)
	if indexErr != nil {
		if indexErr != bleve.ErrorIndexPathDoesNotExist {
			panic(indexErr)
		}

		newMapping, newMappingErr := createMapping()
		if newMappingErr != nil {
			panic(newMappingErr)
		}

		newIndex, newIndexErr := bleve.New(indexPath, newMapping)
		if newIndexErr != nil {
			panic(newIndexErr)
		}

		index = newIndex
	}

	indexer = index
}
