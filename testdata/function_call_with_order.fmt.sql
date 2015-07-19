select
  name,
  array_agg(foo order by
    bar
  )
from
  baz
group by
  name
