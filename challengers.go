package tcppow

import "tcp-pow/challenge"

type ChallengerProvider func() challenge.Challenger

var PingPongProvider = func() challenge.Challenger {
	return challenge.PingPongChallenger{}
}

var HasherProvider = func() challenge.Challenger {
	// bitLen and count may be dynamic depending on current metrics or whatever
	return challenge.NewHasherChallenger(2, 4)
}
