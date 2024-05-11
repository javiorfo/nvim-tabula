use std::convert::Infallible;

use nvim_oxi::{Dictionary, Function, Object};

#[nvim_oxi::module]
fn coagula_rs() -> nvim_oxi::Result<Dictionary> {
    let execute = Function::from_fn(|_: ()| {
        Ok::<_, Infallible>(())
    });

    let dictionary = Dictionary::from_iter([
        ("execute", Object::from(execute)),
    ]);

    Ok(dictionary)
}
