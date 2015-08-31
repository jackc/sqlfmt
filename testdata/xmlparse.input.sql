select xmlparse(document '<?xml version="1.0"?><person><name>John</name></person>'),
xmlparse(content '<?xml version="1.0"?><person><name>John</name></person>'),
xmlparse(content '<?xml version="1.0"?><person><name>John</name></person>' preserve whitespace),
xmlparse(content '<?xml version="1.0"?><person><name>John</name></person>' strip whitespace)
