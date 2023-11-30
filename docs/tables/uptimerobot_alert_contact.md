---
title: "Steampipe Table: uptimerobot_alert_contact - Query UptimeRobot Alert Contacts using SQL"
description: "Allows users to query Alert Contacts in UptimeRobot, specifically the details of each alert contact, providing insights into the alert contact configurations and settings."
---

# Table: uptimerobot_alert_contact - Query UptimeRobot Alert Contacts using SQL

UptimeRobot is a service that provides website monitoring, detecting downtime, and sending alerts. It allows you to monitor your websites, servers, and APIs in real-time and get alerted when they go down. Alert Contacts in UptimeRobot are the recipients of these alerts, which can be configured according to different parameters.

## Table Usage Guide

The `uptimerobot_alert_contact` table provides insights into Alert Contacts within UptimeRobot. As a Site Reliability Engineer, explore alert contact-specific details through this table, including contact type, status, and associated metadata. Utilize it to understand the configuration of alert contacts, such as their value, type, and the thresholds for triggering alerts.

## Examples

### Basic info
Explore which alert contacts are active in UptimeRobot to manage your notifications effectively. This aids in understanding the types and statuses of your alert contacts, ensuring you're always informed about your site's uptime.

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
Discover the segments that consist of Telegram alert contacts within your system, allowing you to assess the status and details of each contact for effective communication management. This query is particularly useful for maintaining and managing alert contacts within a Telegram-based notification system.

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
Identify inactive alert contacts to assess the elements within your system that may not be receiving critical notifications. This is useful for maintaining efficient communication channels and ensuring prompt responses to system alerts.

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