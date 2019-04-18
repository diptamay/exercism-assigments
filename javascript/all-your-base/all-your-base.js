// Convert a number, represented as a sequence of digits in one base, to any other base.
// Implements general base conversion. Given a number in base a, represented as a sequence of digits, convert it to base b.
export const convert = function (digits, baseFrom, baseTo) {
  checkBases(baseFrom, baseTo);
  checkDigits(digits, baseFrom);

  let num = toBase10(digits, baseFrom);
  return toBaseN(num, baseTo);
}

const checkDigits = function (digits, base) {
  if (digits.length === 0 ||
    (digits.length > 1 && digits[0] === 0)) {
    throw new Error("Input has wrong format");
  }
  if (digits.length > 1 && digits.every(d => d === 0)) {
    throw new Error("Input has wrong format");
  }
  if (digits.filter(d => d < 0 || d >= base).length > 0) {
    throw new Error("Input has wrong format");
  }
}

const checkBases = function (baseFrom, baseTo) {
  if (baseFrom < 2 || !Number.isInteger(baseFrom)) {
    throw new Error("Wrong input base");
  }
  if (baseTo < 2 || !Number.isInteger(baseTo)) {
    throw new Error("Wrong output base");
  }
}

const toBase10 = function (digits, base) {
  let d = 1;
  let num = 0;
  for (var i = digits.length - 1; i >= 0; --i) {
    num += d * digits[i];
    d = d * base;
  }
  return num;
}

const toBaseN = function (num, b) {
  let ans = [];
  let q = num;
  let r = 0;
  do {
    r = q % b;
    q = Math.floor(q / b);
    ans.push(r);
  } while (q > 0)
  ans.reverse();
  return ans;
}