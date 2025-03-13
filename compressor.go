package bf

import (
	"slices"
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

	input = strings.ReplaceAll(input, " ", "")  // remove whitespace
	input = strings.ReplaceAll(input, "\n", "") // remove newlines

	sb := strings.Builder{}

	var chart c_num_chart
	chart.c = input[0]

	var charts []c_num_chart

	bf_syntax := []string{">", "<", "+", "-", ".", ",", "[", "]"}

	for i, c := range input {
		if !slices.Contains(bf_syntax, string(c)) {
			continue
		}

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
