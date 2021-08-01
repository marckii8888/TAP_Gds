import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import "bootstrap/dist/css/bootstrap.min.css";
import NavBar from "./components/navbar"
import UrlField from "./components/url_field"
import Footer from "./components/footer"

class App extends React.Component {
  render() {
    return (
      <div className="body">
        <NavBar/>
        <UrlField/>
        <Footer/>
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("root"));