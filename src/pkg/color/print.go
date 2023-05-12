package color

import "fmt"

func print(color string, args ...interface{}) {
	fmt.Print(color)
	fmt.Print(args...)
	fmt.Print(string(Color.Reset))
	fmt.Println()
}

func PrintRed(args ...interface{}) {
	print(Color.Red, args...)
}

func PrintGreen(args interface{}) {
	print(Color.Green, args)
}

func PrintYellow(args ...interface{}) {
	print(Color.Yellow, args...)
}

func PrintBlue(args ...interface{}) {
	print(Color.Blue, args...)
}

func PrintPurple(args ...interface{}) {
	print(Color.Purple, args...)
}

func PrintCyan(args ...interface{}) {
	print(Color.Cyan, args...)
}

func PrintGray(args ...interface{}) {
	print(Color.Gray, args...)
}

func PrintWhite(args ...interface{}) {
	print(Color.White, args...)
}

func PrintError(args ...interface{}) {
	a := []interface{}{"Error: "}
	a = append(a, args...)

	print(Color.Red, a...)
}

func PrintWarning(args ...interface{}) {
	a := []interface{}{"Warning: "}
	a = append(a, args...)

	print(Color.Yellow, a...)
}
