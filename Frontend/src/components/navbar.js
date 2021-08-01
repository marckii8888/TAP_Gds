import Navbar from "react-bootstrap/Navbar";
import Container from "react-bootstrap/Container";
import React from "react";

class NavBar extends React.Component {
  render() {
    return (
      <Navbar expand="lg" className="color-nav" variant="dark">
        <Container>
          <Navbar.Brand>URL Shortener</Navbar.Brand>
          <Navbar.Collapse
            id="responsive-navbar-nav"
            className="justify-content-end"
          >
            <Navbar.Text>
              Done By:{" "}
              <a href="https://github.com/marckii8888">Marcus Cheong</a>
            </Navbar.Text>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    );
  }
}

export default NavBar