package gtre

func post2nfa(postfix []rune) *frag {
	var stack [1000]*frag
	stackp := 0 //Again, Access and pointer increment implemented by index variable

	push := func(f *frag) {
		stack[stackp] = f
		stackp++
	}
	pop := func() *frag {
		stackp--
		return stack[stackp]
	}
    index := 0
    
	for ; index < len(postfix); index++ {
        p := postfix[index]
		switch p {
		case '.': /* catenate */
			e2 := pop()
			e1 := pop()
			e := fragJoin(e1, e2)
			push(e)
		case '|': /* alternate */
			e2 := pop()
			e1 := pop()
			ff := fragAppendSplit(e1, e2)
			push(ff)
		case '?': /* zero or one */
			e := pop()
			ff := fragPrependSplit(e)
            push(ff)
        case '[':
            closeI := index+1
            var charList *stateChar = nil
            var firstChar *stateChar = nil
            for ; postfix[closeI] != ']' ; closeI++ {
                charListNew := &stateChar{}
                if charList == nil {
                    firstChar = charListNew
                } else {
                    charList.next = charListNew
                }
                charListNew.char = &postfix[closeI]
                charList = charListNew
            }
            index = closeI
            
			s := stateRuneList(firstChar)
			f := fragOne(s)
			push(f)
		default:
			s := stateRune(p)
			f := fragOne(s)
			push(f)
		}
	}

	e := pop()

	if stackp != 0 {
		panic("Something bad happened")
	}

	//adds matchState at the end
	m := stateMatch()
	finalFrag := fragJoin(e, fragOne(m))

	// patch(e.out, matchstate)

	// fmt.Printf("nfa: %#v\n", e.start)
	return finalFrag
}