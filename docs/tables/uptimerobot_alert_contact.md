# Table: uptimerobot_alert_contact

UptimeRobot Alert Contact extracts a dataset with general information for a set of contacts used to be alert in case of up/down of the given monitors for a particular user.

These are the supported types:
- 1 - sms
- 2 - e-mail
- 3 - twitter
- 5 - web-hook
- 6 - pushbullet
- 7 - zapier
- 8 - pro-sms
- 9 - pushover
- 11 - slack
- 14 - voice-call
- 15 - splunk
- 16 - pagerduty
- 17 - opsgenie
- 18 - telegram
- 20 - ms-teams
- 21 - google-chat
- 23 - discord

These are the supported statuses:
- 0 - not activated
- 1 - paused
- 2 - active

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
