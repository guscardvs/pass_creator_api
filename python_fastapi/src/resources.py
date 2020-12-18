import asyncio
from typing import List
from pydantic import BaseModel
from aiohttp import ClientSession
from secrets import randbelow
from random import shuffle


LETTERS = "abcdefghijklmnopqrstuvwxyz"
PASSPHASE_SIZE = 5


class PassPhrase(BaseModel):
    passphrase: str


URL = "http://api.datamuse.com/words?sp={letter}*"


def random_item(list_, max_):
    response = list_[randbelow(max_)]
    return response


async def word_request():
    async with ClientSession() as session:
        async with session.get(URL.format(letter=random_item(LETTERS, 26))) as response:
            data = await response.json()
            return random_item(data, len(data))

async def get_passphrase():
    words: List[str] = []
    for cor in asyncio.as_completed([word_request() for _ in range(PASSPHASE_SIZE)]):
        words.append((await cor).get("word"))
    shuffle(words)
    string = " ".join(words)
    return string
    
