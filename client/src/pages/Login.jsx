import React from "react";
import axios from "axios";
import Cookies from "js-cookie";
import { useState } from "react";
import { Link } from "react-router-dom";
import { Navbar } from "../components";

// ตั้งค่า Axios ในระดับแอป
axios.defaults.withCredentials = true;

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [failed, setFailed] = useState(false);

  async function handleLogin() {
    try {
      const response = await fetch("http://127.0.0.1:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        // credentials: "include",
        body: JSON.stringify({ username, password }),
      });
      const data = await response.json();
      console.log("response status : ", response.status);
      if (response.status === 200) {
        Cookies.set("jwt", data, { path: "/" });
        window.location.replace("/");
      } else {
        console.log("SOMETHING WENT WRONG");
        setFailed(true);
      }
    } catch (error) {
      console.error("Error login failed:", error);
      setFailed(true);
    }
  }

  return (
    <>
      <Navbar />
      <div className="container my-3 py-3">
        <h1 className="text-center">Login</h1>
        <hr />
        <div className="row my-4 h-100">
          <div className="col-md-4 col-lg-4 col-sm-8 mx-auto">
            {/* <form> */}
            <div className="my-3">
              <label for="display-4">Username</label>
              <input
                type="text"
                className="form-control"
                id="floatingInput"
                placeholder="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div className="my-3">
              <label for="floatingPassword display-4">Password</label>
              <input
                type="password"
                className="form-control"
                id="floatingPassword"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <div className="my-3">
              <p>
                New Here?{" "}
                <Link
                  to="/register"
                  className="text-decoration-underline text-info"
                >
                  Register
                </Link>{" "}
              </p>
            </div>

            {failed ? (
              <div className="my-3">
                <p>Username or password incorrect ...</p>
              </div>
            ) : (
              ""
            )}

            <div className="text-center">
              <button
                onClick={handleLogin}
                className="my-2 mx-auto btn btn-dark"
                type="submit"
              >
                Login
              </button>
            </div>
            {/* </form> */}
          </div>
        </div>
      </div>
    </>
  );
};

export default Login;
