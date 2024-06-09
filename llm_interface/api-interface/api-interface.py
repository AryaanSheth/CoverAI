from os import getenv
from dotenv import load_dotenv
from groq import Groq

class GroqInterface:
    def __init__(self):
        load_dotenv()
        self.client = Groq(
            api_key = getenv('GROQ_API_KEY')
            )
        self.model = "llama3-8b-8192"
        self.role = "user"

    def get_response(self, message: str):
        chat_completion = self.client.chat.completions.create(
            messages=[
                {
                    "role": self.role,
                    "content": message,
                }
            ],
            model=self.model,
        )
        return chat_completion
    
    def __call__(self, message: str):
        return self.get_response(message)
    
    def __repr__(self):
        return "Type: GroqInterface\nModel: {}\nRole: {}".format(self.model, self.role)
