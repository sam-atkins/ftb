# ftb - a football CLI

## Install

First, get an API token from [www.football-data.org](https://www.football-data.org/). Then run the commands below.

```bash
# Create a config file and add your API token
mkdir -p $HOME/.config/ftb && echo "TOKEN: REPLACE_WITH_YOUR_API_TOKEN" > $HOME/.config/ftb/config.yaml

# run the install command
make build

# confirm the binary is installed and check the help menu for available commands
ftb --help
```
