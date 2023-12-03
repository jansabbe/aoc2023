package day3

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

func (t Token) FindNeighbours(tokens []Token) []Token {
	var result []Token
	for _, potentialNeighbour := range tokens {
		if t.IsNeighbour(potentialNeighbour) {
			result = append(result, potentialNeighbour)
		}
	}
	return result
}
