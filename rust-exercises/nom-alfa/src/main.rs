use crate::alfa::parse;

mod alfa;

fn main() {
    let result = parse(
        "policy ciaociao {
            rule abc123 {
                deny
            }
        }",
    );
    dbg!(result.unwrap());
}
