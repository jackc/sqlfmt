select
  3 > any (select
    generate_series(1, 10)
  )
  ,
  3 > all (select
    generate_series(1, 10)
  )
  ,
  3 > any (array[1, 2, 3, 4]),
  3 operator(>) any (array[1, 2, 3, 4])
