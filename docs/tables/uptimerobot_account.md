---
title: "Steampipe Table: uptimerobot_account - Query UptimeRobot Accounts using SQL"
description: "Allows users to query UptimeRobot Accounts, providing information about the account details and settings."
---

# Table: uptimerobot_account - Query UptimeRobot Accounts using SQL

UptimeRobot is a service that performs website checks to determine their uptime and functionality. It can monitor different types of internet services, including HTTP, HTTPS, SMTP, POP3, and IMAP. UptimeRobot can send alerts about downtime and other technical issues.

## Table Usage Guide

The `uptimerobot_account` table provides insights into the UptimeRobot account settings. As a Site Reliability Engineer (SRE), you can use this table to access account-specific details, including the type of monitoring checks, alert contacts, and monitor intervals. Utilize it to review and manage the monitoring settings and ensure the smooth operation of your internet services.

## Examples

### Basic info
Analyze the settings to understand the limits and intervals of your uptime monitors, as well as identify the users associated with each monitor. This is beneficial for managing resources and ensuring optimal monitoring performance.

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval,
  up_monitors
from
  uptimerobot_account;
```

### Get monitor limits of the account
Explore the limitations and intervals for monitoring within your account. This can help you understand and manage your account's monitoring capacity and frequency.

```sql
select
  email,
  user_id,
  monitor_limit,
  monitor_interval
from
  uptimerobot_account;
```

### List up, down, and paused monitors in the account
Determine the status of your monitors across your account by identifying which ones are active, inactive, or paused. This can help you manage your resources effectively by focusing your attention on monitors that require immediate action.

```sql
select
  email,
  user_id,
  up_monitors,
  down_monitors,
  paused_monitors
from
  uptimerobot_account;
```