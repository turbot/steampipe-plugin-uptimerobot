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
