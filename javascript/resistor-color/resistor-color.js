// Resistors have color coded bands, where each color maps to a number
const colorToCode = new Map([
  ["black", 0],
  ["brown", 1],
  ["red", 2],
  ["orange", 3],
  ["yellow", 4],
  ["green", 5],
  ["blue", 6],
  ["violet", 7],
  ["grey", 8],
  ["white", 9],
]);

// Get all resistor colors
const colorsFromMap = function () {
  let ret = []
  for (let key of colorToCode.keys())
    ret.push(key);
  return ret;
}

// Get all resistor colors
const COLORS = colorsFromMap();

// Given a color find the code
const colorCode = function (color) {
  if (!color || color === "")
    throw new Error("Invalid color");

  return colorToCode.get(color);
};

export { COLORS, colorCode };
