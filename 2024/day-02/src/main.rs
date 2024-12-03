use std::fs;

fn main() {
    let data = fs::read_to_string("input.txt").expect("Failed to read input file");
    println!("{}", part_one(&data));
    println!("{}", part_two(&data));
}

fn part_one(input: &str) -> usize {
    input
        .lines()
        .filter(|line| {
            let numbers: Vec<usize> = line
                .split_whitespace()
                .map(|s| s.parse::<usize>().expect("Failed to parse number"))
                .collect();

            numbers.windows(3).all(|window| {
                if let [w0, w1, w2] = *window {
                    let diff0 = w2.abs_diff(w1);
                    let diff1 = w1.abs_diff(w0);
                    (w2 > w1) == (w1 > w0) && diff0 > 0 && diff0 < 4 && diff1 > 0 && diff1 < 4
                } else {
                    false
                }
            })
        })
        .count()
}

fn part_two(input: &str) -> usize {
    input
        .lines()
        .filter(|line| {
            let numbers: Vec<isize> = line
                .split_whitespace()
                .map(|s| s.parse::<isize>().expect("Failed to parse number"))
                .collect();

            let mut permutations = vec![numbers.clone()];
            for i in 0..numbers.len() {
                let mut reduced = numbers.clone();
                reduced.remove(i);
                permutations.push(reduced);
            }

            permutations.iter().any(|numbers| {
                numbers.windows(3).all(|window| {
                    if let [w0, w1, w2] = *window {
                        let diff0 = w2.abs_diff(w1);
                        let diff1 = w1.abs_diff(w0);
                        (w2 > w1) == (w1 > w0) && diff0 > 0 && diff0 < 4 && diff1 > 0 && diff1 < 4
                    } else {
                        false
                    }
                })
            })
        })
        .count()
}
