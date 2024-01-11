//Julia Rieger
//11/4/23
//main.rs - lucas sequences

fn main() {
    let l = 10; //L value
    let count: i32 = 1; //count or index of the term

    //call recursive func to print next term until program dies
    let _ratio = next(1, l, count);

}

//recursive func prints next term in lucas sequence until program dies
//p: n-2, l: n-1, count: index of p+l in the sequence
fn next(p: i32, l: i32, mut count: i32) -> i32 {
    count += 1;
    let largest = l; //to be printed once threshold is hit
    if (p as i64 +l as i64) > 2147483647 { //if the sum of the last 2 terms will overflow the i32 storage (testing with i64) we must terminate
        println!("The largest term we can reach with 32 bits is {}: {}", count, largest);

        let ratio = largest as f64/(p as f64 +l as f64); //calc ratio
        println!("The ratio of the largest term we can compute with 32 bits and the next term is: {}", ratio);

        return ratio as i32;                         
    }
    //println!("{}: {}", count, p+l); //will print every term
    return next(l, p+l, count);
}