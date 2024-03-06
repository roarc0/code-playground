fn delete_nth(lst: &[u8], n: usize) -> Vec<u8> {
    let mut counts = std::collections::HashMap::new();
    lst.iter()
        .filter(|&num| {
            let c = counts.entry(*num).or_insert(0);
            if *c < n {
                *c += 1;
                true
            } else {
                false
            }
        })
        .cloned()
        .collect()
}

fn validate_pin(pin: &str) -> bool {
    !((pin.len() != 4 && pin.len() != 6) || pin.chars().any(|c| !c.is_ascii_digit()))
}

fn number(bus_stops: &[(i32, i32)]) -> i32 {
    let total = bus_stops
        .iter()
        .fold((0, 0), |acc, v| (acc.0 + v.0, acc.1 + v.1));
    total.0 - total.1
}
