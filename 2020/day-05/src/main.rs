use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let path = "input.txt";
    let file = File::open(path).expect("could not open file");
    let reader = io::BufReader::new(file);
    let lines: Vec<String> = reader.lines().map_while(Result::ok).collect();
    println!("{}", part_one(&lines));
    println!("{}", part_two(&lines))
}

fn part_one(input: &Vec<String>) -> i64 {
    let seats = calculate_seat_ids(input);
    seats.iter().max().copied().unwrap_or(0)
}

fn part_two(input: &Vec<String>) -> i64 {
    let mut seats = calculate_seat_ids(input);
    seats.sort();
    find_missing_number(seats)
}

fn calculate_seat_ids(input: &Vec<String>) -> Vec<i64> {
    let mut seats: Vec<i64> = vec![];
    for seat in input {
        let left = 1;
        let right = 128;

        let row = seat
            .chars()
            .take(7)
            .fold((left, right), |(left, right), c| {
                let mid = left + (right - left) / 2;
                match c {
                    'B' => (mid, right),
                    'F' => (left, mid),
                    _ => (left, right),
                }
            })
            .0;

        let left = 0;
        let right = 8;
        let col = seat
            .chars()
            .skip(7)
            .fold((left, right), |(left, right), c| {
                let mid = left + (right - left) / 2;
                match c {
                    'R' => (mid, right),
                    'L' => (left, mid),
                    _ => (left, right),
                }
            })
            .0;

        let seat_id = row * 8 + col;
        seats.push(seat_id);
    }
    seats
}

fn find_missing_number(nums: Vec<i64>) -> i64 {
    let mut left = 0;
    let mut right = nums.len() as i64 - 1;

    while left <= right {
        let mid = left + (right - left) / 2;
        if nums[mid as usize] == nums[0] + mid {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    nums[0] + left
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = vec![
            String::from("BFFFBBFRRR"),
            String::from("FFFBBBFRRR"),
            String::from("BBFFBBFRLL"),
        ];
        assert_eq!(part_one(&input), 820);
    }

    #[test]
    fn test_find_missing_number() {
        let seats = vec![1, 2, 3, 4, 5, 7, 8, 9, 10];
        assert_eq!(find_missing_number(seats), 6);
    }
}
