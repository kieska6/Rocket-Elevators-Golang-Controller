package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	floor     int
	direction string
	id        int
	status    string
}

func NewFloorRequestButton(_floor int, _direction string) *FloorRequestButton {
	return &FloorRequestButton{
		floor:     _floor,
		direction: _direction,
	}

}
