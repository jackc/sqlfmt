select
  substring('Thomas' from 2 for 3),
  substring('Thomas' from '...$'),
  substring('Thomas' from '%#"o_a#"_' for '#'),
  substring('Thomas', 2, 3),
  substring()
