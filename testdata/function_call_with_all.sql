select name, array_agg(all foo) from baz group by name
