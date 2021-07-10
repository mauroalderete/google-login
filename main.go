package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

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

	flag.Parse()

	if *flagVersion {
		fmt.Println("v1.0.0")
		return
	}

	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	var credential string = path.Join(*flagWorkdir, *flagCredential)
	var token string = path.Join(*flagWorkdir, *flagToken)

	fmt.Println("Checking resources availables...")
	_, err := os.Open(credential)
	if err != nil {
		log.Fatalf("Valid credential file is required for to login %v", err)
		return
	}

	var auth auth_service.IAuthService = nil

	if *flagSpreadsheet {
		auth = &auth_service_spreadsheet.AuthServiceSpreadsheet{}
	}

	if auth == nil {
		log.Fatalln("Flag of type access autorization request is required\nExecute 'login -help' to more information")
		return
	}

	auth.Initialize(credential, token, true)

	fmt.Println("Request a new token...")
	err = auth.RequestToken()
	if err != nil {
		log.Fatalf("[Main] Error al solicitar token %v", err)
		return
	}

	fmt.Println("Testing simple auth...")
	err = auth.Initialize(credential, token, true)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar %v", err)
	}

	err = auth.Authenticate()
	if err != nil {
		log.Fatalf("[Main] Error al autenticar %v", err)
	}

	fmt.Println()
	fmt.Println("Login succefull")
}
