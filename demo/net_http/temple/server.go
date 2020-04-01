package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//字符串模板
func stringHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := "My name is {{ .Name }}. I am {{ .Age }} years old.\n"
	t := template.New("tmpl.html")
	t = template.Must(t.Parse(tmpl))

	data := struct {
		Name string
		Age  int
	}{"golang", 14}

	err := t.Execute(w, data)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//点动作
func dotHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("dot.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	err = t.Execute(w, "Hello World")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//条件动作
func conditionHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("condition.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

//迭代动作
func iteratorHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("iterator.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	err := t.Execute(w, daysOfWeek)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//设置动作
func setHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("set.html"))
	err := t.Execute(w, "hello")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//包含动作
func includeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("include1.html", "include2.html"))
	err := t.Execute(w, "hello")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//嵌套模板/定义动作
func nestedHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("nested.html"))
	err := t.ExecuteTemplate(w, "layout", "world")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

//块动作
func blockHandler(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("block.html"))
	t := template.Must(template.ParseFiles("block.html", "content.html"))
	err := t.ExecuteTemplate(w, "layout", "world")
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

//自定义函数
func funcHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	t := template.New("func.html").Funcs(funcMap)

	t, err := t.ParseFiles("func.html")
	if err != nil {
		log.Fatal("Parse error:", err)
	}
	err = t.Execute(w, time.Now())
	if err != nil {
		log.Fatal("Exeute error:", err)
	}
}

//上下文感知
func contextHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("context.html"))
	err := t.Execute(w, `He saied: <i>"She's alone?"</i>`)
	if err != nil {
		log.Fatal("Exeute error:", err)
	}
}

//防御xss攻击
func xssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t := template.Must(template.ParseFiles("xss_display.html"))
		t.Execute(w, r.FormValue("comment"))                //会对内容进行转义
		t.Execute(w, template.HTML(r.FormValue("comment"))) //不对内容进行转义
	} else {
		t := template.Must(template.ParseFiles("xss_form.html"))
		t.Execute(w, nil)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/string", stringHandler)
	mux.HandleFunc("/dot", dotHandler)
	mux.HandleFunc("/condition", conditionHandler)
	mux.HandleFunc("/iterator", iteratorHandler)
	mux.HandleFunc("/set", setHandler)
	mux.HandleFunc("/include", includeHandler)
	mux.HandleFunc("/nested", nestedHandler)
	mux.HandleFunc("/block", blockHandler)

	mux.HandleFunc("/func", funcHandler)
	mux.HandleFunc("/context", contextHandler)
	mux.HandleFunc("/xss", xssHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
