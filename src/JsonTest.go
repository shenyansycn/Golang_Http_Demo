package main
import (
	"encoding/json"
	"fmt"
)
type AppFeedbackRespone struct {
	State   string `json:"state"`
	Message string `json:"message"`
}
func main() {
	var respone AppFeedbackRespone
	respone.Message = "not support Get Request"
	respone.State = "1"
	b, err := json.Marshal(respone)
	if (err != nil) {
		fmt.Println("json err = ", err)
	} else {
		fmt.Println(string(b))
	}
}