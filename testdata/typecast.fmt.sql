select
  '42'::integer,
  foo::text,
  (foo + bar)::text,
  '3.14'::numeric(8, 2),
  '123.1'::decimal(8, 1),
  '424.234'::dec(8, 3),
  '324.5'::float(20),
  '23.23'::double precision,
  'asdf'::customtype(3),
  '1942'::setof int,
  '{123,34}'::int array[4],
  '{123,34}'::setof int array[4],
  '{123,34}'::int array,
  '{123,34}'::setof int array
from
  baz
