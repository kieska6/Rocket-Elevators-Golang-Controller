package main

type Battery struct {
	ID                        int
	status                    string
	amountOfFloors            int
	amountOfColumns           int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButton

}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	b := Battery{}
	b.ID = id
	b.status = status
	b.amountOfFloors = amountOfFloors
	b.amountOfColumns = amountOfColumns
	b.amountOfBasements = amountOfBasements
	b.columnsList = []Column{}
	b.floorRequestButtonsList = []FloorRequestButton{}

	if amountOfBasements > 0 {
		b.createFloorRequestButtons(amountOfBasements)
		b.createBasementColumn(b.amountOfBasements, amountOfElevatorPerColumn)
		amountOfColumns--
	}

	b.createFloorRequestButtons(amountOfFloors)
	b.createColumns(amountOfColumns, amountOfFloors, amountOfElevatorPerColumn)

	return b
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	var newcolumn Column = b.columnsList[0]
	for _, column := range c.columnsList {
		if column.servedFloorsList.Contains(_requestedFloor){
			newcolumn = column
		}
	}
	return newcolumn 
	
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (Column, Elevator) {
	//requestedFloor int = Int32.Parse(_requestedFloor);
	var column = b.findBestColumn(requestedFloor) //A voir 
	bestElevator Elevator = column.findElevator(1, _direction) // The floor is always 1 because that request is always made from the lobby.
	bestElevator.addNewRequest(1)
	bestElevator.move()
	bestElevator.addNewRequest(requestedFloor)
	bestElevator.move()
	return column, bestElevator
}
func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int){
	servedFloors := []int{}
	floor := -1

	for i := 0; i < amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}

	column Column = new Column(_columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, column)
	columnID++
}
func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int,_amountOfBasements int,_amountOfElevatorPerColumn int){
	amountOfFloorsPerColumn := int(math.Ceil(float64(b.amountOfFloors / amountOfElevatorPerColumn)))
	floor := 1

	for i := 0; i < amountOfColumns; i++ {
		servedFloors := []int{}
		for v := 0; v < amountOfFloorsPerColumn; v++ {

			if floor <= b.amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column Column = new Column(_columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, column)
		columnID++
	}

}
func (b *Battery) createFloorRequestButtons(_amountOfFloors int){
	buttonFloor := 1

	for i := 0; i < amountOfFloors; i++ {
		floorRequestButtons := FloorRequestButton{floorRequestButtonID, "OFF", buttonFloor, "UP"}
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButtons)
		buttonFloor++
		floorRequestButtonID++
	}
}
func (b *Battery) createBasementFloorRequestButtons (_amountOfBasements int){
	buttonFloor := -1

	for i := 0; i < amountOfBasements; i++ {
		floorRequestButtons := FloorRequestButton{floorRequestButtonID, "OFF", buttonFloor, "DOWN"}
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButtons)
		buttonFloor--
		floorRequestButtonID++
	}

}

