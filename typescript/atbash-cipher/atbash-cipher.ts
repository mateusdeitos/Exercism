const ALPHABET = "abcdefghijklmnopqrstuvwxyz";
const CIPHER = "zyxwvutsrqponmlkjihgfedcba";

const ALPHABET_TO_CIPHER = ALPHABET.split("").reduce<Record<string, string>>(
  (acc, char, index) => {
    acc[char] = CIPHER[index];
    return acc;
  },
  {}
);

const CIPHER_TO_ALPHABET = Object.fromEntries(
  Object.entries(ALPHABET_TO_CIPHER).map(([key, value]) => [value, key])
);

export function encode(plainText: string): string {
  const normalizedText = plainText.replace(/[^a-z0-9]/gi, "").toLowerCase();
  const chunks = normalizedText.match(/.{1,5}/g);

  if (!chunks?.length) {
    return "";
  }

  const encoded = chunks
    .map((chunk) =>
      chunk
        .split("")
        .map((char) => ALPHABET_TO_CIPHER[char] ?? char)
        .join("")
    )
    .join(" ");

  return encoded;
}

export function decode(cipherText: string): string {
  return cipherText
    .split(" ")
    .map((word) =>
      word
        .split("")
        .map((char) => CIPHER_TO_ALPHABET[char] ?? char)
        .join("")
    )
    .join("");
}
