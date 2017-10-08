package main

import (
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/utils"
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/backpack"
	"strconv"
	"strings"
)

func main() {
	fileData, err := util.ReadFromFile("C:\\workspace\\bin\\testcase01.txt")
	if err != nil {
		panic(err)
	}
	backPackWeight, err := strconv.ParseFloat(fileData[0], 64)

	items := make([]backpack.Item, len(fileData)-1)

	for i, val := range fileData[1:] {
		split := strings.Split(val, " ")
		weight, _ := strconv.ParseFloat(split[0], 64)
		price, _ := strconv.ParseFloat(split[1], 64)
		items[i] = *backpack.NewItem(i, weight, price)
	}

	population := backpack.GetInitialPopulation(4, len(items))
	backPack := backpack.NewBackPack(backPackWeight, items, population)

	backPack.PrintSolution()
}
