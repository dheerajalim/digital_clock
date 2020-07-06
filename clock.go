package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clock() {

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	seperator := placeholder{
		"  ",
		"░░",
		"  ",
		"░░",
		"  ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine}

	clearscreen()
	for {
		clearscreen()

		currenttiime := time.Now()
		hour, minute, second := currenttiime.Hour(), currenttiime.Minute(), currenttiime.Second()

		clock := [...]placeholder{
			digits[hour/10],
			digits[hour%10],
			seperator,
			digits[minute/10],
			digits[minute%10],
			seperator,
			digits[second/10],
			digits[second%10],
		}
		for i := range digits[0] {
			for _, j := range clock {

				current := j[i]
				if j == seperator && second%2 == 0 {
					current = "  "
				}
				fmt.Printf("%s  ", current)
			}
			fmt.Println()

		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}

func clearscreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

}
func main() {

	clock()
}
