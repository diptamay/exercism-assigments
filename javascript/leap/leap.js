// isLeap returns true/false if a year is a leap year
const isLeap = function (year) {
    if (year < 0) {
        return false;
    }

    if (year % 100 === 0 && year % 400 === 0)
        return true;
    else return year % 4 === 0 && year % 100 !== 0 && year % 400 !== 0;

}

export {isLeap};
