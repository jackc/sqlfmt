select xmlexists('//town[text() = ''Toronto'']' passing '<towns><town>Toronto</town><town>Ottawa</town></towns>'),
xmlexists('//town[text() = ''Toronto'']' passing by ref '<towns><town>Toronto</town><town>Ottawa</town></towns>' by ref)
