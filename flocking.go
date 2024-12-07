package main

import (
	"fmt"
	"math"
	"math/rand"
	"encoding/csv"
	"os"
)

const (
	BIRD_COUNT = 200
	BIRD_SPEED = 1.0
	TIME_STEP = 0.005
	TOTAL_TIME_STEPS = 1000
	BOX_LENGTH = 1.0
	CLOSE_DISTANCE = 0.1
	REPEL_DISTANCE = 0.05
	WALL_DISTANCE = 0.1
	FORCE_TO_CENTER_SCALE = 0.5
	FORCE_REPEL_SCALE = 0.5
	FORCE_ALIGN_SCALE = 0.5
	FORCE_WALL_SCALE = 2.0
	CSV_FILE_TO_WRITE = "trajectory.csv"
)

type Bird struct {
	Position Vector
	Velocity Vector
	ID int
}

func CreateRandomBird(id int) Bird {
	return Bird{
		Position: Vector{
			x: rand.Float64(), 
			y: rand.Float64(),
		}, 
		Velocity: Vector{
			x: rand.Float64(), 
			y: rand.Float64(),
		},
		ID: id,
	}
}

func (bird *Bird) PositionTick() {
	bird.Position.y += TIME_STEP * bird.Velocity.y
	bird.Position.x += TIME_STEP * bird.Velocity.x
}

func (bird1 *Bird) CheckClose(bird2 Bird) (bool,bool) {
	dist := math.Sqrt(
		math.Pow(bird1.Position.x - bird2.Position.x,2) + 
		math.Pow(bird1.Position.y - bird2.Position.y,2))
	return CLOSE_DISTANCE > dist, REPEL_DISTANCE > dist
}

func (bird *Bird) ForceToCenter(closeBirds []Bird) Vector {
	force := Vector{x:0.0, y: 0.0}
	if len(closeBirds) == 0 {
		return force
	}
	for i := range closeBirds {
		force.x += closeBirds[i].Position.x
		force.y += closeBirds[i].Position.y
	}
	force.x /= float64(len(closeBirds))
	force.y /= float64(len(closeBirds))
	force.x -= bird.Position.x
	force.y -= bird.Position.y
	force.x *= FORCE_TO_CENTER_SCALE
	force.y *= FORCE_TO_CENTER_SCALE
	return force
}

func (bird *Bird) ForceRepel(repelBirds []Bird) Vector {
	force := Vector{x:0.0, y: 0.0}
	if len(repelBirds) == 0 {
		return force
	}
	for i := range repelBirds {
		force.x -= (repelBirds[i].Position.x - bird.Position.x)
		force.y -= (repelBirds[i].Position.x - bird.Position.y)
	}
	force.x /= float64(len(repelBirds))
	force.y /= float64(len(repelBirds))
	force.x *= FORCE_REPEL_SCALE
	force.y *= FORCE_REPEL_SCALE
	return force
}

func (bird *Bird) ForceAlign(closeBirds []Bird) Vector {
	force := Vector{x:0.0, y: 0.0}
	if len(closeBirds) == 0 {
		return force
	}
	for i := range closeBirds {
		force.x += closeBirds[i].Velocity.x
		force.y += closeBirds[i].Velocity.y
	}
	force.x /= float64(len(closeBirds))
	force.y /= float64(len(closeBirds))
	force.x -= bird.Velocity.x
	force.y -= bird.Velocity.y
	force.x *= FORCE_ALIGN_SCALE
	force.y *= FORCE_ALIGN_SCALE
	return force
}

func (bird *Bird) ForceWall() Vector {
	force := Vector{x:0.0, y: 0.0}
	if bird.Position.x < WALL_DISTANCE {
		force.x += FORCE_WALL_SCALE
	}
	if bird.Position.y < WALL_DISTANCE {
		force.y += FORCE_WALL_SCALE
	}
	if bird.Position.y > (BOX_LENGTH - WALL_DISTANCE) {
		force.y -= FORCE_WALL_SCALE
	}
	if bird.Position.x > (BOX_LENGTH - WALL_DISTANCE) {
		force.x -= FORCE_WALL_SCALE
	}
	return force
}

func (bird *Bird) GetCloseAndRepelBirds(birdsArray [BIRD_COUNT]Bird) ([]Bird, []Bird) {
	var closeBirds, repelBirds []Bird 
	var close, repel bool
	for i := range birdsArray {
		if bird.ID != birdsArray[i].ID {
			close, repel = bird.CheckClose(birdsArray[i])
			if close {
				closeBirds = append(closeBirds, birdsArray[i])
			}
			if repel {
				repelBirds = append(repelBirds, birdsArray[i])
			}
		}
	}
	return closeBirds, repelBirds
}

func (bird *Bird) VelocityUpdate(birdsArray [BIRD_COUNT]Bird) {
	force := Vector{x:0.0, y:0.0}
	closeBirds, repelBirds := bird.GetCloseAndRepelBirds(birdsArray)

	centerForce := bird.ForceToCenter(closeBirds)
	force.AddAssign(centerForce)
	repelForce := bird.ForceRepel(repelBirds)
	force.AddAssign(repelForce)
	alignForce := bird.ForceAlign(closeBirds)
	force.AddAssign(alignForce)
	wallForce := bird.ForceWall()
	force.AddAssign(wallForce)

	bird.Velocity.AddAssign(force.Scale(TIME_STEP))
}

func WriteOutput(birdsArray [BIRD_COUNT]Bird, fileName string) {
	f, e := os.Create(fileName)
    if e != nil {
        fmt.Println(e)
    }

    writer := csv.NewWriter(f)
	var workingString [][]string
	for i := range birdsArray {
		positionPair := []string {
			fmt.Sprintf("%f",birdsArray[i].Position.x),
			fmt.Sprintf("%f",birdsArray[i].Position.y),
		}
		workingString = append(workingString, positionPair)
	}

    e = writer.WriteAll(workingString)
    if e != nil {
        fmt.Println(e)
    }
}

func RunSimulation(){
	// Initialize
	var birdsArray [BIRD_COUNT]Bird
	for i:=0; i<len(birdsArray); i++ {
		birdsArray[i] = CreateRandomBird(i)
	}
	var csvFileName string
	// Compute Loop
	for i :=0; i<TOTAL_TIME_STEPS; i++{
		for i := range birdsArray {
			birdsArray[i].VelocityUpdate(birdsArray)
		}
		for i := range birdsArray {
			birdsArray[i].PositionTick()
		}
		csvFileName = fmt.Sprintf("out/positions_%03d.csv",i)
		WriteOutput(birdsArray, csvFileName)
	}
}

func main() {
	RunSimulation()
}