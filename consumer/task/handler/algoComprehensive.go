package handler

import (
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AlgoComp struct {
}

type response struct {
	Id   string `json:"id"`
	Rank string `json:"rank"`
}

func postReq(id string) response {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8165/",
		strings.NewReader(fmt.Sprintf("GraphId=%s", id)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	var test response
	ha := string(body)
	err = json.Unmarshal([]byte(ha), &test)
	if err != nil {
		return response{}
	}
	defer resp.Body.Close()
	return test
}

func (h *AlgoComp) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.AlgoRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	id := req.GraphID
	ret := postReq(strconv.FormatInt(id, 10))
	ranks_ := strings.Split(ret.Rank, "\n")
	ranks_ = ranks_[0 : len(ranks_)-1]
	ranks := make([]*types.Rank, 0)
	for _, rank := range ranks_ {
		temp := strings.Split(rank, ",")
		score, _ := strconv.ParseFloat(temp[1], 64)
		nodeId, _ := strconv.ParseInt(temp[0], 10, 64)
		ranks = append(ranks, &types.Rank{
			NodeID: nodeId,
			Score:  score,
		})
	}
	resp := &types.AlgoRankReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Ranks: ranks,
		File:  ret.Id,
	}
	return jsonx.MarshalToString(resp)

}
