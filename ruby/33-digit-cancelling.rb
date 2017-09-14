#!/usr/bin/env ruby

numerator_product, denominator_product = 1, 1

(10..98).each do |n|
  (n+1..99).each do |d|
    next if n % 10 == 0 && d & 10 == 0
    if (n.to_s[0] == d.to_s[1] && n.to_s[1] < d.to_s[0])
      fraction_as_float = Float(n.to_s[1]) / Float(d.to_s[0])
    elsif (n.to_s[1] == d.to_s[0] && n.to_s[0] < d.to_s[1])
      fraction_as_float = Float(n.to_s[0]) / Float(d.to_s[1])
    else
      next
    end
    if (Float(n) / Float(d) == fraction_as_float)
      numerator_product*=n
      denominator_product*=d
    end
  end
end

puts denominator_product / numerator_product
