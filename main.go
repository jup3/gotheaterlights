package main

import (
  "machine"
  "time"
)

func main() {
  // config PINs
  lone := machine.Pin(machine.D13)
  lone.Configure(machine.PinConfig{Mode: machine.PinOutput})
  lone.Low()
  ltwo := machine.Pin(machine.D12)
  ltwo.Configure(machine.PinConfig{Mode: machine.PinOutput})
  ltwo.Low()
  lthree := machine.Pin(machine.D7)
  lthree.Configure(machine.PinConfig{Mode: machine.PinOutput})
  lthree.Low()
  mo := machine.Pin(machine.D2)
  mo.Configure(machine.PinConfig{Mode: machine.PinInput})
  for {
    if mo.Get() == true {
        for i := 0; i <= 30; i++ {
          // marquee
          blink(lone)
          blink(ltwo)
          blink(lthree)
        }
    }
  }
}

func blink(led machine.Pin) {
  led.High()
  time.Sleep(time.Millisecond * 150)
  led.Low()
}
