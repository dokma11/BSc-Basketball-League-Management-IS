<h1 align="center"> Bsc-Basketball-League-Management-IS </h1>

<div align="center">
    <img src="https://img.shields.io/badge/Oracle-F80000?style=for-the-badge&logo=oracle&logoColor=black">
    &nbsp;<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
    &nbsp;<img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white">
    &nbsp;<img src="https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white">
</div>

## Overview:
This repository contains all the materials from my bachelor's thesis. It includes the project specification, 
the database schema, both logical and relational schemas created using Oracle's SQL Developer Data Modeler, 
as well as a complete application. The application features a server-side implemented in Golang and a client-side developed using the Angular framework.

## Content:
- Database schema: The database schema is provided in the `Database_schema.png` file. Additionally, within the `datamodeler` 
directory, you'll find both the logical and relational schema implementations, along with the corresponding DDL file.

- Server side: Server side of the application was implemented using the Go programming language. The implementation is 
provided in the `server` directory.

- Client side: Client side of the application is implemented using the Angular framework. The implementation is provided
in the `client` directory.

- Scraping: To ensure that my database contains accurate and realistic team data, I wrote a web scraping script in Python 
using the `requests` and `bs4.BeautifulSoup` libraries. This script is available in the `bball-league-data-scraper` 
directory. The data was scraped from the website [basketball-reference.com](https://www.basketball-reference.com/).

- System requirements specification: All functional and non-functional requirements are detailed in the
`SRS_za_informacioni_sistem_za_rukovodjenje_clanovima_kosarkaske_lige` file. This document also includes sequence and 
use case diagrams, which can additionally be found in the `Sequence` and `UseCases` directories.

- PL/SQL triggers, sequences and indexes script: All the triggers, sequences and indexes implemented in the project are 
provided in the `DB_Triggers_Sequences_Indexes.sql` file. 

- Insert script: All the insert queries are provided in the `DB_Data_Insert_Script.sql` file.

- Thesis: Provided in the `FTN_BSc_DokmanovicVukasin_Rad_Final.pdf` and `FTN_BSc_DokmanovicVukasin_Rad_Final.docx` files.

- Presentation: Presentation used in the thesis defense is provided in the `FTN_BSc_DokmanovicVukasin.pptx` file.

## Technologies:
- Oracle: Relational database system used for data storage and retrieval.
- Go: Programming language used on the server side.
- Angular: A platform and framework for building single-page client applications using HTML, CSS, and TypeScript.
- TypeScript: A typed superset of JavaScript that compiles to plain JavaScript, adding static types to enhance code quality and development efficiency.
- HTML: The standard markup language used for creating and structuring content on the web.
- CSS: A style sheet language used for describing the presentation and design of web pages, including layout, colors, and fonts
