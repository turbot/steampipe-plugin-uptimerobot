# Table: uptimerobot_account

Give details of the account in the uptimerobot.

## Examples

### Basic info

```sql
select
  email,
  monitor_limit,
  monitor_interval,
  up_monitors
from
  uptimerobot_account;
```

### Get monitor limits of your account

```sql
select
  email,
  monitor_limit
from
  uptimerobot_account;
```

### List accounts that have monitors down

```sql
select
  email,
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
  monitor_limit,
  monitor_interval,
  up_monitors
from
  uptimerobot_account
where
 paused_monitors > 0;
 ```
