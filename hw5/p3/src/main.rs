use std::env;

fn main() {
    let mut args: Vec<String> = env::args().collect();
    //remove misc args
    args.remove(0);
    args.remove(0);

    args.iter().for_each(|s| println!("{s} {}", s.len()));
}