from uvicorn import run
from fastapi import FastAPI

from resources import get_passphrase

app = FastAPI()

@app.get("/")
async def create_passphrase():
    return await get_passphrase()

if __name__=="__main__":
    run("main:app", host="0.0.0.0", port=3000, reload=True)