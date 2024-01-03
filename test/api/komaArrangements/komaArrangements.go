package komaArrangements

import (
	"net/http"
	"fmt"
	"time"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"bytes"
)

type KomaArrangementsResponse struct {
	ID         			uint16
	ArrangementID       uint16
	Arrangement         Arrangement
    KomaID              byte
	Koma                Koma
	PositionID          uint16
	Position            Position
	IsFirstMove         bool
	IsFront             bool
}

type Arrangement struct {
	ID                  uint16
	Name                string
	gorm.Model
}

type Koma struct {
	ID                 byte
	MoveID             byte
	MoveID2            byte
	Name               string
	Name2              string
	BasicColumn
}

type BasicColumn struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

type KomaBase struct {
	ID                 byte
	MoveID             byte
	MoveID2            byte
	Name               string
	Name2              string
}

type Position struct {
	ID               uint16
	Width            uint8
	Height           uint8
	Name             string
	BasicColumn
}

type PositionBase struct {
	ID               uint16
	Number           uint8
	Name             string
}

func TestGet(ok *int, ng *int) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env file was not found\n")
	}

	BASE_URL := os.Getenv("BASE_URL")
	endpoint := fmt.Sprintf("%s/api/komaArrangements/1",BASE_URL)
	resp, err := http.Get(endpoint)
	if err != nil {
		*ng++
		fmt.Printf("Failed at test komaArrangements")
		return
	}
    defer resp.Body.Close()

	if resp.StatusCode != 200 {
		*ng++
	    fmt.Printf("failed at test komaArrangements\n")
		return
	}

	var expected = `[{"id":1,"arrangementId":1,"arrangement":{"id":1,"name":"平手"},"komaId":1,"koma":{"id":1,"moveId":1,"moveId2":1,"name":"ou","name2":""},"positionId":5,"position":{"id":5,"width":5,"height":1,"name":"五一"},"isFirstMove":false,"isFront":true},{"id":2,"arrangementId":1,"arrangement":{"id":1,"name":"平手"},"komaId":3,"koma":{"id":3,"moveId":3,"moveId2":3,"name":"gin","name2":"narigin"},"positionId":7,"position":{"id":7,"width":3,"height":1,"name":"三一"},"isFirstMove":false,"isFront":true}]`

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	rb := buf.String()
	got := string(rb)

	if (got != expected) {
	    *ng++
	    fmt.Printf("failed at test komaArrangements\n")
		return
	}

	*ok++
}
