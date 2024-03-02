# stock command

## Overview

The stock command allows for interaction with the stock market through macros.

## Setup

1. This command relies on UI vision's chrome browser extension. Install this to continue. Install Xmodules from UI vision to allow for runtime macro changes by jobin through file editing.

2. Create one autostart html file (Under Settings > API in UI vision). (api.html)

3. Add the paths to your html file and macros folder in the .jobin.json config file. Modify any configs as desired

4. Update the config by running the install script.

   You are ready to go :)

## buy Subcommand (WIP)

`stock buy [broker] [...tickers]`

Use shorthand for brokers

Available brokers:

| Broker      | shorthand |
| ----------- | --------- |
| Wells Fargo | wf        |
