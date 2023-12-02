use std::io::{self, Write};

pub struct Shell {
    history: Vec<String>,
    reader: io::Stdin,
}

impl Shell {
    pub fn new() -> Shell {
        Shell {
            history: Vec::new(),
            reader: io::stdin(),
        }
    }

    pub fn run(&mut self) {
        let mut input = String::new();

        loop {
            print!("shell> ");
            io::stdout().flush().unwrap();

            input.clear();

            match self.reader.read_line(&mut input) {
                Ok(_) => {
                    let trimmed = input.trim();
                    if trimmed == "exit" {
                        break;
                    }
                    self.history.push(trimmed.to_string());
                    println!("executing: {}", trimmed);
                    println!("history: {:?}", self.history);
                }
                Err(error) => println!("Error: {}", error),
            }
        }
    }
}
