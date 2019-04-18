//Reverse a string
const reverseString = function (str) {
  if (!str || str.length === 0)
    return str;

  let rev = '';
  for (let i = str.length - 1; i >= 0; i--) {
    rev += str.charAt(i);
  }
  return rev
};

export {reverseString};
