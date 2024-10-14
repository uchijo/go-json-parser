package main

type TokenType int

const (
	STRING TokenType = iota
	OTHER
	CONTROL
)

type Token struct {
	Content string
	Type    TokenType
}

func (t Token) String() string {
	switch t.Type {
	case STRING:
		return "STRING:\t" + t.Content
	case OTHER:
		return "OTHER:\t" + t.Content
	case CONTROL:
		return "CONTROL:\t" + t.Content
	}
	panic("it cant be here")
}

type TokenList []Token

func tokenize(input string) TokenList {
	runedInput := []rune(input)

	retval := []Token{}
	var tmp string
	isInString := false
	for i := 0; i < len(runedInput); i++ {
		c := runedInput[i]

		// exit from string
		if isInString && c == '"' {
			retval = append(retval, Token{Type: STRING, Content: tmp})
			tmp = ""
			retval = append(retval, Token{Type: CONTROL, Content: string(c)})
			isInString = false
			continue
		}

		// escape in string
		if isInString && c == '\\' {
			i++
			tmp += string(runedInput[i])
			continue
		}

		// inside string
		if isInString {
			tmp += string(c)
			continue
		}

		// enters string
		if c == '"' {
			isInString = true
		}

		// control outside string
		if isControl(c) {
			if len(tmp) != 0 {
				retval = append(retval, Token{Type: OTHER, Content: tmp})
			}
			retval = append(retval, Token{Type: CONTROL, Content: string(c)})
			tmp = ""
			continue
		}

		// other literal outside string
		tmp += string(c)
	}

	if len(tmp) != 0 {
		retval = append(retval, Token{Type: OTHER, Content: tmp})
	}
	return retval
}

func isControl(c rune) bool {
	return c == '{' ||
		c == '}' ||
		c == ',' ||
		c == '[' ||
		c == ']' ||
		c == '"' ||
		c == ':'
}
