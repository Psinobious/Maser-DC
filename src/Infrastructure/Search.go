package Infrastructure

import(
	"net/http"
	"context"
	"encoding/json"
	"log"
	"bytes"
	
)

func (u *WebServiceHandler) SearchPublications(w http.ResponseWriter, r *http.Request){
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := u.ESC.Search(
		u.ESC.Search.WithContext(context.Background()),
		u.ESC.Search.WithIndex("publications"),
		u.ESC.Search.WithBody(&buf),
		u.ESC.Search.WithTrackTotalHits(true),
		u.ESC.Search.WithPretty(),
	  )
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()


}