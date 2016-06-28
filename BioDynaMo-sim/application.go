package application

import "time"


type SpaceCube struct {
    x, y, w, h float32
    complexity int
}

func (cube SpaceCube) Process() {
    delay := cube.w*cube.h*float32(cube.complexity)
    time.Sleep(time.Duration(delay) * time.Millisecond)
}
