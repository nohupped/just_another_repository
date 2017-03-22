//use std::io;
fn main() {
/*    println!("Hello, world!");
    println!("Guess the number");
    println!("Enter your guess");
    let mut guess = String::new();
//    io::stdin().read_line(&mut guess).expect("Failed to read line");
    println!("You guessed {}", guess);


    let x = 10;
    {
        let y = 20;
        println!("value of x is {} and y  is {}", x, y);
    }
    //println!("value of x is {} and y  is {}", x, y);
    */
    let x = 11;
    println!("value of x is {}", x);
    print_num(x);

    static mut y:i32 = 20;
    unsafe {
        print_till_zero(&mut y);
    }
}

fn print_num(x: i32) {
    static mut COUNT: i32 = 10;
    unsafe {
        COUNT -= 1;

        println!("Number is {}", x);
        while COUNT != 0 {
            print_num(x);
        }
    }

}

fn print_till_zero(x: &mut i32) {
    *x -= 1;
    println!("Variable is {}", *x);
    while *x != 0 {
        print_till_zero(x);
    }
}
