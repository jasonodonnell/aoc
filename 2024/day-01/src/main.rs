use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let file = File::open("input.txt").expect("could not open file");
    let reader = io::BufReader::new(file);
    let lines: Vec<String> = reader.lines().map_while(Result::ok).collect();
    let (mut a, mut b) = from_lines(lines);
    println!("{}", part_one(&mut a, &mut b));
    println!("{}", part_two(&mut a, &mut b));
}

fn from_lines(input: Vec<String>) -> (Vec<i64>, Vec<i64>) {
    input
        .iter()
        .map(|line| {
            let mut parts = line
                .split_whitespace()
                .map(|num| num.parse::<i64>().unwrap());
            (parts.next().unwrap(), parts.next().unwrap())
        })
        .unzip()
}

fn part_one(a: &mut [i64], b: &mut [i64]) -> i64 {
    a.sort();
    b.sort();

    a.iter()
        .enumerate()
        .map(|(index, value)| (value - b[index]).abs())
        .sum()
}

fn part_two(a: &mut [i64], b: &mut [i64]) -> i64 {
    a.sort();
    b.sort();

    let frequency: HashMap<i64, usize> = b.iter_mut().fold(HashMap::new(), |mut map, num| {
        *map.entry(*num).or_insert(0) += 1;
        map
    });

    a.iter()
        .map(|value| value * (*frequency.get(value).unwrap_or(&0)) as i64)
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = vec![
            String::from("3   4"),
            String::from("4   3"),
            String::from("2   5"),
            String::from("1   3"),
            String::from("3   9"),
            String::from("3   3"),
        ];

        let (mut a, mut b) = from_lines(input);
        assert_eq!(part_one(&mut a, &mut b), 11)
    }

    #[test]
    fn part_two_example() {
        let input = vec![
            String::from("3   4"),
            String::from("4   3"),
            String::from("2   5"),
            String::from("1   3"),
            String::from("3   9"),
            String::from("3   3"),
        ];
        let (mut a, mut b) = from_lines(input);
        assert_eq!(part_two(&mut a, &mut b), 31)
    }
}
