package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"            // ebitenパッケージをインポート
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" // ebitenutilパッケージをインポート
	"github.com/hajimehoshi/ebiten/v2/text"       // textパッケージをインポート
	"golang.org/x/image/font/basicfont"           // basicfontパッケージをインポート
)

// ゲーム画面やボール、パドルのサイズ、パドルの移動速度を設定
const (
	screenWidth, screenHeight = 640, 480
	paddleWidth, paddleHeight = 20, 80
	ballSize                  = 20
	paddleSpeed               = 4
)

// Game型の定義
type Game struct {
	ballPositionX, ballPositionY float64 // ボールの位置
	ballDX, ballDY               float64 // ボールの移動速度
	paddle1Y, paddle2Y           float64 // パドルの位置
	player1Score, player2Score   int     // プレイヤーのスコア
}

// Updateはゲームの状態を更新するためのメソッド
func (g *Game) Update() error {
	g.ballPositionX += g.ballDX // ボールのx座標を更新
	g.ballPositionY += g.ballDY // ボールのy座標を更新

	if g.ballPositionX < 0 { // ボールが左の壁に当たった時
		g.ballDX = -g.ballDX // ボールの向きを反転
		g.ballPositionX = 0  // ボールのx座標を0に設定
		g.player2Score++     // プレイヤー2のスコアを加算
		// ボールを中央に戻す
		g.ballPositionX = screenWidth / 2
		g.ballPositionY = screenHeight / 2
	} else if g.ballPositionX > screenWidth-ballSize { // ボールが右の壁に当たった時
		g.ballDX = -g.ballDX                     // ボールの向きを反転
		g.ballPositionX = screenWidth - ballSize // ボールのx座標を右端に設定
		g.player1Score++                         // プレイヤー1のスコアを加算
		// ボールを中央に戻す
		g.ballPositionX = screenWidth / 2
		g.ballPositionY = screenHeight / 2
	}
	if g.ballPositionY < 0 { // ボールが上の壁に当たった時
		g.ballDY = -g.ballDY // ボールの向きを反転
		g.ballPositionY = 0  // ボールのy座標を0に設定
	} else if g.ballPositionY > screenHeight-ballSize { // ボールが下の壁に当たった時
		g.ballDY = -g.ballDY                      // ボールの向きを反転
		g.ballPositionY = screenHeight - ballSize // ボールのy座標を下端に設定
	}

	// パドルの操作
	if ebiten.IsKeyPressed(ebiten.KeyW) { // Wキーが押された時、パドル1を上に動かす
		g.paddle1Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) { // Sキーが押された時、パドル1を下に動かす
		g.paddle1Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) { // Upキーが押された時、パドル2を上に動かす
		g.paddle2Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) { // Downキーが押された時、パドル2を下に動かす
		g.paddle2Y += paddleSpeed
	}

	// ボールとパドルが接触した時の処理
	if (g.ballPositionX < paddleWidth && g.ballPositionY+ballSize > g.paddle1Y && g.ballPositionY < g.paddle1Y+paddleHeight) ||
		(g.ballPositionX+ballSize > screenWidth-paddleWidth && g.ballPositionY+ballSize > g.paddle2Y && g.ballPositionY < g.paddle2Y+paddleHeight) {
		g.ballDX = -g.ballDX // ボールの向きを反転
	}

	return nil
}

// Drawはゲームの描画を行うメソッド
func (g *Game) Draw(screen *ebiten.Image) {
	// パドルとボールの描画
	ebitenutil.DrawRect(screen, 10, g.paddle1Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, screenWidth-30, g.paddle2Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, g.ballPositionX, g.ballPositionY, ballSize, ballSize, color.White)

	// スコアの描画
	score := fmt.Sprintf("Player 1: %d - Player 2: %d", g.player1Score, g.player2Score)
	text.Draw(screen, score, basicfont.Face7x13, screenWidth/2-50, 30, color.White)
}

// Layoutはウィンドウのサイズを設定するメソッド
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// main関数
func main() {
	g := &Game{
		ballPositionX: screenWidth / 2,                 // ボールの初期x座標を設定
		ballPositionY: screenHeight / 2,                // ボールの初期y座標を設定
		ballDX:        2,                               // ボールの初期x方向の速度を設定
		ballDY:        2,                               // ボールの初期y方向の速度を設定
		paddle1Y:      screenHeight/2 - paddleHeight/2, // パドル1の初期位置を設定
		paddle2Y:      screenHeight/2 - paddleHeight/2, // パドル2の初期位置を設定
	}

	ebiten.SetWindowSize(screenWidth, screenHeight) // ウィンドウのサイズを設定
	ebiten.SetWindowTitle("Pong ")                  // ウィンドウのタイトルを設定

	// ゲームの実行
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err) // エラーが発生した場合、プログラムを終了
	}
}
