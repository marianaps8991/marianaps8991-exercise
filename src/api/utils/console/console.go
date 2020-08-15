package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func DisplayData(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
