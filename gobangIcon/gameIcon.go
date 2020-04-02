package gobangIcon

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type ImageStyle int

var (
	chess fyne.Resource
	black fyne.Resource
	white fyne.Resource
)

const (
	BlackChess ImageStyle = iota
	WhiteChess
	Chessboard
)

type GobangIcon struct {
	ImageType ImageStyle
	widget.Icon
}

func (g *GobangIcon) Tapped(p *fyne.PointEvent) {
	if g.ImageType == Chessboard {
		chessBoardClicked(p)
		g.Refresh()
	} else if g.ImageType == BlackChess {
		log.Println("黑棋被单击了，坐标：", p.AbsolutePosition)
	} else if g.ImageType == WhiteChess {
		log.Println("白棋被单击了，坐标：", p.AbsolutePosition)
	}
}

func (g *GobangIcon) TappedSecondary(_ *fyne.PointEvent) {

}

func chessBoardClicked(p *fyne.PointEvent) {
	log.Println("棋盘被单击了，坐标：", p.Position)
}

func newGobangIcon(is ImageStyle, res *fyne.Resource) *GobangIcon {
	img := &GobangIcon{}
	img.ExtendBaseWidget(img)
	img.ImageType = is
	img.SetResource(*res)

	return img
}

func NewChessboard() *GobangIcon {
	return newGobangIcon(Chessboard, &chess)
}

func NewBlackChess() *GobangIcon {
	return newGobangIcon(BlackChess, &black)
}

func NewWhiteChess() *GobangIcon {
	return newGobangIcon(WhiteChess, &white)
}

func init() {
	var err error
	chess, err = fyne.LoadResourceFromPath("./image/chessboard.jpg")
	if err != nil {
		log.Println("读取棋盘图像出错，错误：", err)
		return
	}
	black, err = fyne.LoadResourceFromPath("./image/black.png")
	if err != nil {
		log.Println("读取黑子象出错。")
		return
	}
	white, err = fyne.LoadResourceFromPath("./image/white.png")
	if err != nil {
		log.Println("读取白子象出错。")
		return
	}
}
