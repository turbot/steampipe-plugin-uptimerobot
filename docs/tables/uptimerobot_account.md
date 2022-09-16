# Table: uptimerobot_account

Provide account details of the UptimeRobot.

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

### Get monitor limits of your accounts

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval
from
  uptimerobot_account;
```

### List up/down/paused monitors

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
