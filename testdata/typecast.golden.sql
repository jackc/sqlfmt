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
  'asdf'::myschema.customtype,
  '1942'::setof int,
  '{123,34}'::int array[4],
  '{123,34}'::setof int array[4],
  '{123,34}'::int array,
  '{123,34}'::setof int array,
  'f'::char,
  'fads'::varchar,
  'fads'::char(10),
  'fads'::varchar(10),
  'f'::char,
  'fads'::varchar,
  'fads'::char(10),
  'fads'::varchar(10),
  'f'::char,
  'fads'::varchar,
  'f'::char,
  'fads'::char(10),
  'asdf'::varchar character set sql_text,
  '1'::bit,
  '1010'::bit(4),
  '1010'::varbit,
  '1010'::varbit,
  '00:30:00'::interval hour to minute,
  '00:15:00'::interval(2)
from
  baz
