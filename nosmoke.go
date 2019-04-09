package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	avgCigPrice   = 0.40
	avgMinsPerCig = 5
)

type SaveData struct {
	MoneySaved float64
	NumCigs    int
	Mins       int
}

var (
	usr, _        = user.Current()
	usrHomeDir, _ = filepath.Abs(usr.HomeDir)
	sFilePath     = filepath.Join(usrHomeDir, ".nosmoke.json")
	texts         = []string{
		"Number of cigarettes not smoked",
		"Total time saved by not smoking",
		"Total cash saved by not smoking",
		"nosmoke <break|stats|reset>",
	}
)

func smokeFreeBreak(s *SaveData) {
	s.NumCigs += 1
	s.Mins += avgMinsPerCig
	s.MoneySaved += avgCigPrice
}

func stats(s SaveData) {
	fmt.Printf("%s: %d\n%s: %d\n%s: $%.2f\n",
		texts[0], s.NumCigs,
		texts[1], s.Mins,
		texts[2], s.MoneySaved)
}

func load() SaveData {
	savefile, e := ioutil.ReadFile(sFilePath)
	if e != nil {
		log.Fatal(e)
	}

	var data SaveData
	err := json.Unmarshal(savefile, &data)
	if err != nil {
		log.Fatal(e)
	}

	return data
}

func save(s SaveData) {
	savefileJson, _ := json.Marshal(s)
	err := ioutil.WriteFile(sFilePath, savefileJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func showHelp() {
	fmt.Println(texts[3])
}

func main() {
	args := os.Args[1:]
	s := load()
	if len(args) == 1 {
		switch args[0] {
		case "break":
			smokeFreeBreak(&s)
			save(s)
		case "stats":
			stats(s)
		case "reset":
			s = SaveData{}
			save(s)
		}
	} else {
		showHelp()
	}
}
