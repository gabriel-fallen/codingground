package application

import "time"


type SpaceCube struct {
    x, y, w, h float32
    complexity int
}

func (cube SpaceCube) Process() {
    delay := int(cube.w*cube.h)*cube.complexity
    time.Sleep(delay * time.Millisecond)
}
