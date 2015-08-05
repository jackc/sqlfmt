select '42'::integer, foo::text, (foo+bar)::text, '3.14'::numeric(8,2),
'123.1'::decimal(8,1), '424.234'::dec(8,3),
'324.5'::float(20), '23.23'::double precision,
'asdf'::customtype(3) from baz
