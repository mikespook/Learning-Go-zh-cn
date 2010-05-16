# simple makefile
.PHONY: fig fig/*.svg tab tab/*.tex once

all:	fig tab go.pdf

once:
	xelatex go.tex

go.pdf:	go.tex go-*.tex ex-*/*.tex src/*.go blocksbook.cls go.bib
	xelatex go.tex && bibtex go && makeindex go && xelatex go.tex && xelatex go.tex

fig:	fig/*.svg
	( cd fig; make all )

tab:	tab/*.tex
	( cd tab; make all )

clean:
	rm -f go.lol go.aux *.log map.log go.pdf go.bbl go.blg go.toc go.ind go.lot go.lof
	rm -f go.ilg go.idx go.lgpl missfont.log doc_data.txt go.ex

distclean: clean
	( cd fig; make clean )
	( cd tab; make clean )
	( cd src; make clean )
