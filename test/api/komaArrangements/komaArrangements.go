package komaArrangements

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
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

	var komaArrangements []KomaArrangementsResponse
	err2 := json.NewDecoder(resp.Body).Decode(&komaArrangements)
	if err2 != nil {
    	fmt.Printf("failed to decode json into Type KomaArrangements")
	    *ng++
	    return
	}

	for i, _ := range komaArrangements {
		if fmt.Sprintf("%T", komaArrangements[i].ID) != "uint16" ||
		   fmt.Sprintf("%T", komaArrangements[i].ArrangementID) != "uint16" ||
		   fmt.Sprintf("%T", komaArrangements[i].Arrangement) != "komaArrangements.Arrangement" ||
		   fmt.Sprintf("%T", komaArrangements[i].KomaID) != "uint8" ||
		   fmt.Sprintf("%T", komaArrangements[i].Koma) != "komaArrangements.Koma" &&
		   fmt.Sprintf("%T", komaArrangements[i].PositionID) != "uint16" &&
		   fmt.Sprintf("%T", komaArrangements[i].Position) != "komaArrangements.Position" &&
		   fmt.Sprintf("%T", komaArrangements[i].IsFirstMove) != "bool" &&
		   fmt.Sprintf("%T", komaArrangements[i].IsFront) != "bool" {
	        fmt.Printf("failed at test koma arrangement type check\n")
		    *ng++
		    return
		}
	}

	*ok++
}
