package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	floor     int
	direction string
}

func NewCallButton(_floor int, _direction string) *CallButton {
	return &CallButton{
		floor:     _floor,
		direction: _direction,
	}

}
