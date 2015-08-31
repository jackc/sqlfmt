select extract(year from '2000-01-01 12:34:56'::timestamptz),
extract(month from '2000-01-01 12:34:56'::timestamptz),
extract(day from '2000-01-01 12:34:56'::timestamptz),
extract(hour from '2000-01-01 12:34:56'::timestamptz),
extract(minute from '2000-01-01 12:34:56'::timestamptz),
extract(second from '2000-01-01 12:34:56'::timestamptz),
extract('second' from '2000-01-01 12:34:56'::timestamptz),
extract("second" from '2000-01-01 12:34:56'::timestamptz)
