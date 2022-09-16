# Table: uptimerobot_maintenance_window

Provides details of maintenance window in the UptimeRobot.

These are the supported types:
- 1 - Once
- 2 - Daily
- 3 - Weekly
- 4 - Monthly

These are the supported statuses:
- 0 - paused
- 1 - active

## Examples

### Basic info

```sql
select
  id,
  user,
  friendly_name,
  start_time
from
  uptimerobot_maintenance_window;
```

### List paused maintenance windows

```sql
select
  id,
  friendly_name,
  status
from
  uptimerobot_maintenance_window
where
  status = 0;
```

### List maintenance windows with duration longer than than 60 minutes

```sql
select
  id,
  friendly_name,
  duration
from
  uptimerobot_maintenance_window
where
  duration > 60;
```

### List maintenance windows that are checked on daily basis

```sql
select
  id,
  friendly_name,
  duration
from
  uptimerobot_maintenance_window
where
  type = 2;
```
