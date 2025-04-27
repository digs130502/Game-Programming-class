package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gridWidth    = 80
	gridHeight   = 60
	cellSize     = 10
	minRoomSize  = 6
	maxDepth     = 5
)

type Rect struct {
	X, Y, W, H int
}

type Node struct {
	Partition Rect
	Left      *Node
	Right     *Node
	Room      *Rect
	Color     rl.Color
}

var (
	rooms    []Rect
	hallways []Rect
	root     *Node
	dungeon  [][]rl.Color
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rl.InitWindow(screenWidth, screenHeight, "BSP Dungeon Generation")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	generateDungeon()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyR) {
			generateDungeon()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for y := 0; y < gridHeight; y++ {
			for x := 0; x < gridWidth; x++ {
				rl.DrawRectangle(int32(x*cellSize), int32(y*cellSize), cellSize, cellSize, dungeon[y][x])
			}
		}
		rl.EndDrawing()
	}
}

func generateDungeon() {
	dungeon = make([][]rl.Color, gridHeight)
	for i := range dungeon {
		dungeon[i] = make([]rl.Color, gridWidth)
	}

	root = &Node{Partition: Rect{0, 0, gridWidth, gridHeight}}
	rooms = []Rect{}
	hallways = []Rect{}
	split(root, 0)
	drawPartitions(root)
	connectRooms(root)
}

func split(n *Node, depth int) {
	if depth >= maxDepth || (n.Partition.W < 2*minRoomSize && n.Partition.H < 2*minRoomSize) {
		makeRoom(n)
		return
	}

	horizontal := rand.Float32() < 0.5
	if n.Partition.W > n.Partition.H {
		horizontal = false
	} else if n.Partition.H > n.Partition.W {
		horizontal = true
	}

	if horizontal {
		if n.Partition.H <= minRoomSize*2 {
			makeRoom(n)
			return
		}
		splitY := rand.Intn(n.Partition.H-minRoomSize*2) + n.Partition.Y + minRoomSize
		n.Left = &Node{Partition: Rect{n.Partition.X, n.Partition.Y, n.Partition.W, splitY - n.Partition.Y}}
		n.Right = &Node{Partition: Rect{n.Partition.X, splitY, n.Partition.W, n.Partition.Y + n.Partition.H - splitY}}
	} else {
		if n.Partition.W <= minRoomSize*2 {
			makeRoom(n)
			return
		}
		splitX := rand.Intn(n.Partition.W-minRoomSize*2) + n.Partition.X + minRoomSize
		n.Left = &Node{Partition: Rect{n.Partition.X, n.Partition.Y, splitX - n.Partition.X, n.Partition.H}}
		n.Right = &Node{Partition: Rect{splitX, n.Partition.Y, n.Partition.X + n.Partition.W - splitX, n.Partition.H}}
	}

	split(n.Left, depth+1)
	split(n.Right, depth+1)
}

func makeRoom(n *Node) {
	w := minRoomSize
	h := minRoomSize

	if n.Partition.W > minRoomSize {
		w = rand.Intn(n.Partition.W-minRoomSize+1) + minRoomSize
	} else {
		w = n.Partition.W
	}

	if n.Partition.H > minRoomSize {
		h = rand.Intn(n.Partition.H-minRoomSize+1) + minRoomSize
	} else {
		h = n.Partition.H
	}

	x := n.Partition.X
	y := n.Partition.Y
	if n.Partition.W > w {
		x += rand.Intn(n.Partition.W - w + 1)
	}
	if n.Partition.H > h {
		y += rand.Intn(n.Partition.H - h + 1)
	}

	room := Rect{x, y, w, h}
	n.Room = &room
	rooms = append(rooms, room)
}

func drawPartitions(n *Node) {
	if n == nil {
		return
	}
	if n.Room != nil {
		color := rl.NewColor(uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255)
		n.Color = color
		for y := n.Room.Y; y < n.Room.Y+n.Room.H; y++ {
			for x := n.Room.X; x < n.Room.X+n.Room.W; x++ {
				dungeon[y][x] = color
			}
		}
	}
	drawPartitions(n.Left)
	drawPartitions(n.Right)
}

func connectRooms(n *Node) {
	if n.Left != nil && n.Right != nil && n.Left.Room != nil && n.Right.Room != nil {
		left := *n.Left.Room
		right := *n.Right.Room
		x1 := left.X + left.W/2
		y1 := left.Y + left.H/2
		x2 := right.X + right.W/2
		y2 := right.Y + right.H/2

		if rand.Float32() < 0.5 {
			// Horizontal first
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				dungeon[y1][x] = rl.Gray
			}
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				dungeon[y][x2] = rl.Gray
			}
		} else {
			// Vertical first
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				dungeon[y][x1] = rl.Gray
			}
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				dungeon[y2][x] = rl.Gray
			}
		}
	}
	if n.Left != nil {
		connectRooms(n.Left)
	}
	if n.Right != nil {
		connectRooms(n.Right)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
