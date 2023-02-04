package AuthEmail

import (
	"log"
	"net/smtp"

	app_config "bonhart_auth_service/Config"
)

type EmailMessage struct {
	To      string
	From    string
	FromAlt string
	Subject string
	Body    string
}

func sendEmail(message EmailMessage) error {

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		app_config.MAIL_USERNAME,
		app_config.MAIL_PASSWORD,
		app_config.MAIL_SERVER,
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{message.To}
	msg := []byte(
		"From: " + message.FromAlt + " <" + message.From + ">\r\n" +
			"To: " + message.To + "\r\n" +
			"Subject: " + message.Subject + "\r\n" +
			"\r\n" +
			message.Body + "\r\n")

	err := smtp.SendMail(
		app_config.MAIL_SERVER+":"+app_config.MAIL_PORT,
		auth,
		message.From,
		to,
		msg,
	)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func sendHtmlEmail(message EmailMessage) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		app_config.MAIL_USERNAME,
		app_config.MAIL_PASSWORD,
		app_config.MAIL_SERVER,
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{message.To}
	msg := []byte(
		"From: " + message.FromAlt + " <" + message.From + ">\r\n" +
			"To: " + message.To + "\r\n" +
			"Subject: " + message.Subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"" + "\r\n" +
			"\r\n" +
			message.Body + "\r\n")
	err := smtp.SendMail(
		app_config.MAIL_SERVER+":"+app_config.MAIL_PORT,
		auth,
		message.From,
		to,
		msg,
	)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func SendRecoveryEmail(to string, recovery_url string) error {
	// send email
	email := new(EmailMessage)

	email.To = to
	email.From = app_config.MAIL_USERNAME
	email.FromAlt = app_config.MAIL_FROM
	email.Subject = "Password Recovery"
	email.Body = "Please click the following link to recover your password: " + recovery_url

	return sendEmail(*email)
}
