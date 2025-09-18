let leap_year year =
  if year mod 400 = 0 then true
  else if year mod 4 = 0 then year mod 100 != 0
  else false
