package gtre

import (
    "fmt"
)

func re2post(rep []rune) []rune {
	nalt := 0
	natom := 0
	buf := make([]rune, len(rep)*2)
	var paren [100]struct {
		nalt  int
		natom int
	}
	// `buf` and `parenPos` are looped in original C code by pointer increment
	// but no pointer arithmetic in golang, so index variable
	bufPos := 0
	parenPos := 0
    index := 0
	for ; index < len(rep); index++ {
        re := rep[index]
		switch re {
		case '(':
			if natom > 1 {
				natom--
				buf[bufPos] = '.'
				bufPos++
			}
			if parenPos >= 100 {
				panic("Maximum parenthesis limit crossed")
			}
			paren[parenPos].nalt = nalt
			paren[parenPos].natom = natom
			parenPos++
			nalt = 0
			natom = 0
		case '|':
			if natom == 0 {
				panic("| in invalid position")
			}
			for natom--; natom > 0; natom-- {
				buf[bufPos] = '.'
				bufPos++
			}
			nalt++
		case ')':
			if parenPos == 0 {
				panic(") without matching (")
			}
			if natom == 0 {
				panic(") in invalid position")
			}
			for natom--; natom > 0; natom-- {
				buf[bufPos] = '.'
				bufPos++
			}
			for ; nalt > 0; nalt-- {
				buf[bufPos] = '|'
				bufPos++
			}
			parenPos--
			nalt = paren[parenPos].nalt
			natom = paren[parenPos].natom
			natom++
		case '*', '+', '?':
			if natom == 0 {
				panic(fmt.Sprintf("%c in invalid position", re))
			}
			buf[bufPos] = re
			bufPos++
        case '[':
			if natom > 1 {
				natom--
				buf[bufPos] = '.'
				bufPos++
			}
			buf[bufPos] = '['
			bufPos++
            
            closeI := index+1
            for ; rep[closeI] != ']' ; closeI++ {
    			buf[bufPos] = rep[closeI]
    			bufPos++
            }
            index = closeI
			buf[bufPos] = rep[closeI]
			bufPos++
            natom++
		default:
			if natom > 1 {
				natom--
				buf[bufPos] = '.'
				bufPos++
			}
			buf[bufPos] = re
			bufPos++
			natom++
		}
	}
	if parenPos != 0 {
		panic("unmatched ( or )")
	}
	for natom--; natom > 0; natom-- {
		buf[bufPos] = '.'
		bufPos++
	}
	for ; nalt > 0; nalt-- {
		buf[bufPos] = '|'
		bufPos++
	}
	post := make([]rune, bufPos)
	copy(post, buf)
	return post
}

func printGraph(start *state) {
	def := ""
	nodes := printState(start, &def)
	_ = nodes
	// fmt.Println("digraph {")

	// fmt.Println(nodes)
	// fmt.Println(def)
	// fmt.Println("}")
}

func printState(s *state, def *string) string {
	if s.isMatch {
		return s.Identifier(def) + ";"
	} else if s.isSplit {
		return s.Identifier(def) + " -> " + printState(s.out, def) + "\n" + s.Identifier(def) + " -> " + printState(s.out2, def)
	} else {
		return s.Identifier(def) + " -> " + printState(s.out, def)
	}
}

func (s *state) String() string {
	if s.isMatch {
		return "END"
	} else if s.isSplit {
		return "O"
	} else {
		return string(*s.char.char)
	}
}

func (s *state) Identifier(def *string) string {
	*def += fmt.Sprintf("N%p[label=\"%s\"]\n", s, s.String())
	return fmt.Sprintf("N%p", s)
}