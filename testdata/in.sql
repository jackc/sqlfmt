select 2 in (1,2,3), 2 not in (1,2,3),
  2 in (select generate_series(1,10)), 2 not in (select generate_series(1,10))
