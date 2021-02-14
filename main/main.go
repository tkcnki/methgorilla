package main

import (
	"tkcnki/methgollira/main/model/entity"
	"tkcnki/methgollira/main/router"
)

func main() {
	initUser()
	router.WebRouting()
}

func initUser() {
	entity.Users = []*entity.User{
		&entity.User{
			Id:   "100",
			Name: "山本　リンダ",
			Kana: "ヤマモト　リンダ",
		},
		&entity.User{
			Id:   "200",
			Name: "山本　小リンダ",
			Kana: "ヤマモト　コリンダ",
		},
	}
}
