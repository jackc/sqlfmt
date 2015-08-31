select xmlroot(xmlparse(document '<?xml version="1.1"?><content>abc</content>'), version '1.0', standalone yes),
xmlroot(xmlparse(document '<?xml version="1.1"?><content>abc</content>'), version '1.0', standalone no),
xmlroot(xmlparse(document '<?xml version="1.1"?><content>abc</content>'), version '1.0', standalone no value),
xmlroot(xmlparse(document '<?xml version="1.1"?><content>abc</content>'), version '1.0'),
xmlroot(xmlparse(document '<?xml version="1.1"?><content>abc</content>'), version no value)
