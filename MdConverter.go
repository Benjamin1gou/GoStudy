package main // mainパッケージを宣言します。プログラムのエントリーポイントとなります。

import ( // 必要なパッケージをインポートします。
	"fmt"      // 標準的なフォーマットI/Oを提供します（例えば、出力と文字列フォーマット）。
	"net/http" // HTTPクライアントとサーバの機能を提供します。
	"strings"  // 文字列の操作を提供します。

	"github.com/PuerkitoBio/goquery" // HTMLの解析と操作を行うライブラリです。
)

// handlerはHTTPリクエストを処理する関数です。
func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url") // URLクエリから"url"パラメータを取得します。
	if url == "" {                  // "url"パラメータが存在しない場合、エラーメッセージを返します。
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// goqueryを使って指定されたURLのドキュメントを取得します。
	doc, err := goquery.NewDocument(url)
	if err != nil { // エラーがあった場合、エラーメッセージを返します。
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTML部分を保存するための文字列スライスを定義します。
	var parts []string
	// HTMLドキュメント内の"div.part"要素を探し、それぞれの要素に対して以下の処理を行います。
	doc.Find("div.part").Each(func(i int, s *goquery.Selection) {
		// "var"要素を探し、それぞれの要素に対して以下の処理を行います。
		s.Find("var").Each(func(j int, v *goquery.Selection) {
			// "var"要素を"<span>$...$</span>"要素に置き換えます。
			v.ReplaceWithHtml(fmt.Sprintf("<span>$%s$</span>", v.Text()))
		})
		// 置き換えた結果のHTMLを取得します。
		part, _ := s.Html()
		// 取得したHTML部分をpartsスライスに追加します。
		parts = append(parts, part)
	})

	// partsスライスの各要素を結合して一つのHTML文字列を作成します。
	html := strings.Join(parts, "")

	// HTMLドキュメントを作成し、レスポンスとして返します。
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
			<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
		</head>
		<body>
			%s
		</body>
		</html>
	`, html)
}

// main関数はプログラムのエントリーポイントです。
func main() {
	// ルートパス("/")へのHTTPリクエストをhandler関数で処理するように設定します。
	http.HandleFunc("/", handler)
	// HTTPサーバを起動します。ポート8080をリッスンします。
	http.ListenAndServe(":8080", nil)
}
