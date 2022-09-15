---
organization: Turbot
category: ["website monitoring"]
icon_url: "/images/plugins/turbot/uptimerobot.svg"
brand_color: "  "
display_name: "UptimeRobot"
short_name: "uptimerobot"
description: "Steampipe plugin to query monitors, alert contacts and more from UptimeRobot."
og_description: "Query UptimeRobot with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/uptimerobot-social-graphic.png"
---

# UptimeRobot + Steampipe

[UptimeRobot](https://uptimerobot.com/) is a free tool used to monitor websites.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

  # API Key for your UptimeRobot account
  # Reference: https://uptimerobot.com/api/
  # Env variable: UPTIMEROBOT_API_KEY
  # api_key = "YOUR_UPTIMEROBOT_ACCESS_KEY"
}
```

- `api_key` - API key from UptimeRobot.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-uptimerobot
- Community: [Slack Channel](https://steampipe.io/community/join)
