use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let file = File::open("input.txt").expect("could not open file");
    let reader = io::BufReader::new(file);
    let answers = parse_answers(reader.lines().map_while(Result::ok)).unwrap();
    println!("{}", part_one(&answers));
    println!("{}", part_two(&answers));
}

#[derive(Debug)]
struct Person {
    answers: Vec<char>,
}

fn part_one(answers: &[Vec<Person>]) -> usize {
    answers
        .iter()
        .map(|group| {
            group
                .iter()
                .flat_map(|person| person.answers.iter())
                .collect::<HashSet<_>>()
                .len()
        })
        .sum()
}

fn part_two(answers: &[Vec<Person>]) -> usize {
    answers
        .iter()
        .map(|group| {
            let mut char_count: HashMap<char, usize> = HashMap::new();
            group
                .iter()
                .flat_map(|person| person.answers.iter())
                .for_each(|ch| {
                    *char_count.entry(*ch).or_insert(0) += 1;
                });

            char_count
                .values()
                .filter(|&&count| count == group.len())
                .count()
        })
        .sum()
}

fn parse_answers<I>(lines: I) -> io::Result<Vec<Vec<Person>>>
where
    I: IntoIterator,
    I::Item: AsRef<str>,
{
    let mut answers: Vec<Vec<Person>> = vec![];
    let mut current: Vec<Person> = vec![];

    for line in lines {
        let line = line.as_ref();
        if line.trim().is_empty() {
            if !current.is_empty() {
                answers.push(std::mem::take(&mut current));
            }
        } else {
            let mut person = Person { answers: vec![] };
            for c in line.chars() {
                person.answers.push(c);
            }
            current.push(person);
        }
    }

    if !current.is_empty() {
        answers.push(current);
    }
    Ok(answers)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "\
abc

a
b
c

ab
ac

a
a
a
a

b
		";

        let lines = input.lines();
        let answers = parse_answers(lines).unwrap();
        assert_eq!(part_one(&answers), 11_usize);
    }

    #[test]
    fn part_two_example() {
        let input = "\
abc

a
b
c

ab
ac

a
a
a
a

b
		";

        let lines = input.lines();
        let answers = parse_answers(lines).unwrap();
        assert_eq!(part_two(&answers), 6_usize);
    }
}
