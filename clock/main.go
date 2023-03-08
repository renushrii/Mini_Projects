package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type digit [5]string

	zero := digit{
		"🟫🟫🟫",
		"🟫  🟫",
		"🟫  🟫",
		"🟫  🟫",
		"🟫🟫🟫",
	}

	one := digit{
		"  🟫  ",
		"🟫🟫  ",
		"  🟫  ",
		"  🟫  ",
		"🟫🟫🟫",
	}

	two := digit{
		"🟫🟫🟫",
		"    🟫",
		"🟫🟫🟫",
		"🟫    ",
		"🟫🟫🟫",
	}

	three := digit{
		"🟫🟫🟫",
		"    🟫",
		"🟫🟫🟫",
		"    🟫",
		"🟫🟫🟫",
	}

	four := digit{
		"🟫  🟫",
		"🟫  🟫",
		"🟫🟫🟫",
		"    🟫",
		"    🟫",
	}

	five := digit{
		"🟫🟫🟫",
		"🟫   ",
		"🟫🟫🟫",
		"    🟫",
		"🟫🟫🟫",
	}

	six := digit{
		"🟫🟫🟫",
		"🟫    ",
		"🟫🟫🟫",
		"🟫  🟫",
		"🟫🟫🟫",
	}

	seven := digit{
		"🟫🟫🟫",
		"    🟫",
		"    🟫",
		"    🟫",
		"    🟫",
	}

	eight := digit{
		"🟫🟫🟫",
		"🟫  🟫",
		"🟫🟫🟫",
		"🟫  🟫",
		"🟫🟫🟫",
	}

	nine := digit{
		"🟫🟫🟫",
		"🟫  🟫",
		"🟫🟫🟫",
		"    🟫",
		"🟫🟫🟫",
	}

	colon := digit{
		"    ",
		" 🟫 ",
		"    ",
		" 🟫 ",
		"    ",
	}

	digits := [...]digit{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]digit{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "    "
				}
				fmt.Print(next, "  ")
			}

			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
