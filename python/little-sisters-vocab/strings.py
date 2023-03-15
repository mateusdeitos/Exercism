"""Functions for creating, transforming, and adding prefixes to strings."""


from typing import List


def add_prefix_un(word: str):
	"""Take the given word and add the 'un' prefix.

	:param word: str - containing the root word.
	:return: str - of root word prepended with 'un'.
	"""

	return f"un{word}"


def make_word_groups(vocab_words: List[str]):

	prefix = vocab_words[0]

	words = vocab_words[1:]

	words_prefixed = map(lambda word: f"{prefix}{word}", words)

	return f"{prefix} :: " + " :: ".join(words_prefixed)


def remove_suffix_ness(word: str):
	"""Remove the suffix from the word while keeping spelling in mind.

	:param word: str - of word to remove suffix from.
	:return: str - of word with suffix removed & spelling adjusted.

	For example: "heaviness" becomes "heavy", but "sadness" becomes "sad".
	"""

	suffix = "ness"
	word_without_suffix = word[:-len(suffix)]
	last_char = word_without_suffix[-1]
	if last_char == "i":
		word_without_suffix = word_without_suffix[:-1] + "y"

	return word_without_suffix


def adjective_to_verb(sentence: str, index: int):
	"""Change the adjective within the sentence to a verb.

	:param sentence: str - that uses the word in sentence.
	:param index: int - index of the word to remove and transform.
	:return: str - word that changes the extracted adjective to a verb.

	For example, ("It got dark as the sun set", 2) becomes "darken".
	"""

	word = sentence.split(" ")[index].strip(".,!?")

	return f"{word}en"
