# slack-notifier

## Overview

`slack-notifier` is a simple Go application for sending messages to a Slack channel using a webhook URL.

## Configuration


The application uses the `config` package for configuration management. You can set the configuration using environment variables or command-line flags.

### Configuration Keys
- `SLACK_WEBHOOK`: The Slack webhook URL used to send messages.

### Setting Up Slack Webhook

To send messages to a Slack channel, you need to set up a webhook URL in your Slack workspace. Follow these steps:

1. Go to your Slack workspace.
2. Navigate to "Apps" and search for "Incoming WebHooks".
3. Click on "Add to Slack" and follow the instructions to create a new webhook.
4. Copy the generated webhook URL.

### Setting Up Configuration

#### Environment Variables

You can set the configuration keys as environment variables. For example:

```bash
export SLACK_WEBHOOK="https://your-slack-webhook-url"
```

#### CLI Flags

You can also pass configuration keys as command-line flags when running the application:

```bash
go run cmd/notifier/main.go --slack-webhook="https://your-slack-webhook-url"
```

## Usage

The application initializes a Slack webhook using the configured URL and sends a test message. Below is an example of how the application works:

1. The `SLACK_WEBHOOK` URL is added to the configuration using the `config` package.
2. The `slack.NewWebhook` function is used to create a new webhook instance.
3. The `webhook.Send` method sends a message to the Slack channel.

### Example

```bash
# Run the application
go run cmd/notifier/main.go
```

If the configuration is correct, you should see the following output:

```
Message sent to Slack successfully!
```

If there is an error, it will print:

```
Error sending message to Slack: <error details>
```

## Using `pkg/slack` in Other Go Modules

You can use the `pkg/slack` package in other Go modules to send messages to Slack. Below is an example:

### Steps:
1. Import the `pkg/slack` package in your Go module:
   ```go
   import "github.com/ilho-tiger/slack-notifier/pkg/slack"
   ```

2. Create a new Slack webhook instance:
   ```go
   webhook := slack.NewWebhook("https://your-slack-webhook-url")
   ```

3. Send a message to Slack:
   ```go
   err := webhook.Send("Hello from another Go module!")
   if err != nil {
       fmt.Printf("Error sending message to Slack: %v\n", err)
   } else {
       fmt.Println("Message sent to Slack successfully!")
   }
   ```

### Example Code

```go
package main

import (
	"fmt"
	"github.com/ilho-tiger/slack-notifier/pkg/slack"
)

func main() {
	webhook := slack.NewWebhook("https://your-slack-webhook-url")
	if err := webhook.Send("Hello from another Go module!"); err != nil {
		fmt.Printf("Error sending message to Slack: %v\n", err)
	} else {
		fmt.Println("Message sent to Slack successfully!")
	}
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or have feature requests.