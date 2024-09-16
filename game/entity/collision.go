package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"math"
)

func computeTightBoundingBox(img *ebiten.Image) image.Rectangle {
	bounds := img.Bounds()
	minX, minY := bounds.Max.X, bounds.Max.Y
	maxX, maxY := bounds.Min.X, bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// If no non-transparent pixels are found, return an empty rectangle
	if minX > maxX || minY > maxY {
		return image.Rect(0, 0, 0, 0)
	}

	// Adjust coordinates to be relative to the sub-image's bounds
	adjustedMinX := minX - bounds.Min.X
	adjustedMinY := minY - bounds.Min.Y
	adjustedMaxX := maxX - bounds.Min.X
	adjustedMaxY := maxY - bounds.Min.Y

	return image.Rect(adjustedMinX, adjustedMinY, adjustedMaxX+1, adjustedMaxY+1)
}

func generateCollisionOutline(img *ebiten.Image) []Vec2d {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	binaryImg := make([][]bool, height)
	for y := 0; y < height; y++ {
		binaryImg[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			_, _, _, a := img.At(bounds.Min.X+x, bounds.Min.Y+y).RGBA()
			if a > 0 {
				binaryImg[y][x] = true
			}
		}
	}

	// Find the starting point
	startX, startY := -1, -1
	for y := 0; y < height && startX == -1; y++ {
		for x := 0; x < width && startX == -1; x++ {
			if binaryImg[y][x] {
				startX, startY = x, y
			}
		}
	}
	if startX == -1 {
		return nil // No non-transparent pixels found
	}

	// Trace the contour
	outline := traceContour(binaryImg, startX, startY)
	return outline
}

func traceContour(img [][]bool, startX, startY int) []Vec2d {
	// 8-connected neighborhood offsets
	offsets := [8][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	width, height := len(img[0]), len(img)
	visited := make([][]bool, height)
	for y := 0; y < height; y++ {
		visited[y] = make([]bool, width)
	}

	var outline []Vec2d
	x, y := startX, startY
	for {
		outline = append(outline, Vec2d{X: float64(x), Y: float64(y)})
		visited[y][x] = true

		found := false
		for _, offset := range offsets {
			nx, ny := x+offset[0], y+offset[1]
			if nx >= 0 && nx < width && ny >= 0 && ny < height {
				if img[ny][nx] && !visited[ny][nx] {
					x, y = nx, ny
					found = true
					break
				}
			}
		}
		if !found || (x == startX && y == startY) {
			break
		}
	}
	return outline
}

func simplifyOutline(points []Vec2d, epsilon float64) []Vec2d {
	if len(points) < 3 {
		return points
	}

	var dmax float64
	index := 0
	end := len(points) - 1
	for i := 1; i < end; i++ {
		d := perpendicularDistance(points[i], points[0], points[end])
		if d > dmax {
			index = i
			dmax = d
		}
	}

	if dmax > epsilon {
		recResults1 := simplifyOutline(points[:index+1], epsilon)
		recResults2 := simplifyOutline(points[index:], epsilon)
		return append(recResults1[:len(recResults1)-1], recResults2...)
	} else {
		return []Vec2d{points[0], points[end]}
	}
}

func perpendicularDistance(point, lineStart, lineEnd Vec2d) float64 {
	if lineStart.X == lineEnd.X && lineStart.Y == lineEnd.Y {
		return math.Hypot(point.X-lineStart.X, point.Y-lineStart.Y)
	}
	numerator := math.Abs((lineEnd.Y-lineStart.Y)*point.X - (lineEnd.X-lineStart.X)*point.Y + lineEnd.X*lineStart.Y - lineEnd.Y*lineStart.X)
	denominator := math.Hypot(lineEnd.Y-lineStart.Y, lineEnd.X-lineStart.X)
	return numerator / denominator
}

func transformOutline(outline []Vec2d, position Vec2d, scale float64, rotation float64) []Vec2d {
	transformed := make([]Vec2d, len(outline))
	sinTheta := math.Sin(rotation)
	cosTheta := math.Cos(rotation)
	for i, point := range outline {
		// Apply scaling
		x := point.X * scale
		y := point.Y * scale

		// Apply rotation
		xRot := x*cosTheta - y*sinTheta
		yRot := x*sinTheta + y*cosTheta

		// Apply translation
		transformed[i] = Vec2d{
			X: xRot + position.X,
			Y: yRot + position.Y,
		}
	}
	return transformed
}

func polygonsCollide(poly1, poly2 []Vec2d) bool {
	axes1 := getAxes(poly1)
	axes2 := getAxes(poly2)

	for _, axis := range axes1 {
		if !overlapOnAxis(poly1, poly2, axis) {
			return false // Separating axis found
		}
	}

	for _, axis := range axes2 {
		if !overlapOnAxis(poly1, poly2, axis) {
			return false // Separating axis found
		}
	}

	return true // No separating axis found; polygons collide
}

func getAxes(polygon []Vec2d) []Vec2d {
	axes := make([]Vec2d, len(polygon))
	for i := 0; i < len(polygon); i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]
		edge := Vec2d{X: p2.X - p1.X, Y: p2.Y - p1.Y}
		normal := Vec2d{X: -edge.Y, Y: edge.X}
		length := math.Hypot(normal.X, normal.Y)
		if length != 0 {
			normal.X /= length
			normal.Y /= length
		}
		axes[i] = normal
	}
	return axes
}

func overlapOnAxis(poly1, poly2 []Vec2d, axis Vec2d) bool {
	min1, max1 := projectPolygon(poly1, axis)
	min2, max2 := projectPolygon(poly2, axis)
	return !(max1 < min2 || max2 < min1)
}

func projectPolygon(polygon []Vec2d, axis Vec2d) (float64, float64) {
	min := dotProduct(polygon[0], axis)
	max := min
	for _, point := range polygon[1:] {
		proj := dotProduct(point, axis)
		if proj < min {
			min = proj
		} else if proj > max {
			max = proj
		}
	}
	return min, max
}

func dotProduct(v1, v2 Vec2d) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func drawOutline(screen *ebiten.Image, outline []Vec2d, clr color.Color) {
	if len(outline) < 2 {
		return // Need at least two points to draw an outline
	}

	var path vector.Path

	// Move to the starting point
	path.MoveTo(float32(outline[0].X), float32(outline[0].Y))

	// Draw lines to each subsequent point
	for _, point := range outline[1:] {
		path.LineTo(float32(point.X), float32(point.Y))
	}

	// Close the path to create a complete shape
	path.Close()

	// Set stroke options
	strokeOptions := &vector.StrokeOptions{
		Width: 1.0, // Adjust the line width as needed
	}

	// Generate vertices and indices for the stroke
	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, strokeOptions)

	// Set the color for each vertex
	R, G, B, A := clr.RGBA()
	for i := range vertices {
		vertices[i].ColorR = float32(R) / 255
		vertices[i].ColorG = float32(G) / 255
		vertices[i].ColorB = float32(B) / 255
		vertices[i].ColorA = float32(A) / 255
	}

	// Create a 1x1 white image to use as a texture
	whitePixel := ebiten.NewImage(1, 1)
	whitePixel.Fill(color.White)

	// Draw the outline on the screen
	screen.DrawTriangles(vertices, indices, whitePixel, nil)
}
