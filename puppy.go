package puppy

import "github.com/kkarthikeyan23/dog"

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())

}
