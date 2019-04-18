// Codon to Protein map
const codonProteins = new Map([
  ["AUG", "Methionine"],
  ["UUU", "Phenylalanine"],
  ["UUC", "Phenylalanine"],
  ["UUA", "Leucine"],
  ["UUG", "Leucine"],
  ["UCU", "Serine"],
  ["UCC", "Serine"],
  ["UCA", "Serine"],
  ["UCG", "Serine"],
  ["UAU", "Tyrosine"],
  ["UAC", "Tyrosine"],
  ["UGU", "Cysteine"],
  ["UGC", "Cysteine"],
  ["UGG", "Tryptophan"],
  ["UAA", "STOP"],
  ["UAG", "STOP"],
  ["UGA", "STOP"],
]);

//Translate RNA sequences into proteins
const translate = function (rna) {
  if (!rna || "" === rna || rna.length % 3 !== 0)
    return [];

  let ret = [];
  let idx = 0;
  let stop = false;

  while (idx < rna.length && !stop) {
    const codon = rna.slice(idx, idx + 3);
    const protein = codonProteins.get(codon);
    if (!protein) {
      throw "Invalid codon";
    } else if (protein && protein !== "STOP") {
      ret.push(protein);
      idx = idx + 3;
    } else {
      stop = true;
    }
  }
  return ret;
};

export {translate};
