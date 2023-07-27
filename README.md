![image](https://hub.steampipe.io/images/plugins/turbot/uptimerobot-social-graphic.png)

# UptimeRobot Plugin for Steampipe

Use SQL to query monitors, alert contacts and more from UptimeRobot.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/uptimerobot)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/uptimerobot/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-uptimerobot/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install uptimerobot
```

Configure the API key in `~/.steampipe/config/uptimerobot.spc`:

```hcl
connection "uptimerobot" {
  plugin  = "uptimerobot"
  api_key = "u1857235-592bd3c445thisisafakekey"
}
```

Run steampipe:

```shell
steampipe query
```

Query your account:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-uptimerobot.git
cd steampipe-plugin-uptimerobot
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/uptimerobot.spc
```

Try it!

```
steampipe query
> .inspect uptimerobot
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-uptimerobot/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [UptimeRobot Plugin](https://github.com/turbot/steampipe-plugin-uptimerobot/labels/help%20wanted)
