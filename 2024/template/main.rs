use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let path = "input.txt";
    let file = File::open(&path).expect("could not open file");
    let reader = io::BufReader::new(file);

    let _lines: Vec<Vec<char>> = reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| line.chars().collect())
        .collect();
}

fn part_one() -> i64 {
    0
}

fn part_two() -> i64 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {}

    #[test]
    fn part_two_example() {}
}
