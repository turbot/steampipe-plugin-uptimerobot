# Table: uptimerobot_maintenance_window

Give details of maintenance window in the uptimerobot.

## Reference

https://uptimerobot.com/api/

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

### List the paused maintenance windows
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

### List the maintenance windows that have duartion more than 60 minutes 

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

### List the maintenance windows that are checked on daily basis

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
