# Table: uptimerobot_account

Provide account details of the uptimerobot.

## Reference

https://uptimerobot.com/api/

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

### List accounts that have monitors down

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval,
  down_monitors
from
  uptimerobot_account
where
  down_monitors > 0;
```

### List accounts that have paused monitors

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval,
  paused_monitors
from
  uptimerobot_account
where
 paused_monitors > 0;
```
