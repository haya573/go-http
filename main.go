package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// http.ResponseWriter: クライアントにデータを返すためのオブジェクト。
// *http.Request: クライアントからのリクエスト情報。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// http.ListenAndServe(":8080", nil): ポート 8080 でサーバーを起動します。nil はデフォルトのマルチプレクサを使用することを意味します。
// fmt.Fprintf(w, "Hello, World!"): クライアントに Hello, World! をレスポンスとして送信します。
func main() {
	// 新しいルーターを作成
	r := mux.NewRouter()

	// ルーティングを設定
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	// http.HandleFunc("/", homeHandler)     // ホームページ
	// http.HandleFunc("/about", aboutHandler) // Aboutページ
	// http.HandleFunc("/hello", helloHandler) // "/hello" のハンドラ
	// 静的ファイルを提供
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	fmt.Println("サーバーをポート8080で起動中...")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page.")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // クエリパラメータ "name" を取得
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}