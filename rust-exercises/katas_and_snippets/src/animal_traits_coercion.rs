trait Animal {
    fn speak(&self);
}

#[derive(Debug, Clone)]
struct Cat {
    name: String,
}

impl Animal for Cat {
    fn speak(&self) {
        println!("meow!")
    }
}

#[derive(Debug, Clone)]
struct Dog {
    name: String,
}

impl Animal for Dog {
    fn speak(&self) {
        println!("woof!")
    }
}

#[test]
fn test_animal_traits() {
    let peanut = "peanut".to_owned();
    let oreo = "oreo".to_owned();
    let jax = "jax";

    print_animal_name(&oreo);
    print_animal_name(jax);

    let cat = Box::new(Cat { name: peanut });
    let dog = Box::new(Dog { name: oreo });
    let dog2 = Dog {
        name: jax.to_owned(),
    };

    print_dog(&dog);
    print_dog(&dog2);

    //let animals: Vec<Box<dyn Animal>> = vec![cat, dog];
    let animals: [Box<dyn Animal>; 2] = [cat, dog];

    animal_sounds(&animals);
}

// prefer str because of coercion from String to str
fn print_animal_name(name: &str) {
    println!("{}", name)
}

// prefer having reference to base struct instead of ref to box. box coerced to ref dog
fn print_dog(dog: &Dog) {
    println!("{:?}", dog)
}

// prefer passing slice instead of vec. vec coerced to slice
fn animal_sounds(animals: &[Box<dyn Animal>]) {
    for a in animals {
        a.speak();
    }
}
