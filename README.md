# Butterfly

![Version](https://img.shields.io/badge/v1-OpenSource-3300AA.svg) ![License](https://img.shields.io/badge/license-MPL--2.0-FF6600.svg) ![Platform](https://img.shields.io/badge/base_on-StarStart!-11BAFF.svg) [![Code-Inspector_Score](https://www.code-inspector.com/project/5659/score/svg) ![Code-Inspector_Score](https://www.code-inspector.com/project/5659/status/svg)](https://frontend.code-inspector.com/public/project/5659/butterfly/dashboard)

> [[Violet](https://github.com/star-inc/violet)] + [[Butterfly](https://github.com/star-inc/butterfly)] = {[StarStart!](https://start.starinc.xyz)}

![logo](logo.svg)

The World-Wide-Web crawler for Apache Solr

## Installation

Get the execute file official build via [GitHub Releases](https://github.com/star-inc/butterfly/releases).

No installation required,
but you can put it on the directory your terminal $PATH variable points for convenient.

Eventually,
set up your config file (Command: `butterfly config`),
and move it to the directory where you want to execute `butterfly` command.

## Usage

### Add a website

    butterfly -add <URI>

The `<URI>` is a variable, please fill in the URL you hope to fetch the website at first time.

It will be add into a list that the butterfly will fetch it.

### Show the website list

    butterfly list

It will display the URL list of websites that the butterfly will visit at the first time.

> NOTICE: If the butterfly had been fetched the URL, it will ignore the URL next time.
>  
> But if you hope to re-fetch the website, please delete the database of this website,
> and restart the butterfly again.

### Delete a website from the list

    butterfly -add <URI>

The `<URI>` is the variable that the website URL you want to remove from the list.

### Execute

    butterfly start

The butterfly will fly to fetch the websites in the list.

## Developmet Environment

### Requirement

Butterfly requires [GoLang](https://golang.org/dl) >= 1.13

Please check your `go version` or install the latest version.

- GNU Linux / MacOSX / Unix Like

    Execute this command

    `sh bootstrap.sh`

- Win32 Platform

    We have no plan to support Microsoft Windows,
    but you still can get the go packages Butterfly required.

### Compilation

Just do it.

    go build

The "butterfly" will be the execute file you build on your own self.

> (c) 2020 Star Inc.
