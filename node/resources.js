const axios = require("axios");
const { response } = require("express");
const randomNumber = require("random-number-csprng");

const letters = "abcdefghijklmnopqrstuvwxyz";
const passphraseSize = 5;

class PassPhrase {
  constructor(passphrase) {
    this.passphrase = passphrase;
  }
}

const url = (word) => `http://api.datamuse.com/words?sp=${word}*`;

const randomItem = async (arr, max = null) => {
  return arr[await randomNumber(0, (max || arr.length) - 1)];
};

const wordRequest = async () => {
  let response = await axios.get(url(await randomItem(letters)));
  return await randomItem(response.data);
};

const getPassPhrase = async () => {
  let words = [];
  const insertIntoWords = async () => {
    let {word} = await wordRequest()
    words.push(word);
  };
  const proms = [];
  for (let i = 0; i < passphraseSize; i++) {
    proms.push(insertIntoWords());
  }
  await Promise.all(proms);

  for (let i = words.length - 1; i > 0; i--) {
    const j = await randomNumber(0, i);
    const temp = words[i];
    words[i] = words[j];
    words[j] = temp;
  }
  return new PassPhrase(words.join(" "));
};

module.exports = {
  getPassPhrase: getPassPhrase,
};
