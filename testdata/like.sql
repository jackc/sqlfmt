select foo, bar from baz
where
foo like 'abd%' or foo like 'ada%' escape '!' or
foo not like 'abd%' or foo not like 'ada%' escape '!'
or foo ilike 'efg%' or foo ilike 'ada%' escape '!'
or foo not ilike 'efg%' or foo not ilike 'ada%' escape '!'
