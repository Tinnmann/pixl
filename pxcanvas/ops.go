package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xdiff := currentCoord.Position.X - previousCoord.Position.X
	ydiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xdiff
	pxCanvas.CanvasOffset.Y += ydiff
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	default:
		pxCanvas.PxSize = 10
	}
}
