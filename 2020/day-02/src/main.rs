use std::error::Error;
use std::fs::File;
use std::io::{self, BufRead};

/*
--- Day 2: Password Philosophy ---

Your flight departs in a few days from the coastal airport; the easiest way
down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day.
"Something's wrong with our computers; we can't log in!" You ask if you can
take a look.

Their password database seems to be a little corrupted: some of the passwords
wouldn't have been allowed by the Official Toboggan Corporate Policy that was
in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of
passwords (according to the corrupted database) and the corporate policy when
that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy
indicates the lowest and highest number of times a given letter must appear
for the password to be valid. For example, 1-3 a means that the password must
contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not;
it contains no instances of b, but needs at least 1. The first and third
passwords are valid: they contain one a or nine c, both within the limits of
their respective policies.

How many passwords are valid according to their policies?

--- Part Two ---

While it appears you validated the passwords correctly, they don't seem to be
what the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the
password policy rules from his old job at the sled rental place down the
street! The Official Toboggan Corporate Policy actually works a little
differently.

Each policy actually describes two positions in the password, where 1 means the
first character, 2 means the second character, and so on. (Be careful; Toboggan
Corporate Policies have no concept of "index zero"!) Exactly one of these
positions must contain the given letter. Other occurrences of the letter are
irrelevant for the purposes of policy enforcement.

Given the same example list from above:

    1-3 a: abcde is valid: position 1 contains a and position 3 does not.
    1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
    2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

How many passwords are valid according to the new interpretation of the policies?
*/

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
