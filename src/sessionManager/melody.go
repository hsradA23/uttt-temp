package sessionManager

import "github.com/olahol/melody"

var MelodySession *melody.Melody

func init() {
	MelodySession = melody.New()

}
