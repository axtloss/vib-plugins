// Copyright 2024 axtlos <axtlos@disroot.org>
// SPDX-License-Identifier: GPL-3.0-ONLY

use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use serde::{Deserialize, Serialize};

#[allow(non_snake_case)]
#[derive(Serialize, Deserialize)]
struct PkgModule {
    name: String,
    r#type: String,

    Packages: Option<Vec<String>>,
    #[serde(default)]
    ExtraFlags: Vec<String>
}

#[no_mangle]
pub unsafe extern "C" fn BuildModule(module_interface: *const c_char, recipe_interface: *const c_char) -> *mut c_char {
    let recipe = CStr::from_ptr(recipe_interface);
    let module = CStr::from_ptr(module_interface);
    let cmd = build_module(String::from_utf8_lossy(module.to_bytes()).to_string(), String::from_utf8_lossy(recipe.to_bytes()).to_string());
    let rtrn =  CString::new(cmd).expect("ERROR: CString::new failed");
    rtrn.into_raw()
}

#[warn(non_snake_case)]
fn build_module(module_interface: String, _: String) -> String {
    let module: PkgModule = match serde_json::from_str(&module_interface) {
	Ok(v) => v,
	Err(error) => return format!("ERROR: {}", error),
    };
    
    format!("pkg install -y {} {}", module.ExtraFlags.join(" "), module.Packages.unwrap_or(vec!["".to_string()]).join(" "))
}
