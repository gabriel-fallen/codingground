//package network
package main

import (
    "time"
    "math/rand"
)

const (
    Latency   = 10 * time.Millisecond
    Bandwidth = 1000 // 1Gb ethernet
    HeartbeatInterval = 500 * time.Millisecond // FIXME: what should it be?
    HeartbeatTolerance = 3 // forget about a node after 3 missed heartbeats
    MTBF = 24.0 // Mean Time Between Failures, hours
    TBFDispersion = 5.0
)

// FIXME: interface{} instead of SpaceCube to avoid dependence on application.go?
type WorkerFunc func(workQueue <-chan SpaceCube, resultQueue chan<- SpaceCube)


func SendMesssage(msg interface{}, link chan<- interface{}) {
    time.Sleep(Latency)
    link <- msg
}

func SimpleNode(worker WorkerFunc, failer func(*time.Ticker), workQueue <-chan SpaceCube, resultQueue chan<- SpaceCube) (heartbeat <-chan time.Time) {
    ticker := time.NewTicker(HeartbeatInterval)
    heartbeat = ticker.C
    go worker(workQueue, resultQueue)
    go failer(ticker)
    return
}

func GaussianFailer(ticker *time.Ticker) {
    ttf := rand.NormFloat64() * TBFDispersion + MTBF
    go func() {
        time.Sleep(time.Duration(ttf) * time.Hour)
        ticker.Stop()
    }()
}
