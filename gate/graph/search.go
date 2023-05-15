package graph

import (
	"encoding/json"
	"gate/graph/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Search(query string, limit int, offset int) (*model.SearchResult, error) {
	res, err := http.Get("http://localhost:8081/search.php?q=" + query + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result = model.SearchResult{}
	json.Unmarshal(resBody, &result)
	return &result, nil
}
