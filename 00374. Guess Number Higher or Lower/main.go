package _0374__Guess_Number_Higher_or_Lower

func guessNumber(n int) int {
	bot, top := 1, n
	for {
		v := (bot + top) / 2
		switch guess(v) {
		case 0:
			return v
		case 1:
			bot = v + 1
		case -1:
			top = v - 1
		}
	}
}

func guess(num int) int { return 0 }
