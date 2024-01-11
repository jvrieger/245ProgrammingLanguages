fn main() {
    let vec0 = vec![22, 44, 66];
    let vec1 = fill_vec(vec0);
    println!("{:?}", vec1);
}

fn fill_vec(vec: Vec<i32>) -> Vec<i32> {
    let mut vec = vec;
    vec.push(88);
    vec
}

//add <i32> Generics in function header of line 7 so program will compile
//change println! to vec1 in line 4 because vec0 is owned by fill_vec and can't be borrowed
//make vec mutable in line 8 because we are adding to vec so it must be mutable