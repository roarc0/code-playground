#include <gtest/gtest.h>

#include <string>
#include <vector>

#include "is_anagram.hpp"

TEST(IsAnagram, MultiTest) {
  struct TestInputs {
    std::string word1;
    std::string word2;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"cat", "rat", false},
      {"listen", "silent", true},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.isAnagram(t.word1, t.word2));
  }
}