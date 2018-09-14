package main

import "fmt"
import collector "github.com/sudhir/stockticker/datacollectors"

func main() {
	fmt.Println("hello stock ticker")
	collector.CollectData()
}
