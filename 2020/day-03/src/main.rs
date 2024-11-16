use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let path = "input.txt";
    let file = File::open(&path).expect("could not open file");
    let reader = io::BufReader::new(file);

    let map: Vec<Vec<char>> = reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| line.chars().collect())
        .collect();

    println!("Tree count: {}", part_one(&map));
    println!("Tree count: {}", part_two(&map));
}

fn part_one(map: &Vec<Vec<char>>) -> i64 {
    count_trees(map, 3, 1)
}

fn part_two(map: &Vec<Vec<char>>) -> i64 {
    let slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)];
    slopes
        .iter()
        .map(|&(right, down)| count_trees(map, right, down))
        .product()
}

fn count_trees(map: &Vec<Vec<char>>, right: usize, down: usize) -> i64 {
    let width = map.first().map_or(0, |row| row.len());
    let mut row_idx = 0;
    let mut col_idx = 0;
    let mut tree_count = 0;

    while row_idx < map.len() {
        if map[row_idx][col_idx % width] == '#' {
            tree_count += 1;
        }
        row_idx += down;
        col_idx += right;
    }
    tree_count
}

#[cfg(test)]
mod tests {
    use super::*;

    fn create_map() -> Vec<Vec<char>> {
        let input = "
    ..##.......
    #...#...#..
    .#....#..#.
    ..#.#...#.#
    .#...##..#.
    ..#.##.....
    .#.#.#....#
    .#........#
    #.##...#...
    #...##....#
    .#..#...#.#
    ";

        input
            .lines()
            .filter_map(|line| {
                let line = line.trim();
                if !line.is_empty() {
                    Some(line.chars().collect())
                } else {
                    None
                }
            })
            .collect()
    }

    #[test]
    fn part_one_example() {
        let input = create_map();
        assert_eq!(part_one(&input), 7);
    }

    #[test]
    fn part_two_example() {
        let input = create_map();
        assert_eq!(part_two(&input), 336);
    }
}
