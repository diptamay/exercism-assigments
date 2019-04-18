// Package beer recites the lyrics to that beloved classic, that field-trip favorite: 99 Bottles of Beer on the Wall.
package beer

import (
	"errors"
	"fmt"
)

// verse parts
const (
	BeerBottles = " bottle%s of beer"
	Action      = "Take %s down and pass it around"
)

// constructing verse parts
var (
	bottlesOfBeer    = fmt.Sprintf(BeerBottles, "s")
	bottleOfBeer     = fmt.Sprintf(BeerBottles, "")
	beersOnWall      = bottlesOfBeer + " on the wall"
	beerOnWall       = bottleOfBeer + " on the wall"
	actionMulti      = fmt.Sprintf(Action, "one")
	emptyBeersOnWall = fmt.Sprintf("No more%s, no more%s.", beersOnWall, bottlesOfBeer)
	actionSingle     = fmt.Sprintf(Action, "it")
	verseEmpty       = fmt.Sprintf("%s\nGo to the store and buy some more, 99%s.\n", emptyBeersOnWall, beersOnWall)
)

// Verse recites a specific verse of 99 Bottles of Beer on the Wall.
func Verse(index int) (string, error) {
	var ret string

	if index > 99 || index < 0 {
		return ret, errors.New("Invalid verse number")
	}

	switch {
	case index == 0:
		//No more bottles of beer on the wall, no more bottles of beer.
		//Go to the store and buy some more, 99 bottles of beer on the wall.
		ret = verseEmpty
	case index == 1:
		//1 bottle of beer on the wall, 1 bottle of beer.
		//Take it down and pass it around, no more bottles of beer on the wall.
		ret = fmt.Sprintf("%d%s, %d%s.\n%s, no more%s.\n", index, beerOnWall, index, bottleOfBeer, actionSingle, beersOnWall)
	case index == 2:
		//2 bottles of beer on the wall, 2 bottles of beer.
		//Take one down and pass it around, 1 bottle of beer on the wall.
		ret = fmt.Sprintf("%d%s, %d%s.\n%s, %d%s.\n", index, beersOnWall, index, bottlesOfBeer, actionMulti, index-1, beerOnWall)
	default:
		//3 bottles of beer on the wall, 3 bottles of beer.
		//Take one down and pass it around, 2 bottles of beer on the wall.
		ret = fmt.Sprintf("%d%s, %d%s.\n%s, %d%s.\n", index, beersOnWall, index, bottlesOfBeer, actionMulti, index-1, beersOnWall)
	}
	return ret, nil
}

// Verses recites the verses to that beloved classic, that field-trip favorite: 99 Bottles of Beer on the Wall
// for specific bounds
func Verses(upperBound int, lowerBound int) (string, error) {
	ret := ""

	if upperBound <= lowerBound || upperBound > 99 || lowerBound < 0 {
		return ret, errors.New("Invalid bounds")
	}

	for i := upperBound; i >= lowerBound; i-- {
		str, _ := Verse(i)
		ret += str + "\n"
	}
	return ret, nil
}

// Song recites the lyrics to that beloved classic, that field-trip favorite: 99 Bottles of Beer on the Wall.
func Song() string {
	str, _ := Verses(99, 0)
	return str
}
