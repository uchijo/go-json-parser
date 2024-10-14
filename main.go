package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("input json required.")
		return
	}
	input := os.Args[1]
	fmt.Printf("input: %v\n", input)
	fmt.Printf("tokenized:\n")
	for i, v := range tokenize(input) {
		fmt.Printf("%v, %v\n", i, v)
	}
}

func consumeControl(input []Token, ch rune) ([]Token, error) {
	if len(input) <= 0 {
		return []Token{}, errors.New("failed to consume " + string(ch) + ", token length == 0")
	}
	if input[0].Content != string(ch) {
		return []Token{}, errors.New("failed to consume " + string(ch) + ", expected " + string(ch) + ", but got " + input[0].Content)
	}
	return input[1:], nil
}

func (t Token) isControlOf(ch rune) bool {
	return t.Type == CONTROL && t.Content == string(ch)
}

func takeString(input []Token) (string, []Token, error) {
	if len(input) <= 2 {
		return "", []Token{}, errors.New("token too short")
	}
	if input[0].Type != STRING {
		return "", []Token{}, errors.New("not string")
	}
	return input[0].String(), input[1:], nil
}

// func parseObject(input []Token) (map[string]interface{}, []Token, error) {
// 	tokens, err := consumeControl(input, '{')
// 	retval := map[string]interface{}{}
// 	if err != nil {
// 		panic("something went wrong")
// 	}
// 	for tokens[0].Type == CONTROL && tokens[0].Content == "}" {
// 		tokensLeft, key, val, err := parseObjectEntry(input)
// 		tokens = tokensLeft
// 		if err != nil {
// 			panic("something went wrong")
// 		}
// 		retval[key] = val
// 	}
// 	return retval, tokens, nil
// }

