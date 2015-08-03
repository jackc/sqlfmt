select (array['a', 'b', 'c', foo, bar])[1], quz[42],
(select array['a', 'b', 'c'])[1] from baz
