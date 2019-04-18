//Roman to numeral map
const numbers = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
const letters = ['M', 'CM', 'D', 'CD', 'C', 'XC', 'L', 'XL', 'X', 'IX', 'V', 'IV', 'I'];

//Function to convert from normal numbers to Roman Numerals
const toRoman = function (input) {
  let ret = '';
  let num = input;
  while (num > 0) {
    for (let i = 0; i < numbers.length; i++) {
      if ((num - numbers[i]) >= 0) {
        ret += letters[i];
        num -= numbers[i];
        break;
      }
    }
  }
  return ret;
};

export { toRoman };
