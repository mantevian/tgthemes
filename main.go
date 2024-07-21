package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

var themeVars = [...]string{"name", "black_0", "black_1", "black_2", "accent_1", "accent_2", "gray", "white"}

var androidInput string
var desktopInput string

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("%v not found", path)
	}

	return string(content), nil
}

func colorHexToDec(hex string) string {
	fixedHex := fmt.Sprintf("%v%v", hex[7:9], hex[1:7])

	result, _ := strconv.ParseInt(fixedHex, 16, 64)

	return fmt.Sprint(int32(result))
}

func replaceAllColorsHexToDec(input string) string {
	regex, _ := regexp.Compile("#[0-9A-Fa-f]{8}")

	return regex.ReplaceAllStringFunc(input, func(s string) string {
		return colorHexToDec(s)
	})
}

func main() {
	androidInput, _ = readFile("./android.txt")
	desktopInput, _ = readFile("./desktop.txt")

	e := echo.New()

	e.Static("/static", "./static")

	e.File("/", "./static/index.html")

	e.GET("/android", func(c echo.Context) error {
		result := androidInput
		decimal := false
		if c.QueryParam("decimal") == "true" {
			decimal = true
		}

		for _, v := range themeVars {
			result = strings.ReplaceAll(result, fmt.Sprintf("{%v}", v), c.QueryParam(v))
		}

		if decimal {
			result = replaceAllColorsHexToDec(result)
			result = strings.ReplaceAll(result, " ", "")
			result = strings.ReplaceAll(result, ":", "=")
		}

		return c.String(200, result)
	})

	e.GET("/desktop", func(c echo.Context) error {
		result := desktopInput
		decimal := false
		if c.QueryParam("decimal") == "true" {
			decimal = true
		}

		for _, v := range themeVars {
			result = strings.ReplaceAll(result, fmt.Sprintf("{%v}", v), c.QueryParam(v))
		}

		if decimal {
			result = replaceAllColorsHexToDec(result)
			result = strings.ReplaceAll(result, " ", "")
			result = strings.ReplaceAll(result, ":", "=")
		}

		return c.String(200, result)
	})

	e.Logger.Fatal(e.Start(":8090"))
}
