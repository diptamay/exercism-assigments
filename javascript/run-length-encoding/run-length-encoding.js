//Implements run-length encoding and decoding.

const encode = function (str) {
  if (!str || str.length === 0)
    return str;

  let ret = '';

  let count = 0;
  let prev = '';
  for (let i = 0; i < str.length; i++) {
    let char = str.charAt(i);
    if (prev === char) {
      count++;
      if (i === str.length - 1) {
        if (count > 1) ret += count;
        ret += prev;
      }
    } else {
      if (count > 1) ret += count;
      ret += prev;
      if (i === str.length - 1) ret += char;
      prev = char;
      count = 1;
    }
  }
  return ret;
};

const decode = function (str) {
  if (!str || str.length === 0)
    return str;

  let ret = '';
  let count = '';

  for (let i = 0; i < str.length; i++) {
    let char = str.charAt(i);
    if (!isNaN(parseInt(char, 10))) {
      count += char;
    } else {
      let ctr = parseInt(count, 10)
      if (!isNaN(ctr)) {
        for (let x = 0; x < ctr; x++)
          ret += char;
      } else {
        ret += char;
      }
      count = '';
    }
  }
  return ret;
};

export {encode, decode};
