package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"gitlab.com/rayquen-google/golang/auth/auth_service"
	"gitlab.com/rayquen-google/golang/auth/auth_service_spreadsheet"
)

func main() {
	var flagWorkdir = flag.String("workdir", "./", "folder to read credential and save token. For default is './'")
	var flagCredential = flag.String("credential", "credential.json", "Filename to read credential. For default is 'credential.json'")
	var flagToken = flag.String("token", "token.json", "Filename to save token when login is succefull. For default is 'token.json' ")
	var flagVersion = flag.Bool("version", false, "Version of alMercadito API RESTful Login")
	var flagHelp = flag.Bool("help", false, "Show help information")
	var flagSpreadsheet = flag.Bool("spreadsheet", false, "Request access autorization to Google Spreadsheets")
	var flagNoColor = flag.Bool("no-color", false, "Disable color output")

	if *flagNoColor {
		color.NoColor = true
	}

	flag.Parse()

	if *flagVersion {
		fmt.Println("v1.0.0")
		return
	}

	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	writeTitle("rayquen-google/login")

	var credential string = path.Join(*flagWorkdir, *flagCredential)
	var token string = path.Join(*flagWorkdir, *flagToken)

	writeInfo("Checking resources availables... ")
	_, err := os.Open(credential)
	if err != nil {
		writeFailf("Valid credential file is required for to login %v", err)
		return
	}
	writeConfirm()

	var auth auth_service.IAuthService = nil

	if *flagSpreadsheet {
		auth = &auth_service_spreadsheet.AuthServiceSpreadsheet{}
	}

	if auth == nil {
		writeErrorf("Flag of type access autorization request is required. Execute 'login -help' to more information")
		return
	}

	auth.Initialize(credential, token, true)

	writeInfo("Request a new token... ")
	err = auth.RequestToken()
	if err != nil {
		writeFailf("[Main] Error al solicitar token %v", err)
		return
	}
	writeConfirm()

	writeInfo("Testing simple initialize... ")
	err = auth.Initialize(credential, token, true)
	if err != nil {
		writeFailf("[Main] Error al inicializar %v", err)
		return
	}
	writeConfirm()

	writeInfo("Testing simple authentication... ")
	err = auth.Authenticate()
	if err != nil {
		writeFailf("[Main] Error al autenticar %v", err)
		return
	}
	writeConfirm()
}

func writeErrorf(format string, a ...interface{}) {
	brackets := color.New(color.FgWhite)
	tag := color.New(color.FgRed)
	text := color.New(color.FgHiRed)

	brackets.Print("[")
	tag.Print("ERROR")
	brackets.Print("]")
	text.Printf(" "+format, a...)
}

func writeInfo(message string) {
	text := color.New(color.FgWhite)
	text.Print(message)
}

func writeTitle(message string) {
	color.Cyan(message)
	color.Cyan("")
}

func writeConfirm() {
	brackets := color.New(color.FgWhite)
	tag := color.New(color.FgGreen)

	brackets.Print("[")
	tag.Print("OK")
	brackets.Println("]")
}

func writeFailf(format string, a ...interface{}) {
	brackets := color.New(color.FgWhite)
	tag := color.New(color.FgRed)
	text := color.New(color.FgHiRed)

	brackets.Print("[")
	tag.Print("Fail")
	brackets.Println("]")
	text.Printf("\t"+format, a...)
}
