
const setCharAt = function (str, index, chr) {
  if (index > str.length - 1) return str;
  return str.substr(0, index) + chr + str.substr(index + 1);
};

//Add the numbers to a minesweeper board
const annotate = function (input) {
  let rowLen = input.length;
  for (let i = 0; i < rowLen; i++) {
    let colLen = input[i].length;
    for (let j = 0; j < colLen; j++) {
      if (!input[i][j] || input[i][j] === '' || input[i][j] === ' ') {
        let count = 0;
        if (j > 0 && input[i][j - 1] === '*') count++;
        if (i > 0 && input[i - 1][j] === '*') count++;
        if (i > 0 && j > 0 && input[i - 1][j - 1] === '*') count++;
        if (i < rowLen - 1 && j > 0 && input[i + 1][j - 1] === '*') count++;
        if (i > 0 && j < colLen - 1 && input[i - 1][j + 1] === '*') count++;
        if (i < rowLen - 1 && j < colLen - 1 && input[i + 1][j + 1] === '*') count++;
        if (i < rowLen - 1 && input[i + 1][j] === '*') count++;
        if (j < colLen - 1 && input[i][j + 1] === '*') count++;
        if (count > 0) {
          input[i] = setCharAt(input[i], j, count);
        }
      }
    }
  }
  return input;
};

export { annotate };