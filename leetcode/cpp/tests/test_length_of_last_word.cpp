#include <gtest/gtest.h>

#include <vector>

#include "length_of_last_word.hpp"

TEST(LengthOfLastWord, MultiTest) {
  struct TestInputs {
    std::string str;
    int want;
  };
  std::vector<TestInputs> testInputs = {
      {"a", 1},
      {"Hello World", 5},
      {"Hello    World    ", 5},
      {"   fly me   to   the moon  ", 4},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.lengthOfLastWord(t.str));
  }
}