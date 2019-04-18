// solve returns the earned points in a single toss of a Darts game
const solve = function (x, y) {
  if (isNaN(x) || isNaN(y)) {
    return null;
  } else if (x > 10 || y > 10) {
    return 0;
  } else if (x > 5 || y > 5) {
    return 1;
  } else if (x > 1 || y > 1) {
    return 5;
  } else {
    return 10;
  }
}

export { solve };
