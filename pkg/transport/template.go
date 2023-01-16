package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Templates struct {
	Main            *template.Template
	CteateOrder     *template.Template
	OrderStatus     *template.Template
	MakeOrderChange *template.Template
	Parts           *template.Template
	Works           *template.Template
}

func NewTemplates() Templates {
	var t Templates
	tpl, err := template.ParseFiles("web/html/header.html", "web/html/index.html", "web/html/footer.html")
	t.Main = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть главную страничку")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/createPage.html", "web/html/footer.html")
	t.CteateOrder = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку создания заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/userStatus.html", "web/html/footer.html")
	t.OrderStatus = tpl
	if err != nil {
		fmt.Println(err, "Не удалось открыть страницу состояния заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/makeChangesOrder.html", "web/html/footer.html")
	t.MakeOrderChange = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку изменения заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/parts.html", "web/html/footer.html")
	t.Parts = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку с запчастями")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/works.html", "web/html/footer.html")
	t.Works = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку с работами")
	}

	return t
}
func (t *Templates) indexPage(w http.ResponseWriter, r *http.Request) {
	t.Main.ExecuteTemplate(w, "index", nil)
}
func (t *Templates) newOrderPage(w http.ResponseWriter, r *http.Request) {
	t.CteateOrder.ExecuteTemplate(w, "createPage", nil)
}
func (t *Templates) statusOrderPage(w http.ResponseWriter, r *http.Request) {
	t.OrderStatus.ExecuteTemplate(w, "userStatus", nil)
}
func (t *Templates) makeChangesOrder(w http.ResponseWriter, r *http.Request) {
	t.MakeOrderChange.ExecuteTemplate(w, "makeChangesOrder", nil)
}
func (t *Templates) parts(w http.ResponseWriter, r *http.Request) {
	t.Parts.ExecuteTemplate(w, "parts", nil)
}
func (t *Templates) works(w http.ResponseWriter, r *http.Request) {
	t.Works.ExecuteTemplate(w, "works", nil)
}
