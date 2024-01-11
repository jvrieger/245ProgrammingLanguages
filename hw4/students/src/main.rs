//Julia Rieger
//11/7/23
//main.rs - students

use std::fs::read_to_string;
use std::fmt;

struct Student {
    first: String,
    last: String,
    id: u64,
} //Student

// methods for Student
impl Student {
    //returns valid ID based on first and last name
    fn check_num(&self) -> u64 {
        let fch = self.first.chars().next().unwrap();
        let mut id: u64 = fch as u64 * 100;
        let sch = self.last.chars().next().unwrap();
        id = id + sch as u64;
        return id;
    }
    //checks whether given id is valid id
    fn check_id(&self) -> bool {
        if self.id == self.check_num() {
            return true;
        }
        return false;
    }
}

//overrides fmt Debug to print formatted Student
impl fmt::Debug for Student {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_tuple("")
         .field(&self.first)
         .field(&self.last)
         .field(&self.id)
         .field(&self.check_id())
         .finish()
    }
}

fn main() {
    let contents: Vec<String> = read_lines("students.txt");     //hold all lines from file
    let mut students: Vec<Student> = Vec::new();                //will hold all Student instances to be printed

    let mut i = 0;
    while i < contents.len() {            //for every line in the file
        let temp: &String = &contents[i]; //to deal with ownership
        students.push(build_stu(temp));   //construct and add Student to students
        i = i + 1;
    }
    println!("{:?}", students);           //print final results
}

//constructs Student from line from file
fn build_stu(line: &String) -> Student {
    let parts = line.split(" "); //tokenize
    let collection = parts.collect::<Vec<&str>>();
    
    Student { //construct and return
        first: collection[0].to_string(),
        last: collection[1].to_string(),
        id: collection[2].to_string().parse::<u64>().unwrap(),
    }
}

//returns a vector of all lines in file
fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}
