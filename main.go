package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		home := os.Getenv("HOME")
		err = godotenv.Load(home + "/.config/envault/.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	if os.Getenv(key) == "" {
		log.Fatalf("Error " + key + " not set in .env file")
	}
	return os.Getenv(key)
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func addToFile(name string, content string) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
}

func createFile(path string, name string, content string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	f, err := os.Create(path + name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
}

// TODO: create configuration per repository
func configureVaultEnv(debug bool) {
	home := os.Getenv("HOME")
	println(color.Colorize(color.Blue, "Please enter the Vault address:"))
	vault_addr := getInput()
	println(color.Colorize(color.Blue, "Please enter the Vault Token:"))
	vault_token := getInput()
	println(color.Colorize(color.Blue, "Please enter the Vault KV mount:"))
	vault_mount := getInput()
	println(color.Colorize(color.Blue, "Please enter the Vault KV mount database:"))
	vault_db := getInput()

	createFile(home+"/.config/envault/", ".env", "VAULT_ADDR="+vault_addr+"\nVAULT_TOKEN="+vault_token+"\nVAULT_MOUNT="+vault_mount+"\nVAULT_DB="+vault_db+"\n")
}

func list(debug bool, outputType string) {
	mount := goDotEnvVariable("VAULT_MOUNT")
	db := goDotEnvVariable("VAULT_DB")
	outputOptions := [4]string{"table", "json", "yaml", "pretty"}

	if !itemExists(outputOptions, outputType) {
		outputType = "json"
	}

	command := "vault kv get -mount=" + mount + " -field=data -format " + outputType + " " + db
	if debug {
		fmt.Println(command)
	}
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()

	if err != nil {
		panic(err)
	}

	output := string(data)
	fmt.Println(output)
}

func status(debug bool) {
	command := "vault status"
	if debug {
		fmt.Println(command)
	}
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()

	if err != nil {
		panic(err)
	}

	output := string(data)
	fmt.Println(output)
}

func set(debug bool, key string, val string) {
	mount := goDotEnvVariable("VAULT_MOUNT")
	db := goDotEnvVariable("VAULT_DB")
	command := "vault kv get -mount=" + mount + " -field=" + key + " " + db
	if debug {
		fmt.Println(command)
	}

	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()

	if err == nil {
		var approve string
		println(color.Colorize(color.Blue, "Looks like key: "+key+" exists, to Cancel press C or Enter to continue: "))
		fmt.Scanln(&approve)

		if strings.ToLower(approve) == "c" {
			os.Exit(1)
		}
	}

	command = "vault kv patch -mount=" + mount + " " + db + " " + key + "=" + val

	if debug {
		fmt.Println(command)
	}

	parts = strings.Fields(command)
	data, err = exec.Command(parts[0], parts[1:]...).Output()

	if err != nil {
		panic(err)
	}
	output := string(data)

	if debug {
		fmt.Println(output)
	}
}

func run(debug bool) {
	mount := goDotEnvVariable("VAULT_MOUNT")
	db := goDotEnvVariable("VAULT_DB")
	command := "vault kv get -mount=" + mount + " -field=data -format json " + db
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()

	if err != nil {
		panic(err)
	}

	output := string(data)
	if debug {
		fmt.Println(output)
	}

	f := map[string]interface{}{}
	if err := json.Unmarshal([]byte(output), &f); err != nil {
		panic(err)
	}

	for key, value := range f {
		str := fmt.Sprintf("%v", value)
		os.Setenv(key, str)
		if debug {
			fmt.Println("Key:", key, "value:", value)
		}
	}

	cmd := strings.Join(strings.Split(strings.Join(os.Args[1:], " "), "--")[1:], "--")
	if debug {
		fmt.Println(cmd)
	}

	parts = strings.Fields(cmd)
	datacmd := exec.Command(parts[0], parts[1:]...)
	datacmd.Stdout = os.Stdout
	datacmd.Stdin = os.Stdin
	datacmd.Stderr = os.Stderr
	datacmd.Run()
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func version() {
	println(color.Colorize(color.Cyan, `   
	                            __   __   
    ____   _______  _______   __ __|  |_/  |_ 
  _/ __ \ /    \  \/ /\__  \ |  |  \  |\   __\
  \  ___/|   |  \   /  / __ \|  |  /  |_|  |  
   \___  >___|  /\_/  (____  /____/|____/__|  
	   \/     \/           \/                    envaultVTAG
	`))
}

func main() {

	if !isCommandAvailable("vault") {
		println(color.Colorize(color.Yellow, "Doesn't look like you have vault installed, please follow the instructions here:\n https://www.vaultproject.io/docs/install"))
		return
	}

	envaultV := flag.Bool("v", false, "envault version")
	giDebug := flag.Bool("d", false, "envault debug (verbose)")
	giRun := flag.Bool("run", false, "set env variables and run command after double dash Ex. envault run -- npm run dev")
	giSet := flag.String("set", "", "Set key value secret Ex: envault -set=KEY=VALUE")
	giList := flag.String("list", "", "List key value secrets")
	giConfigure := flag.Bool("configure", false, "create configuration for the vault to be used")
	giStatus := flag.Bool("status", false, "Get vault status")

	flag.Parse()

	if *envaultV {
		version()
		return
	}

	if *giConfigure {
		configureVaultEnv(*giDebug)
		return
	}

	if *giSet != "" {
		s := strings.Split(*giSet, "=")
		if len(s) < 2 {
			println(color.Colorize(color.Yellow, "Wrong arguments received, Please refer to the help for the right usage"))
			return
		}

		key := s[0]
		value := s[1]
		set(*giDebug, key, value)
		return
	}

	if *giList != "" {
		list(*giDebug, *giList)
		return
	}

	if *giStatus {
		status(*giDebug)
		return
	}

	if *giRun {
		run(*giDebug)
		return
	}

	if *giDebug {
		fmt.Println("debug!!")
	}
}
