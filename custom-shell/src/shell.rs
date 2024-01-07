use std::{
    io::{self, Write},
    path::PathBuf,
};

pub struct Shell {
    history: Vec<String>,
    reader: io::Stdin,
    working_dir: PathBuf,
}

impl Shell {
    pub fn new() -> Shell {
        Shell {
            history: Vec::new(),
            reader: io::stdin(),
            working_dir: dirs::home_dir().expect("Failed to get home directory"),
        }
    }

    pub fn run(&mut self) {
        let mut input = String::new();

        loop {
            print!("cshell> ");
            io::stdout().flush().unwrap();

            input.clear();

            match self.reader.read_line(&mut input) {
                Ok(_) => {
                    let trimmed = input.trim();
                    if trimmed == "exit" {
                        break;
                    }
                    self.history.push(trimmed.to_string());
                }
                Err(error) => println!("Error: {}", error),
            }

            let last_cmd = self.history.last().cloned().unwrap();
            let last_cmd: Vec<&str> = last_cmd.split_whitespace().collect();

            let res = match *last_cmd.first().unwrap_or(&"") {
                "cd" => self.cd(last_cmd),
                "pwd" => self.pwd(),
                _ => Err(io::Error::new(io::ErrorKind::Other, "Invalid command")),
            };

            match res {
                Err(err) => println!("{}", err),
                Ok(status) => println!("{}", status),
            }
        }
    }

    fn cd(&mut self, args: Vec<&str>) -> Result<String, io::Error> {
        match args.len() {
            2 => {
                // TODO: actually change dirs
                let jk = 1;
                Ok("ok".to_string())
            }
            _ => Err(io::Error::new(
                io::ErrorKind::Other,
                "Invalid number of arguments",
            )),
        }
    }

    fn pwd(&mut self) -> Result<String, io::Error> {
        match self.working_dir.to_str() {
            Some(path) => Ok(path.to_string()),
            None => Err(io::Error::new(
                io::ErrorKind::Other,
                "Invalid number of arguments",
            )),
        }
    }
}
