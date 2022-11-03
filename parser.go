package gocon

import (
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

const (
	equalsToken = -(iota - scanner.Comment + 1)
	colonToken
	commaToken
	dotToken
	objectStartToken
	objectEndToken
	arrayStartToken
	arrayEndToken
	commentToken
	variableStartToken
)

var tokenLookup = map[rune]int{
	'=': equalsToken,
	':': colonToken,
	',': commaToken,
	'.': dotToken,
	'{': objectStartToken,
	'}': objectEndToken,
	'[': arrayStartToken,
	']': arrayEndToken,
	'#': commentToken,
	'$': variableStartToken,
}

type parser struct {
	scanner        *scanner.Scanner
	keyBreadCrumbs []string
	config         Config
}

func (p *parser) pushKey(key string) {
	p.keyBreadCrumbs = append(p.keyBreadCrumbs, key)
}

func (p *parser) currentKey() string {
	return strings.Join(p.keyBreadCrumbs, ".")
}

func (p *parser) addValue(v ConfigValue) {
	p.config.addValue(p.currentKey(), v)
}

func (p *parser) popKey() {
	currentLen := len(p.keyBreadCrumbs)
	if currentLen > 0 {
		p.keyBreadCrumbs = p.keyBreadCrumbs[:currentLen-1]
	}
}

func newParser(src io.Reader) *parser {
	s := new(scanner.Scanner)
	s.Init(src)
	// s.Error = func(*scanner.Scanner, string) {} // TODO
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	s.Mode = scanner.ScanChars |
		scanner.ScanFloats |
		scanner.ScanIdents |
		scanner.ScanInts |
		scanner.ScanStrings |
		scanner.ScanRawStrings |
		scanner.SkipComments
	// s.Position.Filename = name
	s.IsIdentRune = func(ch rune, i int) bool {
		if unicode.IsLetter(ch) {
			return true
		}
		switch ch {
		case '_':
			return true
		case '-':
			return true
		}
		return i > 0 && unicode.IsDigit(ch)
	}

	return &parser{scanner: s, config: NewConfig()}
}

func Parse(input io.Reader) (*Config, error) {
	parser := newParser(input)
	return parser.parse()
}

func (p *parser) parse() (*Config, error) {
	for {
		r := p.next()
		switch r {
		case scanner.EOF:
		case scanner.Comment: // ingore these comments
		case commentToken:
			p.parseComment()
		case scanner.Ident:
			p.handleKey(p.scanner.TokenText())
		// TODO: Handle includes
		default:
			if val, ok := tokenLookup[r]; ok {
				r = rune(val)
			}
		}
		if r == scanner.EOF {
			break
		}
	}
	return &p.config, nil
}

func (p *parser) next() rune {
	r := p.scanner.Scan()
	if val, ok := tokenLookup[r]; ok {
		return rune(val)
	}
	return r
}

func (p *parser) handleKey(key string) {
	p.pushKey(key)
	r := p.next()
	switch r {
	case equalsToken:
		p.parseValue()
	case colonToken:
		p.parseValue()
	case objectStartToken:
		p.parseObject()
	case arrayStartToken:
		p.parseArray()
	default:
	}
}

func (p *parser) parseValue() {
	r := p.next()
	switch r {
	case variableStartToken:
		p.parseVariable()
	// handle dates, ints, floats, etc. later - for now, just treat all values as strings
	case scanner.Ident:
		fallthrough
	case scanner.String:
		fallthrough
	case scanner.Int:
		fallthrough
	case scanner.Float:
		fallthrough
	case scanner.Char:
		p.addValue(NewConfigString(p.scanner.TokenText()))
		p.popKey()
	default:
	}
}

func (p *parser) parseVariable() {
	variable := "$"
	isEnd := false
	for {
		r := p.next()
		switch r {
		case '?':
			fallthrough
		case objectStartToken:
			fallthrough
		case scanner.Ident:
			// handle errors, like these tokens are out of order
			variable = fmt.Sprintf("%s%s", variable, p.scanner.TokenText())
		case objectEndToken:
			isEnd = true
			variable = fmt.Sprintf("%s%s", variable, p.scanner.TokenText())
		default:
			isEnd = true
		}
		if isEnd {
			break
		}
	}
	p.addValue(NewConfigString(variable))
	p.popKey()
}

func (p *parser) parseObject() {

}

func (p *parser) parseInclude() {

}

func (p *parser) parseComment() {

}

func (p *parser) parseArray() {

}
