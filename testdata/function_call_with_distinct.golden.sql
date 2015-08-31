select
  name,
  array_agg(distinct foo)
from
  baz
group by
  name
