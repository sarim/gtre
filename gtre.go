package gtre

type Gtre struct {
    source []rune
    start *frag
}

func Parse(pattern []rune) *Gtre {
    post := re2post(pattern)
    startFrag := post2nfa(post)
    g := Gtre{pattern, startFrag}
    return &g
}

func (g *Gtre) Match(text []rune) bool {
    startState := g.start.start
    stop := false
    match := false
    traverse(startState, &text, 0, &stop, &match)
    if (match) {
        return true
    }
    return false
}