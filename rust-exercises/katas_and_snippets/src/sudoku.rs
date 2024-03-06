struct Sudoku {
    data: Vec<Vec<u32>>,
}

impl Sudoku {
    fn is_valid(&self) -> bool {
        let n = self.data.len();
        if n < 1 {
            return false;
        }
        let sqrt_n = (n as f64).sqrt() as usize;
        if sqrt_n * sqrt_n != n {
            return false;
        }

        for i in 0..n {
            if let Some(row) = self.data.get(i) {
                if row.len() != n || !is_valid_line(row) {
                    return false;
                }
                let column = &(0..n).map(|j| self.data[j][i]).collect::<Vec<_>>();
                if column.len() != n || !is_valid_line(column) {
                    return false;
                }
            } else {
                return false;
            }
        }

        for i in 0..sqrt_n {
            for j in 0..sqrt_n {
                let mut square = Vec::with_capacity(n);
                for k in 0..sqrt_n {
                    square.extend_from_slice(
                        &self.data[i * sqrt_n + k][j * sqrt_n..(j + 1) * sqrt_n],
                    );
                }
                if !is_valid_line(&square) {
                    return false;
                }
            }
        }

        true
    }
}

fn is_valid_line(v: &[u32]) -> bool {
    let mut seen = vec![0; v.len()];
    !v.iter().any(|&num| {
        num < 1 || num > v.len() as u32 || {
            seen[num as usize - 1] += 1;
            seen[num as usize - 1] > 1
        }
    })
}

#[test]
fn test_sudoku() {
    let good_sudoku_1 = Sudoku {
        data: vec![
            vec![7, 8, 4, 1, 5, 9, 3, 2, 6],
            vec![5, 3, 9, 6, 7, 2, 8, 4, 1],
            vec![6, 1, 2, 4, 3, 8, 7, 5, 9],
            vec![9, 2, 8, 7, 1, 5, 4, 6, 3],
            vec![3, 5, 7, 8, 4, 6, 1, 9, 2],
            vec![4, 6, 1, 9, 2, 3, 5, 8, 7],
            vec![8, 7, 6, 3, 9, 4, 2, 1, 5],
            vec![2, 4, 3, 5, 6, 1, 9, 7, 8],
            vec![1, 9, 5, 2, 8, 7, 6, 3, 4],
        ],
    };

    let good_sudoku_2 = Sudoku {
        data: vec![
            vec![1, 4, 2, 3],
            vec![3, 2, 4, 1],
            vec![4, 1, 3, 2],
            vec![2, 3, 1, 4],
        ],
    };
    assert!(good_sudoku_1.is_valid());
    assert!(good_sudoku_2.is_valid());

    let bad_sudoku_1 = Sudoku {
        data: vec![
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9],
        ],
    };

    let bad_sudoku_2 = Sudoku {
        data: vec![vec![2]],
    };
    assert!(!bad_sudoku_1.is_valid());
    assert!(!bad_sudoku_2.is_valid());

    let sudoku_x = Sudoku {
        data: vec![
            vec![7, 8, 4, 1, 5, 9, 3, 2, 6],
            vec![5, 3, 9, 6, 7, 2, 8, 4, 1],
            vec![6, 1, 2, 4, 3, 8, 7, 5, 9],
            vec![9, 2, 8, 7, 1, 5, 4, 6, 3],
            vec![3, 5, 7, 8, 4, 6, 1, 9, 2],
            vec![4, 6, 1, 9, 2, 3, 5, 8, 7],
            vec![8, 7, 6, 3, 9, 4, 2, 1, 5],
            vec![2, 4, 3, 5, 6, 1, 9, 7, 8],
            vec![1, 9, 5, 2, 8, 7, 6, 1, 4],
        ],
    };
    assert!(!sudoku_x.is_valid());
}
