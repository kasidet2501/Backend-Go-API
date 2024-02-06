import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { NavLink } from "react-router-dom";
import { useSelector } from "react-redux";

const Navbar = () => {
  const [role, setRole] = useState("");

  async function handleLogout() {
    try {
      const response = await fetch("http://127.0.0.1:8080/logout", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
      });
      setRole("");
      Cookies.remove("jwt");
      window.location.replace("/");
    } catch (error) {
      console.error("Error login failed:", error);
    }
  }

  useEffect(() => {
    const getRole = async () => {
      const response = await fetch("http://127.0.0.1:8080/role", {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
          // Authorization : `Bearer ${Cookies.get('jwt')}`
        },
        // credentials: "include",
      });
      const data = await response.json();
      setRole(data);
    };
    getRole();
  }, [role]);

  const state = useSelector((state) => state.handleCart);
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light py-3 sticky-top">
      <div className="container">
        <NavLink className="navbar-brand fw-bold fs-4 px-2" to="/">
          {" "}
          Go Ecommerce
        </NavLink>
        <button
          className="navbar-toggler mx-2"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>

        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav m-auto my-2 text-center">
            <li className="nav-item">
              <NavLink className="nav-link" to="/">
                Home{" "}
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink className="nav-link" to="/product">
                Products
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink className="nav-link" to="/about">
                About
              </NavLink>
            </li>

            {role == "admin" ? (
              <li className="nav-item">
                <NavLink className="nav-link" to="/producttable">
                  Admin
                </NavLink>
              </li>
            ) : (
              ""
            )}
          </ul>
          <div className="buttons text-center">
            {role == "admin" ? (
              <NavLink
                onClick={() => handleLogout()}
                className="btn btn-outline-dark m-2"
              >
                <i className="fa fa-sign-in-alt mr-1"></i> Logout
              </NavLink>
            ) : (
              ""
            )}

            {role == "user" ? (
              <NavLink
                onClick={() => handleLogout()}
                className="btn btn-outline-dark m-2"
              >
                <i className="fa fa-sign-in-alt mr-1"></i> Logout
              </NavLink>
            ) : (
              ""
            )}

            {role == "user" ? (
              <NavLink to="/cart" className="btn btn-outline-dark m-2">
                <i className="fa fa-cart-shopping mr-1"></i> Cart (
                {state.length}){" "}
              </NavLink>
            ) : (
              ""
            )}

            {role == "" ? (
              <NavLink to="/login" className="btn btn-outline-dark m-2">
                <i className="fa fa-sign-in-alt mr-1"></i> Login
              </NavLink>
            ) : (
              ""
            )}

            {role == "" ? (
              <NavLink to="/register" className="btn btn-outline-dark m-2">
                <i className="fa fa-user-plus mr-1"></i> Register
              </NavLink>
            ) : (
              ""
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
