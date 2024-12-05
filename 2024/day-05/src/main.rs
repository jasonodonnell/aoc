#![allow(dead_code)]
use std::collections::HashMap;
use std::fs;

fn main() {
    let data = fs::read_to_string("input.txt").expect("Failed to read input file");
    let parts: Vec<_> = data.split("\n\n").collect();
    let mut order_map: HashMap<i32, Vec<i32>> = HashMap::new();

    for line in parts[0].lines() {
        if let Some((key, value)) = line.split_once('|') {
            let key: i32 = key.parse().unwrap();
            let value: i32 = value.parse().unwrap();
            order_map.entry(key).or_default().push(value);
        }
    }

    let updates: Vec<Vec<i32>> = parts[1]
        .lines()
        .map(|line| {
            line.split(',')
                .map(|num| num.trim().parse::<i32>().unwrap())
                .collect()
        })
        .collect();

    println!("{}", part_one(order_map.clone(), &updates));
    println!("{}", part_two(order_map.clone(), &updates));
}

fn part_one(order_map: HashMap<i32, Vec<i32>>, updates: &[Vec<i32>]) -> i32 {
    updates
        .iter()
        .map(|pages| {
            let mut middle_page = pages[pages.len() / 2];
            let mut pages = pages.clone();
            while let Some(page) = pages.pop() {
                let exists = order_map
                    .get(&page)
                    .map_or(false, |vec| pages.iter().any(|value| vec.contains(value)));

                if exists {
                    middle_page = 0;
                    break;
                }
            }
            middle_page
        })
        .sum::<i32>()
}

fn part_two(order_map: HashMap<i32, Vec<i32>>, updates: &[Vec<i32>]) -> i32 {
    updates
        .iter()
        .filter_map(|pages| {
            let mut original_pages = pages.clone();
            let mut swaps_occurred = false;

            loop {
                let mut pages = original_pages.clone();
                let mut local_swaps = false;

                while let Some(page) = pages.pop() {
                    if let Some(vec) = order_map.get(&page) {
                        if let Some(&value) = pages.iter().find(|value| vec.contains(value)) {
                            if let (Some(page_idx), Some(value_idx)) = (
                                original_pages.iter().position(|&x| x == page),
                                original_pages.iter().position(|&x| x == value),
                            ) {
                                original_pages.swap(page_idx, value_idx);
                                local_swaps = true;
                                break;
                            }
                        }
                    }
                }

                if !local_swaps {
                    break;
                }
                swaps_occurred = true;
            }

            if swaps_occurred {
                Some(original_pages[original_pages.len() / 2])
            } else {
                None
            }
        })
        .sum::<i32>()
}
