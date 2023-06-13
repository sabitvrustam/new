package http

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Templates struct {
	Main            *template.Template
	CteateOrder     *template.Template
	OrderStatus     *template.Template
	MakeOrderChange *template.Template
	CreateInginer   *template.Template
	Parts           *template.Template
	Works           *template.Template
	Orders          *template.Template
	Employees       *template.Template
	log             *logrus.Logger
}

func NewTemplates(log *logrus.Logger) (t Templates) {

	tpl, err := template.ParseFiles("web/html/header.html", "web/html/index.html", "web/html/footer.html")
	t.Main = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть главную страничку")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/create_order.html", "web/html/footer.html")
	t.CteateOrder = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку создания заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/status_order.html", "web/html/footer.html")
	t.OrderStatus = tpl
	if err != nil {
		t.log.Error(err, "Не удалось открыть страницу состояния заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/orders.html", "web/html/footer.html")
	t.Orders = tpl
	if err != nil {
		t.log.Error(err, "Не удалось открыть страницу состояния заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/change_order.html", "web/html/footer.html")
	t.MakeOrderChange = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку изменения заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/inginer.html", "web/html/footer.html")
	t.CreateInginer = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку изменения заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/parts.html", "web/html/footer.html")
	t.Parts = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку с запчастями")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/works.html", "web/html/footer.html")
	t.Works = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку с работами")
	}

	tpl, err = template.ParseFiles("web/html/header.html", "web/html/employees.html", "web/html/footer.html")
	t.Employees = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть страничку с сотрудниками")
	}

	return t
}
func (t *Templates) indexPage(w http.ResponseWriter, r *http.Request) {
	err := t.Main.ExecuteTemplate(w, "index", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) newOrderPage(w http.ResponseWriter, r *http.Request) {
	err := t.CteateOrder.ExecuteTemplate(w, "createOrder", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) statusOrderPage(w http.ResponseWriter, r *http.Request) {
	err := t.OrderStatus.ExecuteTemplate(w, "userStatus", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) ordersPage(w http.ResponseWriter, r *http.Request) {
	err := t.Orders.ExecuteTemplate(w, "orders", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) makeChangesOrder(w http.ResponseWriter, r *http.Request) {
	err := t.MakeOrderChange.ExecuteTemplate(w, "makeChangesOrder", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) inginerPage(w http.ResponseWriter, r *http.Request) {
	err := t.CreateInginer.ExecuteTemplate(w, "createInginer", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) parts(w http.ResponseWriter, r *http.Request) {
	err := t.Parts.ExecuteTemplate(w, "parts", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) works(w http.ResponseWriter, r *http.Request) {
	err := t.Works.ExecuteTemplate(w, "works", nil)
	if err != nil {
		t.log.Error(err)
	}
}

func (t *Templates) employeesPage(w http.ResponseWriter, r *http.Request) {
	err := t.Employees.ExecuteTemplate(w, "employees", nil)
	if err != nil {
		t.log.Error(err)
	}
}
