// ShareWith returns strings depending of value of `name`
const twoFer = function (name) {
  if (!name || name.length === 0)
    name = 'you';
  return `One for ${name}, one for me.`
};

export {twoFer};
