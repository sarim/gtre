package gtre

type stateChar struct {
    char *rune
    next *stateChar
}

type state struct {
	char    *stateChar
	out     *state
	out2    *state
	isMatch bool
	isSplit bool
}

type arrow struct {
	secondOut bool
	s         *state
	next      *arrow
}

type frag struct {
	start     *state
	ends      *arrow
	lastOfEnd *arrow
}
