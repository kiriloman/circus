package main

import (
	"fmt"
	"sync"

	"github.com/kiriloman/circus"
)

const (
	numberOfSimulations = 150

	numberOfBellsRung = 50
	circusWidth       = 30
	circusHeight      = 30
)

func main() {
	expectedValue := simulate(numberOfSimulations)
	fmt.Printf("%.6f\n", expectedValue)
}

// simulate starts up *simulations* number of simulations
// and extracts the number of unoccupied positions from all of them
// performing the calculation of the expected number of unoccupied positions.
func simulate(simulations int) float32 {
	totalUnoccupied := 0
	unoccupiedPositions := make(chan int, simulations)

	var wg sync.WaitGroup

	for i := 0; i < simulations; i++ {
		wg.Add(1)
		go func() {
			run(&wg, unoccupiedPositions)
		}()
	}

	go func() {
		for num := range unoccupiedPositions {
			totalUnoccupied += num
		}
	}()

	wg.Wait()
	close(unoccupiedPositions)

	return float32(totalUnoccupied) / float32(simulations)
}

// run runs one simulations of the given problem.
// It starts a goroutine per flea in which the flea jumps around
// for numberOfBellsRung times.
// It then sends the number of unoccupied positions to the unoccupiedPositions channel.
func run(simulateWG *sync.WaitGroup, unoccupiedPositions chan<- int) {
	positions := make(chan circus.Position, circusWidth*circusHeight)

	var wg sync.WaitGroup

	for x := 0; x < circusWidth; x++ {
		for y := 0; y < circusHeight; y++ {
			wg.Add(1)

			go func(x, y int) {
				flea := circus.NewFlea(x, y)
				flea.JumpAround(numberOfBellsRung, &wg, positions, circusWidth, circusHeight)
			}(x, y)
		}
	}

	// Use a map as a set to store the occupied positions.
	occupiedPositions := map[circus.Position]struct{}{}
	go func() {
		for finalPosition := range positions {
			occupiedPositions[finalPosition] = struct{}{}
		}
	}()

	wg.Wait()
	close(positions)

	// The number of unoccupied positions is the total number of positions
	// minus the number of occupied ones.
	unoccupiedPositions <- circusWidth*circusHeight - len(occupiedPositions)
	simulateWG.Done()
}
