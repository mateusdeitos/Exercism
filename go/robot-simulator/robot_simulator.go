package robot

import (
	"fmt"
	"sync"
)

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N, E, S, W Dir = 0, 1, 2, 3
)

func Right() {
	Step1Robot.Dir++
	if Step1Robot.Dir > W {
		Step1Robot.Dir = N
		return
	}
}

func Left() {
	Step1Robot.Dir--
	if Step1Robot.Dir < 0 {
		Step1Robot.Dir = W
		return
	}
}

func Advance() {
	if Step1Robot.Dir == N {
		Step1Robot.Y++
		return
	}

	if Step1Robot.Dir == S {
		Step1Robot.Y--
		return
	}

	if Step1Robot.Dir == E {
		Step1Robot.X++
		return
	}

	// Step1Robot.Dir == W
	Step1Robot.X--
}

func (d Dir) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

// Step 2
// Define Action type here.
type Action struct {
	Command
}

func StartRobot(command <-chan Command, action chan<- Action) {
	for c := range command {
		action <- Action{c}
	}

	close(action)
}

func Room(extent Rect, robot Step2Robot, action <-chan Action, report chan<- Step2Robot) {
	for a := range action {
		switch a.Command {
		case 'L':
			robot.TurnLeft()
		case 'R':
			robot.TurnRight()
		case 'A':
			pos := robot.Pos
			robot.Advance()

			if !robotWithinBounds(robot, extent) {
				robot.Pos = pos
				continue
			}

		}
	}

	report <- robot
}

func robotWithinBounds(robot Step2Robot, extent Rect) bool {
	pos := robot.Pos
	if pos.Easting < extent.Min.Easting {
		return false
	}

	if pos.Easting > extent.Max.Easting {
		return false
	}

	if pos.Northing < extent.Min.Northing {
		return false
	}

	if pos.Northing > extent.Max.Northing {
		return false
	}

	return true
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	Name string
	Type string
	Command
}

const (
	actionTypeFinished = "finished"
	actionTypeCommand  = "command"
)

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if len(script) == 0 {
		action <- Action3{name, actionTypeFinished, 0}
		return
	}

	for _, c := range script {
		action <- Action3{name, actionTypeCommand, Command(c)}
	}

	action <- Action3{name, actionTypeFinished, 0}
}

type robot struct {
	*Step3Robot
	index int
}

type PositionalMapping struct {
	p map[RU]RU
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	names := make(map[string]*robot)
	posMap := make(map[RU]RU)
	// uniqueLogControl := make(map[string]bool)
	finishedCount := 0
	for i := range robots {
		r := &robots[i]
		if r.Name == "" {
			log <- fmt.Sprintf("%v: robot name not set", r)
			finishedCount++
			continue
		}

		_, exists := names[r.Name]
		if exists {
			log <- fmt.Sprintf("%v: duplicate robot name %q", r, r.Name)
			finishedCount++
			finishedCount++
			continue
		}

		posNorthing, ok := posMap[r.Pos.Easting]
		if ok && posNorthing == r.Pos.Northing {
			log <- fmt.Sprintf("%v: duplicate position %v", r, r.Pos)
			finishedCount++
			finishedCount++
			continue
		}

		if !robotWithinBounds(r.Step2Robot, extent) {
			log <- fmt.Sprintf("%v: robot out of bounds", r)
			finishedCount++
			continue
		}

		names[r.Name] = &robot{r, i}
		posMap[r.Pos.Easting] = r.Pos.Northing
	}

	for finishedCount < len(robots) {
		a := <-action
		r, ok := names[a.Name]
		if !ok {
			finishedCount++
			log <- fmt.Sprintf("%s: %s", a.Name, "unknown robot")

			continue
		}

		if a.Type == actionTypeFinished {
			finishedCount++
			continue
		}

		if !a.Command.IsValid() {
			log <- fmt.Sprintf("%s: invalid command %q", r.Name, a.Command)
			finishedCount++
			continue
		}

		switch a.Command {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'A':
			var m sync.Mutex

			m.Lock()
			pos := r.Pos
			r.Advance()

			posNorthing, ok := posMap[r.Pos.Easting]
			if ok && posNorthing == r.Pos.Northing {
				log <- fmt.Sprintf("%s on position %v: bumped into another robot on position: %v", r.Name, pos, r.Pos)
				r.Pos = pos
				m.Unlock()
				continue
			}

			if !robotWithinBounds(r.Step2Robot, extent) {
				log <- fmt.Sprintf("%s: bump into wall at %v", r.Name, pos)
				r.Pos = pos
				m.Unlock()
				continue
			}

			posMap[r.Pos.Easting] = r.Pos.Northing
			delete(posMap, pos.Easting)
			m.Unlock()
		}
	}

	rep <- robots
}

func uniqueLog(msg string, log chan string, uniqueLogControl map[string]bool) {
	if _, ok := uniqueLogControl[msg]; ok {
		return
	}
	uniqueLogControl[msg] = true
	log <- msg
}
