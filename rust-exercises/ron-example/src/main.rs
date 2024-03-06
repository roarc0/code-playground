use std::collections::HashMap;

use base64::engine::{general_purpose::STANDARD as BASE64, Engine};
use ron::ser::{to_string_pretty, PrettyConfig};
use serde::{Deserialize, Deserializer, Serialize, Serializer};

#[derive(Debug, Deserialize, Serialize, Default)]
struct MyStruct {
    boolean: bool,
    float: f32,
}

#[derive(Debug, Deserialize, Serialize, Default)]
struct Example {
    name: String,
    #[serde(with = "Base64")]
    data: Vec<u8>,
    boolean: bool,
    float: f32,
    map: HashMap<u8, char>,
    nested: MyStruct,
    option: Option<String>,
    tuple: (u32, u32),
}

enum Base64 {}

impl Base64 {
    fn serialize<S: Serializer>(data: &[u8], serializer: S) -> Result<S::Ok, S::Error> {
        serializer.serialize_str(&BASE64.encode(data))
    }

    fn deserialize<'de, D: Deserializer<'de>>(deserializer: D) -> Result<Vec<u8>, D::Error> {
        let base64_str = <&str>::deserialize(deserializer)?;
        BASE64.decode(base64_str).map_err(serde::de::Error::custom)
    }
}

fn main() {
    let x: MyStruct = ron::from_str("(boolean: true, float: 1.23)").unwrap();
    println!("RON: {}", ron::to_string(&x).unwrap());

    let ex = Example {
        name: "hello world".into(),
        data: vec![64, 23, 77, 123, 97, 112, 56, 72, 73, 73, 12],
        nested: x,
        ..Default::default()
    };

    let pretty = PrettyConfig::new()
        .depth_limit(2)
        .struct_names(true)
        .separate_tuple_members(true)
        .enumerate_arrays(true);
    let s = to_string_pretty(&ex, pretty).expect("Serialization failed");

    println!("RON: {}", s);
}
