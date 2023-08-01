package main

import (
	"image/color" //色を操作するためのパッケージ
	"log"         //ログ出力を扱うためのパッケージ
	"math/rand"   //乱数を生成するためのパッケージ
	"time"        //時間に関する操作を扱うためのパッケージ

	"github.com/hajimehoshi/ebiten/v2"            //Ebitenのゲームエンジンのパッケージ
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" //Ebitenのユーティリティパッケージ
)

// 画面のサイズとボールの大きさを定義
const (
	screenWidth  = 640
	screenHeight = 480
	ballSize     = 10
)

// Game はゲームの状態を表します。ボールの位置と速度を保持します。
type Game struct {
	x, y   float64
	vx, vy float64
}

// NewGame は新しいゲーム状態を作成します。ボールの初期速度はランダムです。
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())                 //乱数のシードを現在の時間（ナノ秒）に設定
	vx, vy := rand.Float64()*4-2, rand.Float64()*4-2 //ランダムな速度を生成
	return &Game{vx: vx, vy: vy}                     //新しいゲーム状態を返す
}

// Update はゲームのロジックを更新します。ボールの位置を更新し、壁への衝突を検出します。
// 壁と衝突した際には、ボールの速度をランダムに変更します。
func (g *Game) Update() error {
	g.x += g.vx // X軸方向に速度分動かします。
	g.y += g.vy // Y軸方向に速度分動かします。

	// ボールが壁に衝突したら、反射させるだけでなくランダムな方向に速度を変更
	if g.x <= 0 || g.x+ballSize >= screenWidth {
		g.vx = rand.Float64()*4 - 2 // X軸の速度をランダムに設定
	}
	if g.y <= 0 || g.y+ballSize >= screenHeight {
		g.vy = rand.Float64()*4 - 2 // Y軸の速度をランダムに設定
	}

	return nil // エラーはないのでnilを返す
}

// Draw はゲームの状態を描画します。背景を黒で塗りつぶし、ボールを白で描画します。
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // 画面を黒で塗りつぶします。

	// 矩形（ボール）を描画します。
	ebitenutil.DrawRect(screen, g.x, g.y, ballSize, ballSize, color.RGBA{255, 255, 255, 255})
}

// Layout は画面の解像度を制御します。ここでは固定の画面解像度を返しています。
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight // 定義した画面解像度を返す
}

// main関数はゲームの開始地点です。ウィンドウサイズとタイトルを設定し、ゲームを開始します。
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight) // ウィンドウのサイズを設定
	ebiten.SetWindowTitle("ボールランダム移動")              // ウィンドウのタイトルを設定
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err) // ゲームの実行中にエラーが発生した場合、ログを出力してプログラムを終了
	}
}
