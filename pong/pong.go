package main

// 必要なパッケージをインポート
import (
	"fmt"         // 標準のフォーマットI/O
	"image/color" // 基本的な色のパッケージ
	"log"         // ロギングを行うパッケージ

	"github.com/hajimehoshi/ebiten/v2"            // ゲーム作成のためのライブラリ
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" // ゲーム開発のユーティリティ
	"github.com/hajimehoshi/ebiten/v2/text"       // テキスト描画用のユーティリティ
	"golang.org/x/image/font/basicfont"           // 基本的なフォントのパッケージ
)

// 画面やゲーム要素のサイズを定義
const (
	screenWidth, screenHeight = 640, 480 // 画面の幅と高さ
	paddleWidth, paddleHeight = 20, 80   // パドルの幅と高さ
	ballSize                  = 20       // ボールのサイズ
	paddleSpeed               = 4        // パドルの移動速度
)

// ゲームの状態を表す型を定義
type GameState int

const (
	Starting  GameState = iota // ゲームが開始されていない状態
	Playing                    // ゲームが進行中の状態
	Resetting                  // ボールがゴールした後、リセット中の状態
)

// ゲームのデータを管理する構造体
type Game struct {
	ballPositionX, ballPositionY float64 // ボールのX座標、Y座標
	ballDX, ballDY               float64 // ボールのX方向、Y方向の速度
	paddle1Y, paddle2Y           float64 // プレイヤー1とプレイヤー2のパドルのY座標
	player1Score, player2Score   int     // プレイヤー1とプレイヤー2のスコア

	state            GameState // ゲームの現在の状態
	resetTicker      int       // リセット時のタイマー
	textBlinkTicker  int       // スタートメッセージの点滅用タイマー
	showStartMessage bool      // スタートメッセージを表示するかのフラグ
}

// ゲームのロジックを更新するメソッド
func (g *Game) Update() error {
	switch g.state {
	case Starting:
		// ゲーム開始前の処理
		g.textBlinkTicker++
		if g.textBlinkTicker >= 30 {
			// スタートメッセージの点滅を切り替え
			g.showStartMessage = !g.showStartMessage
			g.textBlinkTicker = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			// Enterキーが押されたらゲームを開始
			g.state = Playing
			g.ballPositionX = screenWidth / 2
			g.ballPositionY = screenHeight / 2
		}
	case Playing:
		// ゲーム進行中の処理
		g.ballPositionX += g.ballDX
		g.ballPositionY += g.ballDY

		// ボールが左の壁に触れた場合
		if g.ballPositionX < 0 {
			g.ballDX = -g.ballDX
			g.ballPositionX = 0
			g.player2Score++
			g.state = Resetting
		} else if g.ballPositionX > screenWidth-ballSize {
			// ボールが右の壁に触れた場合
			g.ballDX = -g.ballDX
			g.ballPositionX = screenWidth - ballSize
			g.player1Score++
			g.state = Resetting
		}

		// ボールが上または下の壁に触れた場合
		if g.ballPositionY < 0 {
			g.ballDY = -g.ballDY
			g.ballPositionY = 0
		} else if g.ballPositionY > screenHeight-ballSize {
			g.ballDY = -g.ballDY
			g.ballPositionY = screenHeight - ballSize
		}

		// 各プレイヤーのパドル操作
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

		// ボールがプレイヤー1のパドルに当たる処理
		if g.ballPositionX < paddleWidth && g.ballPositionY+ballSize > g.paddle1Y && g.ballPositionY < g.paddle1Y+paddleHeight {
			g.ballDX = -g.ballDX
			relPos := g.ballPositionY - g.paddle1Y
			g.ballDY = (relPos/paddleHeight*2 - 1) * 5
		}

		// ボールがプレイヤー2のパドルに当たる処理
		if g.ballPositionX+ballSize > screenWidth-paddleWidth && g.ballPositionY+ballSize > g.paddle2Y && g.ballPositionY < g.paddle2Y+paddleHeight {
			g.ballDX = -g.ballDX
			relPos := g.ballPositionY - g.paddle2Y
			g.ballDY = (relPos/paddleHeight*2 - 1) * 5
		}
	case Resetting:
		// ボールがゴール後のリセット処理
		g.resetTicker++
		if g.resetTicker > 60 {
			// 60フレーム後にゲームを再開
			g.state = Playing
			g.resetTicker = 0
			g.ballPositionX = screenWidth / 2
			g.ballPositionY = screenHeight / 2
		}
	}
	return nil
}

// ゲームの描画を行うメソッド
func (g *Game) Draw(screen *ebiten.Image) {
	if g.state == Starting {
		if g.showStartMessage {
			startText := "PONG - Press Enter to Start"
			text.Draw(screen, startText, basicfont.Face7x13, (screenWidth-len(startText)*7)/2, screenHeight/2, color.White)
		}
		return
	}

	// パドルとボールの描画
	ebitenutil.DrawRect(screen, 10, g.paddle1Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, screenWidth-30, g.paddle2Y, paddleWidth, paddleHeight, color.White)
	ebitenutil.DrawRect(screen, g.ballPositionX, g.ballPositionY, ballSize, ballSize, color.White)

	// スコアの描画
	player1ScoreStr := fmt.Sprintf("%d", g.player1Score)
	player2ScoreStr := fmt.Sprintf("%d", g.player2Score)
	text.Draw(screen, player1ScoreStr, basicfont.Face7x13, screenWidth/2-50, 30, color.White)
	text.Draw(screen, player2ScoreStr, basicfont.Face7x13, screenWidth/2+40, 30, color.White)

	// 画面中央の線を描画
	for i := 0; i < screenHeight; i += 30 {
		ebitenutil.DrawRect(screen, screenWidth/2-1, float64(i), 2, 20, color.White)
	}
}

// ゲームのレイアウト設定を行うメソッド
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// 定義された画面の幅と高さを返す
	return screenWidth, screenHeight
}

// メイン関数
func main() {
	// ゲームの初期設定
	g := &Game{
		ballPositionX:    screenWidth / 2,
		ballPositionY:    screenHeight / 2,
		ballDX:           3,
		ballDY:           2,
		paddle1Y:         (screenHeight - paddleHeight) / 2,
		paddle2Y:         (screenHeight - paddleHeight) / 2,
		state:            Starting,
		showStartMessage: true,
	}

	// ウィンドウのサイズとタイトルを設定
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong")

	// ゲームの実行開始。エラーがあればログに出力。
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
