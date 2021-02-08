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

if [ "$USER" = "root" ]; then
    config_root="/etc"
else
    config_root="$HOME/.config/butterfly"
fi

config_file="config.json"
config_file_sample="config.sample.json"

echo ""
echo "> Welcome to help us improving Butterfly"
echo ""
echo "https://github.com/star-inc/butterfly"

echo ""
echo "Installing Go Packages"
go mod download

if [ ! -d "$config_root" ]; then
    echo "Creating directory \`$config_root\`"
    mkdir -p "$config_root"
fi

if [ ! -f $config_file ]; then
    echo "Coping \`$config_file_sample\` to \`$config_root/$config_file\`"
    cat "$config_file_sample" > "$config_root/$config_file"
fi

echo ""
echo "Let\`s Gopher (>w<)"
echo ""