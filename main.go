package main

import (
  "fmt"
  "syscall/js"
  "strconv"
  "encoding/json"

  "github.com/brettfischl/stats/proportions"
)

func sampleProportion(this js.Value, args []js.Value) interface{} {
  trialsValue := args[0].String()
  successesValue := args[1].String()
  compareValue := args[2].String()

  trials, _ := strconv.ParseFloat(trialsValue, 64)
	successes, _ := strconv.ParseFloat(successesValue, 64)
  compare, _ := strconv.ParseFloat(compareValue, 64)

  prop := proportions.NewSampleProportion(trials, successes, compare)
  prop.Zscores()

  proportion, err := json.Marshal(prop)
  if err != nil {
    fmt.Println(err)
    return nil
  }

  obj := js.Global().Get("JSON").Call("parse", string(proportion))

  return obj
}

func compareProportion(this js.Value, args []js.Value) interface{} {
  trialsValue1 := args[0].String()
	successesValue1 := args[1].String()
  trialsValue2 := args[2].String()
	successesValue2 := args[3].String()

  trials1, _ := strconv.ParseFloat(trialsValue1, 64)
	successes1, _ := strconv.ParseFloat(successesValue1, 64)
  trials2, _ := strconv.ParseFloat(trialsValue2, 64)
	successes2, _ := strconv.ParseFloat(successesValue2, 64)

  p1 := proportions.NewSampleProportion(
    trials1,
    successes1,
    0,
  )

  p2 := proportions.NewSampleProportion(
    trials2,
    successes2,
    0,
  )

  differenceOfProportions := proportions.DifferenceOfProportions{
    S1: p1,
    S2: p2,
  }

  differenceOfProportions.Test()

  proportion, err := json.Marshal(differenceOfProportions)
  if err != nil {
    fmt.Println(err)
    return nil
  }

  obj := js.Global().Get("JSON").Call("parse", string(proportion))

  return obj
}

func sayHi(this js.Value, args []js.Value) interface{} {
	fmt.Println("Hi!")
  return nil
}

func registerCallbacks() {
  js.Global().Set("computeSampleProportion", js.FuncOf(sampleProportion))
	js.Global().Set("sayHi", js.FuncOf(sayHi))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
