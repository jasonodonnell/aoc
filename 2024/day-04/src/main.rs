#![allow(dead_code)]
use std::fs;

fn main() {
    let data = fs::read_to_string("input.txt").expect("Failed to read input file");
    let word_search: Vec<Vec<char>> = data.lines().map(|line| line.chars().collect()).collect();

    println!("{}", part_one(&word_search));
    println!("{}", part_two(&word_search));
}

fn part_one(word_search: &[Vec<char>]) -> usize {
    let rows = word_search.len() as isize;
    let cols = word_search.first().map_or(0, |row| row.len()) as isize;

    let mas: Vec<_> = "MAS"
        .chars()
        .enumerate()
        .map(|(i, c)| ((i + 1) as isize, c))
        .collect();

    let directions = [
        (1, 1),
        (1, 0),
        (1, -1),
        (0, 1),
        (0, -1),
        (-1, 1),
        (-1, 0),
        (-1, -1),
    ];

    let mut count = 0;
    for i in 0..word_search.len() {
        for j in 0..word_search[0].len() {
            if word_search[i][j] == 'X' {
                let (i, j) = (i as isize, j as isize);

                for (dx, dy) in directions.iter() {
                    if mas.iter().all(|(step, c)| {
                        let x = i + dx * step;
                        let y = j + dy * step;
                        in_bounds(x, y, rows, cols) && word_search[x as usize][y as usize] == *c
                    }) {
                        count += 1
                    }
                }
            }
        }
    }
    count
}

fn part_two(word_search: &[Vec<char>]) -> usize {
    let rows = word_search.len() as isize;
    let cols = word_search.first().map_or(0, |row| row.len()) as isize;

    let nw = (-1, -1);
    let ne = (-1, 1);
    let sw = (1, -1);
    let se = (1, 1);

    let mut count = 0;

    for i in 0..rows {
        for j in 0..cols {
            if word_search[i as usize][j as usize] == 'A' {
                let nw_pos = (i + nw.0, j + nw.1);
                let ne_pos = (i + ne.0, j + ne.1);
                let sw_pos = (i + sw.0, j + sw.1);
                let se_pos = (i + se.0, j + se.1);

                if [nw_pos, ne_pos, sw_pos, se_pos]
                    .iter()
                    .all(|&(x, y)| in_bounds(x, y, rows, cols))
                    && is_valid_pair(
                        word_search[nw_pos.0 as usize][nw_pos.1 as usize],
                        word_search[se_pos.0 as usize][se_pos.1 as usize],
                    )
                    && is_valid_pair(
                        word_search[ne_pos.0 as usize][ne_pos.1 as usize],
                        word_search[sw_pos.0 as usize][sw_pos.1 as usize],
                    )
                {
                    count += 1;
                }
            }
        }
    }
    count
}

fn in_bounds(row: isize, col: isize, rows: isize, cols: isize) -> bool {
    0 <= row && row < rows && 0 <= col && col < cols
}

fn is_valid_pair(c1: char, c2: char) -> bool {
    (c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M')
}
