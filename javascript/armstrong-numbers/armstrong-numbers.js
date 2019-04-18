'use strict';

// noOfDigits returns the number of digits in the number
const noOfDigits = function (num) {
  let n = num;
  let count = 0;
  while (n > 0) {
    n = Math.floor(n / 10);
    count += 1;
  }
  return count;
};

// validate determines whether a number is an Armstrong number
const validate = function (num) {
  const numDigits = noOfDigits(num);
  let n = num;
  console.log(`numDigits = ${numDigits}`);
  let armstrong = 0;
  while (n > 0) {
    const remainder = n % 10;
    armstrong += Math.pow(remainder, numDigits);
    n = Math.floor(n / 10);
  }
  return num === armstrong;
};

export {validate};
