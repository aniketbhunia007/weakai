package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/unixpickle/weakai/svm"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	positives := []svm.Sample{
		{0, 0},
		{0, 0.3},
		{0.3, 0},
		{0.1, 0.1},
	}

	negatives := []svm.Sample{
		{1, 1},
		{1, 0.7},
		{0.7, 1},
		{0.9, 0.9},
	}

	problem := &svm.Problem{
		Positives: positives,
		Negatives: negatives,
		Kernel:    svm.LinearKernel,
	}

	solution := svm.RandomlySolve(problem, 100000, 20)
	fmt.Println("Solution from random solver:", solution)

	subgradientSolver := &svm.SubgradientSolver{
		Tradeoff: 0.001,
		Steps:    100000,
		StepSize: 0.001,
	}
	solution = subgradientSolver.Solve(problem)

	fmt.Println("Solution from subgradient solver:", solution)
}