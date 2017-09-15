#!/usr/bin/env ruby
require 'mathn'

def circular_prime?(i)
  nums = i.to_s.chars.to_a

  nums.length.times do
    return false if (!nums.join().to_i.prime?)
    nums.push nums[0]
    nums.shift
  end

  return true
end

answer = 0
Prime.each { |x|
  break if x >= 1000000
  answer += 1 if circular_prime?(x)
}

puts answer
