package server

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/labstack/echo"
)

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

func Start() {
	androidTemplate, err := template.New("android.txt").ParseFiles("presets/android.txt")
	if err != nil {
		fmt.Printf("android preset didn't load: %v\n", err.Error())
	}

	desktopTemplate, err := template.New("desktop.txt").ParseFiles("presets/desktop.txt")
	if err != nil {
		fmt.Printf("desktop preset didn't load: %v\n", err.Error())
	}

	e := echo.New()

	e.Static("/assets", "client/dist/assets")

	e.File("/", "client/dist/index.html")

	e.GET("/android", func(c echo.Context) error {
		theme := Theme{
			Name:        c.QueryParam("name"),
			Background0: c.QueryParam("background_0"),
			Background1: c.QueryParam("background_1"),
			Background2: c.QueryParam("background_2"),
			Accent1:     c.QueryParam("accent_1"),
			Accent2:     c.QueryParam("accent_2"),
			Text0:       c.QueryParam("text_0"),
			Text1:       c.QueryParam("text_1"),
		}

		var b bytes.Buffer
		err = androidTemplate.Execute(&b, theme)
		if err != nil {
			fmt.Printf("android preset didn't execute: %v\n", err.Error())
		}
		result := b.String()

		decimal := c.QueryParam("decimal") == "true"
		if decimal {
			result = replaceAllColorsHexToDec(result)
			result = strings.ReplaceAll(result, " ", "")
			result = strings.ReplaceAll(result, ":", "=")
		}

		return c.String(200, result)
	})

	e.GET("/desktop", func(c echo.Context) error {
		theme := Theme{
			Name:        c.QueryParam("name"),
			Background0: c.QueryParam("background_0"),
			Background1: c.QueryParam("background_1"),
			Background2: c.QueryParam("background_2"),
			Accent1:     c.QueryParam("accent_1"),
			Accent2:     c.QueryParam("accent_2"),
			Text0:       c.QueryParam("text_0"),
			Text1:       c.QueryParam("text_1"),
		}

		var b bytes.Buffer
		err = desktopTemplate.Execute(&b, theme)
		if err != nil {
			fmt.Printf("desktop preset didn't execute: %v\n", err.Error())
		}
		result := b.String()

		decimal := c.QueryParam("decimal") == "true"
		if decimal {
			result = replaceAllColorsHexToDec(result)
			result = strings.ReplaceAll(result, " ", "")
			result = strings.ReplaceAll(result, ":", "=")
		}

		return c.String(200, result)
	})

	e.Logger.Fatal(e.Start(":8090"))
}
