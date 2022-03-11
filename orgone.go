//  ┌─────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
//  │                                                           │
//  │                                                           │
//  │            ┌──────┤                      ┌───┐      ──────┤
//  └─────────┴──┘      ├──────   ├─────────┴──┘   └──┴─────────┘
//                      └─────────┘
//  ·· A       CLI       ORGONE       ENERGY       ACCUMULATOR ··

//  toDo: user specified orgone quantity
//        colors
//        fluffier random seed

package main

import (
   "fmt"
   "math/rand"
   "time"
)

// generates a row of orgone
func orgone() string {
   // a seed for rand
   var source = rand.New(rand.NewSource(time.Now().UnixNano()))
   // a set of shapes
   shapes := [13]string{"█","▌","▐","▖","▗","▘","▙","▚","▛","▜","▝","▞","▟"}
   // a number of columns
   var columns int = 9
   // clear any leftover energy
   var energy string = ""
   // accumulate new energy
   for columns >= 1 {
      shape := source.Intn(12 - 0)
      energy = energy + shapes[shape]
      columns = columns - 1
   }
   return energy
}

// main func
func main() {
   var rows int = 4
   fmt.Println("A quantity of orgone was accumulated.")
   fmt.Println("┌─────────┐")
   for rows >= 1 {
      fmt.Print("│")
      var energy string = orgone()
      fmt.Print(energy)
      fmt.Println("│")
      rows = rows - 1
   }
   fmt.Println("└─────────┘")
}