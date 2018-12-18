package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	fmt.Println("hello")
	DB, err = sqlx.Open(`mysql`, `lkyz:18281578834lk!@tcp(45.40.203.75:3306)/messoutbuy?charset=utf8&parseTime=true`)
	if err != nil {
		panic("连接错误")
	} else {
		fmt.Println("数据库连接成功")
	}
	if err = DB.Ping(); err != nil {
		panic("运行错误")

	}
}

type Seller struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Telphone string `json:"telphone" form:"telphone"`
	Password string `json:"password" form:"password"`
	Logo     string `json:"logo" form:"logo"`
	Adress   string `json:"adress" form:"adress"`
}

type Menu struct {
	Id       int     `json:"id" form:"id"`
	Sellname string  `json:"sellname" form:"sellname"`
	Belongid int     `json:"belongid" form:"belongid"`
	Price    float32 `json:"price" form:"price"`
	Logo     string  `json:"logo" form:"logo"`
}
type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Callnum  string `json:"callnum" form:"callnum"`
	Callpass string `json:"callpass" form:"callpass"`
	Age      int    `json:"age" form:"age"`
	Logo     string `json:"logo" form:"logo"`
}

type Carte struct {
	Id      int    `json:"id" form:"id"`
	Callnum string `json:"callnum" form:"callnum"`
	Adress  string `json:"adress" form:"adress"`
	Flag    int    `json:"flag" form:"flag"`
	Rp      int    `json:"rp" form:"rp"`
}

type Minedan struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Ctime  string `json:"ctime" form:"ctime"`
	Food   string `json:"food" form:"food"`
	Adress string `json:"adress" form:"adress"`
}

func SellerAll() ([]Seller, error) {
	mods := make([]Seller, 0)
	err := DB.Select(&mods, "SELECT * FROM `seller`")
	return mods, err
}
func SellerAdress(adress string) ([]Seller, error) {
	mods := make([]Seller, 0)
	err := DB.Select(&mods, "SELECT * FROM `seller` WHERE adress=?", adress)
	fmt.Println(err)
	return mods, err
}
func SellerId(id int) (Seller, error) {
	mod := Seller{}
	err := DB.Get(&mod, "SELECT * FROM `seller` WHERE id=?", id)
	return mod, err
}

func MenuBelongid(belongid int) ([]Menu, error) {
	mods := make([]Menu, 0)
	err := DB.Select(&mods, "SELECT * FROM `menu` WHERE belongid=?", belongid)
	return mods, err
}
func MenuDel(id int) bool {
	fmt.Println(id + 1)
	result, err := DB.Exec("DELETE FROM `menu` WHERE id=?", id)
	fmt.Println(err)
	affect, _ := result.RowsAffected()
	if affect == 1 {
		return true
	}
	return false
}

//zz
func UserPhone(callnum string) (User, error) {
	mods := User{}
	err := DB.Get(&mods, "SELECT * FROM user where callnum=?", callnum)
	return mods, err
}
func SellerPhone(Phone string) (Seller, error) {
	fmt.Println("kj")
	mods := Seller{}
	fmt.Println(Phone)
	err := DB.Get(&mods, "SELECT * FROM seller where telphone=?", Phone)
	return mods, err
}
func UserAdd(name, callnum, callpass, logo string, age int) bool {
	DB.Exec("INSERT INTO `user` (`name`,`callnum`,`callpass`,`logo`,`age`) VALUES (?,?,?,?,?)", name, callnum, callpass, logo, age)
	return true
}
func SellerAdd(name, telphone, password, adress string) bool {
	DB.Exec("INSERT INTO `seller` (`name`,`telphone`,`password`,`adress`) VALUES (?,?,?,?)", name, telphone, password, adress)
	return true
}

//1127xzz
func Apartment() ([]Carte, error) {
	mods := make([]Carte, 0)
	err := DB.Select(&mods, "SELECT * FROM `carte` where flag=1")
	return mods, err
}

func Outdoor() ([]Carte, error) {
	mods := make([]Carte, 0)
	err := DB.Select(&mods, "SELECT * FROM `carte` where flag=2")
	return mods, err
}

//1217
func Usermine(name string) ([]Minedan, error) {
	fmt.Println(name)
	mods := make([]Minedan, 0)
	err := DB.Select(&mods, "SELECT * FROM `minedan` where name=?", name)
	fmt.Println(err)
	return mods, err
}
