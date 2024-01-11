//Julia Rieger
//11/7/23
//main.rs - stack depth

use rand::Rng;

fn main() {
    //rec_1(1);

    //rec_2(1, 0);

    //let arr: [i32; 100] = [0;100];
    //rec_3(1, arr);

    let mut rng = rand::thread_rng();
    let mut vector: Vec<i128> = vec!();
    let mut i = 0;
    while i < 1000 {
        vector.push(rng.gen::<i128>());
        i  = i + 1;
    }
    rec_4(1, vector);
}

fn rec_1(count: u128) {
    println!("{}", count);
    rec_1(count + 1);
}

fn rec_2(count: u32, param: i128) {
    println!("{}", count);
    rec_2(count + 1, param);
}

fn rec_3(count: u32, param: [i32; 100]) {
    println!("{}", count);
    rec_3(count + 1, param);
}

fn rec_4(count: u32, param: Vec<i128>) {
    println!("{}", count);
    rec_4(count + 1, param);
}

//rec_1: (range over 5 runs)
//i32: ~74822
//u32: 74786 - 74819
//u64: 74786 - 74840
//u128: 65446 - 65497

//rec_2:
//58210
//58208
//58198
//58203
//58172

//rec_3:
//16369
//16358
//16364
//16360
//16370

//rec_4:
//empty vec and 1000 rand nums: (same output)
//47600
//47623
//47630
//47598
//47615

