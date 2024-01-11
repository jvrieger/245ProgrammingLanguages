//Julia Rieger
//11/7/23
//main.rs - perfect numbers

fn main() {
    //let mut numbers: Vec<Vec<i32>> = vec![vec![]];
    //array of vectors to store near/perf numbers: index 0 = -5, index 1 = -4, etc.
    let mut numbers: [Vec<i32>; 11] = [vec![], vec![], vec![], vec![], vec![], vec![], vec![], vec![], vec![], vec![], vec![]];
    let mut i = 6;
    while i <= 50000000 { //set to upper bound value
        let curr = i;
        let sum = sum(factor(curr)); //sum of all factors for i

        //assign and populate cases from -5 to 5 for near perf numbers
        if sum == curr-5 {numbers[0].push(curr);}
        if sum == curr-4 {numbers[1].push(curr);}
        if sum == curr-3 {numbers[2].push(curr);}
        if sum == curr-2 {numbers[3].push(curr);}
        if sum == curr-1 {numbers[4].push(curr);}
        if sum == curr {numbers[5].push(curr);}
        if sum == curr+1 {numbers[6].push(curr);}
        if sum == curr+2 {numbers[7].push(curr);}
        if sum == curr+3 {numbers[8].push(curr);}
        if sum == curr+4 {numbers[9].push(curr);}
        if sum == curr+5 {numbers[10].push(curr);}

        i = i + 1;
    }

    //print output
    let mut count = -5;
    let j = 0;
    for j in j..numbers.len() {
        println!("{}: {:?}", count, numbers[j]);
        count = count + 1;
    }
}

//takes an i32 and returns a vector of all its factors
//including 1, not including itself
fn factor(input: i32) -> Vec<i32>{
    let count = 1;
    let mut vector = vec![];
    for count in count..input {
        if input % count == 0 {
            vector.push(count);
        }
    }
    return vector;
}

//takes an i32 vector and returns the sum of all values
fn sum(factors: Vec<i32>) -> i32 {
    let mut total = 0;
    let i = 0;
    for i in i..factors.len() {
        total = total + factors[i];
    }
    return total;
}