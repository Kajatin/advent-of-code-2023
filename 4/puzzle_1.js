var fs = require("fs");

var puzzle = fs.readFileSync("puzzle_input.txt", "utf8");

var cards = [];
puzzle.split("\n").forEach((line, idx) => {
  const winning_numbers = line
    .split(":")[1]
    .split("|")[0]
    .trim()
    .split(" ")
    .map((n) => {
      return parseInt(n.trim());
    })
    .filter((num) => !Number.isNaN(num));

  const my_numbers = line
    .split(":")[1]
    .split("|")[1]
    .trim()
    .split(" ")
    .map((n) => {
      return parseInt(n.trim());
    })
    .filter((num) => !Number.isNaN(num));

  cards.push({
    idx,
    winning_numbers,
    my_numbers,
  });
});

const summa = cards.reduce((prev, curr) => {
  const winners = curr.winning_numbers.filter((num) =>
    curr.my_numbers.includes(num)
  );
  const points = winners.length == 0 ? 0 : Math.pow(2, winners.length - 1);

  return prev + points;
}, 0);

console.log("Day 4 Puzzle 1: " + summa);
