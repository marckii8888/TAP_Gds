import Navbar from "react-bootstrap/Navbar";
import Container from "react-bootstrap/Container";
import React from "react";

class NavBar extends React.Component {
  render() {
    return (
      <Navbar className="color-nav" variant="dark">
        <Container>
          <Navbar.Brand>TAP GDS Technical Assessment</Navbar.Brand>
        </Container>
      </Navbar>
    );
  }
}

export default NavBar