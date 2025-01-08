package bf

import (
	"strconv"
	"strings"
)

type c_num_chart struct {
	c   byte
	num int
}

func CompressProgram(input string) string {
	if len(input) == 0 {
		return ""
	}

	input = strings.ReplaceAll(input, " ", "") // remove whitespace

	sb := strings.Builder{}

	var chart c_num_chart
	chart.c = input[0]

	var charts []c_num_chart

	for i, c := range input {
		chart.num++

		if i == len(input)-1 || byte(input[i+1]) != byte(c) {
			charts = append(charts, chart)

			if i != len(input)-1 {
				chart = c_num_chart{c: input[i+1]}
			}
		}
	}

	for _, chart := range charts {
		sb.WriteByte(chart.c)
		if chart.num > 1 {
			sb.WriteString(strconv.Itoa(chart.num))
		}
	}

	return sb.String()
}
