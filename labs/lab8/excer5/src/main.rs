fn main() {
    let s1 = vec![1,2,3];
    let mut s2 = vec![4,5,6];
    println!("A {:?}, {:?}", s1, s2);
    g_o(s1);
    println!("B {:?}", s2); //cannot print s1 bc ownership got passed to g_o
    t_a_g(&mut s2);
    println!("E {:?}", s2); //can print s2 because only address got passed, not actual vec
}

fn g_o(mut aa : Vec<i32>)  {             
    aa[0]=4;
    println!("C {:?}", aa);       
}

fn t_a_g(aa : &mut Vec<i32>) { 
    aa[0]=42;
    println!("D {:?}", aa);       

}

//A [1, 2, 3], [4, 5, 6]
//C [4, 2, 3]
//B [4, 5, 6]
//D [42, 5, 6]
//E [42, 5, 6]