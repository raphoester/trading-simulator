package helper

import (
	"encoding/json"
	"fmt"
)

func ConsoleJson(object interface{}) {
	res, _ := json.MarshalIndent(object, "", "    ")
	fmt.Println(string(res))
}
