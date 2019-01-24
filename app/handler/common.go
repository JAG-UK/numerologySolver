package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/JAG-UK/numerologySolver/config"
)

// Basic conversion - A=1, B=2 etc...
var basicCharMap *map[rune]uint32
var basicValueMap *map[uint32][]string

func InitBasicCharMap() {
	basicMap := make(map[rune]uint32)

	basicMap[rune('A')] = 1
	basicMap[rune('B')] = 2
	basicMap[rune('C')] = 3
	basicMap[rune('D')] = 4
	basicMap[rune('E')] = 5
	basicMap[rune('F')] = 6
	basicMap[rune('G')] = 7
	basicMap[rune('H')] = 8
	basicMap[rune('I')] = 9
	basicMap[rune('J')] = 10
	basicMap[rune('K')] = 11
	basicMap[rune('L')] = 12
	basicMap[rune('M')] = 13
	basicMap[rune('N')] = 14
	basicMap[rune('O')] = 15
	basicMap[rune('P')] = 16
	basicMap[rune('Q')] = 17
	basicMap[rune('R')] = 18
	basicMap[rune('S')] = 19
	basicMap[rune('T')] = 20
	basicMap[rune('U')] = 21
	basicMap[rune('V')] = 22
	basicMap[rune('W')] = 23
	basicMap[rune('X')] = 24
	basicMap[rune('Y')] = 25
	basicMap[rune('Z')] = 26

	basicCharMap = &basicMap
}

//TODO this will go away soon, in favour of pre-populating a permanent database
var basicConversionPrecompTable *map[string]uint32
var basicConversionPrecompValues *map[uint32][]string

//PrecompWordList precalculates values for the alpha word list
func PrecompWordList() error {
	//Load the word list
	//Requires a local \n delimited list of words.  `words_alpha` from https://github.com/dwyl/english-words works well.
	conf := config.GetConfig()
	//fmt.Printf("**DBG** Opening %v", conf.WordListPath)
	file, err := os.Open(conf.WordListPath)
	if err != nil {
		fmt.Printf("**DBG** FAILED!!", conf.WordListPath)
		return err
	}
	defer file.Close()

	precompTableS := make(map[string]uint32)
	precompTableV := make(map[uint32][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		a := BasicConversion(s)

		//Stash it
		//fmt.Printf("**DBG** Converted '%v' into %v", s, a)
		precompTableS[s] = a
		precompTableV[a] = append(precompTableV[a], s)
	}

	basicConversionPrecompTable = &precompTableS
	basicConversionPrecompValues = &precompTableV
	return nil
}

func BasicConversion(word string) uint32 {
	var ret uint32 = 0
	var workingCopy = strings.ToUpper(word)
	var letters = []rune(workingCopy)

	for _, r := range letters {
		ret += (*basicCharMap)[r]
	}
	return ret
}

func FindAll(val uint32, vallen uint32) []string {
	var ret []string

	if vallen == 0 {
		return (*basicConversionPrecompValues)[val]
	} else {
		for _, s := range (*basicConversionPrecompValues)[val] {
			if uint32(len(s)) == vallen {
				ret = append(ret, s)
			}
		}
	}

	return ret
}

// RESTful common code
func SuccessResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	SuccessResponse(w, status, map[string]string{"error": message})
}
