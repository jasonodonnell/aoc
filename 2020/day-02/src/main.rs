use std::error::Error;
use std::fs::File;
use std::io::{self, BufRead};

#[derive(Debug)]
struct PasswordPolicy {
    min: i32,
    max: i32,
    letter: char,
}

impl PasswordPolicy {
    fn from_line(line: &str) -> Result<(PasswordPolicy, String), Box<dyn Error>> {
        let parts: Vec<&str> = line.split(" ").collect();
        let min_max: Vec<&str> = parts[0].split("-").collect();
        let min = min_max[0].parse::<i32>()?;
        let max = min_max[1].parse::<i32>()?;

        let letter = parts[1].chars().next().unwrap();
        let password = parts[2].to_string();
        Ok((PasswordPolicy { min, max, letter }, password))
    }

    fn valid_contains(&self, password: String) -> bool {
        let count = password.chars().filter(|&char| char == self.letter).count() as i32;
        count >= self.min && count <= self.max
    }

    fn valid_pos(&self, password: String) -> bool {
        let chars: Vec<char> = password.chars().collect();
        let min_pos_match = chars.get(self.min as usize - 1).unwrap() == &self.letter;
        let max_pos_match = chars.get(self.max as usize - 1).unwrap() == &self.letter;
        min_pos_match ^ max_pos_match
    }
}

fn main() {
    let path = "input.txt";
    let file = File::open(&path).expect("could not open file");
    let reader = io::BufReader::new(file);

    let passwords: Vec<(PasswordPolicy, String)> = reader
        .lines()
        .filter_map(|line| line.ok())
        .filter_map(|line| PasswordPolicy::from_line(&line).ok())
        .collect();

    let count = part_one(&passwords);
    println!("{count}");

    let count = part_two(&passwords);
    println!("{count}")
}

fn part_one(passwords: &Vec<(PasswordPolicy, String)>) -> i32 {
    passwords
        .iter()
        .filter(|(policy, password)| policy.valid_contains(password.to_string()))
        .count() as i32
}

fn part_two(passwords: &Vec<(PasswordPolicy, String)>) -> i32 {
    passwords
        .iter()
        .filter(|(policy, password)| policy.valid_pos(password.to_string()))
        .count() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let inputs = vec![
            ("1-3 a: abcde", true),
            ("1-3 b: cdefg", false),
            ("2-9 c: ccccccccc", true),
        ];

        for (line, expected) in inputs.iter() {
            let (policy, password) = PasswordPolicy::from_line(line).expect("Failed to parse line");
            assert_eq!(
                policy.valid_contains(password),
                *expected,
                "Failed for input: {}",
                line
            );
        }
    }

    #[test]
    fn part_two_example() {
        let inputs = vec![
            ("1-3 a: abcde", true),
            ("1-3 b: cdefg", false),
            ("2-9 c: ccccccccc", false),
        ];

        for (line, expected) in inputs.iter() {
            let (policy, password) = PasswordPolicy::from_line(line).expect("Failed to parse line");
            assert_eq!(
                policy.valid_pos(password),
                *expected,
                "Failed for input: {}",
                line
            );
        }
    }
}
