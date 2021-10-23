package main





type Elevator struct {
	ElevatorID string
    status string
    amountOfFloors int 
    currentFloor int
    direction string
    overweight bool
    floorRequestList []int
    requestList []int
    door Door
	destination int 

}

func NewElevator(_elevatorID string, _status string, _amountOfFloors, _currentFloor int ) *Elevator {
	return &Elevator{
		ElevatorID : _elevatorID,
		status : _status,
		amountOfFloors : _amountOfFloors,
		direction : "",
        overweight : false,
		currentFloor : _currentFloor,
	}
	
}

func (e *Elevator) move() {
	for  len(e.floorRequestList) > 0 {
		e.destination = e.floorRequestList[0]
		e.status = "moving"
		if e.currentFloor < e.destination{
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < e.destination{ // A voir 
				e.currentFloor = e.currentFloor + 1
				e.screenDisplay = e.currentFloor
			}
		}else if currentFloor > e.destination{//
			e.direction = "down"
			e.sortFloorList()
			for currentFloor > e.destination{
				currentFloor = currentFloor - 1
				e.screenDisplay = currentFloor
			}
		}
		e.status = "stopped"
		e.floorRequestList = e.floorRequestList[1:0]
	}
	e.status = "idle"
	
}

func (e *Elevator) operateDoors(){
	e.door.status = "opened"
	if e != "overweight" {
		e.door.status = "closing"
		if e != "obstructed" { 
			e.door.status = "closed"
		} else {
			e.operateDoors()
		}
	} else{
		for e == "overweight" {
			//timeout = 5000 
		e.operateDoors()
	}
}

/*func (e *Elevator) contains(s []int, b int) bool {
	for _, a := range s {
	  if a == b {
		return true
	  }
	}
	return false
  }
func (e *Elevator) addNewRequest(requestedFloor int){

    if!contains(e.floorRequestList, requestedFloor){ ///a voir 
        e.floorRequestList.append(e.floorRequestList, requestedFloor) 
    }
	
    if (e.currentFloor < requestedFloor{
        e.direction = "up"
    }
    if (e.currentFloor > requestedFloor){
        e.direction = "down"
    }
}
func (e *Elevator) sortFloorList() {
    if (e.direction == "up"){
		sort.Ints(e.floorRequestList)
    }
    else{ 
        sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList))) //A voir
    }
}*/