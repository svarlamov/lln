package main

import (
    "os"
    "fmt"
    "math/rand"
)

func main() {
    n := 10000
    maxOutput := 501
    output := "trial,value,running_mean\n"
    currentSet := make([]float64, 0, n)
    sum := 0.0
    for i := 0; i < n; i++ {
        currentSet = append(currentSet, float64(rand.Intn(maxOutput)))
        sum += float64(currentSet[i])
        output += fmt.Sprint(i, ",", currentSet[i], ",", (sum/float64((i+1))), "\n")
    }
    WriteCSVToFile("./output.csv", output)
}

func WriteCSVToFile(filePath string, contents string) {
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.WriteString(contents)
}
