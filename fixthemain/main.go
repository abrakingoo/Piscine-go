package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	ptrDoor.state = "OPEN"
	return true
}

func IsDoorOpen(door Door) {
	PrintStr("is the Door opened ?\n")
}

func IsDoorClose(ptrDoor *Door) {
	PrintStr("is the Door closed ?\n")
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = "CLOSE"
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	IsDoorClose(door)
	IsDoorOpen(*door)
	CloseDoor(door)
}
