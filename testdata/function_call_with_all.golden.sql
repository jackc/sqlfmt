select
  name,
  array_agg(foo)
from
  baz
group by
  name
