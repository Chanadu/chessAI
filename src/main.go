package main

func main() {
	RaylibCreateWindow()
	defer RaylibCloseWindow()

	PreLoop()

	for !RaylibWindowShouldClose() {
		MainGameLoop()
	}
}
