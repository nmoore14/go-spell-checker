import random
import argparse

# Parse command-line arguments
parser = argparse.ArgumentParser()
parser.add_argument('--dictionary', required=True, help='Path to dictionary file')
args = parser.parse_args()

# Open the dictionary text file and read its contents into a list
with open(args.dictionary, 'r') as f:
    dictionary = [word.strip() for word in f.readlines()]

# Generate 5 random sentences using the words from the dictionary
def buildSentences():
    sentences = []
    for i in range(5):
        sentence = ''
        for j in range(random.randint(5, 25)):
            word = random.choice(dictionary)
            sentence += ' ' + word
        sentences.append(sentence + '\n')
    return sentences

# Write the generated sentences to separate text files
for i in range(5):
    sentences = buildSentences()
    with open('./test_{}.txt'.format(i), 'w') as f:
        for k, sentence in enumerate(sentences):
            f.writelines(sentence)
