package main

import (
	"math"
	"strconv"
)

var elevatorID = 1
var floorRequestButtonID = 1
var columnID = 1
var callButtonID = 1

type Column struct {
	ID                       string
	status                   string
	servedFloors             []int
	amountOfFloors           int
	amountOfElevators        int
	isBasement               bool
	elevatorsList            []Elevator
	callButtonsList          []CallButton
	bestElevatorInformations BestElevatorInformations
}

func NewColumn(_id string, _status string, _amountOfFloors int, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	e := Column{}
	e.ID = _id
	e.status = _status
	e.amountOfFloors = _amountOfFloors
	e.amountOfElevators = _amountOfElevators
	e.isBasement = _isBasement
	return e

}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	var elevator = c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1) //Always 1 because the user can only go back to the lobby
	elevator.move()
	return elevator

}
func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		buttonFloor := -1

		for i := 0; i < _amountOfFloors; i++ {
			callButton := CallButton{callButtonID, "OFF", buttonFloor, "UP"}
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {

		buttonFloor := 1

		for i := 0; i < _amountOfFloors; i++ {
			callButton := CallButton{callButtonID, "OFF", buttonFloor, "DOWN"}
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}
}
func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for i := 0; i < _amountOfFloors; i++ {
		var _elevatorID string = strconv.Itoa(elevatorID)
		elevator := Elevator(_elevatorID, "idle", _amountOfFloors)
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}

}

type BestElevatorInformations struct {
	bestElevator Elevator
	bestScore    int
	referenceGap int
}

/*func (c *BestElevatorInformations) NewBestElevatorInformations(_bestElevator Elevator, _bestScore int, _referenceGap int) *BestElevatorInformations {
	c.bestElevator = _bestElevator
	c.bestScore = _bestScore
	c.referenceGap = _referenceGap
	return BestElevatorInformations

}*/
func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {

	bestElevatorInformations := BestElevatorInformations{
		bestElevator: Elevator{}
		bestScore:    6
		referenceGap: 1000000
	}
	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			//The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is at the lobby and has no requests
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is lower than me and is coming up. It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is above me and is coming down. It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is not at the first floor, but doesnt have any request
			} else {
				bestElevatorInformations = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is not available, but still could take the call if nothing better is found
			}
			bestElevator = bestElevatorInformations.bestElevator
			bestScore = bestElevatorInformations.bestScore
			referenceGap = bestElevatorInformations.referenceGap
		}
	} else {
		for _, elevator := range c.elevatorList {

			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is at the same level as me, and is about to depart to the first floor
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is idle and has no requests
			} else {
				bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor) //The elevator is not available, but still could take the call if nothing better is found
			}
			bestElevator = bestElevatorInformations.bestElevator
			bestScore = bestElevatorInformations.bestScore
			referenceGap = bestElevatorInformations.referenceGap
		}
	}
	return bestElevator
}
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestScore int, referenceGap int, bestElevator Elevator, floor int) *BestElevatorInformations {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor - floor)))
	} else if bestScore == scoreToCheck {
		var gap int = int(math.Abs(float64(newElevator.currentFloor - floor)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevatorInformations
}
