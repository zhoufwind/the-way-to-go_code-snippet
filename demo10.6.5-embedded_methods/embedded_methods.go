package main

type Engine interface {
    Start()
    Stop()
}

type Car struct {
    Engine
}

func (c *Car) GoToWorkIn() {
    // get in car
    c.Start()
    // drive to work
    c.Stop()
    // get out of car
}
