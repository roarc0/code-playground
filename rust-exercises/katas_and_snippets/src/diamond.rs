fn step_diamond(grow: bool, spaces: &mut usize, stars: &mut usize) {
    let sign: i32 = match grow {
        true => 1,
        false => -1,
    };
    *spaces = (*spaces as i32 - sign) as usize;
    *stars = (*stars as i32 + 2 * sign) as usize;
}

fn print_diamond(n: i32) -> Option<String> {
    if n <= 0 || n % 2 == 0 {
        return None;
    }
    let mut spaces: usize = (n / 2).try_into().unwrap();
    let mut stars: usize = 1;
    let mut grow: Option<bool> = Some(true);
    let mut diamond = String::from("");
    let mut line = 0;
    while line < n {
        let l: String = [vec![' '; spaces], vec!['*'; stars]]
            .concat()
            .into_iter()
            .collect();
        diamond += l.as_str();

        match grow {
            Some(true) => {
                if spaces == 0 {
                    grow = Some(false);
                    if stars >= 2 {
                        step_diamond(false, &mut spaces, &mut stars);
                    }
                } else {
                    step_diamond(true, &mut spaces, &mut stars);
                }
            }
            Some(false) => {
                if stars <= 2 {
                    grow = None;
                } else {
                    step_diamond(false, &mut spaces, &mut stars);
                }
            }
            None => {}
        };
        diamond += "\n";
        line += 1;
    }
    Some(diamond)
}
