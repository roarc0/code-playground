use itertools::Itertools;

fn count_sheep(n: u32) -> String {
    (1..=n).map(|i| format!("{i:?} sheep...")).collect()
}

fn to_camel_case(text: &str) -> String {
    text.split(|c| c == '-' || c == '_')
        .enumerate()
        .map(|(index, s)| match index {
            0 => s.to_string(),
            _ => s[..1].to_uppercase() + &s[1..],
        })
        .collect()
}

fn rgb(r: i32, g: i32, b: i32) -> String {
    format!(
        "{:02X}{:02X}{:02X}",
        r.clamp(0, 255),
        g.clamp(0, 255),
        b.clamp(0, 255)
    )
}

fn find_smallest_int(arr: &[i32]) -> i32 {
    *arr.iter().min().unwrap_or(&0)
}

fn rps(p1: &str, p2: &str) -> &'static str {
    match (p1, p2) {
        ("paper", "rock") | ("rock", "scissors") | ("scissors", "paper") => "Player 1 won!",
        ("paper", "scissors") | ("rock", "paper") | ("scissors", "rock") => "Player 2 won!",
        ("paper", "paper") | ("rock", "rock") | ("scissors", "scissors") => "Draw!",
        _ => "?",
    }
}

fn paperwork(n: i16, m: i16) -> u32 {
    if n < 0 || m < 0 {
        0
    } else {
        n as u32 * m as u32
    }
}

fn remove_every_other(arr: &[u8]) -> Vec<u8> {
    arr.iter()
        .enumerate()
        .filter(|i| i.0 % 2 == 0)
        .map(|i| *i.1)
        .collect()
}

fn minimum_perimeter(area: u64) -> u64 {
    (1..=((area as f64).sqrt() as u64))
        .filter(|&i| area % i == 0)
        .map(|i| 2 * (i + area / i))
        .min()
        .unwrap_or(0)
}

fn unique_in_order<T>(sequence: T) -> Vec<T::Item>
where
    T: std::iter::IntoIterator,
    T::Item: PartialEq + Clone + std::fmt::Debug,
{
    sequence.into_iter().dedup().collect()
}

fn keypad_alternatives(c: char) -> Option<&'static [char]> {
    match c {
        '0' => Some(&['0', '8']),
        '1' => Some(&['1', '2', '4']),
        '2' => Some(&['2', '1', '3', '5']),
        '3' => Some(&['3', '2', '6']),
        '4' => Some(&['4', '1', '5', '7']),
        '5' => Some(&['5', '2', '4', '6', '8']),
        '6' => Some(&['6', '3', '5', '9']),
        '7' => Some(&['7', '4', '8']),
        '8' => Some(&['8', '0', '5', '7', '9']),
        '9' => Some(&['9', '6', '8']),
        _ => None,
    }
}

fn get_pins(observed: &str) -> Vec<String> {
    observed.chars().filter_map(keypad_alternatives).fold(
        Vec::from([String::from("")]),
        |acc, d| {
            d.iter()
                .flat_map(|i| acc.iter().map(move |c| format!("{}{}", c, i)))
                .collect()
        },
    )
}

fn switcheroo(s: &str) -> String {
    s.chars()
        .map(|c| match c {
            'a' => 'b',
            'b' => 'a',
            _ => c,
        })
        .collect()
}

fn rot13(message: &str) -> String {
    message
        .chars()
        .map(|c| match c {
            'A'..='M' | 'a'..='m' => (c as u8 + 13) as char,
            'N'..='Z' | 'n'..='z' => (c as u8 - 13) as char,
            _ => c,
        })
        .collect()
}

use ::num::BigInt;

fn increment_string(s: &str) -> String {
    let split_index = s
        .chars()
        .rev()
        .position(|c| !c.is_ascii_digit())
        .map(|split_index| s.len() - split_index)
        .unwrap_or(0);

    let num_str = s.get(split_index..).unwrap_or("0");
    let num = num_str.parse::<BigInt>().unwrap_or(BigInt::default()) + BigInt::from(1);

    format!(
        "{}{:0>width$}",
        s.get(..split_index).unwrap_or_default(),
        num.to_string(),
        width = num_str.len()
    )
}

use ::num::integer::gcd;

fn nbr_of_laps(x: u16, y: u16) -> (u16, u16) {
    let gcd = gcd(x, y);
    (y / gcd, x / gcd)
}

fn what_century(year: &str) -> String {
    let c = (year.parse::<f32>().ok().unwrap_or_default() / 100.).ceil() as i32;
    let suffix = match (c % 10) * { !(11..=19).contains(&c) as i32 } {
        1 => "st",
        2 => "nd",
        3 => "rd",
        _ => "th",
    };
    format!("{c}{suffix}")
}

fn disemvowel(s: &str) -> String {
    s.chars()
        .filter(|&c| !c.is_ascii_alphabetic() || !"aeiou".contains(c.to_ascii_lowercase()))
        .collect()
}

fn is_square(n: i64) -> bool {
    n >= 0 && (n as f64).sqrt().fract() == 0.0
}

fn stray(arr: &[u32]) -> u32 {
    if arr[0] != arr[1] && arr[0] != arr[2] {
        return arr[0];
    }
    for i in 1..arr.len() {
        if arr[0] != arr[i] {
            return arr[i];
        }
    }
    unreachable!();
}

use time::{Duration, PrimitiveDateTime as DateTime};

fn after(start: DateTime) -> DateTime {
    start + Duration::seconds(1000000000)
}

pub fn is_armstrong_number(num: u32) -> bool {
    let num_str = num.to_string();
    let num_len = num_str.len() as u32;
    num as u64
        == num_str
            .chars()
            .filter_map(|c| c.to_digit(10))
            .map(|i| (i as u64).pow(num_len))
            .sum()
}
