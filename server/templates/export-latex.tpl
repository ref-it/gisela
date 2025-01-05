\documentclass[
    11pt,
	a4paper,
	landscape,
	titlepage
]{scrartcl}

\usepackage[
	includehead,
    includefoot,
	top=2.5cm,
	left=1cm,
	right=1cm,
	bottom=1cm
]{geometry}

\usepackage[T1]{fontenc}
\usepackage[ngerman]{babel}
\usepackage{lmodern}
\usepackage{paratype}
\renewcommand*\familydefault{\sfdefault}
\usepackage{microtype}
\usepackage{tabularx}
\usepackage{xltabular}
\usepackage{multirow}
\usepackage[table]{xcolor}
\usepackage{graphicx}
\usepackage[iso,ngerman]{isodate}
\usepackage{lastpage}
\usepackage[pdfborder={0 0 0}]{hyperref}
\setkomafont{pagenumber}{\normalfont}

\usepackage[
	headsepline,
	footsepline,
	automark
]{scrlayer-scrpage}
\ihead*{\normalfont Inventarverzeichnis}
\chead*{}
\ohead*{\normalfont Studierendenrat Ilmenau}
\ifoot*{\normalfont\footnotesize\today}
\cfoot*{}
\ofoot*{\normalfont\footnotesize\pagemark~/~\pageref*{LastPage}}

\begin{document}
	
	\begin{titlepage}
		
		~\vspace{-2.6cm}
		\begin{center}
			\thispagestyle{empty}
			~\\[.11cm]
			
			\begin{tabularx}{\textwidth}{@{}X r@{}}
				\huge\textbf{Studierendenrat} & \multirow{2}{5cm}{\hfill\includegraphics[width=2.2cm]{../../img/logo-briefkopf.pdf}}\\[.2cm]
				{\Large der Technischen Universität Ilmenau}
			\end{tabularx}
			\\\vspace{.2cm}\rule{\textwidth}{.25pt}
			
			\vfill
			\textbf{\Huge\sffamily Inventarverzeichnis\\}
			\vspace{1cm}
			{\Large\selectfont\today}
			
		\end{center}
		\vfill
        \def\arraystretch{1.4}
		\rowcolors{2}{gray!25}{white}
		\begin{tabularx}{\textwidth}{|l|c|X|}
			\hline
			{{ if .ExportRequest.FilterDateOfReceipt }}\textbf{Zugangsdatum} & {{ if eq .ExportRequest.RelationDateOfReceipt "≤" }}$\leq${{ else if eq .ExportRequest.RelationDateOfReceipt "≥" }}$\geq${{ else }}{{ .ExportRequest.RelationDateOfReceipt }}{{ end }} & {{ .ExportRequest.DateOfReceipt }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterDateOfRetirement }}\textbf{Abgangsdatum} & {{ if eq .ExportRequest.RelationDateOfRetirement "≤" }}$\leq${{ else if eq .ExportRequest.RelationDateOfRetirement "≥" }}$\geq${{ else }}{{ .ExportRequest.RelationDateOfRetirement }}{{ end }} & {{ .ExportRequest.DateOfRetirement }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterAcquisitionCost }}\textbf{Anschaffungskosten} & {{ if eq .ExportRequest.RelationAcquisitionCost "≤" }}$\leq${{ else if eq .ExportRequest.RelationAcquisitionCost "≥" }}$\geq${{ else }}{{ .ExportRequest.RelationAcquisitionCost }}{{ end }} & {{ .ExportRequest.AcquisitionCost }} €\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterCategories }}\textbf{Kategorien} & & {{ range .ExportRequest.Categories }}{{ index $.Maps.Categories . }}{{ end }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterGroups }}\textbf{Gremien} & & {{ range .ExportRequest.Groups }}{{ index $.Maps.Groups . }}{{ end }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterPlaces }}\textbf{Lagerorte} & & {{ range .ExportRequest.Places }}{{ index $.Maps.Places . }}{{ end }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterRetired }}\textbf{Ausgemustert} & & {{ if .ExportRequest.Retired }}ja{{ else }}nein{{ end }}\\
			\hline{{ end }}
			{{ if .ExportRequest.FilterBorrowable }}\textbf{Ausleihbar} & & {{ if .ExportRequest.Borrowable }}ja{{ else }}nein{{ end }}\\
			\hline{{ end }}
		\end{tabularx}
		
	\end{titlepage}
	
	\setcounter{page}{1}
	
    \def\arraystretch{1.4}
	\rowcolors{2}{gray!25}{white}
    \tiny
	\begin{xltabular}{\textwidth}{|r|>{\raggedright}p{1.5cm}|>{\raggedright}X|>{\raggedright}p{1.5cm}|>{\raggedright}p{1cm}|>{\raggedright}p{1cm}|>{\raggedright}p{1cm}|l|l|l|l|>{\raggedright}p{1.5cm}|r|>{\raggedright}p{0.75cm}|>{\raggedright}p{0.75cm}|l|l|}
		\hline\rowcolor{gray!70} \textbf{ID} & \textbf{Titel} & \textbf{Beschreibung} & \textbf{Zustand} & \textbf{Gremium} & \textbf{Kategorie} & \textbf{Lagerort} & \textbf{Zugang} & \textbf{Abgang} & \textbf{Anz.} & \textbf{Teil von} & \textbf{Lieferant} & \textbf{Ansch.-K.} & \textbf{Projekt} & \textbf{Abrech.} & \textbf{Buchung} & \textbf{Ausl.}\\\hline
        \endhead
        {{ range .InventoryItems }}
            {{ .ItemID }} &
            {{ .Title }} &
            {{ .Description }} &
            {{ .Condition }} &
            {{ index $.Maps.Groups .Group }} &
            {{ index $.Maps.Categories .Category }} &
            {{ index $.Maps.Places .Place }} &
            {{ .DateOfReceipt }} &
            {{ .DateOfRetirement }} &
            {{ .Number }} &
            {{ if ne .ParentItem 0 }}{{ .ParentItem }}{{ end }} &
            {{ .Supplier }} &
            {{ .AcquisitionCost }} € &
            {{ range $index, $projectno := .Project }}{{ if $index }}, {{ end }}\href{https://finanzen.stura-ilmenau.de/projekt/{{ $projectno }} }{ {{ $projectno }} }{{ end }} &
            {{ range $index, $receiptno := .Receipt }}{{ if $index }}, {{ end }}\href{https://finanzen.stura-ilmenau.de/auslagen/{{ $receiptno }} }{ {{ $receiptno }} }{{ end }} &
            {{ .DateOfAccounting }} &
            {{ if .Borrowable }}ja{{ else }}nein{{ end }}\\
            \hline
        {{ end }}
    \end{xltabular}

\end{document}