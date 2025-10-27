package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

const ShellToUse = "bash"

func main() {
	fmt.Println("starting ...")

	e = echo.New()

	e.GET(
		"/",
		HandelFunc(),
	)

	log.Fatal(e.Start(":8888"))
}

type Input struct {
	Str string `json:"str,omitempty"`
}

func HandelFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(Input)
		c.Bind(req)

		// cmd := fmt.Sprintf("ls | grep %s", req.Str)

		// println(cmd)

		out, errOut, err := ShellOut("ls -ltr | grep " + req.Str)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errOut)

		if err != nil {
			println(err.Error())
			return c.JSON(
				http.StatusBadRequest,
				"FAIL",
			)
		}

		return c.JSON(
			http.StatusOK,
			string(out),
		)
	}
}

func ShellOut(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
