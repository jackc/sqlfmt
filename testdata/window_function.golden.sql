select
  foo,
  row_number() over ()
from
  baz
