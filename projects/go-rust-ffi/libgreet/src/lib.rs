use std::ffi::CStr;

#[allow(clippy::not_unsafe_ptr_arg_deref)]
#[no_mangle]
pub extern "C" fn greet(name: *const libc::c_char) {
    if !name.is_null() {
        let name_cstr = unsafe { CStr::from_ptr(name) };
        if let Ok(name_str) = name_cstr.to_str() {
            println!("Hello, {}!", name_str);
            return;
        }
    }
    println!("No valid name provided!");
}

#[cfg(test)]
pub mod test {
    use super::*;
    use std::ffi::CString;
    use std::ptr;

    #[test]
    fn test_null_greet() {
        greet(ptr::null());
    }

    #[test]
    fn test_greet() {
        greet(CString::new("world").unwrap().into_raw());
    }
}
