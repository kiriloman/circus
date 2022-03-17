package circus

import (
	"math/rand"
	"sync"
)

type Position struct {
	X int
	Y int
}

type Direction int

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type Flea struct {
	Position
}

func NewFlea(x, y int) Flea {
	return Flea{Position: Position{X: x, Y: y}}
}

// JumpAround makes the flea to jump to adjacent valid positions *times* times.
// After reaching the final positions after the final random jump,
// the position is sent to the *positionChannel*.
func (flea *Flea) JumpAround(
	times int,
	wg *sync.WaitGroup,
	positionChannel chan<- Position,
	widthBoundary, heightBoundary int,
) {
	// this loop could be further improved by having a goroutine per jump
	// and would require a mutex lock on the flea's position.
	// This would be beneficial for a case with many required jumps.
	for i := 0; i < times; i++ {
		possibleDirections := flea.possibleDirections(widthBoundary, heightBoundary)
		direction := possibleDirections[rand.Intn(len(possibleDirections))]
		flea.move(direction)
	}
	positionChannel <- flea.Position
	wg.Done()
}

// possibleDirections validates in which directions the flea can move.
// It is not possible to move out of bounds of a given circus area.
func (flea *Flea) possibleDirections(heightBoundary, widthBoundary int) []Direction {
	var directions []Direction

	// Can move up
	if flea.Y != 0 {
		directions = append(directions, DirectionUp)
	}

	// Can move right
	if flea.X != widthBoundary-1 {
		directions = append(directions, DirectionRight)
	}

	// Can move down
	if flea.Y != heightBoundary-1 {
		directions = append(directions, DirectionDown)
	}

	// Can move left
	if flea.X != 0 {
		directions = append(directions, DirectionLeft)
	}

	return directions
}

// move moves the flea in a given Direction
func (flea *Flea) move(direction Direction) {
	switch direction {
	case DirectionUp:
		flea.Position.Y--
	case DirectionRight:
		flea.Position.X++
	case DirectionDown:
		flea.Position.Y++
	case DirectionLeft:
		flea.Position.X--
	}
}
