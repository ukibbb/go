package main

func main() {
	s := NewServer(":3000")

	go s.Start()

	select {}

}
