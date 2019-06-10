package main
import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "strconv"
)

// 问题类型
type Question struct {
    Description string `json:"description"`
    ChoiceList []string `json:"choice_list"`
}

// 问题提供的接口类型
type QuestionProvider struct {
    QuestionList []Question `json:"question_list"`
}

func (qp *QuestionProvider) InitFromFile (path string) {
    question_list := `{"question_list":[{"description":"abcde?","choice_list":["A. g","B. f","C. k","D. z"]},
                      {"description":"1 + 1 = ?","choice_list":["A. 2","B. 3","C. 4","D. 5"]},
                      {"description":"1 + 2 = ?","choice_list":["A. -99","B. jkl","C. e","D. 3"]}
                      ]}`
    _ = json.Unmarshal([]byte(question_list), &qp)
    return
}

// 实现http.Handler的接口
func (qp QuestionProvider) ServeHTTP (w http.ResponseWriter, r *http.Request){
    r.ParseForm() // 解析参数
    var question_id = r.Form["question_id"]
    if question_id != nil {
        // fmt.Fprintf(w, "Received question_id is %s\n", question_id)
        qid, err := strconv.Atoi(question_id[0])
        // fmt.Fprintf(w, "Total questions number: %d\n", len(qp.QuestionList))
        if err == nil && qid < len(qp.QuestionList) {
            fmt.Fprintf(w, "Question description: %s\n", qp.QuestionList[qid].Description)
            for _, choice := range(qp.QuestionList[qid].ChoiceList) {
                fmt.Fprintf(w, "%s\n", choice)
            }
        } else {
            fmt.Fprintf(w, "Wrong question_id")
        }
    } else {
        fmt.Fprintf(w, "Error on question_id")
    }
}

func main() {
    qp := &QuestionProvider{}
    qp.InitFromFile("path-to-read")
    err := http.ListenAndServe(":23333", qp)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
