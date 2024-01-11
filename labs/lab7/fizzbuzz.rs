fn main() {
    // Statements here are executed when the compiled binary is called.
    let mut i = 1;

    while i < 101 {
        if (i%3==0) && (i%5==0) {
            println!("FizzBuzz");
        }
        else if i%3==0 {
            println!("Fizz");
        }
        else if i%5==0 {
            println!("Buzz");
        }
        else {
            println!(i);
        }
        i = i + 1;
    }
}