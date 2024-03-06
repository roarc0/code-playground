use itertools::Itertools;

fn keypad_alternatives(d: u32) -> &'static [char] {
    match d {
        0 => &['0', '8'],
        1 => &['1', '2', '4'],
        2 => &['2', '1', '3', '5'],
        3 => &['3', '2', '6'],
        4 => &['4', '1', '5', '7'],
        5 => &['5', '2', '4', '6', '8'],
        6 => &['6', '3', '5', '9'],
        7 => &['7', '4', '8'],
        8 => &['8', '0', '5', '7', '9'],
        9 => &['9', '6', '8'],
        _ => &[],
    }
}

fn get_pins(observed: &str) -> Vec<String> {
    observed
        .chars()
        .filter_map(|c| c.to_digit(10).map(keypad_alternatives))
        .fold(vec![String::new()], |acc, d| {
            d.iter()
                .flat_map(|i| {
                    acc.iter()
                        .map(|c| format!("{}{}", c, i))
                        .collect::<Vec<String>>()
                })
                .collect()
        })
}

#[test]
fn test_get_pins() {
    assert_eq!(
        get_pins("369").iter().sorted().collect::<Vec<&String>>(),
        vec![
            "236", "238", "239", "256", "258", "259", "266", "268", "269", "296", "298", "299",
            "336", "338", "339", "356", "358", "359", "366", "368", "369", "396", "398", "399",
            "636", "638", "639", "656", "658", "659", "666", "668", "669", "696", "698", "699"
        ]
    );
}
