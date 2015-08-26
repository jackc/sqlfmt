select percentile_disc(0.25) within group (order by n) from generate_series(1,10) n
