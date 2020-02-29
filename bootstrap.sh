#!/bin/sh

cat <<EOF
   _____     _   _           ___ _     
  | __  |_ _| |_| |_ ___ ___|  _| |_ _ 
  | __ -| | |  _|  _| -_|  _|  _| | | |
  |_____|___|_| |_| |___|_| |_| |_|_  |
                                  |___|
                                    (c)2020 Star Inc.

[The software licensed under Mozilla Public License Version 2.0]
EOF

configfile=".butterfly_config.json"
configfile_sample="config.sample.json"

echo ""
echo "> Welcome to help us improve Butterfly"
echo ""
echo "https://github.com/star-inc/butterfly"

echo ""
echo "Installing Go Packages"
go get -u github.com/grokify/html-strip-tags-go
go get -u github.com/star-inc/butterfly-solr/solr
go get -u github.com/gocolly/colly/...
go get -u github.com/velebak/colly-sqlite3-storage/colly/sqlite3
go get -u github.com/PuerkitoBio/goquery
go get -u github.com/temoto/robotstxt


if [ ! -f $configfile ]; then
    echo "Coping \`$configfile_sample\` to \`$configfile\`"
    cat $configfile_sample > $configfile
fi

echo ""
echo "Let\`s Gopher (>w<)"
echo ""