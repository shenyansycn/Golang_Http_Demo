package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type AppFeedbackRespone struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Method = ", r.Method)
	fmt.Println("Form = ", r.Form)
	fmt.Println("path = ", r.URL.Path)
	fmt.Println("scheme = " + r.URL.Scheme)
	fmt.Println("key1 = ", r.Form["key1"])
	fmt.Println("UserAgent = ", r.UserAgent())
	for k, v := range r.Form {
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}
	if r.Method == "GET" {
	} else {

	}
	var respone AppFeedbackRespone
	respone.Message = "not support Get Request"
	respone.State = "1"
	b, err := json.Marshal(respone)
	if (err != nil) {
		fmt.Println("json err = ", err)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func appFeedback(w http.ResponseWriter, r *http.Request) {
	go process(w, r)

}
func main() {
	fmt.Println("Start")
	http.HandleFunc("/appFeedback", appFeedback)
	fmt.Println("Linstenering")
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		fmt.Println("ListenerAndServer: ", err)
	} else {
		fmt.Println("start server success")
	}
	fmt.Println("over")
}
