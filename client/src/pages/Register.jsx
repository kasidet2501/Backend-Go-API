import React from "react";
import Cookies from "js-cookie";
import { useState } from "react";
import { Footer, Navbar } from "../components";
import { Link } from "react-router-dom";

const Register = () => {
  const [firstname, setFirstname] = useState("");
  const [lastname, setLastname] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [failed, setFailed] = useState("");

  async function handleRegister() {
    try {
      const response = await fetch("http://127.0.0.1:8080/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        // credentials: "include",
        body: JSON.stringify({
          username: username,
          email: email,
          password: password,
          name: {
            firstname: firstname,
            lastname: lastname,
          },
        }),
      });
      const data = await response.json();
      if (response.status === 200) {
        try {
          const response = await fetch("http://127.0.0.1:8080/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            // credentials: "include",
            body: JSON.stringify({ username, password }),
          });
          const data = await response.json();
          //   Cookies.set("jwt", data, { path: "/" });
          window.location.replace("/product");
        } catch (error) {
          console.error("Error login failed:", error);
        }
      } else {
        console.log("SOMETHING WENT WRONG");
        //   setFailed(true);
      }
    } catch (error) {
      console.error("Error Register failed:", error);
    }
  }

  return (
    <>
      <Navbar />
      <div className="container my-3 py-3">
        <h1 className="text-center">Register</h1>
        <hr />
        <div class="row my-4 h-100">
          <div className="col-md-4 col-lg-4 col-sm-8 mx-auto">
            <form onSubmit={handleRegister}>
              <div class="form my-3">
                <label for="Name">Firstname</label>
                <input
                  type="text"
                  class="form-control"
                  id="firstname"
                  placeholder="Enter Your Firstname"
                  value={firstname}
                  onChange={(e) => setFirstname(e.target.value)}
                  required
                />
              </div>
              <div class="form my-3">
                <label for="Name">Lastname</label>
                <input
                  type="text"
                  class="form-control"
                  id="lastname"
                  placeholder="Enter Your Lastname"
                  value={lastname}
                  onChange={(e) => setLastname(e.target.value)}
                  required
                />
              </div>
              <div class="form my-3">
                <label for="Name">Username</label>
                <input
                  type="text"
                  class="form-control"
                  id="username"
                  placeholder="Enter Your Username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  required
                />
              </div>
              <div class="form my-3">
                <label for="Email">Email address</label>
                <input
                  type="email"
                  class="form-control"
                  id="Email"
                  placeholder="name@example.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>
              <div class="form  my-3">
                <label for="Password">Password</label>
                <input
                  type="password"
                  class="form-control"
                  id="Password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
              </div>
              <div className="my-3">
                <p>
                  Already has an account?{" "}
                  <Link
                    to="/login"
                    className="text-decoration-underline text-info"
                  >
                    Login
                  </Link>{" "}
                </p>
              </div>
              <div className="text-center">
                <button
                  class="my-2 mx-auto btn btn-dark"
                  type="submit"
                  // onClick={handleRegister}
                >
                  Register
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Register;
