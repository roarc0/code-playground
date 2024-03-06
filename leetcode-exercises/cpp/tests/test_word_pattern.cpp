#include <gtest/gtest.h>

#include <string>

#include "word_pattern.hpp"

TEST(WordPattern, MultiTest) {
  struct TestInputs {
    std::string pattern;
    std::string s;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"a", "a", true},
      {"abba", "dog dog dog dog", false},
      {"abba", "dog cat cat dog", true},
      {"abba", "dog cat cat fish", false},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.wordPattern(t.pattern, t.s))
        << "Failed for pattern: " << t.pattern << ", s: " << t.s;
  }
}