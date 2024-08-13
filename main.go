package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("./measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close()

	scanner := bufio.NewScanner(measurements)

	dados := make(map[string]Measurement)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";") //retorna o index onde encontrar ;
		location := rawData[:semicolonIndex]          //le a string da posição zero até semicolonIndex - 1
		rawTemp := rawData[semicolonIndex+1:]         //le a string da posição semicolonIndex + 1 até o final

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}

	// for name, measurement := range dados {
	// 	fmt.Printf("%s: %#+v\n", name, measurement)
	// }

	locations := make([]string, 0, len(dados))

	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")
}
