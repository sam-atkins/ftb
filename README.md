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
task install
```

Confirm the binary is installed and check the help menu for available commands.

```bash
ftb --help
```

## Available Competitions

These are the competitions available from [www.football-data.org](https://www.football-data.org/) free tier.

| Code | Competition Name |
| ---- | ---------------- |
| WC | FIFA World Cup |
| CL | UEFA Champions League |
| BL1 | Bundesliga |
| DED | Eredivisie |
| BSA | Campeonato Brasileiro Série A |
| PD | Primera Division |
| FL1 | Ligue 1 |
| ELC | Championship |
| PPL | Primeira Liga |
| EC | European Championship |
| SA | Serie A |
| PL | Premier League |
| CLI | Copa Libertadores |
