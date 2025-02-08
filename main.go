package main

func main() {
	s := NewServer(":8000")

	go s.Start()

	select {}

}
