fn shout_question(s: &str) -> bool {
    s == s.to_uppercase() && s.chars().any(|c| c.is_alphabetic()) && s.ends_with('?')
}

pub fn reply(message: &str) -> &str {
    let m = message.trim();
    match m {
        "" => "Fine. Be that way!",
        m if shout_question(m) => "Calm down, I know what I'm doing!",
        m if m.ends_with("?") => "Sure.",
        m if m == m.to_uppercase() && m.chars().any(|c| c.is_alphabetic()) => "Whoa, chill out!",
        _ => "Whatever.",
    }
}
