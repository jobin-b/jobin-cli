# stock command

## Overview

The stock command allows for interaction with the stock market through macros.

## Setup

1. This command relies on UI vision's chrome browser extension. Install this to continue. Install Xmodules from UI vision to allow for runtime macro changes by jobin through file editing.

2. Create one autostart html file (Under Settings > API in UI vision).

3. Create a new shortcut with the Target set to the html file. Create one for each macro (Macros must be on physical storage).

   Target:

   `"C:\Program Files\Google\Chrome\Application\chrome.exe" "file:///C:\[API_FILE_PATH]\[API_FILE].html?direct=1&macro=[MACRO_NAME]"`

4. Add the paths to your macros and links (shortcuts) folder in the .jobin.json config file.

5. Update the config by running the install script.

   You are ready to go :)

## buy Subcommand (WIP)

`stock buy [broker] [...tickers]`

Use shorthand for brokers

Available brokers:

| Broker      | shorthand |
| ----------- | --------- |
| Wells Fargo | wf        |
