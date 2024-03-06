use smol::{io, net, prelude::*, Unblock};

fn main() {
    _ = get_page();
    
    smol::block_on(async {
        let (s, ctrl_c) = async_channel::bounded(100);
        ctrlc::set_handler(move || {
            s.try_send(()).ok();
        })
        .unwrap();
        println!("Waiting for Ctrl-C...");
        ctrl_c.recv().await.ok();
        println!("Done!");
    });
}

fn get_page() -> io::Result<()> {
    smol::block_on(async {
        let mut stream = net::TcpStream::connect("example.com:80").await?;
        let req = b"GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n";
        stream.write_all(req).await?;

        let mut stdout = Unblock::new(std::io::stdout());
        io::copy(stream, &mut stdout).await?;
        Ok(())
    })
}
