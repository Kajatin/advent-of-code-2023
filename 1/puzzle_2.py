with open("puzzle_input.txt", "r") as f:
    data = f.readlines()
    data = [x.strip() for x in data]

text_to_int = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

def convert_text_to_numbers(line: str):
    for key, value in text_to_int.items():
        while True:
            s = line.find(key)
            if s < 0:
                break
            line = "".join([line[:s+1], value, line[s+len(key)-2:]])

    return line

def extract_first_and_last_number(line):
    print(line)
    line = convert_text_to_numbers(line)
    print(line)

    numbers = []
    for n in line:
        try:
            numbers.append(int(n))
        except ValueError:
            pass

    print(numbers)
    print(10 * numbers[0] + numbers[-1])
    return numbers

def construct_two_digit_number(numbers):
    return 10 * numbers[0] + numbers[-1]

numbers = [extract_first_and_last_number(line) for line in data]
two_digit_numbers = [construct_two_digit_number(ns) for ns in numbers]
summa = sum(two_digit_numbers)
print("Day 1 Puzzle 2: {}".format(summa))
