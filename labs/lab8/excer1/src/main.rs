fn main() {
    let mut list = vec![1, 2, 3];
    println!("Before defining closure: {:?}", list);
    let mut borrows_mutably = || list.push(7);
    println!("{:?}", list); 
    borrows_mutably();
    println!("After calling closure: {:?}", list);
}

//compiler cannot borrow list as immutable in println!() because it is also borrowed as mutable in line 4
//first borrow happens because list is in closure, || can be seen as parenthesis and list as arguement
//mutable borrow later used in line 6 when closure is called
