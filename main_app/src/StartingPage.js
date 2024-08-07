import React, { useState } from "react";
import logo from "./static/logo1.png";
import bg from "./static/bg1.jpeg";
import { useNavigate } from "react-router-dom";

function StartingPage() {
  const [subscriberCount, setSubscriberCount] = useState(0);
  const navigate = useNavigate();
  const handleSubmit = (event) => {
    event.preventDefault();
    setSubscriberCount(subscriberCount + 1);
    navigate("/main");
  };

  return (
    <main className="w">
      <section>
        <div className="auto-w-pc">
          <h2 className="text-center">Subscribe to CodeDemonz Newsletter</h2>
          <img
            src={logo}
            alt="CodeDemonz Logo"
            className="d-block mx-auto auto-w"
          />
        </div>
        <div>
          <p className="mx-auto w-75">
            The <b>CodeDemonz Newsletter</b> features latest events and articles
            about gaming and technology that suites the interest of the
            developer and audiences.
          </p>

          <form
            action="/submit"
            method="POST"
            className="mx-auto h-75 w-75 centerize"
            onSubmit={handleSubmit}
          >
            <div className="mb-3 clamp-shit">
              <div className="d-flex flex-row">
                <input
                  type="email"
                  name="email"
                  id="email"
                  className="p-2 form-control"
                  aria-describedby="emailHelp"
                  placeholder="Enter your email"
                  required
                />
                <button type="submit" className="p-2 btn btn-primary">
                  Submit
                </button>
              </div>
              <div id="emailHelp" className="form-text text-center">
                {subscriberCount} subscribers are waiting for you to join
              </div>
            </div>
          </form>
        </div>
      </section>
    </main>
  );
}

export default StartingPage;
