A. Today / later today

me to drink water at 3pm

me to stretch at 15:00

me to take a break in 30 minutes

me to stand up and walk in 10 minutes

me to call Mom in 2 hours

me to check emails in 1 hour

me to review PRs later today

me to finish the report later today

me to buy groceries this evening

me to send the invoice tonight

me to follow up with John at 5pm

B. Tomorrow

me to review PRs tomorrow at 09:00

me to review PRs tomorrow

me to review PRs tomorrow morning

me to review PRs tomorrow afternoon

me to review PRs tomorrow evening

me to go for a run tomorrow at 6am

me to pay the electricity bill tomorrow

me to book a dentist appointment tomorrow afternoon

me to clean the kitchen tomorrow evening

me to prep for the meeting tomorrow at 2pm

C. Specific weekdays (no date)

me to review PRs on Monday at 10am

me to review PRs on Friday morning

me to check metrics on Friday afternoon

me to send the weekly update on Friday at 4pm

me to plan next week’s tasks on Sunday evening

me to do admin on Monday morning

me to back up my laptop on Saturday at 11am

me to water the plants on Wednesday

me to go for a long run on Saturday morning

D. Relative dates (“next week”, “next month”, etc.)

me to review PRs next week

me to start planning the release next week

me to check my goals next month

me to renew my subscription next month

me to start performance review notes next quarter

me to revisit this idea next year

E. Relative durations (“in X …”)

me to check on the deploy in 5 minutes

me to test the new feature in 20 minutes

me to stand up and stretch in 45 minutes

me to grab a coffee in 1 hour

me to re-run the script in 2 hours

me to re-check logs in 3 hours

me to follow up with support in 24 hours

F. Recurring reminders (common Slack-style)

me to drink water every hour

me to review PRs every weekday at 09:00

me to check Jira every morning at 10am

me to write a daily summary every weekday at 5pm

me to plan my week every Monday at 9am

me to do a weekly review every Sunday at 7pm

me to back up my code every Friday at 4pm

me to check my budget on the first day of every month

me to review goals on the last day of every month

G. Slightly vague but realistic user phrasing

me to look at this again later

me to think about this more tomorrow

me to revisit this idea next week

me to clean up my tasks sometime tomorrow

me to book flights this weekend

me to look into this after lunch

me to go over these notes after standup

me to check on the deploy after work

me to send feedback before end of day

me to finalise slides before the meeting

2. Behaviour rules for missing or vague times

This is the “spec” your program should follow. No code, just behaviour.

2.1 General principles

If the user gives an explicit time, always use that time (e.g. at 3pm, at 15:00).

If the user gives only a date/relative date (e.g. tomorrow, next week) with no time, choose a sane, consistent default time.

If the user uses fuzzy parts of the day (morning / afternoon / evening / tonight / later today), map those to fixed default times.

The original text (e.g. "tomorrow morning") should still be stored as-is for display, but internally you compute a concrete timestamp.

2.2 Default times for parts of the day

You suggested some; I’ll lock them in explicitly:

Morning ("morning", "tomorrow morning", "Monday morning"):
→ 09:00 local time

Afternoon ("afternoon", "tomorrow afternoon", "Friday afternoon"):
→ 12:00 local time

Evening ("evening", "tomorrow evening", "Friday evening"):
→ 19:00 local time (you can pick 7pm as a reasonable default)

Tonight ("tonight"):
→ 20:00 local time on the same day,
→ If it’s already past, treat “tonight” as tomorrow at 20:00.

After lunch:
→ 14:00 local time (2pm)

After work:
→ 18:00 local time (6pm)

You don't need to overcomplicate; just be consistent.

2.3 “Later today”

If phrase contains “later today” and no explicit time, set reminder for:

Now + 3 hours, using the user’s local time.

Optional safety rule: if now + 3h lands too late (e.g. after 21:00), you can clamp it to 21:00 or move it to tomorrow at 09:00 — up to you, but be consistent.

Examples:

me to review PRs later today → today, now + 3h.

me to finish the report later today → today, now + 3h.

2.4 “Tomorrow” without time

If user says “tomorrow” with no time:

Default to 09:00 tomorrow.

Examples:

me to review PRs tomorrow → tomorrow at 09:00.

me to clean the kitchen tomorrow → tomorrow at 09:00.

If they say “tomorrow morning/afternoon/evening”, use the part-of-day defaults above.

2.5 Specific weekday without time

For phrases like "on Monday", "on Friday", "next Monday", where no time is provided:

Base rule: default to 09:00 on that day.

Examples:

me to water the plants on Wednesday → next Wednesday at 09:00.

me to plan my week on Monday → next Monday at 09:00.

If they say "Monday morning" / "Friday afternoon" / "Sunday evening", use the part-of-day mapping instead of 09:00.

2.6 “Next week” and similar

For phrases like "next week" with no specific day or time:

“next week” → next Monday at 09:00 (start of next week)

“this weekend” → Saturday at 10:00 (first weekend day at a mid-morning time)

“next month” → 1st day of next month at 09:00

“next year” → January 1st next year at 09:00

Examples:

me to review PRs next week → next Monday at 09:00.

me to book flights this weekend → upcoming Saturday at 10:00.

me to check my goals next month → 1st of next month at 09:00.

2.7 “Before end of day” / “before EOD”

For phrases like "before end of day", "before EOD" with no time:

Use a common “end of day” default: 17:00 (5pm) local time today.

If it’s already past 17:00, you can:

Either use now + 1 hour, or

Move it to tomorrow at 17:00 (your call; just be consistent).

Examples:

me to send feedback before end of day → today at 17:00.

me to finalise slides before EOD → today at 17:00.

2.8 Relative durations (“in X …”)

If the phrase clearly says “in N minutes/hours/days”:

Use the current time as the base and add that duration.

"in 2 hours" → now + 2 hours.

"in 30 minutes" → now + 30 minutes.

"in 3 days" → 3 days from now, same time of day.

These are explicit, so no extra default needed.