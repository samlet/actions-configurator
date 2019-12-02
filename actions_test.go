package configurator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestActionsConf(test *testing.T) {
	root := rootConf{}
	dat, err := ioutil.ReadFile("./actions.json")
	check(err)
	err = json.Unmarshal(dat, &root)
	check(err)
	fmt.Println(root.Host.Port)
	// fmt.Println(root.Datastore)

	for k, v := range root.Datastores {
		// fmt.Println(k, v["host"].(string), v["port"].(int))
		fmt.Println(k, v.Host, v.Port)
	}
	for k, v := range root.Actions {
		// fmt.Println(k, v["host"].(string), v["port"].(int))
		fmt.Println(k, v.Name, v.Host, v.Port, v.Items)
	}
}

type tracker struct {
}
type domain struct {
}
type actionReq struct {
	NextAction string  `json:"next_action"`
	Sender     string  `json:"sender_id"`
	Tracker    tracker `json:"tracker"`
	Domain     domain  `json:"domain"`
}

func TestActionsReq(test *testing.T) {
	actionReq := actionReq{}
	dat, err := ioutil.ReadFile("./action_req.json")
	check(err)
	err = json.Unmarshal(dat, &actionReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(actionReq.NextAction)
}

type eventCnt struct {
	Event string `json:"event"`
	Name  string `json:"name"`
	// Value map[string]interface{} `json:"value"`
}
type responseCnt struct {
}
type actionResp struct {
	Events    []eventCnt    `json:"events"`
	Responses []responseCnt `json:"responses"`
}

func TestActionsResp(test *testing.T) {
	actionResp := actionResp{}
	dat, err := ioutil.ReadFile("./action_resp.json")
	check(err)
	err = json.Unmarshal(dat, &actionResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(actionResp.Events[0])
}
