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

  # API key for your UptimeRobot account.
  # Reference: https://uptimerobot.com/api/
  # Can also be set with the UPTIMEROBOT_API_KEY environment variable.
  # api_key = "u1857235-*********592bd3c445"
}
```

- `api_key` - UptimeRobot API key. There are 3 types of api_keys for reaching the data:
  - account-specific api_key which allows using all the API methods on all the monitors of an account.
  - monitor-specific api_keys which allows using only the getMonitors method for the given monitor.
  - read-only api_key which allows fetching data with all the get\* API endpoints.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-uptimerobot
- Community: [Slack Channel](https://steampipe.io/community/join)
