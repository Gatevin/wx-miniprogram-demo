package main
import (
    "fmt"
    "net/http"
    "log"
)

func giveQuestion(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // 解析参数
    var question_id = r.Form["question_id"]
    if question_id != nil {
        fmt.Fprintf(w, "Received question_id is %s", question_id);
    } else {
        fmt.Fprintf(w, "Error on question_id");
    }
}

func main() {
    http.HandleFunc("/", giveQuestion)
    err := http.ListenAndServe(":23333", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
