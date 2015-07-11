select
  case
  when foo = bar then
    7
  when foo > bar then
    42
  else
    1
  end
from
  baz
