package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		str := strings.TrimSpace(s.Text())
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("Please set integer.")
			log.Fatal(err)
		}

		fmt.Printf("Using: %dGB, Total Charge: $%d\n", i, GetTotalCharge(i, samplePricingTable))
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

var samplePricingTable = []*TieredPricing{
	{Usage: DataUsage{Min: 0, Max: 5}, Price: 0},
	{Usage: DataUsage{Min: 6, Max: 10}, Price: 10},
	{Usage: DataUsage{Min: 11, Max: 50}, Price: 8},
	{Usage: DataUsage{Min: 51, Max: 100}, Price: 5},
	{Usage: DataUsage{Min: 101, Max: 1000}, Price: 2},
	{Usage: DataUsage{Min: 1001}, Price: 1},
}

type TieredPricing struct {
	Usage DataUsage
	Price int64
}

type DataUsage struct {
	Min int64
	Max int64
}

func GetTotalCharge(using int64, pricingTable []*TieredPricing) int64 {
	var total int64
	for _, p := range pricingTable {
		min := p.Usage.Min
		// 0 data usage not include price but the others include itself price
		if p.Usage.Min != 0 {
			min--
		}
		if using >= p.Usage.Max && p.Usage.Max != 0 {
			total += (p.Usage.Max - min) * p.Price
		} else if using > min && (using < p.Usage.Max || p.Usage.Max == 0) {
			total += (using - min) * p.Price
		}
	}
	return total
}
