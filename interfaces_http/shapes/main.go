package main

func main() {
	tri := triangle{
		height: 32,
		base:   12,
	}
	sq := square{side: 20}

	printArea(sq)
	printArea(tri)
}
