#[derive(Debug, Clone)]
struct Symbol {
    row: i32,
    col: i32,
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

fn number_is_near_symbol(num: Number, sym: Symbol) -> bool {
    if (num.row - sym.row).abs() > 1 {
        return false;
    }

    if sym.col + 1 < num.col || sym.col > num.col + num.len {
        return false;
    }

    true
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
                });
            }

            insert_number(i as i32, j as i32, line, &mut number_start, &mut numbers);
        }

        insert_number(i as i32, line.len() as i32, line, &mut number_start, &mut numbers);
    }

    let mut summa: i32 = 0;
    for num in numbers {
        // println!("{:?}", num);
        for sym in &symbols {
            if number_is_near_symbol(num.clone(), sym.clone()) {
                summa += num.num;
                break;
            }
        }
    }

    println!("Day 3 Puzzle 1: {}", summa);
}
