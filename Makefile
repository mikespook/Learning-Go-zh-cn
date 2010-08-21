.PHONY: once spell codecheck

all:	go.pdf

go.pdf:	go-*.tex ex-*/*.tex src/*.go tab/*.tex fig/*.tex blocksbook.cls go.bib .fig .tab go.tex
	xelatex go.tex && bibtex go && makeindex go \
	&& xelatex go.tex && xelatex go.tex

.fig:	fig/*.svg
	( cd fig; make all )
	touch .fig

.tab:	
	( cd tab; make all )
	touch .tab

clean:
	rm -f go.lol go.aux *.log map.log go.pdf go.bbl go.blg go.toc go.ind go.lot go.lof go.loe
	rm -f go.ilg go.idx go.lgpl missfont.log doc_data.txt go.ex .fig .tab

distclean: clean
	( cd fig; make clean )
	( cd tab; make clean )
	( cd src; make clean )

spell:
	for i in *.tex ex-*/*.tex; do aspell check $$i; done
once:	
	xelatex go.tex

codecheck:
	bin/go-lstinputlisting.pl ~/git/gobook   *.tex
	bin/go-lstinputlisting.pl ~/git/gobook   ex-*/*.tex
	rm -f *.6 *.8
