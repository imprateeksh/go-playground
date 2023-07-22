package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name      string   `json:"name"`
	Strengths []string `json:"skills"`
	Weakness  []string `json:"weakness"`
	Health    int      `json:"energy"`
}

func main() {
	sFrom := `{"name":"Paul", "skills":["Jaw Breaker", "FootLaunch", "PushAway"], "weakness":["cake", "music"], "energy":100}`
	// Attempt UnMarshalling i.e. JSON representation into Go data structure
	fmt.Println("##################\nJson to Map (Un-Marshalling)\n###################")
	var chkoutput Player
	err := json.Unmarshal([]byte(sFrom), &chkoutput)
	if err != nil {
		fmt.Println("UnMarshalling failed...")
		return
	}
	fmt.Println(chkoutput)

	// Attempt Marshalling i.e. Go data structure into JSON representation.
	fmt.Println("##################\nMap to Json (Marshalling)\n###################")
	outputJson, err := json.MarshalIndent(chkoutput, "", "  ")
	if err != nil {
		fmt.Println("Marshalling failed")
		return
	}
	fmt.Println(string(outputJson))

}
