select
  foo,
  array(select
    bar
  from
    quz
  where
    baz.foo = quz.foo
  )
from
  baz
