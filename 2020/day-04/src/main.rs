use std::fs::File;
use std::io::{self, BufRead};

#[allow(dead_code)]
#[derive(Debug)]
struct Passport {
    birth_year: Option<i32>,      // byr
    issue_year: Option<i32>,      // iyr
    expiration_year: Option<i32>, // eyr
    height: Option<String>,       // hgt
    hair_color: Option<String>,   // hcl
    eye_color: Option<String>,    // ecl
    id: Option<String>,           // pid
    country_id: Option<i32>,      // cid
}

impl Passport {
    fn from_record(record: &str) -> Passport {
        let mut passport = Passport {
            birth_year: None,
            issue_year: None,
            expiration_year: None,
            height: None,
            hair_color: None,
            eye_color: None,
            id: None,
            country_id: None,
        };

        for fields in record.split_whitespace() {
            let mut parts = fields.split(':');
            let key = parts.next().unwrap();
            let value = parts.next().unwrap();

            match key {
                "byr" => passport.birth_year = value.parse().ok(),
                "iyr" => passport.issue_year = value.parse().ok(),
                "eyr" => passport.expiration_year = value.parse().ok(),
                "hgt" => passport.height = Some(value.to_string()),
                "hcl" => passport.hair_color = Some(value.to_string()),
                "ecl" => passport.eye_color = Some(value.to_string()),
                "pid" => passport.id = Some(value.to_string()),
                "cid" => passport.country_id = value.parse().ok(),
                _ => {}
            }
        }
        passport
    }

    fn valid_fields(&self) -> bool {
        self.birth_year.is_some()
            && self.issue_year.is_some()
            && self.expiration_year.is_some()
            && self.height.is_some()
            && self.hair_color.is_some()
            && self.eye_color.is_some()
            && self.id.is_some()
    }

    fn valid_values(&self) -> bool {
        self.valid_birth_year()
            && self.valid_issue_year()
            && self.valid_expiration_year()
            && self.valid_height()
            && self.valid_hair_color()
            && self.valid_eye_color()
            && self.valid_pid()
    }

    fn valid_birth_year(&self) -> bool {
        match self.birth_year {
            Some(byr) if byr >= 1920 && byr <= 2002 => true,
            _ => false,
        }
    }

    fn valid_issue_year(&self) -> bool {
        match self.issue_year {
            Some(iyr) if iyr >= 2010 && iyr <= 2020 => true,
            _ => false,
        }
    }

    fn valid_expiration_year(&self) -> bool {
        match self.expiration_year {
            Some(eyr) if eyr >= 2020 && eyr <= 2030 => true,
            _ => false,
        }
    }

    fn valid_height(&self) -> bool {
        if let Some(ref hgt) = self.height {
            if hgt.ends_with("cm") {
                hgt.trim_end_matches("cm")
                    .parse::<i32>()
                    .map_or(false, |value| value >= 150 && value <= 193)
            } else if hgt.ends_with("in") {
                hgt.trim_end_matches("in")
                    .parse::<i32>()
                    .map_or(false, |value| value >= 59 && value <= 76)
            } else {
                false
            }
        } else {
            false
        }
    }

    fn valid_hair_color(&self) -> bool {
        if let Some(ref hcl) = self.hair_color {
            hcl.starts_with('#') && hcl.len() == 7 && hcl.chars().skip(1).all(|c| c.is_digit(16))
        } else {
            false
        }
    }

    fn valid_eye_color(&self) -> bool {
        if let Some(ref ecl) = self.eye_color {
            matches!(
                ecl.as_str(),
                "amb" | "blu" | "brn" | "gry" | "grn" | "hzl" | "oth"
            )
        } else {
            false
        }
    }

    fn valid_pid(&self) -> bool {
        self.id
            .as_ref()
            .map(|pid| pid.len() == 9 && pid.chars().all(|c| c.is_digit(10)))
            .unwrap_or(false)
    }
}

fn main() -> io::Result<()> {
    let file = File::open("input.txt").expect("could not open file");
    let reader = io::BufReader::new(file);
    let records = parse_records(reader.lines().map(|line| line.unwrap()))?;

    println!("{}", part_one(&records));
    println!("{}", part_two(&records));
    Ok(())
}

fn parse_records<I>(lines: I) -> io::Result<Vec<String>>
where
    I: IntoIterator,
    I::Item: AsRef<str>,
{
    let mut records: Vec<String> = vec![];
    let mut current_record = String::new();

    for line in lines {
        let line = line.as_ref();
        if line.trim().is_empty() {
            if !current_record.is_empty() {
                records.push(std::mem::take(&mut current_record));
            }
        } else {
            if !current_record.is_empty() {
                current_record.push(' ');
            }
            current_record.push_str(&line);
        }
    }

    if !current_record.is_empty() {
        records.push(current_record);
    }

    Ok(records)
}

fn part_one(records: &Vec<String>) -> usize {
    records
        .iter()
        .filter(|record| Passport::from_record(record).valid_fields())
        .count()
}

fn part_two(records: &Vec<String>) -> usize {
    records
        .iter()
        .filter_map(|record| {
            let passport = Passport::from_record(record);
            if passport.valid_fields() {
                Some(passport)
            } else {
                None
            }
        })
        .filter(|passport| passport.valid_values())
        .count()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "\
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

cid:62
pid:651390352
hcl:#efcc98
iyr:2018
eyr:2027
ecl:brn
hgt:66in
byr:1953

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
";

        let lines = input.lines();
        let records = parse_records(lines).unwrap();
        assert_eq!(part_one(&records), 4)
    }

    #[test]
    fn part_two_example() {
        let input = "\
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
";

        let lines = input.lines();
        let records = parse_records(lines).unwrap();
        assert_eq!(part_two(&records), 4)
    }
}
