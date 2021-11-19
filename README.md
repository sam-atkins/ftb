[![build](https://github.com/sam-atkins/ftb/actions/workflows/ci.yml/badge.svg)](https://github.com/sam-atkins/ftb/actions/workflows/ci.yml)

# ftb - a football CLI

## Install

First, get an API token from [www.football-data.org](https://www.football-data.org/).

Create a config file and add your API token.

```bash
mkdir -p $HOME/.config/ftb && echo "TOKEN: REPLACE_WITH_YOUR_API_TOKEN" > $HOME/.config/ftb/config.yaml
```

Run the install command.

```bash
make install
```

Confirm the binary is installed and check the help menu for available commands.

```bash
ftb --help
```
