package scanner

type Token struct {
	Type          int    //token type
	Value         string //token value
	Line          int
	Col           int
	Pos           int
	Endpos        int
	Nlb           bool
	Attributes    []string
	CommentsBefor []*Token //comments
}

const (
	TokenOperator = iota
	TokenKeyword
	TokenPunc
	TokenNumber
	TokenString
	TokenLineComment
	TokenMultComment
	TokenRegexp
	TokenName
	TokenAtom
)
