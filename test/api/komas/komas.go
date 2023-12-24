package komas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

type KomaResponse struct {
	ID      byte
	MoveID  byte
	MoveID2 byte
	Name    string
	Name2   string
}

func TestGet(ok *int, ng *int) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env file was not found\n")
	}

	BASE_URL := os.Getenv("BASE_URL")
	endpoint := fmt.Sprintf("%s/api/koma", BASE_URL)
	resp, err := http.Get(endpoint)
	if err != nil {
		*ng++
		fmt.Printf("Failed at test komas")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		*ng++
		fmt.Printf("failed at test koma\n")
		return
	}

	var komas []KomaResponse
	json.NewDecoder(resp.Body).Decode(&komas)
	for i, _ := range komas {
		if fmt.Sprintf("%T", komas[i].Name) != "string" ||
		   fmt.Sprintf("%T", komas[i].Name2) != "string" ||
		   fmt.Sprintf("%T", komas[i].MoveID) != "uint8" ||
		   fmt.Sprintf("%T", komas[i].MoveID2) != "uint8" {
		    fmt.Printf("failed at test koma type check\n")
			*ng++
			return
		}
	}

	*ok++
}
