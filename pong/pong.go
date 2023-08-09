package main

// 必要なパッケージをインポート
import (
	"fmt"                              // フォーマット関連のユーティリティ
	"image/color"                      // 色関連のユーティリティ
	"log"                             // ロギング関連のユーティリティ

	"github.com/hajimehoshi/ebiten/v2" // ゲーム開発ライブラリ
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" // ゲーム開発のユーティリティ
	"github.com/hajimehoshi/ebiten/v2/text" // テキスト描画のユーティリティ
	"golang.org/x/image/font/basicfont" // ベーシックなフォントのユーティリティ
)

// ゲームの定数を設定
const (
	screenWidth, screenHeight = 640, 480 // 画面の大きさ
	paddleWidth, paddleHeight = 20, 80   // パドルの大きさ
	ballSize                  = 20       // ボールの大きさ
	paddleSpeed               = 4        // パドルの速度
	maxScore                  = 5        // 勝利に必要なスコア
)

// ゲームの状態を示す型を定義
type GameState int

const (
	Starting GameState = iota // ゲーム開始前の状態
	Playing                   // ゲームプレイ中の状態
	Resetting                 // スコア後のリセット状態
	Winner                    // 勝者が決定した状態
)

// ゲームのデータを保持する構造体を定義
type Game struct {
	ballPositionX, ballPositionY float64  // ボールの位置
	ballDX, ballDY               float64  // ボールの移動量
	paddle1Y, paddle2Y           float64  // 2つのパドルの位置
	player1Score, player2Score   int      // 2プレイヤーのスコア

	state            GameState   // 現在のゲーム状態
	resetTicker      int         // リセット用のカウンタ
	textBlinkTicker  int         // スタートメッセージの点滅用カウンタ
	showStartMessage bool        // スタートメッセージを表示するかのフラグ
}

// ゲームのロジックを更新する関数
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		// Escキーが押されたらゲームをリセット
		g.state = Starting
		g.player1Score = 0
		g.player2Score = 0
		g.ballPositionX = screenWidth / 2
		g.ballPositionY = screenHeight / 2
		return nil
	}

	switch g.state {
	case Starting:
		// ゲーム開始前の処理
		g.textBlinkTicker++
		if g.textBlinkTicker >= 30 {
			// 30フレームごとにメッセージを点滅
			g.showStartMessage = !g.showStartMessage
			g.textBlinkTicker = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			// Enterキーでゲーム開始
			g.state = Playing
		}
	case Playing:
		// ゲームプレイ中の処理
		g.ballPositionX += g.ballDX
		g.ballPositionY += g.ballDY

		// 左壁にボールが当たった場合の処理
		if g.ballPositionX < 0 {
			g.ballDX = -g.ballDX
			g.ballPositionX = 0
			g.player2Score++
			if g.player2Score >= maxScore {
				// プレイヤー2が勝利
				g.state = Winner
			} else {
				// ボールを中央にリセット
				g.state = Resetting
			}
		} else if g.ballPositionX > screenWidth-ballSize {
			// 右壁にボールが当たった場合の処理
			g.ballDX = -g.ballDX
			g.ballPositionX = screenWidth - ballSize
			g.player1Score++
			if g.player1Score >= maxScore {
				// プレイヤー1が勝利
				g.state = Winner
			} else {
				// ボールを中央にリセット
				g.state = Resetting
			}
		}

		// 上壁または下壁にボールが当たった場合の処理
		if g.ballPositionY < 0 {
			g.ballDY = -g.ballDY
			g.ballPositionY = 0
		} else if g.ballPositionY > screenHeight-ballSize {
			g.ballDY = -g.ballDY
			g.ballPositionY = screenHeight - ballSize
		}

		// パドルの操作
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

		// ボールがパドル1に当たった場合の処理
		if g.ballPositionX < paddleWidth && g.ballPositionY+ballSize > g.paddle1Y && g.ballPositionY < g.paddle1Y+paddleHeight {
			g.ballDX = -g.ballDX
			relPos := g.ballPositionY - g.paddle1Y
			g.ballDY = (relPos/paddleHeight*2 - 1) * 5
		}

		// ボールがパドル2に当たった場合の処理
		if g.ballPositionX+ballSize > screenWidth-paddleWidth && g.ballPositionY+ballSize > g.paddle2Y && g.ballPositionY < g.paddle2Y+paddleHeight {
			g.ballDX = -g.ballDX
			relPos := g.ballPositionY - g.paddle2Y
			g.ballDY = (relPos/paddleHeight*2 - 1) * 5
		}
	case Resetting:
		// ボールのリセット処理
		g.resetTicker++
		if g.resetTicker > 60 {
			g.state = Playing
			g.resetTicker = 0
			g.ballPositionX = screenWidth / 2
			g.ballPositionY = screenHeight / 2
		}
	case Winner:
		// 勝者が決定した場合の処理
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.state = Starting
			g.player1Score = 0
			g.player2Score = 0
		}
	}
	return nil
}

// ゲームの描画を行う関数
func (g *Game) Draw(screen *ebiten.Image) {
	if g.state == Starting && g.showStartMessage {
		// ゲーム開始前のメッセージを描画
		startText := "PONG - Press Enter to Start"
		text.Draw(screen, startText, basicfont.Face7x13, (screenWidth-len(startText)*7)/2, screenHeight/2, color.White)
		return
	}

	if g.state == Winner {
		// 勝者のメッセージを描画
		var winText string
		if g.player1Score >= maxScore {
			winText = "Player 1 Wins!"
		} else {
			winText = "Player 2 Wins!"
		}
		text.Draw(screen, winText, basicfont.Face7x13, (screenWidth-len(winText)*7)/2, screenHeight/2, color.White)
		return
	}

	// ゲームの要素を描画
	ebitenutil.DrawRect(screen, 10, g.paddle1Y, paddleWidth, paddleHeight, color.White) // パドル1
	ebitenutil.DrawRect(screen, screenWidth-30, g.paddle2Y, paddleWidth, paddleHeight, color.White) // パドル2
	ebitenutil.DrawRect(screen, g.ballPositionX, g.ballPositionY, ballSize, ballSize, color.White) // ボール

	// スコアを描画
	player1ScoreStr := fmt.Sprintf("%d", g.player1Score)
	player2ScoreStr := fmt.Sprintf("%d", g.player2Score)
	text.Draw(screen, player1ScoreStr, basicfont.Face7x13, screenWidth/2-50, 30, color.White) // プレイヤー1のスコア
	text.Draw(screen, player2ScoreStr, basicfont.Face7x13, screenWidth/2+40, 30, color.White) // プレイヤー2のスコア

	// 画面中央の線を描画
	for i := 0; i < screenHeight; i += 30 {
		ebitenutil.DrawRect(screen, screenWidth/2-1, float64(i), 2, 20, color.White)
	}
}

// ゲームのレイアウトを設定する関数
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// ゲームのメイン関数
func main() {
	// ゲームの初期化
	g := &Game{
		ballPositionX:   screenWidth / 2,
		ballPositionY:   screenHeight / 2,
		ballDX:          3,
		ballDY:          2,
		paddle1Y:        (screenHeight - paddleHeight) / 2,
		paddle2Y:        (screenHeight - paddleHeight) / 2,
		state:           Starting,
		showStartMessage: true,
	}

	// ウィンドウの設定
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong")

	// ゲームの実行
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err) // エラーが発生した場合はログに出力して終了
	}
}
