from .mysql_utils import MysqlConnection, MYSQL_CONFIG
from models import ExpertPersonalData

class ExpertsRepository:
    def __init__(self, config: MYSQL_CONFIG=None):
        if not config:
            config = MYSQL_CONFIG.createFromEnv()
        self.config = config
    
    def getExpertPersonalData(self, expert_id: int) -> tuple[ExpertPersonalData, None]:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT name, email, username, minute_price FROM experts WHERE id = ?"
            params = (expert_id,)
            cursor.execute(sql, params)
            result = cursor.fetchone()
            if result:
                return ExpertPersonalData(*result), None
            return None, None