package blevedemo

import (
	"encoding/json"
	"net/http"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
)

func Search(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(400)
		rw.Write([]byte("Invalid method"))
		return
	}

	rawQuery := req.URL.Query().Get("q")
	searchReq := bleve.NewSearchRequest(query.NewQueryStringQuery(rawQuery))
	// Set result size
	searchReq.Size = 30
	// Enable highlighting
	searchReq.Highlight = bleve.NewHighlight()

	result, resultErr := indexer.Search(searchReq)
	if resultErr != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(resultErr.Error()))
		return
	}

	res, resErr := json.Marshal(result.Hits)
	if resErr != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(resErr.Error()))
		return
	}

	rw.WriteHeader(200)
	rw.Write(res)
}
