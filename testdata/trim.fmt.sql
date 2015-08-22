select
  trim(both 'x' from 'xBobxx'),
  trim(leading 'x' from 'xBobxx'),
  trim(trailing 'x' from 'xBobxx'),
  trim(both from 'xBobxx', 'x'),
  trim(leading from 'xBobxx', 'x'),
  trim(trailing from 'xBobxx', 'x'),
  trim(from 'xBobxx', 'x'),
  trim(from 'xBobxx'),
  trim('xBobxx', 'x'),
  trim('xBobxx')
