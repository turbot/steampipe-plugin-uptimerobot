---
title: "Steampipe Table: uptimerobot_monitor - Query UptimeRobot Monitors using SQL"
description: "Allows users to query Monitors in UptimeRobot, providing a comprehensive view of the status, type, and response times of the monitors. It also provides information on the alert contacts associated with each monitor."
---

# Table: uptimerobot_monitor - Query UptimeRobot Monitors using SQL

UptimeRobot is a service that provides website monitoring, alerting, and status pages. It allows users to monitor HTTP(s), ping, port and check keywords. This service helps to minimize the downtime by providing timely alerts.

## Table Usage Guide

The `uptimerobot_monitor` table provides insights into the monitors within UptimeRobot. As a Site Reliability Engineer, you can explore monitor-specific details through this table, including status, type, response times, and associated alert contacts. Utilize it to uncover information about monitors, such as their uptime percentage, the type of monitor (HTTP, keyword, ping, port), and the alert contacts associated with each monitor.

## Examples

### Basic Info
Discover the segments that are being monitored by UptimeRobot by identifying their unique identifiers, friendly names, URLs, and types. This is beneficial for gaining insights into the variety and scope of your monitoring efforts.

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
Explore which monitors are currently paused, allowing you to identify potential issues or areas for review in your uptime monitoring system. This could be useful in maintaining optimal performance and minimizing downtime.

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
Discover the segments that are being monitored by the heartbeat monitoring system, which is crucial for ensuring the smooth functioning of your network by identifying any irregularities or disruptions in the network's heartbeat signals.

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
Identify the contact details linked to a specific monitor alert to understand who will be notified in case of any issues or disruptions. This can be useful in ensuring the right individuals or teams are kept informed and can respond promptly to any potential problems.

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
Explore the log details of a specific monitor to understand its performance history and identify potential issues. This query is beneficial for troubleshooting and maintaining optimal system performance.

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