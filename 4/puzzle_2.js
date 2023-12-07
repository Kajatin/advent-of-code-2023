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
    count: 1,
  });
});

cards.forEach((card, idx, cards) => {
  const winners = card.winning_numbers.filter((num) =>
    card.my_numbers.includes(num)
  );

  for (let index = idx + 1; index < idx + winners.length + 1; index++) {
    cards[index].count += card.count;
  }
});

const summa = cards.reduce((prev, curr) => {
  return prev + curr.count;
}, 0);

console.log("Day 4 Puzzle 2: " + summa);
