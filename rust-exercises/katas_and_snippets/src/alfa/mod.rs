use nom::{
    branch::alt,
    bytes::complete::{is_not, tag, tag_no_case},
    character::complete::{alpha1, alphanumeric1, multispace0, multispace1},
    combinator::{map, map_res, recognize},
    error::ParseError,
    multi::{many0, many0_count, many1},
    sequence::{delimited, pair},
    IResult,
};

#[derive(Debug)]
pub enum Node<'a> {
    Namespace(Namespace<'a>),
    PolicySet(PolicySet<'a>),
    Policy(Policy<'a>),
}

impl<'a> Node<'a> {
    // Methods for easy access to the inner data
    pub fn as_namespace(&self) -> Option<&Namespace<'a>> {
        if let Node::Namespace(namespace) = self {
            Some(namespace)
        } else {
            None
        }
    }

    pub fn as_policy_set(&self) -> Option<&PolicySet<'a>> {
        if let Node::PolicySet(policy_set) = self {
            Some(policy_set)
        } else {
            None
        }
    }

    pub fn as_policy(&self) -> Option<&Policy<'a>> {
        if let Node::Policy(policy) = self {
            Some(policy)
        } else {
            None
        }
    }
}

#[derive(Debug)]
pub struct Policy<'a> {
    id: &'a str,
    rules: Vec<Rule<'a>>,
}

#[derive(Debug)]
pub struct PolicySet<'a> {
    id: &'a str,
    policies: Vec<Policy<'a>>,
}

#[derive(Debug)]
pub struct Namespace<'a> {
    id: &'a str,
    nodes: Vec<Node<'a>>,
}

#[derive(Debug)]
pub enum RuleResult {
    Permit,
    Deny,
}

#[derive(Debug)]
pub struct Rule<'a> {
    id: &'a str,
    result: RuleResult,
}

pub fn parse(input: &str) -> IResult<&str, Node> {
    node(input)
}

fn namespace_node(input: &str) -> IResult<&str, Node> {
    let (input, id) = parse_named_identifier(input, "namespace")?;
    let (input, _) = ws(tag("{"))(input)?;
    let (input, nodes) = nodes(input)?;
    let (input, _) = ws(tag("}"))(input)?;
    Ok((input, Node::Namespace(Namespace { id, nodes })))
}

fn nodes(input: &str) -> IResult<&str, Vec<Node<'_>>> {
    many0(node)(input)
}

fn node(input: &str) -> IResult<&str, Node> {
    alt((namespace_node, policy_set_node, policy_node))(input)
}

fn policy_set_node(input: &str) -> IResult<&str, Node> {
    let (input, id) = parse_named_identifier(input, "policy_set")?;
    let (input, _) = ws(tag("{"))(input)?;
    let (input, policies) = policies(input)?;
    let (input, _) = ws(tag("}"))(input)?;
    Ok((input, Node::PolicySet(PolicySet { id, policies })))
}

fn policies(input: &str) -> IResult<&str, Vec<Policy<'_>>> {
    many1(policy)(input)
}

fn policy_node(input: &str) -> IResult<&str, Node> {
    let (input, policy) = policy(input)?;
    Ok((input, Node::Policy(policy)))
}

fn policy(input: &str) -> IResult<&str, Policy> {
    let (input, id) = parse_named_identifier(input, "policy")?;
    let (input, _) = ws(tag("{"))(input)?;
    let (input, rules) = rules(input)?;
    let (input, _) = ws(tag("}"))(input)?;
    Ok((input, Policy { id, rules }))
}

fn rules(input: &str) -> IResult<&str, Vec<Rule<'_>>> {
    many1(rule)(input)
}

fn rule(input: &str) -> IResult<&str, Rule> {
    let (input, id) = parse_named_identifier(input, "rule")?;
    let (input, _) = ws(tag("{"))(input)?;

    let (input, result) = map_res(alt((tag("permit"), tag("deny"))), |res: &str| match res {
        "permit" => Ok(RuleResult::Permit),
        "deny" => Ok(RuleResult::Deny),
        _ => Err(nom::Err::Error(nom::error::Error::new(
            input,
            nom::error::ErrorKind::Tag,
        ))),
    })(input)?;
    let (input, _) = ws(tag("}"))(input)?;
    Ok((input, Rule { id, result }))
}

fn parse_named_identifier<'a>(input: &'a str, name: &'a str) -> IResult<&'a str, &'a str> {
    let (input, _) = ws(comment)(input)?;
    let (input, _) = tag_no_case(name)(input)?;
    let (input, _) = multispace1(input)?;
    identifier(input)
}

fn identifier<'a>(input: &'a str) -> IResult<&'a str, &'a str> {
    recognize(pair(
        alt((alpha1, tag("_"))),
        many0_count(alt((alphanumeric1, tag("_")))),
    ))(input)
}

fn ws<'a, F: 'a, O, E: ParseError<&'a str>>(
    inner: F,
) -> impl FnMut(&'a str) -> IResult<&'a str, O, E>
where
    F: Fn(&'a str) -> IResult<&'a str, O, E>,
{
    delimited(multispace0, inner, multispace0)
}

pub fn comment<'a>(i: &'a str) -> IResult<&'a str, &'a str> {
    alt((
        map(pair(tag("//"), is_not("\n\r")), |(_, c)| c), // For // end-of-line comments
        delimited(tag("/*"), is_not("*/"), tag("*/")),    // For /* */ block comments
        multispace0,
    ))(i)
}

#[test]
fn test_parse() {
    let result = parse(
        "
        namespace root {
        /* test_policy:
           is a policy that always returns permit
        */
        policy test_policy {
            rule a {
                permit
            }
        }
        namespace sf {
            policy_set hello {
                policy test_policy1 {
                    // this is my rule a
                    rule b {
                        permit
                    }
                }
                policy test_policy2 {
                    // this is my rule a
                    rule a {
                        deny
                    }
                }
            }
        }
    }",
    );
    dbg!(result.unwrap());
}
