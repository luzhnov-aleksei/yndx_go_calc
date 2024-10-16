package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() { 
	var (
		stat        = make(map[int]float64)
		operator    string
		id          int
		temp        float64
		sumTemp     float64
		parts       []string
	)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		parts = strings.Split(input, " ")

		operator = parts[0]
        if operator == "!" {
            break
        } else if operator == "?"{
			if len(stat) == 0 {
				fmt.Printf("0.000000000\n")
			} else {
				average := sumTemp / float64(len(stat))
				fmt.Printf("%.9f\n", average)
			}
			continue
		}


		id, err = strconv.Atoi(parts[1])
		if err != nil {
			continue 
		}

		if operator == "-" || operator == "−" { 
			if val, exists := stat[id]; exists {
				sumTemp -= val
				delete(stat, id)
			}
			continue
		}

		
		if len(parts) < 3 {
			continue
		}
		temp, err = strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue 
		}

		if operator == "+" {
			if _, exists := stat[id]; exists {
				sumTemp -= stat[id]
			}
			stat[id] = temp
			sumTemp += temp
		} else if operator == "~" {
			if val, exists := stat[id]; exists {
				sumTemp -= val
			}
			stat[id] = temp
			sumTemp += temp
		}	
	}
}
