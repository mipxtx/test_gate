package graph

import (
	"encoding/json"
	"fmt"
	"gate/graph/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UpstreamResult struct {
	Items []model.Item `json:"items"`
}

func Multiget(ids []int) ([]*model.Item, error) {
	reqest := "http://localhost:8082/find.php?a=1"
	for _, v := range ids {
		reqest += "&ids[]=" + strconv.Itoa(v)
	}
	res, err := http.Get(reqest)
	resBody, err := ioutil.ReadAll(res.Body)

	var resp = UpstreamResult{}
	json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	var result = []*model.Item{}

	for i := range resp.Items {
		result = append(result, &resp.Items[i])
	}
	fmt.Println(result)
	return result, nil
}
