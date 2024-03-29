---
organization: Turbot
category: ["website monitoring"]
icon_url: "/images/plugins/turbot/uptimerobot.svg"
brand_color: "#3BD771"
display_name: "UptimeRobot"
short_name: "uptimerobot"
description: "Steampipe plugin to query monitors, alert contacts and more from UptimeRobot."
og_description: "Query UptimeRobot with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/uptimerobot-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# UptimeRobot + Steampipe

[UptimeRobot](https://uptimerobot.com/) is a free tool used to monitor websites.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

Get UptimeRobot account details:

```sql
select
  email,
  monitor_limit,
  monitor_interval,
  up_monitors
from
  uptimerobot_account;
```

```
+---------------------+---------------+------------------+-------------+
| email               | monitor_limit | monitor_interval | up_monitors |
+---------------------+---------------+------------------+-------------+
| niharika@turbot.com | 50            | 5                | 3           |
+---------------------+---------------+------------------+-------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/uptimerobot/tables)**

## Get started

### Install

Download and install the latest UptimeRobot plugin:

```bash
steampipe plugin install uptimerobot
```

### Configuration

Installing the latest uptimerobot plugin will create a config file (`~/.steampipe/config/uptimerobot.spc`) with a single connection named `uptimerobot`:

```hcl
connection "uptimerobot" {
  plugin       = "uptimerobot"

  # API key for your UptimeRobot account.
  # We recommend creating a read-only api_key, for more information on the
  # different types, please see https://uptimerobot.com/api/.
  # Can also be set with the UPTIMEROBOT_API_KEY environment variable.
  # api_key = "u1857235-592bd3c445thisisafakekey"
}
```

- `api_key` - UptimeRobot API key. Can also be set with the `UPTIMEROBOT_API_KEY` environment variable.


