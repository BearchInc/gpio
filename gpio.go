package main

import "github.com/stianeikeland/go-rpio"
import "time"
import "fmt"

func main() {
  if err := rpio.Open(); err != nil {
    panic(err)
  }

  pin := rpio.Pin(21)
  pin.Input()

  for {
    time.Sleep(1 * time.Second)
    fmt.Print(pin.Read())
  }
}
