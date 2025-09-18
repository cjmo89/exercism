fn shout_question(s: &str) -> bool {
    s.chars().all(|c| c == c.to_ascii_uppercase())
        && s.chars().any(|c| c.is_alphabetic())
        && s.chars().last() == Some('?')
}

pub fn reply(message: &str) -> &str {
    let m = message.trim();
    match m {
        "" => "Fine. Be that way!",
        m if shout_question(m) => "Calm down, I know what I'm doing!",
        m if m.chars().last() == Some('?') => "Sure.",
        m if m == m.to_uppercase() && m.chars().any(|c| c.is_alphabetic()) => "Whoa, chill out!",
        _ => "Whatever.",
    }
}
