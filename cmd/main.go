//https://habr.com/ru/articles/1006316/

package main

import "knigga/server"

func main() {
	server := server.NewServer(":3000")
	server.Start()
}
