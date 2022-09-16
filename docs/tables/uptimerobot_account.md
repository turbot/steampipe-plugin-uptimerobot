# Table: uptimerobot_account

UptimeRobot acounts can be used to monitor websites, keywords, pings, ports, and more.

## Examples

### Basic info

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval,
  up_monitors
from
  uptimerobot_account;
```

### Get monitor limits of the account

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval
from
  uptimerobot_account;
```

### List up, down, and paused monitors in the account

```sql
select
  email,
  user_id,
  up_monitors,
  down_monitors,
  paused_monitors
from
  uptimerobot_account;
```
