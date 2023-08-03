package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// 画面の幅と高さ
	screenWidth, screenHeight = 640, 480
	// パドルの幅と高さ
	paddleWidth, paddleHeight = 20, 80
	// ボールのサイズ
	ballSize = 20
	// パドルの速度
	paddleSpeed = 4
)

// ゲームの状態を管理する構造体
type Game struct {
	ballPositionX, ballPositionY float64 // ボールの位置
	ballDX, ballDY               float64 // ボールの速度
	paddle1Y, paddle2Y           float64 // パドルの位置
}

// ゲームの状態を更新する関数
func (g *Game) Update() error {
	g.ballPositionX += g.ballDX
	g.ballPositionY += g.ballDY

	// ボールが左端または右端に到達した場合、ボールの方向を反転させる
	if g.ballPositionX < 0 {
		g.ballDX = -g.ballDX
		g.ballPositionX = 0
	} else if g.ballPositionX > screenWidth-ballSize {
		g.ballDX = -g.ballDX
		g.ballPositionX = screenWidth - ballSize
	}
	// ボールが上端または下端に到達した場合、ボールの方向を反転させる
	if g.ballPositionY < 0 {
		g.ballDY = -g.ballDY
		g.ballPositionY = 0
	} else if g.ballPositionY > screenHeight-ballSize {
		g.ballDY = -g.ballDY
		g.ballPositionY = screenHeight - ballSize
	}

	// プレイヤーの操作によりパドルを移動させる
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.paddle1Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.paddle1Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.paddle2Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.paddle2Y += paddleSpeed
	}

	// ボールとパドルが衝突した場合、ボールの方向を反転させる
	if (g.ballPositionX < paddleWidth && g.ballPositionY+ballSize > g.paddle1Y && g.ballPositionY < g.paddle1Y+paddleHeight) ||
		(g.ballPositionX+ballSize > screenWidth-paddleWidth && g.ballPositionY+ballSize > g.paddle2Y && g.ballPositionY < g.paddle2Y+paddleHeight) {
		g.ballDX = -g.ballDX
	}

	return nil
}

// 画面を描画する関数
func (g *Game) Draw(screen *ebiten.Image) {
	// パドルとボールを描画する
	ebitenutil.DrawRect(screen, 10, g.paddle1Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, screenWidth-30, g.paddle2Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, g.ballPositionX, g.ballPositionY, ballSize, ballSize, color.White)
}

// 画面のレイアウトを設定する関数
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// ゲームの初期状態を設定
	g := &Game{
		ballPositionX: screenWidth / 2,
		ballPositionY: screenHeight / 2,
		ballDX:        2,
		ballDY:        2,
		paddle1Y:      screenHeight/2 - paddleHeight/2,
		paddle2Y:      screenHeight/2 - paddleHeight/2,
	}

	// ウィンドウのサイズとタイトルを設定
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong")

	// ゲームを実行する
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
