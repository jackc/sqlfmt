select name, array_agg(foo) filter (where a=b) from baz group by name
