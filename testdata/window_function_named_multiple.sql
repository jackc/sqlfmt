select foo, row_number() over w from baz window w as (partition by quz), w2 as (w order by abc)
