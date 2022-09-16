# Table:

UptimeRobot Monitors checks whether our website is down or not. It checks after configured interval time so we can get an instant notification in case things go wrong.

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

### List the monitors with a particular name

```sql
select
  id,
  friendly_name,
  url
from
  uptimerobot_monitor
where 
  friendly_name = 'test';
```

### List the monitors with status paused 

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

### List monitors with type Heartbeat

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
