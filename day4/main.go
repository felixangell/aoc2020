package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var entryKeys = map[string]bool{"byr": true, "iyr": true, "eyr": true, "hgt": true, "hcl": true, "ecl": true, "pid": true, "cid": true}
var eyeCol = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

type record map[string]string

func isValidValue(key, val string) (bool, string) {
	if len(val) == 0 {
		return false, key
	}

	switch key {
	case "hgt":
		m, _ := regexp.MatchString(`[0-9]+(in|cm)`, val)
		heightStr, unit := val[:len(val)-2], val[len(val)-2:]
		if !m {
			return false, key
		}

		height, _ := strconv.ParseInt(heightStr, 10, 32)

		switch unit {
		case "cm":
			return height >= 150 && height <= 193, key
		case "in":
			return height >= 59 && height <= 76, key
		}
		return m, key

	case "byr":
		iv, _ := strconv.ParseInt(val, 10, 64)
		return iv >= 1920 && iv <= 2002, key
	case "iyr":
		iv, _ := strconv.ParseInt(val, 10, 64)
		return iv >= 2010 && iv <= 2020, key
	case "eyr":
		iv, _ := strconv.ParseInt(val, 10, 64)
		return iv >= 2020 && iv <= 2030, key

	case "hcl":
		m, _ := regexp.MatchString(`#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`, val)
		return m && len(val) == 7, key
	case "ecl":
		_, ok := eyeCol[val]
		return ok, key
	case "pid":
		m, _ := regexp.MatchString(`[0-9]+`, val)
		return m && len(val) == 9, key

	case "cid":
		return true, ""
	}

	return false, key
}

func isValidRecord(r record) bool {
	missing := 0
	for key := range entryKeys {
		if _, ok := r[key]; !ok {
			if strings.Compare(key, "cid") == 0 {
				continue
			}
			missing++
		}
	}

	for k, v := range r {
		if ok, bk := isValidValue(k, v); !ok {
			log.Println("bad record", bk, r[bk])
			return false
		}
	}

	return missing == 0
}

func countValidRecords(lines []string) (validRecords int) {
	lastRecord := record{}
	for _, line := range lines {
		if len(line) == 0 {
			if isValidRecord(lastRecord) {
				validRecords++
			}
			lastRecord = record{}
			continue
		}

		fields := strings.Fields(line)
		for _, field := range fields {
			parts := strings.Split(field, ":")
			key, val := parts[0], parts[1]
			lastRecord[key] = val
		}
	}
	return validRecords
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// lol
	entries := strings.ReplaceAll(string(data), "\n", "!")

	lines := strings.Split(entries, "!")
	goodRecords := countValidRecords(lines)
	fmt.Println(goodRecords, "good records")
}
