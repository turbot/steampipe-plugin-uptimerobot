# Table: uptimerobot_alert_contact

Uptime Robot Alert Contact extracts a dataset with general informations for a set of contacts used to be alert in case of up/down of the given monitors for a particular user.

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

### Get alert contact with type telegram

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
