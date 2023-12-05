with open("puzzle_input.txt", "r") as f:
    data = f.readlines()
    data = [x.strip() for x in data]

def extract_first_and_last_number(line):
    numbers = []
    for n in line:
        try:
            numbers.append(int(n))
        except ValueError:
            pass

    return numbers

def construct_two_digit_number(numbers):
    return 10 * numbers[0] + numbers[-1]

numbers = [extract_first_and_last_number(line) for line in data]
two_digit_numbers = [construct_two_digit_number(ns) for ns in numbers]
summa = sum(two_digit_numbers)
print("Day 1 Puzzle 1: {}".format(summa))
