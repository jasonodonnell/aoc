use std::fs;

fn main() {
    let data = fs::read_to_string("input.txt").expect("Failed to read input file");
    println!("{}", part_one(&data));
    println!("{}", part_two(&data))
}

#[derive(Debug)]
enum ParseState {
    Start,
    M,
    MU,
    MUL,
    OpenParen,
    NumberOne,
    Comma,
    NumberTwo,
}

fn parse_commands(input: &str, toggle_process: bool) -> Vec<&str> {
    let mut state: ParseState = ParseState::Start;
    let chars = input.chars().enumerate();
    let mut commands: Vec<&str> = vec![];
    let mut start: usize = 0;
    let mut process = true;

    for (i, c) in chars {
        if toggle_process {
            if let Some(slice) = input.get(i..i + 5) {
                if slice == "don't" {
                    process = false;
                    continue;
                }
            }
            if let Some(slice) = input.get(i..i + 2) {
                if slice == "do" {
                    process = true;
                    continue;
                }
            }
        }

        if process {
            state = match (state, c) {
                (ParseState::Start, 'm') => {
                    start = i;
                    ParseState::M
                }
                (ParseState::M, 'u') => ParseState::MU,
                (ParseState::MU, 'l') => ParseState::MUL,
                (ParseState::MUL, '(') => ParseState::OpenParen,
                (ParseState::OpenParen, c) if c.is_numeric() => ParseState::NumberOne,
                (ParseState::NumberOne, c) if c.is_numeric() => ParseState::NumberOne,
                (ParseState::NumberOne, ',') => ParseState::Comma,
                (ParseState::Comma, c) if c.is_numeric() => ParseState::NumberTwo,
                (ParseState::NumberTwo, c) if c.is_numeric() => ParseState::NumberTwo,
                (ParseState::NumberTwo, ')') => {
                    commands.push(&input[start..=i]);
                    start = 0;
                    ParseState::Start
                }
                _ => {
                    start = 0;
                    ParseState::Start
                }
            }
        }
    }
    commands
}

fn calculate_sum(commands: &[&str]) -> i32 {
    commands
        .iter()
        .filter_map(|command| {
            command
                .strip_prefix("mul(")
                .and_then(|s| s.strip_suffix(')'))
                .and_then(|inner| inner.split_once(','))
                .and_then(|(num1, num2)| num1.parse::<i32>().ok().zip(num2.parse::<i32>().ok()))
                .map(|(n1, n2)| n1 * n2)
        })
        .sum()
}

fn part_one(input: &str) -> i32 {
    let commands = parse_commands(input, false);
    calculate_sum(&commands)
}

fn part_two(input: &str) -> i32 {
    let commands = parse_commands(input, true);
    calculate_sum(&commands)
}
