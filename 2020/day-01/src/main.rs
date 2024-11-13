use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let path = "input.txt";
    let file = File::open(&path).expect("could not open file");
    let reader = io::BufReader::new(file);

    let numbers: Vec<i32> = reader
        .lines()
        .filter_map(|line| line.ok())
        .filter_map(|line| line.parse::<i32>().ok())
        .collect();

    println!("{}", part_one(numbers.clone()).unwrap());
    println!("{}", part_two(numbers).unwrap());
}

fn part_one(mut numbers: Vec<i32>) -> Option<i32> {
    numbers.sort_unstable();
    find_target_sum(&numbers, 2020)
        .map(|(a, b)| a.checked_mul(b))
        .flatten()
}

fn part_two(mut numbers: Vec<i32>) -> Option<i32> {
    numbers.sort_unstable();
    for (i, num) in numbers.iter().enumerate() {
        let target: i32 = 2020 - num;
        if let Some((a, b)) = find_target_sum(&numbers[i + 1..], target) {
            return Some(a * b * num);
        } else {
            continue;
        }
    }
    None
}

fn find_target_sum(numbers: &[i32], target: i32) -> Option<(i32, i32)> {
    let mut left = 0;
    let mut right = numbers.len().checked_sub(1)?;

    while left < right {
        let sum = numbers[left].checked_add(numbers[right])?;
        match sum.cmp(&target) {
            std::cmp::Ordering::Equal => return Some((numbers[left], numbers[right])),
            std::cmp::Ordering::Less => left += 1,
            std::cmp::Ordering::Greater => right -= 1,
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = vec![1721, 979, 366, 299, 675, 1456];
        assert_eq!(514579, part_one(input).unwrap());
    }

    #[test]
    fn part_two_example() {
        let input = vec![1721, 979, 366, 299, 675, 1456];
        assert_eq!(241861950, part_two(input).unwrap());
    }
}
