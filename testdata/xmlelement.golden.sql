select
  xmlelement(name foo),
  xmlelement(name foo, xmlattributes('bar' as baz)),
  xmlelement(name foo, xmlattributes(bar, baz)),
  xmlelement(name foo, xmlattributes('bar' as baz), 'bo', 'dy'),
  xmlelement(name foo, 'bo', 'dy')
