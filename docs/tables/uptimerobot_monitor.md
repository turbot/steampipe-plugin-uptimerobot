# Table: uptimerobot_monitor

UptimeRobot Monitors checks whether our website is down or not. It checks repeatedly after configured interval time so we can get an instant notification in case things go wrong.

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
