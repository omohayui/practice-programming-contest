package main

import (
	"math"
	"testing"
)

func TestGetTotalCharge(t *testing.T) {
	cases := []struct {
		name         string
		using        int64
		pricingTable []*TieredPricing
		exp          int64
	}{
		{
			name:  "normal",
			using: 16,
			pricingTable: []*TieredPricing{
				{Usage: DataUsage{Min: 0, Max: 5}, Price: 0},
				{Usage: DataUsage{Min: 6, Max: 10}, Price: 10},
				{Usage: DataUsage{Min: 11, Max: 50}, Price: 8},
				{Usage: DataUsage{Min: 51, Max: 100}, Price: 5},
				{Usage: DataUsage{Min: 101, Max: 1000}, Price: 2},
				{Usage: DataUsage{Min: 1001}, Price: 1},
			},
			exp: 98,
		},
		{
			name:  "without free usage",
			using: 16,
			pricingTable: []*TieredPricing{
				{Usage: DataUsage{Min: 0, Max: 5}, Price: 1},
				{Usage: DataUsage{Min: 6, Max: 10}, Price: 10},
				{Usage: DataUsage{Min: 11, Max: 50}, Price: 8},
				{Usage: DataUsage{Min: 51, Max: 100}, Price: 5},
				{Usage: DataUsage{Min: 101, Max: 1000}, Price: 2},
				{Usage: DataUsage{Min: 1001}, Price: 1},
			},
			exp: 103,
		},
		{
			name:  "just 1000GB",
			using: 1000,
			pricingTable: []*TieredPricing{
				{Usage: DataUsage{Min: 0, Max: 5}, Price: 0},
				{Usage: DataUsage{Min: 6, Max: 10}, Price: 10},
				{Usage: DataUsage{Min: 11, Max: 50}, Price: 8},
				{Usage: DataUsage{Min: 51, Max: 100}, Price: 5},
				{Usage: DataUsage{Min: 101, Max: 1000}, Price: 2},
				{Usage: DataUsage{Min: 1001}, Price: 1},
			},
			exp: 2420,
		},
		{
			name:  "use bigint",
			using: math.MaxInt32 + 1,
			pricingTable: []*TieredPricing{
				{Usage: DataUsage{Min: 0, Max: 5}, Price: 0},
				{Usage: DataUsage{Min: 6, Max: 10}, Price: 10},
				{Usage: DataUsage{Min: 11, Max: 50}, Price: 8},
				{Usage: DataUsage{Min: 51, Max: 100}, Price: 5},
				{Usage: DataUsage{Min: 101, Max: 1000}, Price: 2},
				{Usage: DataUsage{Min: 1001}, Price: 1},
			},
			exp: 2420 + math.MaxInt32 + 1 - 1000,
		},
	}

	for _, v := range cases {
		if res := GetTotalCharge(v.using, v.pricingTable); res != v.exp {
			t.Fatalf("TestGetTotalCharge failed. case: %s, exp: %d, res: %d", v.name, v.exp, res)
		}
	}
}
