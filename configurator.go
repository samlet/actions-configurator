package configurator

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type server struct {
	Name  string   `json:"name"`
	Host  string   `json:"host"`
	Port  int      `json:"port"`
	Items []string `json:"items"`
}

type rootConf struct {
	Host       server   `json:"host"`
	Datastores []server `json:"datastore"`
	Actions    []server `json:"actions"`
}

type manager struct {
	state string
	Root  rootConf
}

var singleton *manager
var once sync.Once

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetConfigurator() *manager {
	once.Do(func() {
		root := rootConf{}
		dat, err := ioutil.ReadFile("./actions.json")
		check(err)
		err = json.Unmarshal(dat, &root)
		check(err)
		singleton = &manager{state: "off", Root: root}
	})
	return singleton
}
func (sm *manager) GetState() string {
	return sm.state
}
func (sm *manager) SetState(s string) {
	sm.state = s
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
func (sm *manager) GetHostForAction(action string) *server {
	for _, v := range sm.Root.Actions {
		// fmt.Println(k, v["host"].(string), v["port"].(int))
		if Contains(v.Items, action) {
			return &v
		}
	}
	return nil
}
func (sm *manager) GetDefaultServer() *server {
	return &sm.Root.Actions[0]
}
