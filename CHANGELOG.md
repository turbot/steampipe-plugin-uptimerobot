## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#27](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/27))
- Recompiled plugin with Go version `1.21`. ([#27](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/27))

## v0.2.0 [2023-04-11]

_Dependencies_

Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#18](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/18))

## v0.1.0 [2023-01-19]

_Enhancements_

- Added columns `alert_contacts` and `logs` to `uptimerobot_monitor` table. ([#15](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/15))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.9](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v419-2022-11-30) which fixes hydrate function caching for aggregator connections. ([#16](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/16))

## v0.0.3 [2022-11-11]

_Bug fixes_

- Fixed `ssl` column not returning data in `uptimerobot_monitor` table.
- Removed unused columns `alert_contacts` and `logs` in `uptimerobot_monitor` table.

## v0.0.2 [2022-11-11]

_Bug fixes_

- Fixed paging in `uptimerobot_alert_contact`, `uptimerobot_maintenance_window` and `uptimerobot_monitor` tables. ([#13](https://github.com/turbot/steampipe-plugin-uptimerobot/pull/13))
- Fixed `create_date_time` column returning incorrect date in `uptimerobot_monitor` table.

## v0.0.1 [2022-09-16]

_What's new?_

- New tables added
  - [uptimerobot_account](https://hub.steampipe.io/plugins/turbot/uptimerobot/tables/uptimerobot_account)
  - [uptimerobot_alert_contact](https://hub.steampipe.io/plugins/turbot/uptimerobot/tables/uptimerobot_alert_contact)
  - [uptimerobot_maintenance_window](https://hub.steampipe.io/plugins/turbot/uptimerobot/tables/uptimerobot_maintenance_window)
  - [uptimerobot_monitor](https://hub.steampipe.io/plugins/turbot/uptimerobot/tables/uptimerobot_monitor)
