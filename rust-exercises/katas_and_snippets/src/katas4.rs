fn fold_array(arr: &[i32], runs: usize) -> Vec<i32> {
    if runs == 0 || arr.len() < 2 {
        return arr.to_owned();
    }
    let length = arr.len();
    let mid = length / 2;
    let mut ret = arr[..mid].to_owned();
    for i in 0..mid {
        ret[i] += arr[length - 1 - i];
    }
    if length % 2 == 1 {
        ret.push(arr[mid]);
    }
    fold_array(ret.as_slice(), runs - 1)
}

fn fold_array2(arr: &[i32], runs: usize) -> Vec<i32> {
    if runs == 0 || arr.len() < 2 {
        return arr.to_owned();
    }
    let len = arr.len();
    let ret = arr[..len / 2]
        .iter()
        .enumerate()
        .map(|(i, &x)| x + arr[len - i - 1])
        .chain(if len % 2 == 1 {
            Some(arr[len / 2])
        } else {
            None
        })
        .collect::<Vec<i32>>();
    fold_array(ret.as_slice(), runs - 1)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic() {
        let input = [1, 2, 3, 4, 5];
        assert_eq!(fold_array(&input, 1), [6, 6, 3]);
        assert_eq!(fold_array(&input, 2), [9, 6]);
        assert_eq!(fold_array(&input, 3), [15]);
        let input = [-9, 9, -8, 8, 66, 23];
        assert_eq!(fold_array(&input, 1), [14, 75, 0]);
        let input: [i32; 0] = [];
        assert_eq!(fold_array(&input, 1), []);
        let input: [i32; 1] = [1];
        assert_eq!(fold_array(&input, 1), [1]);
    }
}
