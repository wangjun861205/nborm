package nbcolor

import "fmt"

func addColor(content interface{}, color color) string {
	return fmt.Sprintf("%s%v%s", color, content, noColor)
}

func Black(content interface{}) string {
	return addColor(content, black)
}

func Red(content interface{}) string {
	return addColor(content, red)
}

func Green(content interface{}) string {
	return addColor(content, green)
}

func Brown(content interface{}) string {
	return addColor(content, brown)
}

func Blue(content interface{}) string {
	return addColor(content, blue)
}

func Purple(content interface{}) string {
	return addColor(content, purple)
}

func Cyan(content interface{}) string {
	return addColor(content, cyan)
}

func LightGray(content interface{}) string {
	return addColor(content, lightGray)
}

func DarkGray(content interface{}) string {
	return addColor(content, darkGray)
}

func LightRed(content interface{}) string {
	return addColor(content, lightRed)
}

func LightGreen(content interface{}) string {
	return addColor(content, lightGreen)
}

func Yellow(content interface{}) string {
	return addColor(content, yellow)
}

func LightBlue(content interface{}) string {
	return addColor(content, lightBlue)
}

func LightPurple(content interface{}) string {
	return addColor(content, lightPurple)
}

func LightCyan(content interface{}) string {
	return addColor(content, lightCyan)
}

func White(content interface{}) string {
	return addColor(content, white)
}
