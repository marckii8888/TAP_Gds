import React from "react";
import Container from "react-bootstrap/Container";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import axios from "axios";

class UrlField extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      longurl: "",
      shorturl: "",
    };
  }

  shortenUrl = async (e) => {
    e.preventDefault();
    let payload = { original_url: this.state.longurl };
    let res = await axios.post("http://127.0.0.1:8081/api/shorten", payload, {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });
    let data = await res.data;
    this.setState({ shorturl: data.message });
  };

  handleInputChange = (e) => {
    this.setState({ [e.target.name]: e.target.value });
  };

  render() {
    return (
      <Container fluid className="container" id="pagebody">
        <Row>
          <h1>Title</h1>
        </Row>
        <Row>
          <Form onSubmit={this.shortenUrl}>
            <Row>
              <Col md={{ span: 5, offset: 3 }}>
                <FloatingLabel
                  controlId="floatingInput"
                  label="Long URL"
                  className="mb-3"
                >
                  <Form.Control
                    type="url"
                    name="longurl"
                    placeholder="example.com"
                    onChange={this.handleInputChange}
                  />
                </FloatingLabel>
              </Col>
              <Col xs="auto">
                <Button id="submitBtn" type="submit">
                  Shorten!
                </Button>
              </Col>
            </Row>
          </Form>
        </Row>
        <Row xs="auto">
          {this.state.shorturl && (
            <div className="mx-auto">
              <p className="mb-0">Your shortened url below:</p>
              <Form.Control
                id="para"
                type="text"
                placeholder={this.state.shorturl}
                readOnly
              />
              <Button
                onClick={() => {
                  navigator.clipboard.writeText(this.state.shorturl);
                }}
              >
                Copy
              </Button>
            </div>
          )}
        </Row>
      </Container>
    );
  }
}

export default UrlField;
