select foo, row_number() over w from baz window w as (partition by quz order by abc)
