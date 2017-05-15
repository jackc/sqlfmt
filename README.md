# sqlfmt

## Installation

```console
$ go get github.com/jackc/sqlfmt/...
$ which sqlfmt
$GOPATH/bin/sqlfmt
```

## Usage

  - You can either:

    + Provide the path to one or more SQL files as command line arguments:

      ```console
      $ sqlfmt testdata/select_where.input.sql
      select
        foo,
        bar
      from
        baz
      where
        foo > 5
        and bar < 2
      ```

    + Or, directly provide the SQL string via stdin:

      ```console
      $ echo "select * from users" | sqlfmt
      select
        *
      from
        users
      ```

      ```console
      $ sqlfmt < testdata/like.input.sql
      select
        foo,
        bar
      from
        baz
      where
        foo like 'abd%'
        or foo like 'ada%' escape '!'
        or foo not like 'abd%'
        or foo not like 'ada%' escape '!'
        or foo ilike 'efg%'
        or foo ilike 'ada%' escape '!'
        or foo not ilike 'efg%'
        or foo not ilike 'ada%' escape '!'
      ```

  - View [testdata](./testdata) for more examples.
