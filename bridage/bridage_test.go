package bridage

import "testing"

func TestCircle_Draw(t *testing.T) {
	red := Circle{}
	red.Constructor(10,10,100,&RedCircle{})
	red.Draw()

	yellow := Circle{}
	yellow.Constructor(20,20,200,&YellowCircle{})
	yellow.Draw()
}