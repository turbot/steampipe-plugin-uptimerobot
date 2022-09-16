# Table: uptimerobot_alert_contact

UptimeRobot Alert Contact extracts a dataset with general information for a set of contacts used to be alert in case of up/down of the given monitors for a particular user.

## Examples

### Basic info

```sql
select
  id,
  friendly_name,
  type,
  status,
  value
from
  uptimerobot_alert_contact;
```

### List Telegram alert contacts

```sql
select
  id,
  friendly_name,
  type,
  status,
  value
from
  uptimerobot_alert_contact
where
  type = 18;
```

### List inactive alert contacts

```sql
select
  id,
  friendly_name,
  type,
  status,
  value
from
  uptimerobot_alert_contact
where
  status = 0;
```
