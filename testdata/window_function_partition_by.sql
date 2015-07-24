select foo, row_number() over (partition by quz) from baz
