fn array_diff<T: PartialEq>(a: Vec<T>, b: Vec<T>) -> Vec<T> {
    a.into_iter().filter(|e| !b.contains(e)).collect()
}

#[test]
fn test_array_diff() {
    assert_eq!(array_diff(vec![1, 2, 3, 4], vec![2, 3, 6]), vec![1, 4]);
}

fn square_sum(vec: Vec<i32>) -> i32 {
    vec.iter().map(|x| x * x).sum()
}

#[test]
fn test_square_sum() {
    assert_eq!(square_sum(vec![1, 2, 3, 4]), 30);
}

fn digitize(n: u64) -> Vec<u8> {
    n.to_string()
        .chars()
        .rev()
        .map(|c| match c.to_digit(10) {
            Some(c) => c as u8,
            None => 0,
        })
        .collect()
}

fn fake_bin(s: &str) -> String {
    s.chars()
        .map(|c| match c.to_digit(10).unwrap_or(0) >= 5 {
            true => '1',
            false => '0',
        })
        .collect()
}

fn printer_error(s: &str) -> String {
    format!("{}/{}", s.chars().filter(|c| c > &'m').count(), s.len())
}

fn binary_slice_to_number(slice: &[u32]) -> u32 {
    let s: String = slice.iter().map(|&x| x.to_string()).collect();
    u32::from_str_radix(s.as_str(), 2).unwrap_or(0)
}

fn are_you_playing_banjo(name: &str) -> String {
    match &name[0..1] {
        "r" | "R" => format!("{name} plays banjo"),
        _ => format!("{name} does not play banjo"),
    }
}

use itertools::Itertools;
use regex::Regex;

fn camel_casing_splitter(s: &str) -> String {
    let regex = Regex::new(r#"([A-Z])([a-z0-9]+)"#).unwrap();
    return String::from(regex.replace_all(s, " $1$2"));
}

#[test]
fn test_camel_case_splitter() {
    assert_eq!(
        camel_casing_splitter("camelCasingTest"),
        "camel Casing Test"
    );
}

fn bouncing_ball(h: f64, bounce: f64, window: f64) -> i32 {
    if h < 0. || !(0. ..h).contains(&window) || !(0. ..1.).contains(&bounce) {
        return -1;
    }
    let next_h = bounce * h;
    match next_h > window {
        true => 2 + bouncing_ball(next_h, bounce, window),
        false => 1,
    }
}

fn count_duplicates(text: &str) -> u32 {
    let dups: Vec<(char, usize)> = text
        .chars()
        .map(|c| c.to_ascii_lowercase())
        .sorted()
        .group_by(|&c| c)
        .into_iter()
        .map(|(key, group)| (key, group.count()))
        .filter(|&item| item.1 > 1)
        .collect();
    dups.len() as u32
}

fn spin_words(words: &str) -> String {
    words
        .split_whitespace()
        .map(|s| match s.len() >= 5 {
            true => s.chars().rev().collect(),
            false => s.to_string(),
        })
        .collect::<Vec<_>>()
        .join(" ")
}

#[test]
fn test_spin_words() {
    assert_eq!(
        spin_words("Hello world this is a wonderful day This is a test"),
        "olleH dlrow this is a lufrednow day This is a test"
    );
}
