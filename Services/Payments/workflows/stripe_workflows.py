from datetime import datetime, timedelta
import Config as service_config
import repository
import models
import stripe
import json

stripe.api_key = service_config.LOCAL_CONFIG["stripe_pk"]

def createLineItem(product_name:str, price_in_cents:int ,extra_parametes: dict[str, any]) -> dict[str, any]:
    line_item = {
        "price_data": {
            "currency": extra_parametes.get("currency", None) or "mxn",
            "unit_amount": price_in_cents, # price in cents
            "product_data": {
                "name": product_name,
            }
        },
        "quantity": 1 or extra_parametes.get("quantity", None),
    }
    
    if image := extra_parametes.get("image", None):
        line_item["price_data"]["product_data"]["images"] = [image]

    if description := extra_parametes.get("description", None):
        line_item["description"] = description
        
    if tax_rates := extra_parametes.get("tax_rates", None):
        line_item["tax_rates"] = tax_rates
        
    # if furute me wants to add discounts, this is the place. its key is "discounts" and body is a list of discounts each of this is a dict with the following keys: 'amount_off', 'description'
    
    return line_item

def createCheckoutSession(line_items: list[dict[str, any]], success_url: str, cancel_url: str, extra_params: dict[str, any]) -> stripe.checkout.Session:
    
    assert "identifier" in extra_params, "identifier is required in extra_params"
    
    # create checkout session that expires in 4 minutes
    checkout_session = stripe.checkout.Session.create(
        payment_method_types=["card"],
        line_items=line_items,
        mode="payment",
        success_url=success_url,
        cancel_url=cancel_url,
        customer_email=extra_params.get("customer_email", None),
        client_reference_id=extra_params.get("identifier", None),
        metadata=extra_params.get("metadata", [])
    )
    
    return checkout_session

def createAppointmentCheckoutSession(appointment: models.Appointment, success_url:str, cancel_url:str, cache_hash: str) -> stripe.checkout.Session:
    extra_params = {
        "identifier": cache_hash,
        "customer_email": appointment.customer_email,
        "metadata": appointment.toDict()
    }
    
    print(f"extra_params: {extra_params}")
    
    expert_data, err = repository.experts.getExpertPersonalData(appointment.expert)
    if err:
        print(f"Error while getting expert personal data: {err}")
        return None
    
    price_in_cents = int((appointment.durationMinutes() * expert_data.MinutePrice) * 100)
    
    line_items = [
        createLineItem(f"Cita con {expert_data.name}", price_in_cents, extra_params)
    ]
    
    success_url += f"?session_id={{CHECKOUT_SESSION_ID}}"
    
    return createCheckoutSession(line_items, success_url, cancel_url, extra_params)

def retriveCheckoutSession(session_id: str) -> tuple[stripe.checkout.Session, Exception]:
    checkout_session = None
    err = None
    
    try:
        checkout_session = stripe.checkout.Session.retrieve(session_id)
    except Exception as e:
        err = e

    return checkout_session, err

def retrivePaymentIntent(checkout_session: stripe.checkout.Session) -> tuple[stripe.PaymentIntent, Exception]:
    payment_intent = None
    err = None
    
    try:
        payment_intent = stripe.PaymentIntent.retrieve(checkout_session.payment_intent)
    except Exception as e:
        err = e

    return payment_intent, err

def retriveCharge(checkout_session: stripe.checkout.Session, payment_intent: stripe.PaymentIntent) -> tuple[stripe.Charge, Exception]:
    err = None
    if not payment_intent:
        payment_intent, err = retrivePaymentIntent(checkout_session)
        if err:
            return None, err

    charge = None
    
    try:
        charge = stripe.Charge.retrieve(payment_intent.latest_charge)
    except Exception as e:
        err = e
    
    return charge, err


def isSessionPaid(checkout_session: stripe.checkout.Session) -> bool:
    return checkout_session.payment_status == "paid"

def refoundSession(checkout_session: stripe.checkout.Session, reason: str=None) -> tuple[stripe.Refund, Exception]:
    charge = retriveCharge(checkout_session)
    available_reason = reason != None
    reason = reason or "No reason provided"
    
    refund = stripe.Refund.create(charge=charge.id, metadata={"reason": reason, "available": available_reason})
    payment, err = repository.payments.getPaymentBySession(checkout_session.id)
    if err:
        # TODO: in the checkout session there is enought information to create a payment, consider  inserting it the missing payment entry
        return None, err
    
    payment.stripe_refund_id = refund.id
    payment.status = models.PaymentStatus.REFUNDED

def registerPayment(checkout_session: stripe.checkout.Session, appointment: models.Appointment) -> Exception:
    
    # Check if session is paid
    if not isSessionPaid(checkout_session):
        return registerUnpayment(checkout_session, appointment)
    
    payment = models.createEmptyPayment(appointment.id, checkout_session.amount_total)
    payment.stripe_session_id = checkout_session.id
    
    payment_intent, err = retrivePaymentIntent(checkout_session)
    if not err:
        payment.stripe_payment_intent_id = payment_intent.id
    
    charge, err = retriveCharge(checkout_session, payment_intent)
    if not err:
        payment.stripe_charge_id = charge.id
    
    payment.status = models.PaymentStatus.SUCCEDED
    
    err = repository.payments.insertPayment(payment)
    
def registerUnpayment(checkout_session: stripe.checkout.Session, appointment: models.Appointment) -> Exception:
    if checkout_session.status == "open":
        return Exception("Session is still open")
    
    payment = models.createEmptyPayment(appointment.id, checkout_session.amount_total)
    appointment.status = models.AppointmentStatus.CANCELLED
    payment.appointment = appointment.id
    payment.stripe_session_id = checkout_session.id
    
    payment_intent, err = retrivePaymentIntent(checkout_session)
    if not err:
        payment.stripe_payment_intent_id = payment_intent.id
    
    payment.status = models.PaymentStatus.FAILED
    
    err = repository.payments.insertPayment(payment)
    return err