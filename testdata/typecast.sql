select '42'::integer, foo::text, (foo+bar)::text, '3.14'::numeric(8,2),
'123.1'::decimal(8,1), '424.234'::dec(8,3),
'324.5'::float(20), '23.23'::double precision,
'asdf'::customtype(3), '1942'::setof int,
'{123,34}'::int array[4], '{123,34}'::setof int array[4],
'{123,34}'::int array, '{123,34}'::setof int array,
'f'::character, 'fads'::character varying,
'fads'::character(10), 'fads'::character varying(10),
'f'::char, 'fads'::char varying,
'fads'::char(10), 'fads'::char varying(10),
'f'::national character, 'fads'::national character varying,
'f'::nchar, 'fads'::nchar(10),
'asdf'::varchar character set sql_text,
'1'::bit, '1010'::bit(4), '1010'::bit varying, '1010'::varbit

 from baz
