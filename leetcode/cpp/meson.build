project('leetcodes', 'cpp',
  version : '1.0.0',
  license : 'MIT',
  meson_version : '>= 0.50.0',
  default_options : ['warning_level=3', 'buildtype=debug', 'cpp_std=c++23']
)

sources = []

subdir('src')
subdir('tests')

gtest_dep = dependency('gtest')

executable(
  'leetcode-tests',
  sources,
  include_directories: 'include',
  dependencies: [gtest_dep]
)
