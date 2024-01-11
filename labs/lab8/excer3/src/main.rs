fn main() {
    let vec0 = vec![22, 44, 66];
    let mut vec1 = fill_vec(vec0);
    println!("{:?}",vec1);
}

fn fill_vec(mut vec: Vec<i32>) -> Vec<i32> {
    vec.push(88);
    vec
}

//add generics for Vec to make the compiler happy, it must know what type the vec is storing
//make vec parameter (line 7) mutable in the function header: then 88 can be pushed on