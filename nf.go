package gutils

import (
	"fmt"
	"strconv"
)

func FormatFloat(f float64, valFmt string) string {
	_, err := strconv.Atoi(valFmt)
	if err == nil {
		return fmt.Sprintf("%."+valFmt+"f", f)
	} else {
		if len(valFmt) == 2 {
			if valFmt[1] == '%' {
				n, err := strconv.Atoi(string(valFmt[0]))
				if err != nil {
					n = 2
				}
				return fmt.Sprintf("%."+strconv.Itoa(n)+"f", f) + "%"
			}
		}
		//default %.2f
		return fmt.Sprintf("%.2f", f)

	}
}
