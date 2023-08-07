### このゲームについて
このコードは、古典的な「PONG」ゲームの実装です。2人のプレイヤーがパドルを操作して、画面中央を移動するボールを打ち返すことを目的としています。

### 使われている主要なパッケージ
- `fmt`: 標準のフォーマットI/Oを提供します。
- `image/color`: 基本的な色を扱うためのパッケージです。
- `log`: ロギングのためのパッケージです。
- `github.com/hajimehoshi/ebiten/v2`: ゲーム作成のためのライブラリです。
- `github.com/hajimehoshi/ebiten/v2/ebitenutil`: ゲーム開発に便利なユーティリティを提供します。
- `github.com/hajimehoshi/ebiten/v2/text`: テキスト描画用のユーティリティです。
- `golang.org/x/image/font/basicfont`: 基本的なフォントを提供するパッケージです。

### ゲームの主な機能と挙動
1. ゲームの状態は「開始前(Starting)」、「進行中(Playing)」、「リセット中(Resetting)」の3つの状態を持ちます。
2. ゲームは開始前の状態から始まり、Enterキーを押すと進行中の状態に移行します。
3. ボールが左または右の壁に触れると、リセット中の状態に移行し、ボールの位置が中央に戻されます。
4. プレイヤーはW, Sキーで左のパドルを、上矢印, 下矢印キーで右のパドルを操作します。

[![](https://mermaid.ink/img/pako:eNqdU01PwkAQ_SvNnusf6IET3NSL8WJ6WemqRGmxtBpCPHQ3KCoGNYhRIQgaxc8QQRPx68eMLfAv7LZAYgJJ9fZ25u17M5OZNIpqCkESSpJVk6hREo7hRR3HZVVWBQFYFeglsGavWHOOGhOhENAnYO_AziVhCH0K0FegXz74_ig52X1PYUAa87ez8-Jkdu1s2SlV7FxxtGlkPmYQFdiNFy9ywG65yBXQC-7Itj38zPO0DVadV8K5WWB33fqD_Xj6R-WxrYF1Pczaj5VuNceVR-r8bnk2oWCD2AfvYDXtLbfI4-Dzcc5aTrER0Cis4_V_2uTzncJbQJtJnNJMI6CRK9sr15z7qp1h3a-Cvdfi5GnNIIK2RnRhZE3P9Lu9CdanO3J_eKLXm9i3BnrYeW0A3XEFwTpBIooTPY5jirvNaS4vI2OJxImMJBcqWF-WkaxuuDxsGtpMSo0iydBNIiLTU-9vPpIW8EpyGI0oMUPTh0HiPaf8m_FOR0QJrM5p2uDjxg8aidEt?type=png)](https://mermaid.live/edit#pako:eNqdU01PwkAQ_SvNnusf6IET3NSL8WJ6WemqRGmxtBpCPHQ3KCoGNYhRIQgaxc8QQRPx68eMLfAv7LZAYgJJ9fZ25u17M5OZNIpqCkESSpJVk6hREo7hRR3HZVVWBQFYFeglsGavWHOOGhOhENAnYO_AziVhCH0K0FegXz74_ig52X1PYUAa87ez8-Jkdu1s2SlV7FxxtGlkPmYQFdiNFy9ywG65yBXQC-7Itj38zPO0DVadV8K5WWB33fqD_Xj6R-WxrYF1Pczaj5VuNceVR-r8bnk2oWCD2AfvYDXtLbfI4-Dzcc5aTrER0Cis4_V_2uTzncJbQJtJnNJMI6CRK9sr15z7qp1h3a-Cvdfi5GnNIIK2RnRhZE3P9Lu9CdanO3J_eKLXm9i3BnrYeW0A3XEFwTpBIooTPY5jirvNaS4vI2OJxImMJBcqWF-WkaxuuDxsGtpMSo0iydBNIiLTU-9vPpIW8EpyGI0oMUPTh0HiPaf8m_FOR0QJrM5p2uDjxg8aidEt)

### ゲームの実行方法
1. 必要なパッケージをインストールします。
```bash
go get github.com/hajimehoshi/ebiten/v2
```
2. ゲームを実行します。
```bash
go run [ファイル名].go
```

### 注意点
- `ebiten`ライブラリはアクティブに開発が進められているため、最新のバージョンを使用する際は公式ドキュメントを確認してください。
