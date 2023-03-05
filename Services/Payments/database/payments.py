from .mysql_utils import MysqlConnection, MYSQL_CONFIG
from models import Payment

class PaymentsRepository:
    def __init__(self, config: MYSQL_CONFIG=None):
        if not config:
            config = MYSQL_CONFIG.createFromEnv()
        self.config = config
    
    def insertPayment(self, payment: Payment) -> Exception:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "INSERT INTO payments (appointment, amount, currency, stripe_charge_id, stripe_customer_id, stripe_payment_intent_id, stripe_session_id, stripe_refund_id, description, status, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
            params = (payment.appointment, payment.amount, payment.currency, payment.stripe_charge_id, payment.stripe_customer_id, payment.stripe_payment_intent_id, payment.stripe_session_id, payment.stripe_refund_id, payment.description, payment.status.value, payment.created_at)
            cursor.execute(sql, params)
            conn.commit()
            return None
    
    def updatePayment(self, payment: Payment) -> Exception:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "UPDATE payments SET appointment=?, amount=?, currency=?, stripe_charge_id=?, stripe_customer_id=?, stripe_payment_intent_id=?, stripe_session_id=?, stripe_refund_id=?, description=?, status=?, created_at=? WHERE id=?"
            params = (payment.appointment, payment.amount, payment.currency, payment.stripe_charge_id, payment.stripe_customer_id, payment.stripe_payment_intent_id, payment.stripe_session_id, payment.stripe_refund_id, payment.description, payment.status.value, payment.created_at, payment.id)
            cursor.execute(sql, params)
            conn.commit()
            return None
    
    def getAllPayments(self) -> tuple[list[Payment], Exception]:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT id, appointment, amount, currency, stripe_charge_id, stripe_customer_id, stripe_payment_intent_id, stripe_session_id, stripe_refund_id, description, status, created_at FROM payments"
            cursor.execute(sql)
            result = cursor.fetchall()
            
            err = None if result else Exception("No payments found")
            
            return [Payment(*row) for row in result], err
    
    def getPaymentBySession(self, session_id: str) -> tuple[Payment, Exception]:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT id, appointment, amount, currency, stripe_charge_id, stripe_customer_id, stripe_payment_intent_id, stripe_session_id, stripe_refund_id, description, status, created_at FROM payments WHERE stripe_session_id=?"
            cursor.execute(sql, (session_id,))
            result = cursor.fetchone()
            
            err = None if result else Exception("No payment found")
            payment = None
            
            if not err:
                payment = Payment(*result)

            return payment, err
        
    def getPaymentByAppointment(self, appointment_id: int) -> tuple[Payment, Exception]:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT id, appointment, amount, currency, stripe_charge_id, stripe_customer_id, stripe_payment_intent_id, stripe_session_id, stripe_refund_id, description, status, created_at FROM payments WHERE appointment=?"
            cursor.execute(sql, (appointment_id,))
            result = cursor.fetchone()
            
            err = None if result else Exception("No payment found")
            payment = None
            
            if not err:
                payment = Payment(*result)

            return payment, err
                