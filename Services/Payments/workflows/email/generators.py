"""
    Email body generators
"""
from email.mime.multipart import MIMEMultipart
from email.mime.image import MIMEImage
from email.mime.text import MIMEText
import string
import Config as service_config
import repository
import models
import os

def loadLogo() -> tuple[bytes, str]:
    email_module_path = os.path.dirname(os.path.abspath(__file__)) # this will look like: /home/runner/work/Expert-Appointment-Booking-System/Expert-Appointment-Booking-System/workflows/email
    image_filename = 'logo.png'
    
    with open(f'{email_module_path}/resources/logo.png', 'rb') as logo_file:
        logo = logo_file.read()
        
    return logo, image_filename

def successfulPaymentEmail(appointment:models.Appointment, payment:models.Payment, checkout_id: str) -> tuple[str, str]:
    expert_personal_info, err = repository.experts.getExpertPersonalData(appointment.expert)
    if err:
        print(f'Error getting expert personal data: {err}')
        return None, None

    
    email_module_path = os.path.dirname(os.path.abspath(__file__)) # this will look like: /home/runner/work/Expert-Appointment-Booking-System/Expert-Appointment-Booking-System/workflows/email
    
    with open(f'{email_module_path}/templates/successful_payment.html', 'r') as html_file:
        email_html = html_file.read()
        template = string.Template(email_html)
        
    plain_text = f"""
        Querida {appointment.customer_email},
        
        Tu pago por {payment.amount}MXN ha sido recibido, tu cita con {expert_personal_info.name} a sido agendada con éxito.
        Gracias por tu confianza.
    """
    
    
    msg = MIMEMultipart('alternative')
    msg['Subject'] = 'Pago exitoso'
    msg['From'] = service_config.MAIL_FROM
    msg['To'] = appointment.customer_email
    
    
    html_body = template.substitute(customer_name=appointment.customer_email, total_price="{price:,.2f}".format(price=payment.amount/100), expert_name=expert_personal_info.name)
    msg.attach(MIMEText(html_body, 'html'))
    
    
    return msg, plain_text
        
    