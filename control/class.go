package control

import (
	"encoding/json"
	"engineering1/model"
	"engineering1/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("index.html")
	w.Write(buf)
}
func JoinView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/joinin.html")
	w.Write(buf)
}

func RiderView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/rider.html")
	w.Write(buf)
}

func WaimaiView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/waimai.html")
	w.Write(buf)
}
func WaimaimenuView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/waimaimenu.html")
	w.Write(buf)
}
func SellerAll(w http.ResponseWriter, r *http.Request) {
	mods, _ := model.SellerAll()
	buf := util.NewResult(200, "商家", mods)

	w.Write(buf)
}

func SellerAdress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adress := r.FormValue("adress")
	mods, _ := model.SellerAdress(adress)
	buf := util.NewResult(200, "查询成功", mods)
	w.Write(buf)
}

func Sellerid(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	mod, _ := model.SellerId(id)
	buf := util.NewResult(200, "查询成功", mod)
	w.Write(buf)
}

func MenuBelongid(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Belongid, _ := strconv.Atoi(r.FormValue("Belongid"))
	mods, _ := model.MenuBelongid(Belongid)
	buf := util.NewResult(200, "查询成功", mods)
	w.Write(buf)
}

func MenuDel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("kid"))
	fmt.Println(id)
	if err != nil {
		// w.Write([]byte(`删除数据错误,请重试`))
		// return
		buf := util.NewResult(300, "删除的数据有错误，请重试")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
		return
	}
	ok := model.MenuDel(id)
	// fmt.Println(ok)
	if ok {
		buf := util.NewResult(200, "删除成功")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
		// w.Write([]byte(`删除成功`))
	} else {
		// w.Write([]byte(`删除失败`))
		buf := util.NewResult(300, "删除失败")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	}
}

//xzz
func AdduserView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/usercreat.html")
	w.Write(buf)
}

func UserloginView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/userlogin.html")
	w.Write(buf)
}
func SellerloginView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/sellerlogin.html")
	w.Write(buf)
}
func AddsellerView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/sellercreat.html")
	w.Write(buf)
}
func SellerManageView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/sellerManage.html")
	w.Write(buf)
}
func Useradd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.Form.Get("kname")
	if name == "" {
		w.Write([]byte(`名字不能为空`))
		return
	}
	callnum := r.FormValue("kcallnum")
	callpass := r.FormValue("kcallpass")
	logo := r.FormValue("klogo")
	age, _ := strconv.Atoi(r.FormValue("kage"))

	ok := model.UserAdd(name, callnum, callpass, logo, age)

	if ok {
		w.Write([]byte(`添加成功`))
	} else {
		w.Write([]byte(`添加失败`))
	}

}
func Selleradd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.Form.Get("kname")
	if name == "" {
		w.Write([]byte(`名字不能为空`))
		return
	}
	telphone := r.FormValue("ktelphone")
	password := r.FormValue("kpassword")
	logo := r.FormValue("klogo")
	adress := r.FormValue("kadress")

	ok := model.SellerAdd(name, telphone, password, logo, adress)

	if ok {
		w.Write([]byte(`添加成功`))
	} else {
		w.Write([]byte(`添加失败`))
	}

}
func UserPhone(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	callnum := r.FormValue("callnum")
	mod, _ := model.UserPhone(callnum)
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func SellerPhone(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	callnum := r.FormValue("telphone")
	mod, _ := model.SellerPhone(callnum)
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

//1127xzz
func Apartment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mod, _ := model.Apartment()
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func Outdoor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mod, _ := model.Outdoor()
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
func ApartmentView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/apartment.html")
	w.Write(buf)
}
func OutdoorView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/outdoor.html")
	w.Write(buf)
}
