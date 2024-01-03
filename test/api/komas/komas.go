package komas

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"bytes"
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

	var expected = `[{"id":1,"moveId":1,"moveId2":1,"name":"ou","name2":""},{"id":2,"moveId":2,"moveId2":2,"name":"kin","name2":""},{"id":3,"moveId":3,"moveId2":3,"name":"gin","name2":"narigin"}]`

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	rb := buf.String()
	got := string(rb)

	if (got != expected) {
		*ng++
		fmt.Printf("Failed at test komas")
		return
	}


	*ok++
}
