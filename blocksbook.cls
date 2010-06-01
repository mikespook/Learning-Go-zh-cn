\NeedsTeXFormat{LaTeX2e}[1995/12/01]
\ProvidesClass{blocksbook}[2006/06/10 v0.1
        Layout with blocks]
\RequirePackage{ifthen}
\RequirePackage{scalefnt}


%% supported options:
%%      draft - place DRAFT in the left hand side and give to memoir
%%      serif - don't use sans serif fonts
%%      headbold - display the head also in bold
%%      bold - use also bold in description, etc

%% Class Options Defaults
\newcommand{\headplaceholder}{}
\newcommand{\myheadplaceholder}{}
\newcommand{\myfamily}{\sffamily}
\newcommand{\myseries}{\bfseries}
\newcommand{\myheadseries}{}
\newcommand{\mytitlepage}{}
\newcounter{bottomline}
\setcounter{bottomline}{0}

\DeclareOption{draft}{
        \renewcommand{\myheadplaceholder}
        {\myfamily\bfseries\headplaceholder{}\normalfont\scshape{} draft}
        \PassOptionsToClass{draft}{memoir}
}
\DeclareOption{serif}{
        \renewcommand{\myfamily}{\rmfamily}
}
\DeclareOption{headbold}{
        \renewcommand{\myheadseries}{\bfseries}
}
\DeclareOption{bold}{
        \renewcommand{\myheadseries}{\myseries}
}
\DeclareOption{bottomline}{
        \setcounter{bottomline}{1}
}

\DeclareOption{titlepagenr}{
	\renewcommand{\mytitlepage}{
		\addtolength{\fboxsep}{0.5ex}
			\framebox{\myfamily\thepage}
	}
}

%% Pass all unknown options to memoir
\DeclareOption*{\PassOptionsToClass{\CurrentOption}{memoir}}
\ProcessOptions
\LoadClass{memoir}

\setcounter{secnumdepth}{2}     
%%
\makepagestyle{blocks}
\aliaspagestyle{chapter}{empty}
\aliaspagestyle{part}{empty}
\setlength{\headwidth}{\textwidth}
  \addtolength{\headwidth}{\marginparsep}
  \addtolength{\headwidth}{\marginparwidth}

\makepsmarks{blocks}{
  \let\@mkboth\markboth
  \def\chaptermark##1{\markboth{\myheadseries##1}{\myfamily##1}}  %% left & right marks
  \def\sectionmark##1{\markright{                             %% right mark
    \ifnum \c@secnumdepth>\z@
      \normalfont\myfamily\myheadseries\thesection{ }
    \fi
    \normalfont\myheadseries\myfamily##1}}
}

\makerunningwidth{blocks}{\headwidth}
\makeevenhead{blocks}{
        \addtolength{\fboxsep}{0.5ex}
        \framebox{\myfamily\thepage}
        % hack to get spacing correct
        \vspace{-0.1mm}
        \hrule
}{
         \raisebox{0.2ex}{\parbox{\headwidth}{ 
         \hspace{1em}
         \ifthenelse{\value{page} > 9}%
                {\hspace{1em}}{}
         \ifthenelse{\value{page} > 99}%
                {\hspace{1em}}{}
         \ifthenelse{\value{page} > 999}%
                {\hspace{1em}}{}
         { }\normalfont\leftmark %
         \vspace{1.1mm}
         }}
}{\myheadplaceholder}

\makeoddhead{blocks}{\myheadplaceholder}
{
         \raisebox{0.2ex}{\parbox{\headwidth}{ 
         \raggedleft
         \normalfont\rightmark %
         \hspace{1.2em}
         \ifthenelse{\value{page} > 9}%
                {\hspace{1em}}{}
         \ifthenelse{\value{page} > 99}%
                {\hspace{1em}}{}
         \ifthenelse{\value{page} > 999}%
                {\hspace{1em}}{}
         \vspace{1.1mm}
         } }
}
{
        \addtolength{\fboxsep}{0.5ex}
        \framebox{\myfamily{\thepage}} 
        %% hack to get spacing correct
        \vspace{-0.14mm}
        \hspace{-1.2mm}
	\hrule
%%	\rule{\linewidth}{0.1mm}
}
\ifthenelse{\value{bottomline} = 1}%
{\makefootrule{blocks}{\headwidth}{\normalrulethickness}{\footruleskip}}%
{}
%% centreert de header wat meer als ik dit uit zet, anders wordt ie
%% echt breed
%%\makeheadposition{blocks}{flushright}{flushleft}{flushright}{flushleft}
\addtolength{\headsep}{\headsep}

%% chapter style
\makechapterstyle{blocks}{%
  \renewcommand{\chapterheadstart}{\vspace{2cm}}
  \renewcommand{\printchaptername}{}
  \renewcommand{\chapternamenum}{}
  \renewcommand{\printchapternum}{%
        \addtolength{\fboxsep}{2ex} %% \numtoName
        \framebox{\myseries\myfamily\scalefont{6.0}{\thechapter}}
}
  \renewcommand{\afterchapternum}{\hskip 5mm}
  \renewcommand{\printchaptertitle}[1]{%
    \normalfont\raggedright\HUGE{{\myfamily##1}}%
    \chaptermark{\myfamily##1}%
    }
  \renewcommand{\afterchaptertitle}{%
    \vskip\onelineskip \vskip\onelineskip}
}

%% part style
\makepagestyle{part}
\renewcommand{\partname}{}
\renewcommand{\partnumfont}{\HUGE\myseries\myfamily}
\renewcommand{\parttitlefont}{\HUGE\myseries\myfamily}
\renewcommand{\printparttitle}[1]{\vspace{10ex}\parttitlefont #1}
\renewcommand{\printpartnum}{%
        \raggedright
        \addtolength{\fboxsep}{5ex}
        \framebox{\partnumfont\thepart}
}

%% bibliography style
\renewcommand{\bibsection}{%
  \chapter{\bibname}
  \bibmark
  \ifnobibintoc\else
    \phantomsection
    \addcontentsline{toc}{chapter}{\bibname}
  \fi
\prebibhook}

%% index style
\makeatletter
\renewenvironment{theindex}{%
  \if@twocolumn
    \@restonecolfalse
  \else
    \@restonecoltrue
  \fi
  \ifonecolindex
    \onecolumn
    \chapter*{\indexname}
    \preindexhook
  \else
    \setlength{\columnseprule}{\indexrule}%
    \setlength{\columnsep}{\indexcolsep}%
%%    \twocolumn[\@makeschapterhead{\indexname} %% disabled
%%              \preindexhook]%
  \fi
  \indexmark
  \ifnoindexintoc\else
    \phantomsection
    \addcontentsline{toc}{chapter}{\indexname}%
  \fi
  \thispagestyle{chapter}\parindent\z@
  \parskip\z@ \@plus .3\p@\relax
  \let\item\@idxitem}%
  {\if@restonecol\onecolumn\else\twocolumn\fi}
\makeatother

%% title page style - always odd... XXX
\makepagestyle{tblocks}
\makeoddfoot{tblocks}{}{\mytitlepage}{}

%% section style
\setsecheadstyle{\raggedright\myfamily\Large}
  \setbeforesecskip{-\onelineskip}
  \setaftersecskip{\onelineskip}

%% subsection
\setsubsecheadstyle{\sethangfrom{\noindent ##1}\raggedright\myfamily\slshape}
%%\setsubsecheadstyle{\raggedleft\myfamily\thesubsection\raggedright\myfamily\hskip 1em}
  \setbeforesubsecskip{-\onelineskip}
    \setaftersubsecskip{\onelineskip}

%% subsubsection
\setsubsubsecheadstyle{\sethangfrom{\noindent ##1}\raggedright\myfamily\slshape}
  \setbeforesubsubsecskip{-\onelineskip}
    \setaftersubsubsecskip{\onelineskip}

%% paragraph
\setparaheadstyle{\sethangfrom{\noindent ##1}\raggedright\myfamily\slshape}

%% description
\renewcommand*{\descriptionlabel}[1]{\hspace\labelsep
       \normalfont\rmfamily\textit{#1}\hskip\textwidth}

%% tableofcontents
\renewcommand{\cftpartfont}{\myfamily\myseries}
\renewcommand{\cftchapterfont}{\myfamily\myseries}

%% set the options for the document
\captiontitlefont{\small\myfamily}  %% small? TODO
\captionnamefont{\myfamily}
\pagestyle{blocks}
\chapterstyle{blocks}
\aliaspagestyle{title}{tblocks}
