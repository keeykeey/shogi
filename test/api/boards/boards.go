package boards

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"bytes"
)

type Board struct {
	ID                  byte
	BoardTop            uint16
	BoardLeft           uint16
	SquareHeightCount   byte
	SquareWidthCount    byte
	SquareHeightLen     uint16
	SquareWidthLen      uint16
	LineWidth           byte
}

func TestGet(ok *int, ng *int) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env file was not found\n")
	}

	BASE_URL := os.Getenv("BASE_URL")
	endpoint := fmt.Sprintf("%s/api/board/1", BASE_URL)
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

	var expected = `{"ID":1,"boardTop":65,"boardLeft":41,"squareHeightCount":9,"squareWidthCount":9,"squareHeightLen":30,"squareWidthLen":30,"lineWidth":2}`

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

