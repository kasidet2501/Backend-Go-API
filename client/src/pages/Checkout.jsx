import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { Footer, Navbar } from "../components";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
const Checkout = () => {
  const state = useSelector((state) => state.handleCart);
  const [total, setTotal] = useState("");

  const [fname, setFname] = useState("");
  const [lname, setLname] = useState("");
  const [email, setEmail] = useState("");

  const [address, setAddress] = useState("");

  let componentMounted = true;

  useEffect(() => {
    const getUser = async () => {
      const response1 = await fetch("http://127.0.0.1:8080/username", {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
      });
      const dataUsername = await response1.json();

      const response2 = await fetch(
        `http://127.0.0.1:8080/user/${dataUsername}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            authorization: `Bearer ${Cookies.get("jwt")}`,
          },
        }
      );
      if (componentMounted) {
        const data = await response2.json();
        setFname(data.name.firstname);
        setLname(data.name.lastname);
        setEmail(data.email);
      }
      return () => {
        componentMounted = false;
      };
    };
    getUser();
  }, [setTotal]);

  // async function handleSubmit() {
  const handleFormSubmit = async (event) => {
    try {
      const response = await fetch("http://127.0.0.1:8080/username", {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
      });
      const dataUsername = await response.json();

      const priceParse = parseFloat(total);
      const values = {
        username: dataUsername,
        firstname: fname,
        lastname: lname,
        email: email,
        address: address,
        carts: state,
        price: priceParse,
      };

      await fetch(`http://127.0.0.1:8080/order/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
        body: JSON.stringify(values),
      });
      window.location.replace("http://localhost:3000/product");
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const ShowCheckout = () => {
    let subtotal = 0;
    let shipping = 30.0;
    let totalItems = 0;
    state.map((item) => {
      return (subtotal += item.price * item.qty);
    });

    state.map((item) => {
      return (totalItems += item.qty);
    });
    return (
      <>
        <div className="col-md-5 col-lg-4 order-md-last">
          <div className="card mb-4">
            <div className="card-header py-3 bg-light">
              <h5 className="mb-0">Order Summary</h5>
            </div>
            <div className="card-body">
              <ul className="list-group list-group-flush">
                <li className="list-group-item d-flex justify-content-between align-items-center border-0 px-0 pb-0">
                  Products ({totalItems})<span>${Math.round(subtotal)}</span>
                </li>
                <li className="list-group-item d-flex justify-content-between align-items-center px-0">
                  Shipping
                  <span>${shipping}</span>
                </li>
                <li className="list-group-item d-flex justify-content-between align-items-center border-0 px-0 mb-3">
                  <div>
                    <strong>Total amount</strong>
                  </div>
                  <span>
                    <strong>
                      ${Math.round(subtotal + shipping)}
                      {setTotal(Math.round(subtotal + shipping))}
                    </strong>
                  </span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </>
    );
  };
  return (
    <>
      <Navbar />
      <div className="container my-3 py-3">
        <h1 className="text-center">Checkout</h1>
        <hr />
        <div className="container py-5">
          <div className="row my-4">
            <ShowCheckout />
            <div className="col-md-7 col-lg-8">
              <div className="card mb-4">
                <div className="card-header py-3">
                  <h4 className="mb-0">Billing address</h4>
                </div>
                <div className="card-body">
                  <form
                    className="needs-validation"
                    onSubmit={handleFormSubmit}
                  >
                    <div className="row g-3">
                      <div className="col-sm-6 my-1">
                        <label for="firstName" className="form-label">
                          First name
                        </label>
                        <input
                          type="text"
                          className="form-control"
                          id="firstName"
                          value={fname}
                          placeholder=""
                          onChange={(e) => setFname(e.target.value)}
                          required
                        />
                        <div className="invalid-feedback">
                          Valid first name is required.
                        </div>
                      </div>

                      <div className="col-sm-6 my-1">
                        <label for="lastName" className="form-label">
                          Last name
                        </label>
                        <input
                          type="text"
                          className="form-control"
                          id="lastName"
                          value={lname}
                          placeholder=""
                          onChange={(e) => setLname(e.target.value)}
                          required
                        />
                        <div className="invalid-feedback">
                          Valid last name is required.
                        </div>
                      </div>

                      <div className="col-12 my-1">
                        <label for="email" className="form-label">
                          Email
                        </label>
                        <input
                          type="email"
                          className="form-control"
                          id="email"
                          value={email}
                          placeholder="you@example.com"
                          onChange={(e) => setEmail(e.target.value)}
                          required
                        />
                        <div className="invalid-feedback">
                          Please enter a valid email address for shipping
                          updates.
                        </div>
                      </div>

                      <div className="col-12 my-1">
                        <label for="address" className="form-label">
                          Address
                        </label>
                        <input
                          type="text"
                          className="form-control"
                          id="address"
                          value={address}
                          placeholder="1234 Main St"
                          onChange={(e) => setAddress(e.target.value)}
                          required
                        />
                        {console.log("Address : ", address)}
                        <div className="invalid-feedback">
                          Please enter your shipping address.
                        </div>
                      </div>
                    </div>

                    <hr className="my-4" />

                    <button className="w-100 btn btn-primary " type="submit">
                      Continue to checkout
                    </button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Checkout;
