# Age Calculator Bot 🤖
A slack bot developed using Go to calculate age with implemented functionality for event handling. The age is determined using the current year in the program and given the year of birth (YOB).

You will be able to create this slack bot only if you are an admin for the workspace. Given that you are an admin for the workspace, navigate to api.slack.com/apps, and create an app from scratch.

Follow these steps within settings:
1. Socket Mode: name token and generate token
2. Event Subscriptios: Enable events, subscribe to bot events — app_mention, im_history_changed, message.im, message.channels, message.mpim
3. OAuth & Permissions: Add permissions - chat:write, chat:write.public, channels:read, im:read, im:write, mpim:read, mpim:write; install to workspace

xapp represents app token and xoxb represents bot token — replace with generated xapp and xoxb within .Setenv() calls.
