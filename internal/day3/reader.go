package day3

import (
	"bufio"
	"io"
	"unicode"
)

type TokenType int

const (
	NUMBER TokenType = iota
	SYMBOL
	EOF
)

type Token struct {
	Type       TokenType
	Line       int
	StartIndex int
	EndIndex   int
	Value      string
	IntValue   int
}

func (t Token) IsNeighbour(neighbour Token) bool {
	if t.Line == neighbour.Line {
		return neighbour.EndIndex == t.StartIndex-1 || t.EndIndex+1 == neighbour.StartIndex
	}
	if t.Line == neighbour.Line-1 || t.Line == neighbour.Line+1 {
		begin := t.StartIndex - 1
		end := t.EndIndex + 1
		return (neighbour.StartIndex <= begin && neighbour.EndIndex >= end) ||
			(neighbour.EndIndex >= begin && neighbour.EndIndex <= end) ||
			(neighbour.StartIndex >= begin && neighbour.StartIndex <= end)
	}
	return false
}

type Scanner struct {
	reader   *bufio.Reader
	line     int
	position int
}

func NewScanner(reader io.Reader) Scanner {
	return Scanner{
		reader:   bufio.NewReader(reader),
		line:     0,
		position: 0,
	}
}

func (r *Scanner) Tokens() []Token {
	var tokens []Token
	for token := r.Token(); token.Type != EOF; token = r.Token() {
		tokens = append(tokens, token)
	}
	return tokens
}
func (r *Scanner) Token() Token {
	for {
		position := r.position
		char, err := r.readNext()

		if err != nil {
			return Token{
				Type:       EOF,
				Line:       r.line,
				StartIndex: position,
				EndIndex:   position,
			}
		}

		if char != '.' && !unicode.IsDigit(char) && !unicode.IsSpace(char) {
			return Token{
				Type:       SYMBOL,
				Line:       r.line,
				StartIndex: position,
				EndIndex:   position,
				Value:      string(char),
			}
		}

		if unicode.IsDigit(char) {
			value := int(char - '0')
			startIndex := position
			endIndex := position
			for {
				if !r.isNextDigit() {
					break
				}
				nextChar, err := r.readNext()
				if err != nil {
					break
				}
				value = (value * 10) + int(nextChar-'0')
				endIndex += 1
			}

			return Token{
				Type:       NUMBER,
				Line:       r.line,
				StartIndex: startIndex,
				EndIndex:   endIndex,
				IntValue:   value,
			}
		}

		if char == '\n' {
			r.line += 1
			r.position = 0
		}

	}

}

func (r *Scanner) isNextDigit() bool {
	next, err := r.reader.Peek(1)
	return err == nil && unicode.IsDigit(rune(next[0]))
}

func (r *Scanner) readNext() (rune, error) {
	nextChar, _, err := r.reader.ReadRune()
	if err != nil {
		return 0, err
	}
	r.position += 1
	return nextChar, nil
}