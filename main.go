package main

import (
	"fmt"
	"log"
	"net/http"
)

// func MiddleWare(fra http.Handler) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("yesss")

// 		fra.ServeHTTP(w, r)
// 	}
// }

// func MiddleWareV2(fra http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("yesss")

// 		fra.ServeHTTP(w, r)
// 	}
// }

// nije nuzno ni da se zove interface jer i onako pozivam funkciju serverHTTP
// func MiddleWare34(fra http.HandlerFunc) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("my Frined")

// 		fra.ServeHTTP(w, r)
// 	})
// }

func MiddleWareLog(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("log")

	next(w, r)
}

func MiddleWareAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("auth")

	next(w, r)
}

type MiddleWareFn func(http.ResponseWriter, *http.Request, http.HandlerFunc)

func MiddleWareChainV2(finalHandler http.HandlerFunc, middlewareFns ...MiddleWareFn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler := finalHandler

		for i := len(middlewareFns) - 1; i >= 0; i-- {
			currentHandler := handler // Čuvamo trenutni handler
			handler = func(w http.ResponseWriter, r *http.Request) {
				middlewareFns[i](w, r, currentHandler) // Pozivamo trenutni middleware
			}
		}

		handler(w, r) // Na kraju pozivamo finalHandler
	}
}

// func MiddleWareChain(finalHandler http.Handler, middlewareFns ...func(http.Handler) http.Handler) http.Handler {
// 	for i := len(middlewareFns) - 1; i >= 0; i-- {
// 		finalHandler = middlewareFns[i](finalHandler)
// 	}
// 	return finalHandler
// }

func main() {
	// f := func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("yessss is it top")
	// }

	r := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("yessss is it top23s")

		fmt.Println("zasto se ovo izvrsi????")
	}

	mux := http.NewServeMux()

	mux.Handle("/more", MiddleWareChainV2(r, MiddleWareAuth, MiddleWareLog))

	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Println(err)
	}
}
