package utils

func BuildShareLinkHTML(shareLink, yourName string) string {
	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Share Link</title>
			<style>
				body { font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 0; }
				.container { background-color: #ffffff; margin: 20px auto; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); max-width: 600px; }
				.header { text-align: center; padding: 10px 0; background-color: #0078d4; color: white; border-radius: 10px 10px 0 0; }
				.content { padding: 20px; }
				.content p { margin: 0 0 20px; }
				.link-button { display: inline-block; padding: 10px 20px; margin: 10px 0; font-size: 16px; color: white; background-color: #0078d4; text-decoration: none; border-radius: 5px; }
				.footer { text-align: center; padding: 10px 0; font-size: 12px; color: #666; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Nextcloud Share</h1>
				</div>
				<div class="content">
					<p>Hello,</p>
					<p>You have been invited to access a shared file on Nextcloud.</p>
					<p>Click the button below to view the shared content:</p>
					<a href="` + shareLink + `" class="link-button">View Shared Content</a>
					<p>If the button does not work, you can copy and paste the following link into your browser:</p>
					<p><a href="` + shareLink + `">` + shareLink + `</a></p>
					<p>Best regards,<br>"` + yourName + `"</p>
				</div>
				<div class="footer">
					<p>© 2024 Ryan Lake. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>
	`
}