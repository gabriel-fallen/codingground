//package application
package main

import (
    "time"
)


type SpaceCube struct {
    x, y, w, h float32
    complexity int
}

func (cube SpaceCube) Process() {
    delay := cube.w*cube.h*float32(cube.complexity)
    time.Sleep(time.Duration(delay) * time.Millisecond)
}

func SimpleWorker(workQueue <-chan SpaceCube, resultQueue chan<- SpaceCube) {
    for cube := range workQueue {
        cube.Process() // here we sleep simulating busy work
        resultQueue <- cube
    }
    // FIXME: close(resultQueue)
}
