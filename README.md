# Flea circus

A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell? Give your answer rounded to six decimal places.

## Solution

The solution is based on the fact that it is not necessary to wait for a bell to ring if the total number of bells to ring is known.
This knowledge allows us to simply make all the fleas jump N times in parallel instead of waiting for the slowest flea to finish 1 jump.

### Implementation details
1. The fleas all jump in parallel N times, where N is the number of rings of the bell;
2. When using multiple simulations, they are also run in parallel.
3. There are 4 values with which you can play in `main.go`:
   
    - numberOfSimulations: the number of simulations to run;
    - numberOfBellsRung: the number of times to ring the bell;
    - circusWidth: the width of the circus;
    - circusHeight: the height of the circus.
    
4. The circus grid can be imagined as a XY axis where the top left corner is the position `(0,0)` and the bottom right corner is the position `(circusHeight-1, circusWidth-1)`;
5. Finally, there is a simple benchmarking example in `main_test.go` that benchmarks the `simulate()` function.

### Running 
``go run cmd/circus/main.go``
