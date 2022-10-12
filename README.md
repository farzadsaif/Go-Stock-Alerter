# Go-Stock-Alerter
A simple stock alerter using the Go programming language

This stock alerter requires a Gmail account with "less secure app access" turned on. Currently Gmail does not allow for "less secure app access" for personal accounts. You must have a Google Workspace account to use "less secure app access".

If you would like to use your own SMTP email server, you can change the host and the port in the email function in the source code.

When you run the script you will be presented with some questions:

**Enter Stock Threshold**: Enter the percentage change at which the stock alerter will trigger. So for example, if you would like the alert to trigger when the stock falls by 2%, you would input -2%.

**Enter Stock Ticker**: Enter your ticker. For example, AAPL.

**Enter From Email**: Enter the username@gmail.com of the email which the alert will send from. This must be a Gmail account with "less secure app access" turned on.

**Enter Password**: Enter the password of the email which the alert will send from.

**Enter to Email**: Enter the email which the alert will send to. If you want to send an email to yourself, just type your email again.

**Enter SMS Number @ SMS gateway**: Enter the 10 digit number@gateway which you will send an SMS alert to. For example, if the carrier is Verizon Wireless, then you would type 12345678900@vtext.com. [Here is a list of SMS gateways](https://avtech.com/articles/138/list-of-email-to-sms-addresses/)

Press enter to run the stock alerter.
