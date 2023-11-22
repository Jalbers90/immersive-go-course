package main

import (
	"bufio"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Player struct {
	Name      string `json:"name" csv:"name"`
	HighScore int32  `json:"high_score" csv:"high score"`
}

func main() {
	fmt.Println("hello")
	// players, err := ReadJSON()
	// players, err := ReadJSONTxt()
	// players, err := ReadRepeatedJSON()
	// players, err := ReadCSV()
	players, err := ReadBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("players ::: ", players)
	low := Lowest(players)
	high := Highest(players)
	fmt.Printf("Lowest score player ::: Name: %s ::: Score %d\n", low.Name, low.HighScore)
	fmt.Printf("Highest score player ::: Name: %s ::: Score %d\n", high.Name, high.HighScore)
}

func ReadBinary() ([]Player, error) {
	// here we are reading byte by byte with file.Read() using []byte
	// however we could have used a bufio.Reader to help clean things up a bit.
	// Using things like binary.Read(reader, byteOrder, &highscore) ... reader.ReadString('\x00')
	var players []Player
	// f, err := os.Open("examples/custom-binary-be.bin")
	f, err := os.Open("examples/custom-binary-le.bin")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	endian := make([]byte, 2)
	_, err = f.Read(endian)
	if err != nil {
		return nil, err
	}

	for {
		p := Player{}
		score := make([]byte, 4)
		_, err = f.Read(score)
		if err == io.EOF {
			// if file is read properly we must find EOF at this point
			return players, nil
		}
		if err != nil {
			return nil, err
		}
		if endian[0] == '\xFE' && endian[1] == '\xFF' {
			p.HighScore = int32(binary.BigEndian.Uint32(score))
		} else {
			p.HighScore = int32(binary.LittleEndian.Uint32(score))
		}

		nameBuf := make([]byte, 1)
		name := ""
		for {
			_, err := f.Read(nameBuf)
			if err != nil {
				fmt.Println("hey")
				return nil, err
			}
			if nameBuf[0] == 0 {
				break
			}
			c := string(nameBuf[0])
			name += c
		}
		p.Name = name
		players = append(players, p)
	}
}

func ReadCSV() ([]Player, error) {
	f, err := os.Open("examples/data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	var players []Player
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if line[0] == "name" && line[1] == "high score" {
			continue
		}

		hs, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		p := Player{
			Name:      line[0],
			HighScore: int32(hs),
		}
		players = append(players, p)
	}
	return players, nil
}

func ReadRepeatedJSON() ([]Player, error) {
	var players []Player
	f, err := os.Open("examples/repeated-json.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		if string(line[0]) == "#" {
			continue
		}
		var p Player
		err = json.Unmarshal(line, &p)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, nil
}

func ReadJSONTxt() ([]Player, error) {
	file, err := os.ReadFile("examples/json.txt")
	if err != nil {
		return nil, err
	}
	var players []Player
	err = json.Unmarshal(file, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func ReadJSON() ([]Player, error) {
	file, err := os.ReadFile("examples/json.json")
	if err != nil {
		return nil, err
	}
	var players []Player
	err = json.Unmarshal(file, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func Lowest(players []Player) Player {
	var lowest Player
	for _, player := range players {
		if lowest.Name == "" {
			lowest = player
		} else {
			if lowest.HighScore > player.HighScore {
				lowest = player
			}
		}
	}
	return lowest
}

func Highest(players []Player) Player {
	var highest Player
	for _, player := range players {
		if highest.Name == "" {
			highest = player
		} else {
			if highest.HighScore < player.HighScore {
				highest = player
			}
		}
	}
	return highest
}
