#[derive(Debug, Clone)]
struct Symbol {
    row: i32,
    col: i32,
    symbol: char,
}

#[derive(Debug, Clone)]
struct Number {
    row: i32,
    col: i32,
    len: i32,
    num: i32,
}

fn insert_number(i: i32, j: i32, line: &str, number_start: &mut Option<i32>, numbers: &mut Vec<Number>) {
    if let Some(n_start) = *number_start {
        numbers.push(Number {
            row: i,
            col: n_start,
            len: j - n_start,
            num: line[n_start as usize .. j as usize].parse::<i32>().unwrap(),
        });
    }

    *number_start = None;
}

fn gear_ratio(numbers: Vec<Number>, sym: Symbol) -> i32 {
    if sym.symbol != '*' {
        return 0;
    }

    let mut numbers_adjacent: Vec<Number> = Vec::new();
    for num in numbers {
        if (num.row - sym.row).abs() > 1 {
            continue;
        }

        if sym.col + 1 < num.col || sym.col > num.col + num.len {
            continue;
        }

        numbers_adjacent.push(num.clone());
    }

    if numbers_adjacent.len() != 2 {
        return 0;
    }

    numbers_adjacent[0].num * numbers_adjacent[1].num
}

fn main() {
    let puzzle = std::fs::read_to_string("puzzle_input.txt").unwrap();

    let mut symbols: Vec<Symbol> = Vec::new();
    let mut numbers: Vec<Number> = Vec::new();

    for (i, line) in puzzle.lines().enumerate() {
        let mut number_start: Option<i32> = None;
        let numbers_char = vec!['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
        for (j, char_) in line.chars().enumerate() {
            if numbers_char.contains(&char_) {
                if number_start.is_none() {
                    number_start = Some(j as i32);
                }
                continue;
            }

            if char_ != '.' {
                symbols.push(Symbol {
                    row: i as i32,
                    col: j as i32,
                    symbol: char_,
                });
            }

            insert_number(i as i32, j as i32, line, &mut number_start, &mut numbers);
        }

        insert_number(i as i32, line.len() as i32, line, &mut number_start, &mut numbers);
    }

    let mut summa: i32 = 0;
    for sym in symbols {
        summa += gear_ratio(numbers.clone(), sym.clone());
    }

    println!("Day 3 Puzzle 2: {}", summa);
}
