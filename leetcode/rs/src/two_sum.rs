/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Examples:
nums = [2,7,11,15], target = 9 => [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
nums = [3,2,4], target = 6 => [1,2]
nums = [3,3], target = 6 => [0,1]
*/

use std::collections::HashMap;

struct Solution {}

impl Solution {
    // pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    //     for i in 0..nums.len() {
    //         for j in 0..nums.len() {
    //             if nums[i] + nums[j] == target {
    //                 return vec![i as i32, j as i32];
    //             }
    //         }
    //     }
    //     vec![]
    // }

    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = HashMap::new();
        for (i, n) in nums.iter().enumerate() {
            let complement = target - n;
            if let Some(&j) = m.get(&complement) {
                return vec![j as i32, i as i32];
            }
            m.insert(n, i);
        }
        vec![]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct TestInputs {
        nums: Vec<i32>,
        target: i32,
        want: Vec<i32>,
    }

    #[test]
    fn two_sums_multi_test() {
        let test_inputs = vec![TestInputs {
            nums: vec![2, 7, 11, 15],
            target: 9,
            want: vec![0, 1],
        }];

        for t in test_inputs {
            assert_eq!(t.want, Solution::two_sum(t.nums, t.target));
        }
    }
}
