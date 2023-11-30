---
title: "Steampipe Table: uptimerobot_maintenance_window - Query UptimeRobot Maintenance Windows using SQL"
description: "Allows users to query Maintenance Windows in UptimeRobot, providing critical insights into scheduled maintenance periods for monitored websites or services."
---

# Table: uptimerobot_maintenance_window - Query UptimeRobot Maintenance Windows using SQL

UptimeRobot is a website monitoring service that checks your website's availability at regular intervals. Maintenance Windows in UptimeRobot represent scheduled periods of downtime for maintenance on monitored websites or services. These windows allow for planned interruptions in service without impacting uptime statistics.

## Table Usage Guide

The `uptimerobot_maintenance_window` table provides insights into Maintenance Windows within UptimeRobot. As a site reliability engineer or DevOps professional, you can explore details about scheduled maintenance periods through this table, including start and end times, duration, and affected monitors. Utilize it to plan and track maintenance activities, ensuring minimal disruption to your services and preserving the accuracy of your uptime statistics.

## Examples

### Basic info
Explore which maintenance windows are currently active on your UptimeRobot account. This is useful for understanding when your website or service may be temporarily unavailable due to scheduled maintenance.

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
Discover the segments that have paused maintenance windows, allowing you to assess the elements within your system that may need attention or rescheduling. This is useful for understanding the current operational status and planning future maintenance activities.

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
Discover the segments that have maintenance windows extending beyond an hour. This information can be useful for identifying potential periods of downtime or service interruptions.

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
Discover the segments that undergo daily maintenance checks, helping you to understand the frequency and duration of these routine operations. This can assist in scheduling tasks and avoiding potential disruptions.

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