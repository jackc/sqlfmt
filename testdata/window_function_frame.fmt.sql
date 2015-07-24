select
  foo,
  row_number() over (range unbounded preceding),
  row_number() over (rows unbounded preceding),
  row_number() over (range between unbounded preceding and 3 following),
  row_number() over (rows between unbounded preceding and 3 following),
  row_number() over (range current row),
  row_number() over (rows current row),
  row_number() over (range between 2 preceding and unbounded following),
  row_number() over (rows between 2 preceding and unbounded following)
from
  baz
