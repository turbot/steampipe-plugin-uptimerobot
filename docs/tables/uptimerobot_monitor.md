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

### Get alert contact details of a particular monitor

```sql
select
  id,
  friendly_name,
  jsonb_pretty(alert_contacts) as alert_contacts
from
  uptimerobot_monitor
where
  id = '793508639';
```

### Get log details of a particular monitor

```sql
select
  id,
  friendly_name,
  jsonb_pretty(logs) as logs
from
  uptimerobot_monitor
where
  id = '793508639';
```
