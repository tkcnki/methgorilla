package router

import (
	"fmt"
	"net/http"
	"tkcnki/methgollira/main/usecase"

	"github.com/gorilla/mux"
)

// WebRouting  is rooting
func WebRouting() error {
	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc({ エンドポイント }, { レスポンス関数 }).Methods({ リクエストメソッド（複数可能） })
	router.HandleFunc("/rinda", usecase.HelloRinda)
	router.HandleFunc("/rinda/{id}", usecase.GetSingleUser).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
