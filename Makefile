# This file copie and modified from https://github.com/miekg/learninggo/blob/master/Makefile
MMARK=mmark

all: learninggo.html

.PHONY: learninggo.html
learninggo.html:
	$(MMARK) -html -head inc/head.html -css inc/book.css learninggo.md > learninggo.html

.PHONY: learninggo.xml
learninggo.xml:
	$(MMARK) learninggo.md > learninggo.xml

.PHONY: learninggo.txt
learninggo.txt: learninggo.xml
	xml2rfc --v3 learninggo.xml

.PHONY: learninggo-2.xml
learninggo-2.xml:
	$(MMARK) -2 learninggo.md > learninggo-2.xml

.PHONY: learninggo-2.txt
learninggo-2.txt: learninggo-2.xml
	@# Using -n because it doesn't fully validate (RFC7749 is a limited format).
	xml2rfc -n learninggo-2.xml

.PHONY: ast
ast:
	$(MMARK) -ast learninggo.md

.PHONY: test
test:
	$(MMARK) -html learninggo.md

clean:
	rm -f learninggo.html
