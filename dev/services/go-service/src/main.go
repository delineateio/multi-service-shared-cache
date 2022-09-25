package main

func main() {
	s := newServer()
	s.addGet("/healthz", healthzAdapter)
	s.addPost("/increment", writeAdapter)
	s.addGet("/count", readAdapter)
	s.start()
}
