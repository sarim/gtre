package gtre

func traverse(s *state, text *[]rune, pos int, stop *bool, isMatch *bool) {
	for {
		// another thread found match
		if *stop {
			break
		}

    	if s.isSplit {
    		traverse(s.out2, text, pos, stop, isMatch)
    		s = s.out
        } else if s.isMatch {
			if pos == len(*text) {
				*stop = true
				*isMatch = true
                // fmt.Printf("Match Found %q\n", *text)
				break
			} else {
				// regex matched ^pattern but not $
				break
			}
		} else {
			if pos == len(*text) {
				break
			} else {
                if *s.char.char == (*text)[pos] {
    				s = s.out
    				pos++
                } else if s.char.next != nil /*&& matchCharClass(s.char, text, &pos)*/ {
                    sc := s.char
                    for sc = sc.next; sc != nil ; sc = sc.next {
                        if *sc.char == (*text)[pos] {
            				s = s.out
            				pos++
                            break
                        }
                    }
                    if sc == nil {
                        break
                    }    				
                } else {
                    //nothing matches, abandon
                    break
                }
            }
		}
	}
}

func matchCharClass(sc *stateChar, text *[]rune, pos *int) bool {
    for sc = sc.next; sc != nil ; sc = sc.next {
        if *sc.char == (*text)[*pos] {
            return true
        }
    }
    return false
}