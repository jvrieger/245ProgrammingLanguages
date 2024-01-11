use std::env;

fn main() {
    let mut args: Vec<String> = env::args().collect();
    //remove misc args
    args.remove(0);
    args.remove(0);

    let mut fin: Vec<&String> = vec![];

    args.iter().for_each(|s| if s.chars().next().unwrap().is_numeric() {fin.push(s);println!("{:?}", fin);});

    
}