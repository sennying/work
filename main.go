package main

import (
	"engineering1/control"
	"fmt"
	"net/http"
)

func mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//中间件的逻辑
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set(`Access-Control-Allow-Origin`, `*`)
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/sellerimg/", http.StripPrefix("/sellerimg/", http.FileServer(http.Dir("sellerimg/"))))

	//视图
	//首页
	http.HandleFunc(`/`, control.IndexView)
	http.HandleFunc(`/index`, control.IndexView)

	//商家登录注册
	http.HandleFunc(`/joinin`, control.JoinView)
	//骑手界面
	http.HandleFunc(`/rider`, control.RiderView)
	//用户登录界面
	http.HandleFunc(`/userlogin`, control.UserloginView)
	//用户注册界面
	http.HandleFunc(`/usercreat`, control.AdduserView)
	//商家登录界面
	http.HandleFunc(`/sellerlogin`, control.SellerloginView)
	//商家注册界面
	http.HandleFunc(`/sellercreat`, control.AddsellerView)
	//骑手公寓接单界面
	http.HandleFunc(`/apartment`, control.ApartmentView)
	//骑手户外接单界面
	http.HandleFunc(`/outdoor`, control.OutdoorView)
	//点外卖主界面
	http.HandleFunc(`/waimai`, control.WaimaiView)
	//商家菜单界面
	http.HandleFunc(`/waimaimenu`, control.WaimaimenuView)
	//商家管理界面
	http.HandleFunc(`/sellerManage`, control.SellerManageView)

	//api
	//用户api
	//添加用户
	http.Handle(`/api/user/creat`, mid(http.HandlerFunc(control.Useradd)))
	//电话查询用户
	http.Handle(`/api/user/Phone`, mid(http.HandlerFunc(control.UserPhone)))

	//商家api
	//查询所有商家
	http.Handle(`/api/seller/all`, mid(http.HandlerFunc(control.SellerAll)))
	//地址查询商家
	http.Handle(`/api/seller/adress`, mid(http.HandlerFunc(control.SellerAdress)))
	//id查询商家
	http.Handle(`/api/seller/id`, mid(http.HandlerFunc(control.Sellerid)))
	//电话查询商家
	http.Handle(`/api/seller/Phone`, mid(http.HandlerFunc(control.SellerPhone)))
	//添加商家
	http.Handle(`/api/seller/creat`, mid(http.HandlerFunc(control.Selleradd)))

	//
	http.Handle(`/api/rider/apartment`, mid(http.HandlerFunc(control.Apartment)))
	http.Handle(`/api/rider/outdoor`, mid(http.HandlerFunc(control.Outdoor)))

	//商品api
	//删除商品
	http.Handle(`/api/menu/del`, mid(http.HandlerFunc(control.MenuDel)))
	//属于商家查询商品
	http.Handle(`/api/menu/Belongid`, mid(http.HandlerFunc(control.MenuBelongid)))

	fmt.Println("hello")
	http.ListenAndServe(":8080", nil)

}
