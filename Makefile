.PHONY: once

all:	go.pdf

once:	
	xelatex go.tex

go.pdf:	go-*.tex ex-*/*.tex src/*.go tab/*.tex fig/*.tex blocksbook.cls go.bib .fig .tab go.tex
	xelatex go.tex && bibtex go && makeindex go && xelatex go.tex && xelatex go.tex

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
