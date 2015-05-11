package gtre

func stateRune(c rune) *state {
	s := state{}
	s.char = &stateChar{&c, nil}
	return &s
}

func stateRuneList(c *stateChar) *state {
	s := state{}
	s.char = c
	return &s
}

func stateSplit() *state {
	s := state{}
	s.isSplit = true
	return &s
}

func stateMatch() *state {
	s := state{}
	s.isMatch = true
	return &s
}

func fragOne(s *state) *frag {
	//TODO: implement a trash storage
	// and put orphaned frag and arrow there to recycle
	// so that we don't have to allocate so often

	f := frag{}

	f.start = s

	a := arrow{false, s, nil}
	f.ends = &a
	f.lastOfEnd = &a

	return &f
}

func fragJoin(f1 *frag, f2 *frag) *frag {
	//all out arrow from f1 goes to f2->start
	for _arrow := f1.ends; _arrow != nil; _arrow = _arrow.next {
		if _arrow.secondOut {
			_arrow.s.out2 = f2.start
		} else {
			_arrow.s.out = f2.start
		}
	}

	//after combining, final frag's out arrows = f2's arrows
	f1.ends = f2.ends
	f1.lastOfEnd = f2.lastOfEnd

	//TODO: delete/recycle f2
	return f1
}

func fragPrependSplit(f *frag) *frag {
	splitState := stateSplit()
	splitState.out = f.start

	f.start = splitState
	a := arrow{true, splitState, nil}

	f.lastOfEnd.next = &a
	f.lastOfEnd = &a

	return f
}

func fragAppendSplit(f1 *frag, f2 *frag) *frag {
	splitState := stateSplit()
	splitState.out = f1.start
	splitState.out2 = f2.start

	f1.start = splitState

	f1.lastOfEnd.next = f2.ends
	f1.lastOfEnd = f2.lastOfEnd

	//TODO: delete/recycle f2
	return f1
}