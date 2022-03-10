//  ┌─────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
//  │                                                           │
//  │                                                           │
//  │            ┌──────┤                      ┌───┐      ──────┤
//  └─────────┴──┘      ├──────   ├─────────┴──┘   └──┴─────────┘
//                      └─────────┘
//  ·· A       CLI       ORGONE       ENERGY       ACCUMULATOR ··

package main
import (
   "fmt"
   "math/rand"
   "time"
)
func main() {
   source := rand.New(rand.NewSource(time.Now().UnixNano()))
   shapes := [13]string{"█","▌","▐","▖","▗","▘","▙","▚","▛","▜","▝","▞","▟"}
   columns := 9
   rows := 4
   fmt.Println("┌─────────┐")
   for rows >= 1 {
      column := columns
      fmt.Print("│")
      for column >= 1 {
         shape := source.Intn(12 - 0)
         fmt.Print(shapes[shape])
         column = column - 1
      }
      fmt.Println("│")
      rows = rows - 1
   }
   fmt.Println("└─────────┘")
}