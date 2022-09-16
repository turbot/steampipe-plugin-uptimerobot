# Table: uptimerobot_monitor

UptimeRobot Monitors checks whether our website is down or not. It checks repeatedly after configured interval time so we can get an instant notification in case things go wrong.

These are the supported types:
- 1 - HTTP(s)
- 2 - Keyword
- 3 - Ping
- 4 - Port
- 5 - Heartbeat

These are the supported statuses:
- 0 - paused
- 1 - not checked yet
- 2 - up
- 8 - seems down
- 9 - down

## Examples

### Basic Info

```sql
select
  id,
  friendly_name,
  url,
  type
from
  uptimerobot_monitor;
```

### List paused monitors

```sql
select
  id,
  friendly_name,
  url
from
  uptimerobot_monitor
where
  status = 0;
```

### List heartbeat monitors

```sql
select
  id,
  type,
  url
from
  uptimerobot_monitor
where
  type = 5;
```
