from email.mime.multipart import MIMEMultipart
from email.message import EmailMessage
from email.mime.text import MIMEText
import Config as service_config
from . import generators
import smtplib
import models
import ssl

# def _sendEmail(msg: EmailMessage) -> Exception:
#     err: Exception = None
#     print("Attempting to send email to: ", msg['To'])
#     try: 
#         context = ssl.create_default_context()
#         print("Entering context")
#         with smtplib.SMTP_SSL(service_config.MAIL_SERVER, service_config.MAIL_PORT, context=context) as server:
#             server.login(service_config.MAIL_USERNAME, service_config.MAIL_PASSWORD)
#             err = server.sendmail(service_config.MAIL_USERNAME, msg['To'], msg.as_string())
#             print("Email errs: ", err)
#     except Exception as e:
#         err = e
    
#     print(f"{err=}")
#     return err

def _sendEmail(msg: EmailMessage) -> Exception:
    err: Exception = None
    print("Attempting to send email to: ", msg['To'])
    try: 
        context = ssl.create_default_context()
        context.options |= ssl.OP_NO_TLSv1 | ssl.OP_NO_TLSv1_1
        context.set_ciphers('HIGH:!aNULL:!eNULL:!EXPORT:!CAMELLIA:!DES:!MD5:!PSK:!RC4')
        with smtplib.SMTP(service_config.MAIL_SERVER, service_config.MAIL_PORT) as server:
            server.ehlo()
            server.starttls(context=context)
            server.ehlo()
            server.login(service_config.MAIL_USERNAME, service_config.MAIL_PASSWORD)
            err = server.sendmail(service_config.MAIL_USERNAME, msg['To'], msg.as_string())
            print("Email errs: ", err)
    except Exception as e:
        err = e
    
    print(f"{err=}")
    return err

# FIXME: Some data in the email is not being replaced

def sendSuccessfulPaymentEmail(appointment:models.Appointment, payment:models.Payment, checkout_id: str) -> Exception:
    email_message, email_plain_text = generators.successfulPaymentEmail(appointment, payment, checkout_id)
    print("Sending email to: ", email_message['To'])
    return _sendEmail(email_message)
    
    