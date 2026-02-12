use std::io::{self, BufRead};

fn compress(line: String) {
    let mut chars = line.chars().peekable();

    while let Some(curr) = chars.next() {
        let mut cnt = 1;

        while let Some(&next) = chars.peek() {
            if next == curr {
                chars.next();
                cnt += 1;
            } else {
                break;
            }
        }
        if cnt > 2 {
            print!("{}{}", cnt, curr);
        } else {
            for _ in 0..cnt {
                print!("{}", curr);
            }
        }
    }
    println!();
}

fn main() {
    let stdin = io::stdin();

    for lines in stdin.lock().lines() {
        match lines {
            Ok(line) => {
                compress(line);
            }
            Err(error) => {
                eprintln!("Error reading line: {}", error);
                break;
            }
        }
    }
}
