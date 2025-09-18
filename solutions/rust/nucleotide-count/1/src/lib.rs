use std::{char, collections::HashMap};

fn valid_nucleotide(c: char) -> bool {
    let nucleotides = ['A', 'C', 'G', 'T'];
    nucleotides.contains(&c)
}

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !valid_nucleotide(nucleotide) {
        return Err(nucleotide);
    }
    let mut counter = 0;
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c);
        }
        if c == nucleotide {
            counter += 1;
        }
    }
    Ok(counter)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut map: HashMap<char, usize> = [('A', 0), ('G', 0), ('T', 0), ('C', 0)]
        .into_iter()
        .collect();
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c);
        }
        map.entry(c).and_modify(|counter| *counter += 1);
    }
    Ok(map)
}
