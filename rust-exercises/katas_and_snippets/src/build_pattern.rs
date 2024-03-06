#![allow(dead_code, unused_variables, non_camel_case_types)]

#[derive(Debug, Clone)]
pub struct TLSCert {
    key: String,
    cert: String,
}

type ms = u32;

#[derive(Debug)]
pub struct Server {
    host: String,
    port: u16,
    tls: Option<TLSCert>,
    hot_reload: bool,
    timeout: ms,
}

impl Server {
    fn new(host: String, port: u16) -> ServerBuilder {
        ServerBuilder {
            host,
            port,
            tls: None,
            hot_reload: None,
            timeout: None,
        }
    }
}

pub struct ServerBuilder {
    host: String,
    port: u16,
    tls: Option<TLSCert>,
    hot_reload: Option<bool>,
    timeout: Option<ms>,
}

impl ServerBuilder {
    fn tls(&mut self, tls: TLSCert) -> &mut Self {
        self.tls = Some(tls);
        self
    }
    fn hot_reload(&mut self) -> &mut Self {
        self.hot_reload = Some(true);
        self
    }

    fn timeout(&mut self, timeout: ms) -> &mut Self {
        self.timeout = Some(timeout);
        self
    }
    fn build(&mut self) -> Server {
        Server {
            host: self.host.clone(),
            port: self.port,
            tls: self.tls.clone(),
            hot_reload: self.hot_reload.unwrap_or_default(),
            timeout: self.timeout.unwrap_or(2000),
        }
    }
}

#[derive(Debug, Builder)]
pub struct Server2 {
    host: String,
    port: u16,
    tls: Option<TLSCert>,
    hot_reload: bool,
    timeout: ms,
}

#[test]
fn test_build_pattern() {
    let host = "localhost".to_owned();
    let port: u16 = 8080;
    let tls = TLSCert {
        key: "...".to_owned(),
        cert: "...".to_owned(),
    };

    let basic_server: Server = Server::new(host.clone(), port).build();
    let tls_server: Server = Server::new(host.clone(), port).tls(tls.clone()).build();
    let advanced_server: Server = Server::new(host.to_owned(), port)
        .tls(tls.clone())
        .hot_reload()
        .timeout(4000)
        .build();

    let server2 = Server2Builder::default()
        .host(host.to_owned())
        .port(port)
        .tls(Some(tls.clone()))
        .timeout(5000)
        .hot_reload(false)
        .build()
        .unwrap();

    println!("{:?}", server2);
}
