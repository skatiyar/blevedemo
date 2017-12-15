package blevedemo

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/analysis/token/apostrophe"
	"github.com/blevesearch/bleve/analysis/token/camelcase"
	"github.com/blevesearch/bleve/analysis/token/lowercase"
	"github.com/blevesearch/bleve/analysis/token/stop"
	"github.com/blevesearch/bleve/analysis/tokenizer/letter"
	"github.com/blevesearch/bleve/analysis/tokenizer/unicode"
	"github.com/blevesearch/bleve/analysis/tokenmap"
	"github.com/blevesearch/bleve/mapping"
)

const indexPath = "sites.index"

var indexer bleve.Index

type Page struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

func (p Page) Type() string {
	return "page"
}

func createMapping() (*mapping.IndexMappingImpl, error) {
	newMapping := bleve.NewIndexMapping()

	if mapErr := newMapping.AddCustomTokenMap("tagsMap", map[string]interface{}{
		"type": tokenmap.Name,
		"tokens": []interface{}{
			"top",
			"Top",
		},
	}); mapErr != nil {
		return nil, mapErr
	}

	// create custom token filter
	if tagsFilterErr := newMapping.AddCustomTokenFilter("tagsFilter", map[string]interface{}{
		"type":           stop.Name,
		"stop_token_map": "tagsMap",
	}); tagsFilterErr != nil {
		return nil, tagsFilterErr
	}

	// create custom text analyzer
	if textAnalyzerErr := newMapping.AddCustomAnalyzer("textAnalyzer", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": unicode.Name,
		"token_filters": []string{
			apostrophe.Name,
			lowercase.Name,
			camelcase.Name,
		},
	}); textAnalyzerErr != nil {
		return nil, textAnalyzerErr
	}

	// create custom text analyzer
	if tagsAnalyzerErr := newMapping.AddCustomAnalyzer("tagsAnalyzer", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": letter.Name,
		"token_filters": []string{
			lowercase.Name,
			"tagsFilter",
		},
	}); tagsAnalyzerErr != nil {
		return nil, tagsAnalyzerErr
	}

	// create new page mapping
	pageMapping := mapping.NewDocumentMapping()

	textFieldMapping := mapping.NewTextFieldMapping()
	textFieldMapping.Analyzer = "textAnalyzer"
	textFieldMapping.Store = true
	textFieldMapping.Index = true
	textFieldMapping.IncludeTermVectors = true
	textFieldMapping.IncludeInAll = true

	tagsFieldMapping := mapping.NewTextFieldMapping()
	tagsFieldMapping.Analyzer = "tagsAnalyzer"
	tagsFieldMapping.Store = true
	tagsFieldMapping.Index = true
	tagsFieldMapping.IncludeTermVectors = true
	tagsFieldMapping.IncludeInAll = true

	pageMapping.AddFieldMappingsAt("title", textFieldMapping)
	pageMapping.AddFieldMappingsAt("content", textFieldMapping)
	pageMapping.AddFieldMappingsAt("tags", tagsFieldMapping)

	newMapping.AddDocumentMapping("snippet", pageMapping)

	if mappingErr := newMapping.Validate(); mappingErr != nil {
		return nil, mappingErr
	}

	return newMapping, nil
}

// initialise
func Init() error {
	index, indexErr := bleve.Open(indexPath)
	if indexErr != nil {
		if indexErr != bleve.ErrorIndexPathDoesNotExist {
			return indexErr
		}

		newMapping, newMappingErr := createMapping()
		if newMappingErr != nil {
			return newMappingErr
		}

		newIndex, newIndexErr := bleve.New(indexPath, newMapping)
		if newIndexErr != nil {
			return newIndexErr
		}

		index = newIndex
	}

	indexer = index

	return nil
}
