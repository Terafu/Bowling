package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0
	spare := false

	if len(game) != 10 {
		erreur := fmt.Errorf("There is more or less than 10 throws")
		return 0, erreur
	}

	for i := 0; i < len(game); i++ {
		if game[i].firstThrow+game[i].secondThrow > 10 {
			erreur := fmt.Errorf("A throw give a score > 10")
			return 0, erreur
		} else if game[i].firstThrow < 0 || game[i].secondThrow < 0 {

			erreur := fmt.Errorf("A throw is negative")
			return 0, erreur
		} else if game[i].firstThrow+game[i].secondThrow == 10 {

			spare = true
			score = score + game[i].firstThrow + game[i].secondThrow
		} else if spare == true {
			score = score + (game[i].firstThrow+game[i].secondThrow)*2
			spare = false
		} else {

			score = score + game[i].firstThrow + game[i].secondThrow
		}
	}
	return score, nil
}
