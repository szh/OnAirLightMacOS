#!/usr/bin/env bash

./build.sh

cp -a "../../bin/macos/" "$HOME/onairlight/"
chmod a+x "$HOME/onairlight/client"

PLIST_TARGET="$HOME/Library/LaunchAgents/onairlight.loginscript.plist"
cp "launch.plist" "$PLIST_TARGET"
sed -i "" "s/_username_/$USER/" "$PLIST_TARGET"
launchctl unload "$PLIST_TARGET"
launchctl load "$PLIST_TARGET"
launchctl start szh/OnAirLightMacOs
