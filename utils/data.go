package utils

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func CleanData(data string) string {
	var sb []string
	heads := strings.Split(data, ";")
	for _, head := range heads {
		split := strings.Split(head, ",")
		start, err := strconv.ParseFloat(split[1], 32)
		end, err := strconv.ParseFloat(split[2], 32)
		if err != nil {
			log.Println("Error parsing data ", err)
		}
		if strings.Count(head, ".") > 0 {
			for floatItem := start; floatItem <= end; {
				if floatItem != math.Round(floatItem) {
					sb = append(sb, fmt.Sprintf("%.1f", floatItem))
				} else {
					sb = append(sb, fmt.Sprintf("%.1f", floatItem))
				}
				floatItem = math.Round((floatItem+0.1)*10) / 10
			}
		} else {
			for floatItem := start; floatItem <= end; floatItem++ {
				sb = append(sb, fmt.Sprintf("%.0f", floatItem))
			}
		}
	}
	return strings.Join(sb, ",")
}

func XmlToJson(data string) string {
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>` + data)
	json, err := xj.Convert(xml)
	if err != nil {
		log.Println("Error parsing data ", err)
	}

	return json.String()
}
