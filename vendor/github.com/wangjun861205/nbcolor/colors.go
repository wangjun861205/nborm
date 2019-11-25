package nbcolor

type color string

const (
	noColor     color = "\033[0m"
	black       color = "\033[0;30m"
	red         color = "\033[0;31m"
	green       color = "\033[0;32m"
	brown       color = "\033[0;33m"
	blue        color = "\033[0;34m"
	purple      color = "\033[0;35m"
	cyan        color = "\033[0;36m"
	lightGray   color = "\033[0;37m"
	darkGray    color = "\033[1;30m"
	lightRed    color = "\033[1;31m"
	lightGreen  color = "\033[1;32m"
	yellow      color = "\033[1;33m"
	lightBlue   color = "\033[1;34m"
	lightPurple color = "\033[1;35m"
	lightCyan   color = "\033[1;36m"
	white       color = "\033[1;37m"
)
