\NeedsTeXFormat{LaTeX2e}[1995/12/01]
\ProvidesClass{blocksbook}[2006/06/10 v0.1
        Layout with blocks]
\RequirePackage{ifthen}
\RequirePackage{scalefnt}


\newenvironment{display}{\def\FrameCommand{\hskip\parindent}%%
\MakeFramed{\advance\hsize-\width\FrameRestore}%%
\vspace*{-2ex}\small\begin{alltt}}%
{\end{alltt}\vspace*{-2ex}\endMakeFramed}

\newenvironment{lbar}[1][]{%
\def\FrameCommand{\rightskip=\parindent\hskip\parindent\vrule width 1pt \hspace{10pt}}%
\MakeFramed{\rightskip=\parindent\advance\hsize-\width\FrameRestore\noindent\hskip-0.6ex}%
{%
\ifthenelse{\equal{#1}{}}{}{\hskip0.6ex\normalfont\myfamily\myheadseries{#1}\\\\}}%
}%
{\endMakeFramed}

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

\newcommand{\typename}{}
\newlength{\gnat}

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

\setsecnumdepth{chapter} %% nice :-)

\makepagestyle{blocks}
\aliaspagestyle{chapter}{empty}
\aliaspagestyle{part}{empty}
\setlength{\headwidth}{\textwidth}
  \addtolength{\headwidth}{\marginparsep}
  \addtolength{\headwidth}{\marginparwidth}

\makepsmarks{blocks}{
  \let\@mkboth\markboth
  \def\chaptermark##1{\markboth{\myfamily \typename{} \thechapter: ##1}{\myfamily##1}}  %% left & right marks
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
	 \settowidth{\gnat}{\thepage}
	 \hspace{\gnat}
         \hspace{1em}
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
	 \settowidth{\gnat}{\thepage}
	 \hspace{\gnat}
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
  \renewcommand{\chapterheadstart}{\vspace*{-5em}}
  \renewcommand{\printchaptername}{}
  \renewcommand{\typename}{\chaptername}
  \renewcommand{\chapternamenum}{}
%%  \renewcommand{\printchapternonum}{%
%%        \addtolength{\fboxsep}{2ex}
%%        \framebox{\myseries\myfamily\scalefont{6.0}{*}}
%%	\afterchapternum
%%}
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
% appendix style
\addtodef{\appendix}{}{\renewcommand{\typename}{\appendixname}}

%% part style
\makepagestyle{part}
\renewcommand{\partname}{}
\renewcommand{\partnumfont}{\myseries\myfamily}
\renewcommand{\parttitlefont}{\myseries\myfamily}
\renewcommand{\printparttitle}[1]{\vspace{10ex}\parttitlefont\scalefont{3.0}{#1}}
\renewcommand{\printpartnum}{%
        \raggedright
        \addtolength{\fboxsep}{5ex}
        \framebox{\partnumfont\scalefont{6.0}{\thepart}}
}

%% bibliography style
\renewcommand{\bibsection}{%
  \chapter{\bibname}
  \bibmark
  \ifnobibintoc\else
    \phantomsection
    \addcontentsline{toc}{chapter}{\bibname}
  \fi
\prebibhook
\markboth{\myfamily \typename{} \thechapter: \bibname}{\myfamily\bibname}
}

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
\setsecheadstyle{\raggedright\myfamily\LARGE}
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
