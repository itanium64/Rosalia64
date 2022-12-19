pub fn format_qualifying_predicate(qp: u64) -> String {
    if qp == 0 {
        return String::from("     ")
    } 

    let mut formatted  = format!("(p{})", qp);

    //if number isn't 2 digits add another space
    //so that single digit numbers and stuff don't indent stuff wrong
    if formatted.len() < 5 {
        formatted.push_str(&String::from(" "))
    }

    return formatted
}